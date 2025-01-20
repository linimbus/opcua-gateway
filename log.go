package main

import (
	"encoding/json"
	"fmt"
	"path/filepath"

	"github.com/astaxie/beego/logs"
)

func ApplicationLogInit(config ApplicationConfig) error {
	cfg := config.LogConfig
	if cfg.Filename == "" {
		cfg.Filename = "runlog.log"
	}
	cfg.Filename = filepath.Join(config.LogPath, cfg.Filename)
	value, err := json.Marshal(&cfg)
	if err != nil {
		fmt.Printf("json.Marshal failed: %v", err)
		return err
	}
	err = logs.SetLogger(logs.AdapterFile, string(value))
	if err != nil {
		fmt.Printf("logs.SetLogger failed: %v", err)
		return err
	}
	logs.EnableFuncCallDepth(true)
	logs.SetLogFuncCallDepth(3)
	return nil
}
