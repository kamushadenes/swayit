package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var home, _ = os.UserHomeDir()

var cfgFile string

var SwayItConfig *Config

func NewConfig() *Config {
	var c Config

	c.Paths = &ConfigPaths{
		Cache:   path.Join(home, ".cache", "swayit"),
		LastRun: path.Join(home, ".cache", "swayit", "last_run"),
		Tmp:     path.Join(os.TempDir(), "swayit"),
		Output:  path.Join(os.TempDir(), "swayit", "output"),
		Data:    path.Join(os.TempDir(), "swayit", "data"),
		Flags:   path.Join(os.TempDir(), "swayit", "flags"),
	}

	c.Chess = &ConfigChess{}
	c.Power = &ConfigPower{}
	c.Weather = &ConfigWeather{}
	c.Mail = &ConfigMail{}
	c.Prometheus = &ConfigPrometheus{}
	c.Crypto = &ConfigCrypto{}
	c.ExternalIP = &ConfigExternalIP{}
	c.Fan = &ConfigFan{}
	c.Forex = &ConfigForex{}
	c.IntelGPU = &ConfigIntelGPU{}
	c.Itau = &ConfigItau{}
	c.Khal = &ConfigKhal{}
	c.NordVPN = &ConfigNordVPN{}
	c.Nubank = &ConfigNubank{}
	c.OpsGenie = &ConfigOpsGenie{}
	c.Todoist = &ConfigTodoist{}

	return &c
}

func ReadCobraConfig(cmd *cobra.Command, args []string) {
	ReadConfig()
}

func ReadConfig() {
	if cfgFile == "" {
		cfgFile = filepath.Join(home, ".config", "swayit.yml")
	}

	data, err := ioutil.ReadFile(cfgFile)
	if err == nil {
		err = yaml.Unmarshal(data, &SwayItConfig)
		if err != nil {
			fmt.Printf("error reading config file: %s\n", err.Error())
			os.Exit(1)
		}
	} else {
		SwayItConfig = NewConfig()
		err = SwayItConfig.Save()
		if err != nil {
			fmt.Printf("error reading config file: %s\n", err.Error())
			os.Exit(1)
		}
	}
}

func init() {
	ReadConfig()
}
