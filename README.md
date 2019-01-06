# confita-sugar [![Build Status](https://travis-ci.org/thekondor/confita-sugar.svg?branch=master)](https://travis-ci.org/thekondor/confita-sugar) :: [![Go Report Card](https://goreportcard.com/badge/github.com/thekondor/confita-sugar)](https://goreportcard.com/report/github.com/thekondor/confita-sugar)
Set of extensions for Confita configuration loading library to ease the development of (mostly) 12-factor apps.

# Documentation

_(with more examples)_ could be found at [![GoDoc](https://godoc.org/github.com/thekondor/confita-sugar?status.svg)](http://godoc.org/github.com/thekondor/confita-sugar/)

# Motivation

[confita](https://github.com/heetch/confita) package provides with a great tool to load an application's configuration from multiple sources; and not cluttered with lot of options which makes easy to get started with the library. In spite of this, some use cases could be simplified. 

Thanks to the good abstractions provided, that is possible to simplify these use cases without changing Confita's source code at all. That is the reason #1. The reason #2 is that corresponding PR(s) might take time (_or possibly could be not approved by the authors for reasons_) but the covered use cases were required in the real-life project; the most stable approach is extend 3rd-party logic with own decorators. That is why `confita-sugar` was engineered.

# Status

- The implementation is considered stable; `confita-sugar` has already been succesfully used in real-life applications;
- API _might_ be broken without further notice. 

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
  
  This use case is not supported by vanilla `Confita` out of the box without `confita-sugar`. The workaround in the past:
  ```go
  backends := []confita.Backend{ confita_file.NewBackend("/path/to/config.json") }
  if isFileExists("/path/to/another-config.yaml") {
    backends = append(backends, confita_file.NewBackend("/path/to/another-config.yaml"))
  }
  loader := confita.NewLoader(backends...)
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
       prefix_util.NewEnvBackend("MY_COMPANY")
       ...
  )
  ```
  Since using the helper, that is possible to provide a configuration to the built app using prefixed environment variables:
  ```shell
  MY_COMPANY_SERVER_HOST=example.com MY_COMPANY_SERVER_PORT=8080 ./server-app.bin
  ```
  
# License

The library is released under the MIT license. See LICENSE file.
