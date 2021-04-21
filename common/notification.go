package common

import (
	"fmt"
	"github.com/martinlindhe/notify"
)

func Notify(title string, text string) {
	notify.Notify("SwayIT", fmt.Sprintf("SwayIT - %s", title), text, "")
}

func Alert(title string, text string) {
	notify.Alert("SwayIT", fmt.Sprintf("SwayIT - %s", title), text, "")
}
