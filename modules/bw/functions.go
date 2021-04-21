package bw

import (
	"encoding/json"
	"fmt"
	"github.com/kamushadenes/swayit/common"
	"github.com/kamushadenes/swayit/config"
	"os"
	"os/exec"
	"strings"
	"time"
)

func maskPassword(pwd string) string {
	var masked []string
	
	masked = append(masked, string(pwd[0]))
	for i := 1; i <= len(pwd)-2; i++ {
		masked = append(masked, "*")
	}
	masked = append(masked, string(pwd[len(pwd)-1]))
	
	return strings.Join(masked, "")
}

func formatItemsWofi(tasks []*Item) []string {
	var formatted []string

	for _, item := range tasks {
		formatted = append(formatted, fmt.Sprintf("<b>%s</b>\n%s", item.Name, item.Login.Username))
	}

	return formatted
}

func GetItems() ([]*Item, error) {
	var items []*Item
	cached, expired, err := common.GetCache(module.GetSlug(), "items.json", time.Duration(config.SwayItConfig.BW.MaxAge)*time.Minute)
	if err == nil && !expired {
		err = json.Unmarshal(cached, &items)
		if err == nil {
			return items, nil
		}
	}
	
	common.Notify("BW", "Refreshing items")

	cmd := exec.Command(config.SwayItConfig.BW.Command, "list", "items")
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, fmt.Sprintf("BW_SESSION=%s", config.SwayItConfig.BW.SessionToken))
	
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	
	err = json.Unmarshal(out, &items)
	
	b, err := json.Marshal(&items)
	if err != nil {
		return nil, err
	}

	err = common.SetCache(module.GetSlug(), "items.json", b)

	return items, err
}