package common

import "encoding/json"

type WaybarOutput struct {
	Text    string `json:"text"`
	Alt     string `json:"alt,omitempty"`
	Class   string `json:"class,omitempty"`
	Tooltip string `json:"tooltip,omitempty"`
}

func (w *WaybarOutput) ToJSON() string {
	b, err := json.Marshal(w)
	if err == nil {
		return string(b)
	}

	return ""
}
