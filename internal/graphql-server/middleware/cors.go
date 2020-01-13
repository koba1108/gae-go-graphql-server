package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"time"
)

type CorsConfig struct {
	Allow struct {
		Protocols []string `yaml:"protocols"`
		Methods   []string `yaml:"methods"`
		Headers   []string `yaml:"headers"`
		Domains   []string `yaml:"domains"`
	}
}

var config *CorsConfig

func init() {
	if config == nil {
		setConfig()
	}
}

func CORS() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     config.AllowOrigins(),
		AllowMethods:     config.Allow.Headers,
		AllowHeaders:     config.Allow.Methods,
		AllowCredentials: true,
		AllowWebSockets:  true,
		AllowFiles:       true,
		MaxAge:           12 * time.Hour,
	})
}

func setConfig() {
	buf, err := ioutil.ReadFile("configs/cors.yml")
	if err != nil {
		panic(err)
	}
	if err := yaml.Unmarshal(buf, &config); err != nil {
		panic(err)
	}
}

func (c *CorsConfig) AllowOrigins() []string {
	var origins []string
	for _, protocol := range c.Allow.Protocols {
		for _, domain := range c.Allow.Domains {
			origins = append(origins, protocol + domain)
		}
	}
	return origins
}
