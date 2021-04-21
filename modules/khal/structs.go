package khal

import (
	"fmt"
	"strings"
	"time"
)

type EventSorter []*Event

func (evs EventSorter) Len() int      { return len(evs) }
func (evs EventSorter) Swap(i, j int) { evs[i], evs[j] = evs[j], evs[i] }
func (evs EventSorter) Less(i, j int) bool {
	idt, _ := evs[i].GetStartTime()
	jdt, _ := evs[j].GetStartTime()
	return idt.Before(jdt)
}

type Event struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	StartDate   string `json:"start-date"`
	StartTime   string `json:"start-time"`
	EndDate     string `json:"end-date"`
	EndTime     string `json:"end-time"`
	Location    string `json:"location"`
}

func (ev *Event) GetMeetingLink() (string, error) {
	desc := ev.Location + ev.Description

	if strings.Contains(desc, "https://") {
		m := meetRegex.FindString(desc)
		if m != "" {
			return m, nil
		}
		m = zoomRegex.FindString(desc)
		if m != "" {
			return m, nil
		}
	}

	return "", fmt.Errorf("no link found")
}

func (ev *Event) GetStartDate() (time.Time, error) {
	return time.Parse("2006-01-02", ev.StartDate)
}

func (ev *Event) GetEndDate() (time.Time, error) {
	return time.Parse("2006-01-02", ev.EndDate)
}

func (ev *Event) GetStartTime() (time.Time, error) {
	dayDt, err := ev.GetStartDate()
	if err != nil {
		return dayDt, err
	}

	startTime := ev.StartTime
	if startTime == "" {
		startTime = "00:00"
	}

	return time.Parse("2006-01-02 15:04:05", fmt.Sprintf("%s %s:00", dayDt.Format("2006-01-02"), startTime))
}

func (ev *Event) GetEndTime() (time.Time, error) {
	dayDt, err := ev.GetEndDate()
	if err != nil {
		return dayDt, err
	}

	endTime := ev.EndTime
	if endTime == "" {
		endTime = "00:00"
	}

	return time.Parse("2006-01-02 15:04:05", fmt.Sprintf("%s %s:00", dayDt.Format("2006-01-02"), endTime))
}
