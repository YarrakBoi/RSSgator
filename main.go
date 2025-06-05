package main

import (
	"RSSgator/internal/config"
	"fmt"
	"os"
)

func main() {
	readCfg, err := config.ReadConfig()
	if err != nil {
		fmt.Println("Error reading config:", err)
		os.Exit(1)
	}

	fmt.Printf("Initial config: %+v\n", readCfg)

	s := config.State{
		Cfg : readCfg,
	}

	cmds := config.Commands{
		CmdNames: make(map[string]func(*config.State, config.Command) error),
	}

	cmds.Register("login", config.HandlerLogin)

	if len(os.Args) < 2 {
		fmt.Println("Require more argument")
		os.Exit(1)
	}
	
	cmd := config.Command{
		Name : os.Args[1],
		Args : os.Args[2:],
	}

	fErr := cmds.Run(&s, cmd)

	if fErr != nil {
		fmt.Println(fErr)
		os.Exit(1)
	}

	fmt.Printf("Updated config: %+v\n", readCfg)

}