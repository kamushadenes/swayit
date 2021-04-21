package externalIp

import (
	"fmt"
	"github.com/kamushadenes/swayit/common"
	"io/ioutil"
	"net/http"
	"strings"
)

func run(w *common.WaybarOutput) error {
	resp, err := http.Get("https://checkip.amazonaws.com")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	w.Text = fmt.Sprintf("\uF0AC %s", strings.TrimSpace(string(b)))

	return nil
}
