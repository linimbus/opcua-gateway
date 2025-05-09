package main

import (
	"os/exec"

	"github.com/astaxie/beego/logs"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func OpenBrowserWeb(url string) {
	cmd := exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	err := cmd.Run()
	if err != nil {
		logs.Error("run cmd fail, %s", err.Error())
	}
}

func AboutAction() {
	var ok *walk.PushButton
	var about *walk.Dialog
	var err error

	_, err = Dialog{
		AssignTo:      &about,
		Title:         "About",
		Icon:          walk.IconInformation(),
		MinSize:       Size{Width: 200, Height: 200},
		Font:          DefaultFont(),
		DefaultButton: &ok,
		Layout:        VBox{},
		Children: []Widget{
			Composite{
				Layout: VBox{},
				Children: []Widget{
					TextLabel{
						Text: "Program Version: " + VersionGet(),
					},
					TextLabel{
						Text: "Technical support. " + CompanyGet(),
					},
				},
			},
			Composite{
				Layout: HBox{},
				Children: []Widget{
					HSpacer{},
					PushButton{
						Text:      "OK",
						OnClicked: func() { about.Cancel() },
					},
					HSpacer{},
				},
			},
		},
	}.Run(mainWindow)

	if err != nil {
		logs.Error(err.Error())
	}
}
