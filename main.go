package main

import (
	"User/model"
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
)

func main() {
	var cfg model.Config
	if _, err := toml.DecodeFile("config.toml", &cfg); err != nil {
		log.Fatalln(err)
	}
	fmt.Println(cfg)
}
