package common

import (
	"io/ioutil"
	"strings"
)

func IsACConnected() bool {
	c, _ := ioutil.ReadFile("/sys/class/power_supply/AC/online")

	switch strings.TrimSpace(string(c)) {
	case "1":
		return true
	case "0":
		return false
	}

	return true
}

