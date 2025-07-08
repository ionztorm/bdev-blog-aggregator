package main

import (
	"fmt"
	"gator/internal/config"
)

func main() {

	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
	}

	cfg.SetUser("Leon")

	cfg, err = config.Read()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%v", cfg)
}
