package main

import (
	"kv-storage/internal/app/core"
	"kv-storage/internal/pkg/config"
	"kv-storage/internal/pkg/logger"
)

func main() {
	err := core.StartApp(core.Params{
		Port:   config.Core.Port,
		Prefix: config.Core.Prefix,
	})
	if err != nil {
		core.StopApp()
		logger.Error(err.Error())
	}
}
