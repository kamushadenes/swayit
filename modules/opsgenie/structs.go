package opsgenie

import "time"

type GetScheduleResponse struct {
	Data struct {
		Description string `json:"description"`
		Enabled     bool   `json:"enabled"`
		ID          string `json:"id"`
		Name        string `json:"name"`
		OwnerTeam   struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"ownerTeam"`
		Rotations []struct {
			ID           string `json:"id"`
			Length       int    `json:"length"`
			Name         string `json:"name"`
			Participants []struct {
				ID       string `json:"id"`
				Type     string `json:"type"`
				Username string `json:"username"`
			} `json:"participants"`
			StartDate       time.Time `json:"startDate"`
			TimeRestriction struct {
				Restrictions []struct {
					EndDay    string `json:"endDay"`
					EndHour   int    `json:"endHour"`
					EndMin    int    `json:"endMin"`
					StartDay  string `json:"startDay"`
					StartHour int    `json:"startHour"`
					StartMin  int    `json:"startMin"`
				} `json:"restrictions"`
				Type string `json:"type"`
			} `json:"timeRestriction"`
			Type string `json:"type"`
		} `json:"rotations"`
		Timezone string `json:"timezone"`
	} `json:"data"`
	RequestId string  `json:"requestId"`
	Took      float64 `json:"took"`
}

type ListScheduleResponse struct {
	Data []struct {
		Description string `json:"description"`
		Enabled     bool   `json:"enabled"`
		ID          string `json:"id"`
		Name        string `json:"name"`
		OwnerTeam   struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"ownerTeam"`
		Rotations []interface{} `json:"rotations"`
		Timezone  string        `json:"timezone"`
	} `json:"data"`
	Expandable []string `json:"expandable"`
	RequestId  string   `json:"requestId"`
	Took       float64  `json:"took"`
}

type ListOnCallResponse struct {
	Data struct {
		Parent struct {
			Enabled bool   `json:"enabled"`
			ID      string `json:"id"`
			Name    string `json:"name"`
		} `json:"_parent"`
		OnCallRecipients []string `json:"onCallRecipients"`
	} `json:"data"`
	RequestId string  `json:"requestId"`
	Took      float64 `json:"took"`
}
