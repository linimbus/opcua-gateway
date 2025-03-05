package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

var configStatus *walk.StatusBarItem

func StatusConfig(config string) {
	if configStatus != nil {
		configStatus.SetText(config)
		configStatus.SetWidth(20 + len(config)*6)
	}
}

func StatusBarInit() []StatusBarItem {
	return []StatusBarItem{
		{
			Icon:  ICON_Status,
			Text:  "Configuration Files",
			Width: 160,
		},
		{
			AssignTo: &configStatus,
			Text:     "Not loaded",
		},
	}
}
