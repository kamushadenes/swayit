package config

import (
	"github.com/creasty/defaults"
	"gopkg.in/yaml.v2"
)

func (c *ConfigChess) UnmarshalYAML(unmarshal func(interface{}) error) error {
	defaults.Set(c)

	type plain ConfigChess
	if err := unmarshal((*plain)(c)); err != nil {
		return err
	}

	return nil
}

func (c *ConfigCrypto) UnmarshalYAML(unmarshal func(interface{}) error) error {
	defaults.Set(c)

	type plain ConfigCrypto
	if err := unmarshal((*plain)(c)); err != nil {
		return err
	}

	return nil
}

func (c *ConfigPower) UnmarshalYAML(unmarshal func(interface{}) error) error {
	defaults.Set(c)

	type plain ConfigPower
	if err := unmarshal((*plain)(c)); err != nil {
		return err
	}

	return nil
}

func (c *ConfigPrometheus) UnmarshalYAML(unmarshal func(interface{}) error) error {
	defaults.Set(c)

	type plain ConfigPrometheus
	if err := unmarshal((*plain)(c)); err != nil {
		return err
	}

	return nil
}

func (c *ConfigWeather) UnmarshalYAML(unmarshal func(interface{}) error) error {
	defaults.Set(c)

	type plain ConfigWeather
	if err := unmarshal((*plain)(c)); err != nil {
		return err
	}

	return nil
}

func (c *ConfigMail) UnmarshalYAML(unmarshal func(interface{}) error) error {
	defaults.Set(c)

	type plain ConfigMail
	if err := unmarshal((*plain)(c)); err != nil {
		return err
	}

	return nil
}

func (c *ConfigExternalIP) UnmarshalYAML(unmarshal func(interface{}) error) error {
	defaults.Set(c)

	type plain ConfigExternalIP
	if err := unmarshal((*plain)(c)); err != nil {
		return err
	}

	return nil
}

func (c *ConfigForex) UnmarshalYAML(unmarshal func(interface{}) error) error {
	defaults.Set(c)

	type plain ConfigForex
	if err := unmarshal((*plain)(c)); err != nil {
		return err
	}

	return nil
}

func (c *ConfigIntelGPU) UnmarshalYAML(unmarshal func(interface{}) error) error {
	defaults.Set(c)

	type plain ConfigIntelGPU
	if err := unmarshal((*plain)(c)); err != nil {
		return err
	}

	return nil
}

func (c *ConfigItau) UnmarshalYAML(unmarshal func(interface{}) error) error {
	defaults.Set(c)

	type plain ConfigItau
	if err := unmarshal((*plain)(c)); err != nil {
		return err
	}

	return nil
}

func (c *ConfigKhal) UnmarshalYAML(unmarshal func(interface{}) error) error {
	defaults.Set(c)

	type plain ConfigKhal
	if err := unmarshal((*plain)(c)); err != nil {
		return err
	}

	return nil
}

func (c *ConfigNordVPN) UnmarshalYAML(unmarshal func(interface{}) error) error {
	defaults.Set(c)

	type plain ConfigNordVPN
	if err := unmarshal((*plain)(c)); err != nil {
		return err
	}

	return nil
}

func (c *ConfigNubank) UnmarshalYAML(unmarshal func(interface{}) error) error {
	defaults.Set(c)

	type plain ConfigNubank
	if err := unmarshal((*plain)(c)); err != nil {
		return err
	}

	return nil
}

func (c *ConfigOpsGenie) UnmarshalYAML(unmarshal func(interface{}) error) error {
	defaults.Set(c)

	type plain ConfigOpsGenie
	if err := unmarshal((*plain)(c)); err != nil {
		return err
	}

	return nil
}

func (c *ConfigTodoist) UnmarshalYAML(unmarshal func(interface{}) error) error {
	defaults.Set(c)

	type plain ConfigTodoist
	if err := unmarshal((*plain)(c)); err != nil {
		return err
	}

	return nil
}

func (c *ConfigFan) UnmarshalYAML(unmarshal func(interface{}) error) error {
	defaults.Set(c)

	type plain ConfigFan
	if err := unmarshal((*plain)(c)); err != nil {
		return err
	}

	return nil
}

func (c *ConfigPaths) UnmarshalYAML(unmarshal func(interface{}) error) error {
	defaults.Set(c)

	type plain ConfigPaths
	if err := unmarshal((*plain)(c)); err != nil {
		return err
	}

	return nil
}

func (c *ConfigEditor) UnmarshalYAML(unmarshal func(interface{}) error) error {
	defaults.Set(c)

	type plain ConfigEditor
	if err := unmarshal((*plain)(c)); err != nil {
		return err
	}

	return nil
}

func (c *ConfigSSH) UnmarshalYAML(unmarshal func(interface{}) error) error {
	defaults.Set(c)

	type plain ConfigSSH
	if err := unmarshal((*plain)(c)); err != nil {
		return err
	}

	return nil
}

func (c *Config) UnmarshalYAML(unmarshal func(interface{}) error) error {
	defaults.Set(c)

	type plain Config
	if err := unmarshal((*plain)(c)); err != nil {
		return err
	}

	return nil
}

func (c *Config) ToYAML() (string, error) {
	data, err := yaml.Marshal(c)

	return string(data), err
}
