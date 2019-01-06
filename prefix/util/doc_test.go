package util

// This file is a part of github.com/thekondor/confita-sugar package.

import (
	"context"
	"fmt"
	confita "github.com/heetch/confita"
	"log"
	"os"
)

func ExampleNewEnvBackend() {
	type Config struct {
		Password string `config:"password"`
	}

	loader := confita.NewLoader(
		/*prefix_util.*/ NewEnvBackend("MY_APP"),
	)

	cfg := Config{Password: "default password"}
	loadAndPrintCfg := func(to *Config) {
		if err := loader.Load(context.Background(), to); nil != err {
			log.Fatalf("Failed to load app configuration: %+v", err)
		}
		fmt.Printf("Password: %s\n", to.Password)
	}

	loadAndPrintCfg(&cfg)

	// Set vendored variable with a prefix MY_APP from the system environment
	os.Setenv("MY_APP_PASSWORD", "password from env")
	loadAndPrintCfg(&cfg)
	// Output:
	// Password: default password
	// Password: password from env
}
