package main

import (
	"github.com/astaxie/beego/logs"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func boxAction(from walk.Form, title string, icon *walk.Icon, message string, done func()) {
	var dlg *walk.Dialog
	var acceptPB, cancelPB *walk.PushButton

	var children []Widget

	logs.Info("box action title: %s message: %s", title, message)

	if done != nil {
		children = []Widget{
			HSpacer{},
			PushButton{
				AssignTo: &cancelPB,
				Text:     "Accept",
				OnClicked: func() {
					dlg.Accept()
					done()
					logs.Info("box action accept")
				},
			},
			HSpacer{},
			PushButton{
				AssignTo: &cancelPB,
				Text:     "Cancel",
				OnClicked: func() {
					dlg.Cancel()
					logs.Info("box action cancel")
				},
			},
			HSpacer{},
		}
	} else {
		children = []Widget{
			HSpacer{},
			PushButton{
				AssignTo: &cancelPB,
				Text:     "OK",
				OnClicked: func() {
					dlg.Accept()
					logs.Info("box action accept")
				},
			},
			HSpacer{},
		}
	}

	_, err := Dialog{
		AssignTo:      &dlg,
		Title:         title,
		Icon:          icon,
		MinSize:       Size{Width: 210, Height: 120},
		Font:          DefaultFont(),
		DefaultButton: &acceptPB,
		CancelButton:  &cancelPB,
		Layout:        VBox{},
		Children: []Widget{
			VSpacer{},
			TextLabel{
				Text:    message,
				MinSize: Size{Width: 180, Height: 100},
			},
			VSpacer{},
			Composite{
				Layout:   HBox{},
				Children: children,
			},
		},
	}.Run(from)

	if err != nil {
		logs.Error("boxAction: %s", err.Error())
	}
}

func ErrorBoxAction(form walk.Form, message string) {
	walk.MsgBox(form, "Error", message, walk.MsgBoxOK)
}

func InfoBoxAction(form walk.Form, message string) {
	walk.MsgBox(form, "Information", message, walk.MsgBoxOK)
}

func ConfirmBoxAction(form walk.Form, message string, done func()) {
	boxAction(form, "Confirmation", walk.IconWarning(), message, done)
}
