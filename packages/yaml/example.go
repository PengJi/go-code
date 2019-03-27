package main

import (
	"fmt"
	"io/ioutil"

	"github.com/alecthomas/gometalinter/_linters/src/gopkg.in/yaml.v2"
)

type Nginx struct {
	Port int `yaml:"Port"`
	LogPath string `yaml:LogPath`
	Path string `yaml:"Path"`
}

type Config struct {
	Name string `yaml:"SiteName"`
	Addr string `yaml:"SiteAddr"`
	Https string `yaml:"Https"`
	SiteNginx Nginx `yaml:"Nginx"`
}

func main(){
	var setting Config
	config, err := ioutil.ReadFile("./example.yaml")
	if err != nil {
		fmt.Print(err)
	}
	yaml.Unmarshal(config, &setting)

	fmt.Println(setting.Name)
	fmt.Println(setting.Addr)
	fmt.Println(setting.Https)
	fmt.Println(setting.SiteNginx.Port)
	fmt.Println(setting.SiteNginx.LogPath)
	fmt.Println(setting.SiteNginx.Path)
}
