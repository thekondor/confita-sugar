# confita-sugar [![GoDoc](https://godoc.org/github.com/thekondor/confita-sugar?status.svg)](http://godoc.org/github.com/thekondor/confita-sugar/)
Helpers for Confita configuration loading library to ease the development of 12 factor apps.

# Usage

## Example: optional file with a configuration

  ```go
    type Config struct {
		Login        string `config:"login"`
		Password     string `config:"password"`
		SessionToken string `config:"session_token"`
	}
	cfg := Config{SessionToken: "default session token"}

	loader := confita.NewLoader(
		maybe_util.NewFileBackend("local-config.yml"),
		confita_backend_file.NewBackend("main-config.json"),
	)

	// If `local-config.yml` is not available for any IO reason, the flow continues to execute;
  // no `log.Fatalf()` happens here ever but for non-existing `main-config.json`.
	if err := loader.Load(context.Background(), &cfg); nil != err {
		log.Fatalf("Failed to load app configuration: %+v", err)
	}
  ```
  
 ## Example: configuration over prefixed environment variables

  ```go
  type Config struct {
   Server struct {
       Host string `config:"server_host"`
       Port uint32 `config:"server_port"`
   }
  }
  ...
  loader := confita.NewLoader(
       ...
       prefixed_env.NewDefaultBackend("MY_COMPANY")
       ...
  )
  ```
  Since using the helper, that is possible to provide a configuration to the built app using prefixed environment variables:
  ```shell
  MY_COMPANY_SERVER_HOST=example.com MY_COMPANY_SERVER_PORT=8080 ./server-app.bin
  ```
