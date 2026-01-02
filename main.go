package main


import (
	"fmt"

	"github.com/pjjimiso/gator/internal/config"
)


func main() {
	fmt.Println("Hello, world!")
	cfg, err := config.Read()
	if err != nil {
		fmt.Errorf("failed to read config file: %v", err)
	}

	err = cfg.SetUser()
	if err != nil { 
		fmt.Errorf("failed to update config file: %v", err)
	}
}
