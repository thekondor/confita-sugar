package util

import (
	"context"
	confita "github.com/heetch/confita"
	confita_backend_file "github.com/heetch/confita/backend/file"
	"log"
)

func ExampleNewFileBackend() {
	type Config struct {
		Login        string `config:"login"`
		Password     string `config:"password"`
		SessionToken string `config:"session_token"`
	}
	cfg := Config{SessionToken: "default session token"}

	loader := confita.NewLoader(
		/*maybe_util.*/ NewFileBackend("local-config.yml"),
		confita_backend_file.NewBackend("main-config.json"),
	)

	// If `local-config.yml` is not available for any IO reason, the flow continues to execute; no `log.Fatalf()` happens here ever but for non-existing `main-config.json`
	if err := loader.Load(context.Background(), &cfg); nil != err {
		log.Fatalf("Failed to load app configuration: %+v", err)
	}

}
