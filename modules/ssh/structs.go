package ssh

type HostEntry struct {
	Hostname string `yaml:"Hostname"`
	User     string `yaml:"User"`
}

type HostsFile struct {
	Hosts map[string]*HostEntry `yaml:"hosts"`
}
