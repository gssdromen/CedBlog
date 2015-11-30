package CedStruct

import (
	"CedBlog/CedUtils"
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	BaseUrl     string `json:"base_url"`
	BasePath    string `json:"base_path"`
	ConfigName  string `json:"config_name"`
	ConfigDir   string `json:"config_dir"`
	TemplateDir string `json:"template_dir"`
	PublicDir   string `json:"public_dir"`
	ContentDir  string `json:"content_dir"`
	Favicon     string `json:"favicon"`
}

const ConfigPath = "cb_config.json"

func (c *Config) New() error {
	c.BaseUrl = "127.0.0.1:9099"
	c.BasePath = "/Users/cedricwu/Document/Go/src/CedBlog/"
	c.TemplateDir = "templates/"
	c.PublicDir = "public/"
	c.ContentDir = "content/"
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
		defer file.Close()
		data, err := ioutil.ReadAll(file)
		if err != nil {
			return nil
		}
		err = json.Unmarshal(data, c)
		if err != nil {
			return nil
		}
		return c
	} else {
		c.New()
		c.Read()
	}
	return c
}
