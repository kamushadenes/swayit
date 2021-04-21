package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type ConfigChess struct {
	Enabled           bool   `yaml:"enabled"`
	SuspendOnBattery  bool   `yaml:"suspendOnBattery"`
	Interval          int64  `yaml:"interval" default:"60"`
	IntervalOnBattery int64  `yaml:"intervalOnBattery" default:"300"`
	URL               string `yaml:"url" default:"https://api.chess.com/pub"`
	Username          string `yaml:"username"`
}

type ConfigCrypto struct {
	Enabled           bool   `yaml:"enabled"`
	SuspendOnBattery  bool   `yaml:"suspendOnBattery"`
	Interval          int64  `yaml:"interval" default:"900"`
	IntervalOnBattery int64  `yaml:"intervalOnBattery" default:"3600"`
	CMCApiKey         string `yaml:"apiKey"`
	BaseCurrency      string `yaml:"baseCurrency" default:"USD"`
	Coins             []struct {
		Symbol      string `yaml:"symbol"`
		Icon        string `yaml:"icon"`
		TooltipOnly bool   `yaml:"tooltipOnly"`
	} `yaml:"coins"`
}

type ConfigPower struct {
	Enabled           bool   `yaml:"enabled"`
	SuspendOnBattery  bool   `yaml:"suspendOnBattery"`
	Interval          int64  `yaml:"interval" default:"1"`
	IntervalOnBattery int64  `yaml:"intervalOnBattery" default:"5"`
	Battery           string `yaml:"battery" default:"BAT0"`
}

type ConfigPrometheus struct {
	Enabled           bool   `yaml:"enabled"`
	SuspendOnBattery  bool   `yaml:"suspendOnBattery"`
	Interval          int64  `yaml:"interval" default:"60"`
	IntervalOnBattery int64  `yaml:"intervalOnBattery" default:"300"`
	URL               string `yaml:"url"`
	GrafanaURL        string `yaml:"grafanaUrl"`
	Username          string `yaml:"username"`
	Password          string `yaml:"password"`
	Query             string `yaml:"query" default:"instance_job_severity:probe_success:mean5m"`
}

type ConfigWeather struct {
	Enabled           bool   `yaml:"enabled"`
	SuspendOnBattery  bool   `yaml:"suspendOnBattery"`
	Interval          int64  `yaml:"interval" default:"3600"`
	IntervalOnBattery int64  `yaml:"intervalOnBattery" default:"3600"`
	URL               string `yaml:"url" default:"https://wttr.in/?format=j1"`
	FallbackURL       string `yaml:"fallbackUrl" default:"https://wttr.in/Sao_Paulo?format=j1"`
}

type ConfigMail struct {
	Enabled           bool              `yaml:"enabled"`
	SuspendOnBattery  bool              `yaml:"suspendOnBattery"`
	Interval          int64             `yaml:"interval" default:"5"`
	IntervalOnBattery int64             `yaml:"intervalOnBattery" default:"60"`
	MuPath            string            `yaml:"muPath" default:"/usr/bin/mu"`
	MainQuery         string            `yaml:"mainQuery" default:"maildir:/.*.INBOX/ AND flag:unread AND NOT flag:trashed"`
	Queries           map[string]string `yaml:"queries"`
}

type ConfigExternalIP struct {
	Enabled           bool  `yaml:"enabled"`
	SuspendOnBattery  bool  `yaml:"suspendOnBattery"`
	Interval          int64 `yaml:"interval" default:"60"`
	IntervalOnBattery int64 `yaml:"intervalOnBattery" default:"300"`
}

type ConfigForex struct {
	Enabled                bool   `yaml:"enabled"`
	SuspendOnBattery       bool   `yaml:"suspendOnBattery"`
	Interval               int64  `yaml:"interval" default:"3600"`
	IntervalOnBattery      int64  `yaml:"intervalOnBattery" default:"3600"`
	OpenExchangeRatesAppID string `yaml:"apiKey"`
	BaseCurrency           string `yaml:"baseCurrency" default:"USD"`
	BaseSymbol             string `yaml:"baseSymbol" default:"$"`
	Coins                  []struct {
		Symbol      string `yaml:"symbol"`
		Icon        string `yaml:"icon"`
		TooltipOnly bool   `yaml:"tooltipOnly"`
	} `yaml:"coins"`
}

type ConfigIntelGPU struct {
	Enabled           bool   `yaml:"enabled"`
	SuspendOnBattery  bool   `yaml:"suspendOnBattery"`
	Interval          int64  `yaml:"interval" default:"5"`
	IntervalOnBattery int64  `yaml:"intervalOnBattery" default:"30"`
	Command           string `yaml:"command" default:"/usr/bin/intel_gpu_top"`
	SudoCommand       string `yaml:"sudoCommand" default:"/usr/bin/sudo"`
}

type ConfigItau struct {
	Enabled           bool  `yaml:"enabled"`
	SuspendOnBattery  bool  `yaml:"suspendOnBattery"`
	Interval          int64 `yaml:"interval" default:"900"`
	IntervalOnBattery int64 `yaml:"intervalOnBattery" default:"3600"`
}

type ConfigKhal struct {
	Enabled           bool     `yaml:"enabled"`
	SuspendOnBattery  bool     `yaml:"suspendOnBattery"`
	Interval          int64    `yaml:"interval" default:"10"`
	IntervalOnBattery int64    `yaml:"intervalOnBattery" default:"60"`
	DaysAhead         int      `yaml:"daysAhead" default:"10"`
	IgnoredEvents     []string `yaml:"ignoredEvents"`
	StripClockwise    bool     `yaml:"stripClockwise" default:"true"`
}

type ConfigNordVPN struct {
	Enabled           bool   `yaml:"enabled"`
	SuspendOnBattery  bool   `yaml:"suspendOnBattery"`
	Interval          int64  `yaml:"interval" default:"1"`
	IntervalOnBattery int64  `yaml:"intervalOnBattery" default:"5"`
	Command           string `yaml:"command" default:"nordvpn"`
}

type ConfigNubank struct {
	Enabled           bool  `yaml:"enabled"`
	SuspendOnBattery  bool  `yaml:"suspendOnBattery"`
	Interval          int64 `yaml:"interval" default:"900"`
	IntervalOnBattery int64 `yaml:"intervalOnBattery" default:"3600"`
}

type ConfigOpsGenie struct {
	Enabled           bool   `yaml:"enabled"`
	SuspendOnBattery  bool   `yaml:"suspendOnBattery"`
	Interval          int64  `yaml:"interval" default:"300"`
	IntervalOnBattery int64  `yaml:"intervalOnBattery" default:"900"`
	Email             string `yaml:"email"`
	Profiles          []struct {
		Name      string   `yaml:"name"`
		Token     string   `yaml:"token"`
		Schedules []string `yaml:"schedules"`
	} `yaml:"profiles"`
}

type ConfigTodoist struct {
	Enabled           bool     `yaml:"enabled"`
	SuspendOnBattery  bool     `yaml:"suspendOnBattery"`
	Interval          int64    `yaml:"interval" default:"1"`
	IntervalOnBattery int64    `yaml:"intervalOnBattery" default:"5"`
	Command           string   `yaml:"command" default:"todoist"`
	ExcludedProjects  []string `yaml:"excludedProjects"`
}

type ConfigFan struct {
	Enabled           bool  `yaml:"enabled"`
	SuspendOnBattery  bool  `yaml:"suspendOnBattery"`
	Interval          int64 `yaml:"interval" default:"1"`
	IntervalOnBattery int64 `yaml:"intervalOnBattery" default:"5"`
}

type ConfigSSH struct {
	Enabled           bool     `yaml:"enabled"`
	SuspendOnBattery  bool     `yaml:"suspendOnBattery"`
	Interval          int64    `yaml:"interval" default:"1"`
	IntervalOnBattery int64    `yaml:"intervalOnBattery" default:"5"`
	Terminal          string   `yaml:"terminal" default:"kitty"`
	TerminalArgs      []string `yaml:"terminalArgs"`
	Command           string   `yaml:"command" default:"ssh"`
	DigitalOcean      struct {
		Contexts []struct {
			Name  string `yaml:"name"`
			Token string `yaml:"token"`
		} `yaml:"contexts"`
	} `yaml:"digitalOcean"`
}

type ConfigBW struct {
	Enabled           bool   `yaml:"enabled"`
	SuspendOnBattery  bool   `yaml:"suspendOnBattery"`
	Interval          int64  `yaml:"interval" default:"1"`
	IntervalOnBattery int64  `yaml:"intervalOnBattery" default:"5"`
	Command           string `yaml:"command" default:"bw"`
	MaxAge            int64  `yaml:"maxAge" default:"900"`
	SessionToken      string `yaml:"sessionToken"`
}

type ConfigPaths struct {
	Cache    string `yaml:"cache"`
	LastRun  string `yaml:"lastRun"`
	Tmp      string `yaml:"tmp"`
	Output   string `yaml:"output"`
	Data     string `yaml:"data"`
	Flags    string `yaml:"flags"`
	Notes    string `yaml:"notes"`
	Snippets string `yaml:"snippets"`
}

type ConfigEditor struct {
	Terminal string   `yaml:"terminal" default:"kitty"`
	Command  string   `yaml:"command" default:"vim"`
	Args     []string `yaml:"args"`
}

type Config struct {
	Chess      *ConfigChess      `yaml:"chess"`
	Paths      *ConfigPaths      `yaml:"paths"`
	Power      *ConfigPower      `yaml:"power"`
	Weather    *ConfigWeather    `yaml:"weather"`
	Mail       *ConfigMail       `yaml:"mail"`
	Prometheus *ConfigPrometheus `yaml:"prometheus"`
	Crypto     *ConfigCrypto     `yaml:"crypto"`
	ExternalIP *ConfigExternalIP `yaml:"externalIp"`
	Fan        *ConfigFan        `yaml:"fan"`
	Forex      *ConfigForex      `yaml:"forex"`
	IntelGPU   *ConfigIntelGPU   `yaml:"intelGpu"`
	Itau       *ConfigItau       `yaml:"itau"`
	Khal       *ConfigKhal       `yaml:"khal"`
	NordVPN    *ConfigNordVPN    `yaml:"nordvpn"`
	Nubank     *ConfigNubank     `yaml:"nubank"`
	OpsGenie   *ConfigOpsGenie   `yaml:"opsgenie"`
	Todoist    *ConfigTodoist    `yaml:"todoist"`
	Editor     *ConfigEditor     `yaml:"editor"`
	SSH        *ConfigSSH        `yaml:"ssh"`
	BW         *ConfigBW         `yaml:"bw"`
}

func (c *Config) Save() error {
	d, err := yaml.Marshal(c)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(cfgFile, d, 0644)

	return err
}
