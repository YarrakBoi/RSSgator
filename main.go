package main

import (
	"RSSgator/internal/config"
	"fmt"
	"os"
)

func main() {
	cfg, err := config.ReadConfig()
	if err != nil {
		fmt.Println("Error reading config:", err)
		os.Exit(1)
	}

	fmt.Printf("Initial config: %+v\n", cfg)

	err = config.SetConfig("Kev")
	if err != nil {
		fmt.Println("Error setting config:", err)
		os.Exit(1)
	}

	cfg, err = config.ReadConfig()
	if err != nil {
		fmt.Println("Error reading config after set:", err)
		os.Exit(1)
	}

	fmt.Printf("Updated config: %+v\n", cfg)
}