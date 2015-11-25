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

func Vaild() bool {

}

func Exist() bool {

}

func New() *Config {

}

func Read() *Config {

}
