package common

import (
	"github.com/kamushadenes/swayit/config"
	"io/ioutil"
	"os"
	"path"
	"time"
)

func SetCache(slug string, name string, data []byte) error {
	folder := path.Join(config.SwayItConfig.Paths.Cache, slug)
	if _, err := os.Stat(folder); os.IsNotExist(err) {
		err = os.MkdirAll(folder, 0755)
		if err != nil {
			return err
		}
	}
	
	fpath := path.Join(folder, name)

	return ioutil.WriteFile(fpath, data, 0644)
}

func GetCache(slug string, name string, maxAge time.Duration) ([]byte, bool, error) {
	folder := path.Join(config.SwayItConfig.Paths.Cache, slug)
	fpath := path.Join(folder, name)

	if f, err := os.Stat(fpath); err == nil {
		now := time.Now()
		mod := f.ModTime()
		diff := now.Sub(mod)

		data, err := ioutil.ReadFile(fpath)
		return data, diff > maxAge, err
	} else {
		return nil, true, err
	}
}