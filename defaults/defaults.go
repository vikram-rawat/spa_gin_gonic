package defaults

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

// Config struct to hold connection values
type Config struct {
	Database struct {
		Development struct {
			Driver   string
			Host     string
			Port     int
			User     string
			Password string
			Dbname   string
			Sslmode  string
		}
	}
}

// GetYaml for recording basic
func GetYaml() Config {
	conf := Config{}
	file, err := ioutil.ReadFile("connection.yaml")
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(file, &conf)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return conf
}
