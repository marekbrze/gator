package main

import (
	"fmt"

	"github.com/marekbrze/gator/internal/config"
)

func main() {
	config, err := config.Read()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Config test", config)
}
