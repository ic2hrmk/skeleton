package base

import (
	"flag"
	"log"

	"github.com/joho/godotenv"
)

//
// Default env. file for testing
//
const defaultEnvFile = "test.env"

const (
	hostEnvVar = "HOST_ADDRESS"
)

func init() {
	var envFile string

	flag.StringVar(&envFile, "env", defaultEnvFile, "Env. file")
	flag.Parse()

	if err := godotenv.Load(defaultEnvFile); err != nil {
		log.Fatal(err)
	}

	log.Println("ENV FILE LOADED:", envFile)
}
