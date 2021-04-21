package khal

import (
	"encoding/json"
	"fmt"
	"github.com/kennygrant/sanitize"
	"html"
	"os/exec"
	"sort"
	"strings"
	"time"

	"github.com/kamushadenes/swayit/common"
	"github.com/kamushadenes/swayit/config"
)

func run(w *common.WaybarOutput) error {
	var today = time.Now()
	var todayStr = today.Format("2006-01-02")
	var tomorrow = today.AddDate(0, 0, 1)
	var nextAhead = today.AddDate(0, 0, config.SwayItConfig.Khal.DaysAhead)

	var nextAheadStr = nextAhead.Format("2006-01-02")

	cmd := exec.Command("khal", "list", "-df", "", "-f", eventFormat, "now", nextAheadStr)

	output, err := cmd.Output()
	if err != nil {
		return err
	}

	var events []*Event

	var organizedEvents = make(map[string][]*Event)

	rawEvents := strings.Split(string(output), eventDelimiter)

	for _, ev := range rawEvents {
		fields := strings.Split(ev, descriptionDelimiter)
		if len(fields) != 3 {
			continue
		}

		var event Event
		err = json.Unmarshal([]byte(fields[0]), &event)
		if err != nil {
			continue
		}

		event.Title = strings.TrimSpace(fields[1])
		if config.SwayItConfig.Khal.StripClockwise {
			event.Title = strings.TrimSuffix(event.Title, " (via Clockwise)")
		}
		event.Title = html.EscapeString(event.Title)

		event.Description = strings.TrimSpace(fields[2])

		events = append(events, &event)
	}

	for _, ev := range events {
		organizedEvents[ev.StartDate] = append(organizedEvents[ev.StartDate], ev)
	}

	if td, ok := organizedEvents[todayStr]; ok {
		_ = common.SaveData(module.GetSlug(), "currentMeetingLink", "")
		cnt := 0
		for _, ev := range td {
			var evName string
			if ev.StartDate == "" && ev.EndTime == "" {
				if cnt < len(td) {
					cnt++
					continue
				}
				evName = fmt.Sprintf("All Day - %s", ev.Title)
			} else {
				evName = fmt.Sprintf("%s-%s - %s", ev.StartTime, ev.EndTime, ev.Title)
			}
			w.Text = fmt.Sprintf("\uF073 %s", evName)

			w.Tooltip = fmt.Sprintf("<b>Next Meeting</b>\n%s", evName)

			_ = common.SaveData(module.GetSlug(), "currentMeeting", fmt.Sprintf("%s - %s", todayStr, evName))
			_ = common.SaveData(module.GetSlug(), "currentMeetingMarkdown", fmt.Sprintf("%sT%s:00_%s.md", todayStr, ev.StartTime, strings.Trim(sanitize.Name(ev.Title), "-")))
			_ = common.SaveData(module.GetSlug(), "currentMeetingHeader", fmt.Sprintf("# Meeting: %s - %s - %s", todayStr, ev.StartTime, ev.Title))

			if link, err := ev.GetMeetingLink(); err == nil {
				_ = common.SaveData(module.GetSlug(), "currentMeetingLink", link)
				w.Tooltip += fmt.Sprintf("\n%s", link)
			}

			break
		}
	}

	var keys []string

	for day := range organizedEvents {
		keys = append(keys, day)
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	for _, day := range keys {
		evs := organizedEvents[day]
		var dayName string
		dayDt, err := time.Parse("2006-01-02", day)
		if err != nil {
			continue
		}
		if dayDt.Year() == today.Year() && dayDt.Month() == today.Month() && dayDt.Day() == today.Day() {
			dayName = "Today"
		} else if dayDt.Year() == tomorrow.Year() && dayDt.Month() == tomorrow.Month() && dayDt.Day() == tomorrow.Day() {
			dayName = "Tomorrow"
		} else {
			dayName = dayDt.Format("Monday")
		}

		w.Tooltip += fmt.Sprintf("\n\n<b>%s, %s</b>", dayName, day)

		evsSorted := EventSorter(evs)
		sort.Sort(evsSorted)

		for _, ev := range evsSorted {
			var evName string

			if ev.StartTime == "" && ev.EndTime == "" {
				evName = fmt.Sprintf("All Day - %s", ev.Title)
			} else {
				evName = fmt.Sprintf("%s-%s - %s", ev.StartTime, ev.EndTime, ev.Title)
			}

			w.Tooltip += fmt.Sprintf("\n%s", evName)
		}
	}

	return nil
}
