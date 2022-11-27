/*
Reads configuration from and places it in Config structure.
Configuration are readed from yml files and environmental variables.
*/
package config

import (
	"os"
	"log"
	"gopkg.in/yaml.v2"
	"github.com/kelseyhightower/envconfig"
)

// Structure include all needed variables
type Config struct {
    Server struct {
        Port string `yaml:"port", envconfig:"SERVER_PORT"`
        Host string `yaml:"host", envconfig:"SERVER_HOST"`
    } `yaml:"server"`
    Db struct {
        Username string `yaml:"user", envconfig:"DB_USERNAME"`
        Password string `yaml:"pass", envconfig:"DB_PASSWORD"`
				Host string `yaml:"host", envconfig:"DB_HOST"`
				Port string `yaml:"port", envconfig:"DB_PORT"`
				Dbname string `yaml:"dbname", envconfig:"DB_DBNAME"`
    } `yaml:"database"`
}

// Parses configs ("shared" and "appName") and env.
// Returns Config structure or error.
func ParseConfig() (Config , error) {
	var cfg Config
	log.Println("Reading shared config")
	f, err := os.Open("configs/shared.yml")
	if err != nil {
		return cfg, err
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		return cfg, err
	}

  log.Println("Reading app config")
  f2, err := os.Open("configs/admin.yml")
  if err != nil {
    return cfg, err
  }
  defer f2.Close()

  decoder = yaml.NewDecoder(f2)
  err = decoder.Decode(&cfg)
  if err != nil {
    return cfg, err
  }

  err = envconfig.Process("", &cfg) 
  if err != nil { 
		return cfg, err
  }
	return cfg, nil
}

