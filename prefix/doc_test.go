package prefix

// This file is a part of github.com/thekondor/confita-sugar package.

import (
	"context"
	"fmt"
	confita "github.com/heetch/confita"
	confita_backend_env "github.com/heetch/confita/backend/env"
	"log"
	"os"
)

func ExampleWithPrefix() {
	type Config struct {
		Token string `config:"token"`
	}

	loader := confita.NewLoader(
		/*vendor.*/ WithPrefix("MY_COMPANY", confita_backend_env.NewBackend()),
	)

	cfg := Config{Token: "default token"}
	loadAndPrintCfg := func(to *Config) {
		if err := loader.Load(context.Background(), to); nil != err {
			log.Fatalf("Failed to load app configuration: %+v", err)
		}
		fmt.Printf("Token: %s\n", to.Token)
	}

	loadAndPrintCfg(&cfg)

	// Set vendored variable with a prefix MY_COMPANY from the system environment
	os.Setenv("MY_COMPANY_TOKEN", "token from env")
	loadAndPrintCfg(&cfg)
	// Output:
	// Token: default token
	// Token: token from env
}
