package main

import (
	"fmt"

	"github.com/marekbrze/gator/internal/config"
)

func main() {
	configFile, err := config.Read()
	if err != nil {
		fmt.Println(err)
	}

	if err := configFile.SetUser(); err != nil {
		fmt.Println(err)
	}
	configFile, err = config.Read()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(configFile.DBURL, configFile.CurrentUserName)
}
