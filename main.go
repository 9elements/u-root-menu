package main

import (
	"fmt"

	"github.com/9elements/u-root-menu/pkg/config"
)

var Version = "0.0.0-dev"

func main() {
	fmt.Println("Version:", Version)
	cfg, err := config.LoadConfig()
}
