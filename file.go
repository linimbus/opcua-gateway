package main

import (
	"os"
	"path/filepath"
)

var DEFAULT_HOME string

func RunlogDirGet() string {
	dir := filepath.Join(DEFAULT_HOME, "runlog")
	_, err := os.Stat(dir)
	if err != nil {
		os.MkdirAll(dir, 0644)
	}
	return dir
}

func IconDirGet() string {
	dir := filepath.Join(DEFAULT_HOME, "icon")
	_, err := os.Stat(dir)
	if err != nil {
		os.MkdirAll(dir, 0644)
	}
	return dir
}

func ConfigDirGet() string {
	dir := filepath.Join(DEFAULT_HOME, "config")
	_, err := os.Stat(dir)
	if err != nil {
		os.MkdirAll(dir, 0644)
	}
	return dir
}

func appDataDir() string {
	datadir := os.Getenv("APPDATA")
	if datadir == "" {
		datadir = os.Getenv("CD")
	}
	if datadir == "" {
		datadir = ".\\"
	} else {
		datadir = filepath.Join(datadir, "OpcuaGatewayWindows")
	}
	return datadir
}

func FileInit() {
	dir := appDataDir()
	_, err := os.Stat(dir)
	if err != nil {
		err = os.MkdirAll(dir, 0644)
		if err != nil {
			DEFAULT_HOME = ".\\"
			return
		}
	}
	DEFAULT_HOME = dir
}
