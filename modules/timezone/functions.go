package timezone

import (
	"fmt"
	"github.com/kamushadenes/swayit/common"
	"io/ioutil"
	"os"
	"runtime"
	"strings"
)

// https://socketloop.com/tutorials/golang-display-list-of-timezones-with-gmt

var zoneDirs = map[string][]string{
	"android":   {"/system/usr/share/zoneinfo/"},
	"darwin":    {"/usr/share/zoneinfo/"},
	"dragonfly": {"/usr/share/zoneinfo/"},
	"freebsd":   {"/usr/share/zoneinfo/"},
	"linux":     {"/usr/share/zoneinfo/", "/etc/zoneinfo"},
	"netbsd":    {"/usr/share/zoneinfo/"},
	"openbsd":   {"/usr/share/zoneinfo/"},
	"solaris":   {"/usr/share/lib/zoneinfo/"},
}

func InSlice(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}

func ReadTZDir(zoneDir, path string) []string {
	var timeZones []string

	files, _ := ioutil.ReadDir(zoneDir + path)
	for _, f := range files {
		if f.Name() != strings.ToUpper(f.Name()[:1])+f.Name()[1:] {
			continue
		}
		if f.IsDir() {
			timeZones = append(timeZones, ReadTZDir(zoneDir, path+"/"+f.Name())...)
		} else {
			timeZones = append(timeZones, (path + "/" + f.Name())[1:])
		}
	}

	return timeZones
}

func ListTimeZones() []string {
	if _, ok := zoneDirs[runtime.GOOS]; !ok {
		common.Logger.Error().Err(fmt.Errorf("unsupported platform")).Msg("an error has occurred")
		os.Exit(0)
	}

	var timeZones []string

	for _, zoneDir := range zoneDirs[runtime.GOOS] {
		timeZones = append(timeZones, ReadTZDir(zoneDir, "")...)
	}
	
	return timeZones
}
