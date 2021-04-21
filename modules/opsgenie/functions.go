package opsgenie

import (
	"fmt"
	"github.com/kamushadenes/swayit/common"
	"github.com/kamushadenes/swayit/config"
)

func getHeaders(genieKey string) map[string]string {
	return map[string]string{"Authorization": fmt.Sprintf("GenieKey %s", genieKey)}
}

func listOnCall(schedule string, genieKey string) (*ListOnCallResponse, error) {
	url := common.BuildUrl(apiUrl, "/schedules/", schedule, "on-calls")
	url += "?flat=true"

	var resp ListOnCallResponse

	err := common.GetJsonWithHeaders(url, &resp, getHeaders(genieKey))
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func listSchedules(genieKey string) (*ListScheduleResponse, error) {
	url := common.BuildUrl(apiUrl, "/schedules/")

	var resp ListScheduleResponse

	err := common.GetJsonWithHeaders(url, &resp, getHeaders(genieKey))
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func getSchedule(schedule string, genieKey string) (*GetScheduleResponse, error) {
	url := common.BuildUrl(apiUrl, "/schedules/", schedule)

	var resp GetScheduleResponse

	err := common.GetJsonWithHeaders(url, &resp, getHeaders(genieKey))
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func run(w *common.WaybarOutput) error {
	callCnt := 0

	w.Tooltip = fmt.Sprintf("<b>Source:</b> OpsGenie\n<b>Last Update:</b> %s", module.GetLastRun())

	for _, p := range config.SwayItConfig.OpsGenie.Profiles {
		schedules, err := listSchedules(p.Token)
		if err != nil {
			return err
		}
		for _, sched := range schedules.Data {
			w.Tooltip += fmt.Sprintf("\n\n<b>%s</b>\n", sched.Name)
			onCall, err := listOnCall(sched.ID, p.Token)
			if err != nil {
				return err
			}
			for _, v := range onCall.Data.OnCallRecipients {
				if v == config.SwayItConfig.OpsGenie.Email {
					w.Tooltip += fmt.Sprintf("%s <b>(on-call)</b>\n", v)
					callCnt++
				} else {
					w.Tooltip += fmt.Sprintf("%s\n", v)
				}
			}
			if len(onCall.Data.OnCallRecipients) == 0 {
				w.Tooltip += "No Participants"
			}
		}
	}

	if callCnt > 0 {
		w.Text = fmt.Sprintf("\uf095 %d", callCnt)
	}

	return nil
}
