package main

import (
	"AuthService/internal/config"
	"fmt"
)

func main() {
	cfg := config.MustLoadConfig()
	fmt.Println(cfg)
}
