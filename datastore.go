package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/astaxie/beego/logs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func DataStoreDialog(from walk.Form, config *Config) {
	var dlg *walk.Dialog
	var address, username, password, database *walk.LineEdit
	var port, expired *walk.NumberEdit
	var testPB, acceptPB, cancelPB *walk.PushButton
	var enableCB *walk.CheckBox

	sqlConfig := config.Datastore

	_, err := Dialog{
		AssignTo:      &dlg,
		Title:         "MYSQL Database Configuration",
		Icon:          walk.IconInformation(),
		MinSize:       Size{Width: 500, Height: 150},
		Size:          Size{Width: 500, Height: 150},
		Font:          DefaultFont(),
		DefaultButton: &acceptPB,
		CancelButton:  &cancelPB,
		Layout:        VBox{},
		Children: []Widget{
			Composite{
				Layout: Grid{Columns: 4},
				Children: []Widget{
					Label{
						Text: "Database Address:",
					},
					LineEdit{
						Text:     sqlConfig.Address,
						AssignTo: &address,
						OnEditingFinished: func() {
							sqlConfig.Address = address.Text()
						},
					},
					Label{
						Text: "Database Port:",
					},
					NumberEdit{
						AssignTo:    &port,
						Value:       float64(sqlConfig.Port),
						ToolTipText: "1~65535",
						MaxValue:    65535,
						MinValue:    1,
						OnValueChanged: func() {
							sqlConfig.Port = int(port.Value())
						},
					},

					Label{
						Text: "Username:",
					},
					LineEdit{
						Text:     sqlConfig.UserName,
						AssignTo: &username,
						OnEditingFinished: func() {
							sqlConfig.UserName = username.Text()
						},
					},
					Label{
						Text: "Password:",
					},
					LineEdit{
						Text:     sqlConfig.PassWord,
						AssignTo: &password,
						OnEditingFinished: func() {
							sqlConfig.PassWord = password.Text()
						},
					},

					Label{
						Text: "Database Name:",
					},
					LineEdit{
						Text:     sqlConfig.DataBase,
						AssignTo: &database,
						OnEditingFinished: func() {
							sqlConfig.DataBase = database.Text()
						},
					},
					Label{
						Text: "Data Expired Days:",
					},
					NumberEdit{
						AssignTo:    &expired,
						Value:       float64(sqlConfig.Expired),
						ToolTipText: "1~365",
						MaxValue:    365,
						MinValue:    1,
						OnValueChanged: func() {
							sqlConfig.Expired = int(expired.Value())
						},
					},
					HSpacer{},
					CheckBox{
						AssignTo: &enableCB,
						Text:     "Enable",
						Checked:  sqlConfig.Enable,
						OnCheckedChanged: func() {
							sqlConfig.Enable = enableCB.Checked()
						},
					},
					HSpacer{},
					PushButton{
						AssignTo: &testPB,
						Text:     "Connectivity Test",
						OnClicked: func() {
							testPB.SetEnabled(false)
							go func() {
								result := DataStoreTest(sqlConfig)
								InfoBoxAction(dlg, "Test results:"+result)
								testPB.SetEnabled(true)
							}()
						},
					},
				},
			},
			VSpacer{},
			Composite{
				Layout: HBox{},
				Children: []Widget{
					HSpacer{},
					PushButton{
						AssignTo: &acceptPB,
						Text:     "Accept",
						OnClicked: func() {
							config.UpdateDataBase(sqlConfig)
							dlg.Accept()
							logs.Info("data store dialog accept")
						},
					},
					HSpacer{},
					PushButton{
						AssignTo: &cancelPB,
						Text:     "Cancel",
						OnClicked: func() {
							dlg.Cancel()
							logs.Info("data store dialog cancel")
						},
					},
					HSpacer{},
				},
			},
		},
	}.Run(from)

	if err != nil {
		logs.Error("DataStoreDialog: %s", err.Error())
	}
}

func DataStoreTest(config DataStoreConfig) string {
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/", config.UserName, config.PassWord, config.Address, config.Port)
	db, err := sql.Open("mysql", url)
	if err != nil {
		return "The test failed for a reason:" + err.Error()
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return "The test failed for a reason:" + err.Error()
	}
	return "Test Passed"
}
