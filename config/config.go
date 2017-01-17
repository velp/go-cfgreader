package config
 
import (
  "gopkg.in/yaml.v2"
  "io/ioutil"
  "fmt"
)
 
type СliConfig struct {
  Username string  `yaml:"user"`
  Password string  `yaml:"password"`
  Url      string  `yaml:"url"`
}
 
func check(e error) {
    if e != nil {
        panic(e)
    }
}
 
func NewCliConfig(filepath string) (*СliConfig, error) {
  dat, err := ioutil.ReadFile(filepath)
  check(err)
  fmt.Println(string(dat))
 
  fmt.Println("Now, parsing yaml file to struct!")
 
  res := СliConfig{}
 
  err = yaml.Unmarshal(dat, &res)
  check(err)
  fmt.Printf("--- t:\n%v\n\n", res)
 
  fmt.Println(res.Username)
  return &res, nil
}