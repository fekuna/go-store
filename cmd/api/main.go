package main

import (
	"log"
	"os"

	"github.com/fekuna/go-store/config"
	"github.com/fekuna/go-store/internal/server"
	"github.com/fekuna/go-store/pkg/utils"
)

// @title Go Example REST API
// @version 1.0
// @description Example Golang REST API
// @contact_name Alfan Almunawar
// @contact_url https://github.com/fekuna
// @contact_email almunawar.alfan@gmail.com
// @BasePath /api/v1
func main() {
	log.Println("Starting api server")

	configPath := utils.GetConfigPath(os.Getenv("config"))

	cfgFile, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}

	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		log.Fatalf("ParseConfig: %v", err)
	}

	s := server.NewServer(cfg)
	if err = s.Run(); err != nil {
		log.Fatal(err)
	}
}