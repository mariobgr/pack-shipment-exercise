package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

var EnvConfig envVars

type envVars struct {
	AvailablePackSizes []int
}

// LoadConfigContinuously reloads environment values on preset interval (configurable on startup)
func LoadConfigContinuously(done chan bool) {
	// read config for the first time
	loadConfig()

	// default update interval
	updateInterval := 60

	// override with value from .env if any
	if i, err := strconv.Atoi(os.Getenv("UPDATE_INTERVAL_SECONDS")); err == nil {
		updateInterval = i
	}

	// start ticker on the set interval
	ticker := time.NewTicker(time.Duration(updateInterval) * time.Second)

	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				loadConfig()
			}
		}
	}()
}

// loadConfig overrides any existing ENV values with the ones found in .env file
func loadConfig() {
	err := godotenv.Overload(".env")
	if err != nil {
		panic("error loading .env file")
	}

	packSizes := os.Getenv("CONFIG_PACK_SIZES")
	if packSizes == "" {
		panic("no pack sizes specified")
	}

	EnvConfig.AvailablePackSizes = make([]int, 0)
	for _, size := range strings.Split(packSizes, ",") {
		intSize := strToInt(size)
		EnvConfig.AvailablePackSizes = append(EnvConfig.AvailablePackSizes, intSize)
	}
}

// strToInt converts string value to integer or panics on error
func strToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Sprintf("cannot convert string value %s to int - %s", s, err))
	}
	return i
}
