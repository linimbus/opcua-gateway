package main

import (
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/lxn/walk"
)

var notify *walk.NotifyIcon

func NotifyAction() {
	if notify == nil {
		NotifyInit()
	}
	mainWindow.SetVisible(false)
}

func NotifyExit() {
	if notify == nil {
		return
	}
	notify.Dispose()
	notify = nil
	logs.Info("notify exit")
}

var lastCheck time.Time

func NotifyInit() {
	var err error

	notify, err = walk.NewNotifyIcon(mainWindow)
	if err != nil {
		logs.Error("new notify icon fail, %s", err.Error())
		return
	}

	err = notify.SetIcon(ICON_Main)
	if err != nil {
		logs.Error("set notify icon fail, %s", err.Error())
		return
	}

	err = notify.SetToolTip(AppNameGet())
	if err != nil {
		logs.Error("set notify icon fail, %s", err.Error())
		return
	}

	exitBut := walk.NewAction()
	err = exitBut.SetText("Shutdown Program")
	if err != nil {
		logs.Error("notify new action fail, %s", err.Error())
		return
	}

	exitBut.Triggered().Attach(func() {
		ConfirmBoxAction(mainWindow, "Do you want to shutdown the program?", CloseWindows)
		logs.Info("notify triggered exit")
	})

	showBut := walk.NewAction()
	err = showBut.SetText("Show Window")
	if err != nil {
		logs.Error("notify new action fail, %s", err.Error())
		return
	}

	showBut.Triggered().Attach(func() {
		logs.Info("notify windows set visible true")
		mainWindow.SetVisible(true)
	})

	if err := notify.ContextMenu().Actions().Add(showBut); err != nil {
		logs.Error("notify add action fail, %s", err.Error())
		return
	}

	if err := notify.ContextMenu().Actions().Add(exitBut); err != nil {
		logs.Error("notify add action fail, %s", err.Error())
		return
	}

	notify.MouseUp().Attach(func(x, y int, button walk.MouseButton) {
		if button != walk.LeftButton {
			return
		}
		now := time.Now()
		if now.Sub(lastCheck) < time.Second {
			logs.Info("notify windows set visible true")

			mainWindow.SetVisible(true)
		}
		lastCheck = now
	})

	notify.SetVisible(true)

	logs.Info("notify init success")
}
