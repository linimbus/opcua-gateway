package main

import (
	"path/filepath"

	"github.com/astaxie/beego/logs"
	"github.com/lxn/walk"
)

func IconLoadFromBox(filename string, size walk.Size) *walk.Icon {
	body, err := Asset(filename)
	if err != nil {
		logs.Warning("load icon file failed, %s", err.Error())
		return walk.IconApplication()
	}

	filename = filepath.Join(IconDirGet(), filename)

	err = SaveToFile(filename, body)
	if err != nil {
		logs.Warning("save icon file failed, %s", err.Error())
		return walk.IconApplication()
	}

	icon, err := walk.NewIconFromFileWithSize(filename, size)
	if err != nil {
		logs.Error("new icon from file failed, %s", filename)
		return walk.IconApplication()
	}
	return icon
}

var ICON_Main *walk.Icon
var ICON_Start *walk.Icon
var ICON_Stop *walk.Icon
var ICON_Status *walk.Icon

func IconInit() {
	ICON_Main = IconLoadFromBox("main.ico", walk.Size{Width: 64, Height: 64})
	ICON_Start = IconLoadFromBox("start.ico", walk.Size{Width: 64, Height: 64})
	ICON_Stop = IconLoadFromBox("stop.ico", walk.Size{Width: 64, Height: 64})
	ICON_Status = IconLoadFromBox("status.ico", walk.Size{Width: 16, Height: 16})
}
