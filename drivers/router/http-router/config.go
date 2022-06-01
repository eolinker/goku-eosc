package http_router

import (
	"github.com/eolinker/apinto/plugin"
	router_http "github.com/eolinker/apinto/router/router-http"
	"github.com/eolinker/apinto/service"
	"github.com/eolinker/eosc"
)

//DriverConfig http路由驱动配置
type DriverConfig struct {
	//Driver  string                    `json:"driver" yaml:"driver"`
	Listen  int                       `json:"listen" yaml:"listen" title:"port" desc:"使用端口" default:"0"`
	Method  []string                  `json:"method" yaml:"method" enum:"GET,POST,PUT,DELETE,PATH,HEAD,OPTIONS"`
	Host    []string                  `json:"host" yaml:"host"`
	Rules   []DriverRule              `json:"rules" yaml:"rules"`
	Target  eosc.RequireId            `json:"target" yaml:"target" skill:"github.com/eolinker/apinto/service.service.IService"`
	Disable bool                      `json:"disable" yaml:"disable"`
	Plugins map[string]*plugin.Config `json:"plugins" yaml:"plugins"`
}

//DriverRule http路由驱动配置Rule结构体
type DriverRule struct {
	Location string            `json:"location" yaml:"location"`
	Header   map[string]string `json:"header" yaml:"header"`
	Query    map[string]string `json:"query" yaml:"query"`
}

//Config http路由配置结构体
type Config struct {
	name   string
	port   int
	rules  []router_http.Rule
	host   []string
	target service.IService
}

//Cert http路由驱动配置证书Cert结构体
type Cert struct {
	Key string `json:"key"`
	Crt string `json:"crt"`
}
