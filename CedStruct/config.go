package CedStruct

import (
	"CedBlog/CedUtils"
	"encoding/json"
	"os"
)

type Config struct {
	BaseUrl     string `json:"base_url"`
	BasePath    string `json:"base_path"`
	ConfigName  string `json:"config_name"`
	ConfigDir   string `json:"config_dir"`
	TemplateDir string `json:"template_dir"`
	PublicDir   string `json:"public_dir"`
	Favicon     string `json:"favicon"`
}

const ConfigPath = "cb_config.json"

func (c *Config) New() error {
	c.BaseUrl = "127.0.0.1:9099"
	c.BasePath = "/Users/cedricwu/Document/Go/src/CedBlog/"
	c.TemplateDir = "templates/"
	c.PublicDir = "public/"
	str, err := json.Marshal(*c)
	if err != nil {
		return err
	}
	err = c.save(str)
	if err != nil {
		return err
	}
	return nil
}

func (c *Config) save(bytes []byte) error {
	file, err := os.OpenFile(ConfigPath, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(bytes)
	if err != nil {
		return err
	}
	return nil
}

func (c *Config) Read() *Config {
	if CedUtils.Exist(ConfigPath) {
		file, err := os.Open(ConfigPath)
		if err != nil {
			return nil
		}
		data := make([]byte, 10)
		_, err = file.Read(data)
		if err != nil {
			return nil
		}
		x, y := json.Unmarshal(data, c)
		if x == nil {
			return nil
		}
		return c
	} else {
		c.New()
		c.Read()
	}
}
