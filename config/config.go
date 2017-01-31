package config
 
import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)
 
type Config struct {
	Username string  `yaml:"user"`
	Password string  `yaml:"password"`
	Url      string  `yaml:"url"`
}

func Parse(filepath string) (*Config, error) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	res := Config{}
	err = yaml.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}