package bw

import (
	"fmt"
	"github.com/kamushadenes/swayit/config"
	"os"
	"os/exec"
	"time"
)

type Item struct {
	Object         string      `json:"object"`
	Id             string      `json:"id"`
	OrganizationId interface{} `json:"organizationId"`
	FolderId       string      `json:"folderId"`
	Type           int         `json:"type"`
	Name           string      `json:"name"`
	Favorite       bool        `json:"favorite"`
	Login          struct {
		Uris []struct {
			Match interface{} `json:"match"`
			URI   string      `json:"uri"`
		} `json:"uris"`
		Username             string      `json:"username"`
		Password             interface{} `json:"password"`
		PasswordRevisionDate interface{} `json:"passwordRevisionDate"`
	} `json:"login"`
	CollectionIds []interface{} `json:"collectionIds"`
	RevisionDate  time.Time     `json:"revisionDate"`
}

func (item *Item) GetTOTP() (string, error) {
	cmd := exec.Command(config.SwayItConfig.BW.Command, "get", "totp", item.Id)
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, fmt.Sprintf("BW_SESSION=%s", config.SwayItConfig.BW.SessionToken))
	
	out, err :=  cmd.Output()
	
	return string(out), err
}

func (item *Item) ToWofi() []string {
	var items []string

	items = append(items, fmt.Sprintf("<b>Copy Password</b>\n%s", maskPassword(item.Login.Password.(string))))
	totp, err := item.GetTOTP()
	if err == nil {
		items = append(items, fmt.Sprintf("<b>Copy TOTP</b>\n%s", totp))
	}
	items = append(items, fmt.Sprintf("<b>Copy Username</b>\n%s", item.Login.Username))
	
	for _, uri := range item.Login.Uris {
		items = append(items, fmt.Sprintf("<b>Copy URL</b>\n%s", uri.URI))
	}
	
	return items
}