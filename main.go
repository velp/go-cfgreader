package main
 
import (
  "fmt"

  "github.com/velp/go-cfgreader/config"
)

func main() {
  cfg, err := config.NewCliConfig("./test.yaml")
  if err != nil {
    panic(err)
  }
  fmt.Printf("Config: %#v", cfg)
}