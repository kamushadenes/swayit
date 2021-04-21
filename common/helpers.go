package common

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/kamushadenes/swayit/config"
)

func BuildUrl(u *url.URL, paths ...string) string {
	nu := *u
	paths = append([]string{nu.Path}, paths...)
	nu.Path = path.Join(paths...)
	return nu.String()
}

func GetJsonWithHeaders(u string, i interface{}, headers map[string]string) error {
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return err
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(i)
	if err != nil {
		return err
	}

	return nil
}

func GetJson(u string, i interface{}) error {
	resp, err := http.Get(u)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(i)
	if err != nil {
		return err
	}

	return nil
}

func SaveLastRun(slug string) {
	_ = os.MkdirAll(config.SwayItConfig.Paths.LastRun, 0755)
	fname := path.Join(config.SwayItConfig.Paths.LastRun, slug)

	now := time.Now()

	_ = ioutil.WriteFile(fname, []byte(now.Format("2006-01-02 15:04:05")), 0644)
}

func GetLastRun(slug string) string {
	fname := path.Join(config.SwayItConfig.Paths.LastRun, slug)

	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return time.Now().String()
	}

	return string(b)
}

func ListDirectories(opath string) ([]string, error) {
	var folders []string

	err := filepath.Walk(opath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				name := strings.TrimPrefix(strings.TrimPrefix(path, opath), "/")
				if name == "" {
					return nil
				}
				folders = append(folders, name)
			}
			return nil
		})

	return folders, err
}

func ListFiles(opath string) ([]string, error) {
	var files []string
	
	all, err := ioutil.ReadDir(opath)
	if err != nil {
		return nil, err
	}

	for _, f := range all {
		if !f.IsDir() {
			files = append(files, f.Name())
		}
	}

	return files, err
}

func ListFilesRecurse(opath string) ([]string, error) {
	var files []string

	err := filepath.Walk(opath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				name := strings.TrimPrefix(strings.TrimPrefix(path, opath), "/")
				if name == "" {
					return nil
				}
				files = append(files, name)
			}
			return nil
		})

	return files, err
}
