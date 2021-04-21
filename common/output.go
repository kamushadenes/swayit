package common

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/kamushadenes/swayit/config"
	"github.com/mitchellh/go-ps"
	"io/ioutil"
	"os"
	"os/signal"
	"path"
	"syscall"
	"time"
)

func GetOutput(slug string) (string, error) {
	fname := path.Join(config.SwayItConfig.Paths.Output, slug)

	b, err := ioutil.ReadFile(fname)

	return string(b), err
}

func GetOutputFollow(slug string) error {
	out, err := GetOutput(slug)
	if err != nil {
		Logger.Error().Err(err).Msg("an error has occurred")
	} else {
		fmt.Println(out)
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer watcher.Close()

	fname := path.Join(config.SwayItConfig.Paths.Output, slug)

	ticker := time.NewTicker(500 * time.Millisecond)

	ppid := os.Getppid()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan,
		syscall.SIGCHLD,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	exitChan := make(chan int)

	go func() {
		for {
			select {
			case <-exitChan:
				return
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					b, err := ioutil.ReadFile(fname)
					if err != nil {
						return
					}
					fmt.Println(string(b))
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				Logger.Error().Err(err).Msg("an error has occurred")
			case <-signalChan:
				exitChan <- 0
				return
			case <-ticker.C:
				if proc, err := ps.FindProcess(ppid); proc == nil && err == nil {
					exitChan <- 1
					return
				}
			}
		}
	}()

	err = watcher.Add(fname)
	if err != nil {
		return err
	}

	code := <-exitChan
	os.Exit(code)

	return nil
}

func WriteOutput(slug string, data string) error {
	_ = os.MkdirAll(config.SwayItConfig.Paths.Output, 0755)

	fname := path.Join(config.SwayItConfig.Paths.Output, slug)

	b, err := ioutil.ReadFile(fname)
	if err == nil {
		if data == string(b) {
			return nil
		}
	}

	return ioutil.WriteFile(fname, []byte(data), 0644)
}
