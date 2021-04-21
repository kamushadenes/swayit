package common

import (
	"github.com/kamushadenes/swayit/config"
	"io/ioutil"
	"os"
	"path"
)

func SaveData(slug string, name string, data string) error {
	_ = os.MkdirAll(path.Join(config.SwayItConfig.Paths.Data, slug), 0755)

	fname := path.Join(config.SwayItConfig.Paths.Data, slug, name)

	return ioutil.WriteFile(fname, []byte(data), 0644)
}

func GetData(slug string, name string) string {
	fname := path.Join(config.SwayItConfig.Paths.Data, slug, name)

	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return ""
	}

	return string(b)
}

