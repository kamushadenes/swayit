package ssh

import (
	"bytes"
	"context"
	"fmt"
	"github.com/digitalocean/godo"
	"github.com/kamushadenes/swayit/common"
	"github.com/kamushadenes/swayit/config"
	"github.com/kevinburke/ssh_config"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"
	"sort"
	"strings"
)

func DropletList(ctx context.Context, client *godo.Client) ([]godo.Droplet, error) {
	// create a list to hold our droplets
	var list []godo.Droplet

	// create options. initially, these will be blank
	opt := &godo.ListOptions{}
	for {
		droplets, resp, err := client.Droplets.List(ctx, opt)
		if err != nil {
			return nil, err
		}

		// append the current page's droplets to our list
		list = append(list, droplets...)

		// if we are at the last page, break out the for loop
		if resp.Links == nil || resp.Links.IsLastPage() {
			break
		}

		page, err := resp.Links.CurrentPage()
		if err != nil {
			return nil, err
		}

		// set the page we want for the next request
		opt.Page = page + 1
	}

	return list, nil
}

func UpdateHosts() error {
	err := UpdateDigitalOceanHosts()
	if err != nil {
		return err
	}
	
	return nil
}

func UpdateDigitalOceanHosts() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	
	var byRegion = make(map[string][]godo.Droplet)

	for _, v := range config.SwayItConfig.SSH.DigitalOcean.Contexts {
		client := godo.NewFromToken(v.Token)

		droplets, err := DropletList(context.Background(), client)
		if err != nil {
			return err
		}

		for k := range droplets {
			byRegion[droplets[k].Region.Slug] = append(byRegion[droplets[k].Region.Slug], droplets[k])
		}
	}

	for region := range byRegion {
		var hostsFile HostsFile
		hostsFile.Hosts = make(map[string]*HostEntry)

		for _, droplet := range byRegion[region] {
			ip, err := droplet.PublicIPv4()
			if err != nil {
				common.Logger.Error().Err(err).Msg("an error has occurred")
			}
			hostsFile.Hosts[droplet.Name] = &HostEntry{
				Hostname: ip,
				User:     "root",
			}
		}
		b, err := yaml.Marshal(&hostsFile)
		if err != nil {
			return err
		}
		err = ioutil.WriteFile(path.Join(home, ".ssh", "digitalocean", fmt.Sprintf("%s.yml", region)), b, 0644)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetHosts() ([]string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadFile(path.Join(home, ".ssh", "config"))
	if err != nil {
		return nil, err
	}

	cnf, err := ssh_config.Decode(bytes.NewReader(b))
	if err != nil {
		return nil, err
	}

	var hosts = make(map[string]bool)

	for _, v := range cnf.Hosts {
		for _, p := range v.Patterns {
			host := p.String()
			if !strings.Contains(host, "*") {
				if strings.Contains(host, "@") {
					host = strings.Split(host, "@")[1]
				}
				if host != "github.com" {
					hosts[host] = true
				}
			}
		}
	}

	var hostsStr []string

	for k := range hosts {
		hostsStr = append(hostsStr, k)
	}
	sort.Slice(hostsStr, func(i, j int) bool {
		return hostsStr[i] < hostsStr[j]
	})

	return hostsStr, nil
}
