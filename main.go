package main


import (
	"fmt"

	"github.com/pjjimiso/gator/internal/config"
)


func main() {
	fmt.Println("Hello, world!")
	cfg, err := config.Read()
	if err != nil {
		fmt.Errorf("unable to get config: %v", err)
	}

	fmt.Println(cfg)
}
