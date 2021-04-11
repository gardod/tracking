# Go Tracking Client

We have a pretty streight forward one command CLI written in go. We are using [spf13/cobra](github.com/spf13/cobra) as a CLI library and [spf13/viper](github.com/spf13/viper) for config.

Structure is similar to the HTTP service. We hold the business logic in Service, while Repository is a wrapper for outside services.

## Usage
```bash
go build
./client --help
```
