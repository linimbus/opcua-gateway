package main

import (
	"bytes"
	"database/sql"
	"fmt"

	"github.com/astaxie/beego/logs"
	_ "github.com/go-sql-driver/mysql"
)

type ColumnInfo struct {
	Name    string
	Comment string
}

type DataSave struct {
	expired   int
	database  string
	db        *sql.DB
	tableInfo map[string][]ColumnInfo
}

func ExecuteUpdate(db *sql.DB, sql string) error {
	_, err := db.Exec(sql)
	if err != nil {
		logs.Error("ExecuteUpdate SQL[%s] failed, %s", sql, err.Error())
		return err
	}
	// logs.Info("ExecuteUpdate SQL[%s] success", sql)
	return nil
}

func ExecuteQuery(db *sql.DB, sql string) (*sql.Rows, error) {
	rows, err := db.Query(sql)
	if err != nil {
		logs.Error("ExecuteQuery SQL[%s] failed, %s", sql, err.Error())
		return nil, err
	}
	// logs.Info("ExecuteQuery SQL[%s] success", sql)
	return rows, nil
}

func TableCheck(db *sql.DB, database, tableName string) bool {
	sql := fmt.Sprintf("SELECT COUNT(*) FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_SCHEMA = '%s' AND TABLE_NAME = '%s'", database, tableName)

	rows, err := ExecuteQuery(db, sql)
	if err != nil {
		return false
	}
	defer rows.Close()

	if !rows.Next() {
		logs.Error("TableCheck Row Next failed")
		return false
	}

	var number int
	err = rows.Scan(&number)
	if err != nil {
		logs.Error("TableCheck Row Scan failed, %s", err.Error())
		return false
	}

	return number > 0
}

func ColumnCompare(newColumns, oldColumns []ColumnInfo) []ColumnInfo {
	var columns []ColumnInfo
	for _, newColumn := range newColumns {
		found := false
		for _, oldColumn := range oldColumns {
			if newColumn.Name == oldColumn.Name {
				found = true
				break
			}
		}
		if !found {
			columns = append(columns, newColumn)
		}
	}
	return columns
}

func TableInfo(db *sql.DB, database, tableName string) ([]ColumnInfo, error) {
	sql := fmt.Sprintf("SELECT COLUMN_NAME, COLUMN_COMMENT FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_SCHEMA = '%s' AND TABLE_NAME = '%s'", database, tableName)

	rows, err := ExecuteQuery(db, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columns := make([]ColumnInfo, 0)
	for rows.Next() {
		var name, comment string
		err = rows.Scan(&name, &comment)
		if err != nil {
			logs.Error("TableInfo Row Scan failed, %s", err.Error())
			return nil, err
		}
		columns = append(columns, ColumnInfo{Name: name, Comment: comment})
	}
	return columns, nil
}

func TableCreate(db *sql.DB, database, tableName string, columns []ColumnInfo) error {
	var sql bytes.Buffer

	sql.WriteString(fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s.%s (", database, tableName))
	sql.WriteString(" id INT PRIMARY KEY AUTO_INCREMENT, ")
	sql.WriteString(" timestamp DATETIME DEFAULT CURRENT_TIMESTAMP ")

	for _, column := range columns {
		sql.WriteString(", ")
		sql.WriteString(fmt.Sprintf("`%s` TEXT COMMENT '%s'", column.Name, column.Comment))
	}
	sql.WriteString(")")

	return ExecuteUpdate(db, sql.String())
}

func TableAlter(db *sql.DB, database, tableName string, newColumns []ColumnInfo) error {
	var sql bytes.Buffer

	sql.WriteString(fmt.Sprintf("ALTER TABLE %s.%s ", database, tableName))

	for i, column := range newColumns {
		sql.WriteString(fmt.Sprintf("ADD `%s` TEXT COMMENT '%s'", column.Name, column.Comment))
		if i+1 != len(newColumns) {
			sql.WriteString(", ")
		} else {
			sql.WriteString("; ")
		}
	}

	return ExecuteUpdate(db, sql.String())
}

func TableWrite(db *sql.DB, database, tableName string, columns []ColumnInfo, values []string) error {
	var sql bytes.Buffer

	sql.WriteString(fmt.Sprintf("INSERT INTO %s.%s (", database, tableName))

	for i, column := range columns {
		sql.WriteString(fmt.Sprintf("`%s`", column.Name))
		if i+1 != len(columns) {
			sql.WriteString(", ")
		}
	}

	sql.WriteString(") VALUES (")

	for i, value := range values {
		if value == "" {
			sql.WriteString("NULL")
		} else {
			sql.WriteString(fmt.Sprintf("'%s'", value))
		}
		if i+1 != len(values) {
			sql.WriteString(", ")
		}
	}

	sql.WriteString(")")

	return ExecuteUpdate(db, sql.String())
}

func NewDataSave(cfg DataStoreConfig) (*DataSave, error) {
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/", cfg.UserName, cfg.PassWord, cfg.Address, cfg.Port)
	db, err := sql.Open("mysql", url)
	if err != nil {
		logs.Error("CreateDataSave: %s", err.Error())
		return nil, err
	}

	err = ExecuteUpdate(db, fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", cfg.DataBase))
	if err != nil {
		db.Close()
		return nil, err
	}

	dbSave := &DataSave{
		expired:   cfg.Expired,
		database:  cfg.DataBase,
		db:        db,
		tableInfo: make(map[string][]ColumnInfo, 0)}

	logs.Info("CreateDataSave create %v success", cfg)
	return dbSave, nil
}

func (d *DataSave) Close() {
	logs.Info("DataSave ready to close")

	err := d.db.Close()
	if err != nil {
		logs.Error("DataSave.Close: %s", err.Error())
	}
	d.db = nil
}

func (d *DataSave) TableWrite(tableName string, values []string) error {
	columns, ok := d.tableInfo[tableName]
	if !ok {
		return fmt.Errorf("DataSave.TableWrite: %s not init", tableName)
	}

	if len(columns) != len(values) {
		return fmt.Errorf("DataSave.TableWrite: %s columns[%d] != values[%d]", tableName, len(columns), len(values))
	}

	return TableWrite(d.db, d.database, tableName, columns, values)
}

func (d *DataSave) TableInit(tableName string, columns []ColumnInfo) error {
	if !TableCheck(d.db, d.database, tableName) {
		err := TableCreate(d.db, d.database, tableName, columns)
		if err != nil {
			return err
		}
	} else {
		oldColumns, err := TableInfo(d.db, d.database, tableName)
		if err != nil {
			return err
		}
		newColumns := ColumnCompare(columns, oldColumns)
		if len(newColumns) > 0 {
			err = TableAlter(d.db, d.database, tableName, newColumns)
			if err != nil {
				return err
			}
		}
	}
	logs.Info("DataSave.TableInit: %s Columns %d success", tableName, len(columns))
	d.tableInfo[tableName] = columns

	return nil
}

func (d *DataSave) TableExpired(enable bool) error {
	err := ExecuteUpdate(d.db, "SET GLOBAL event_scheduler = ON;")
	if err != nil {
		return err
	}

	err = ExecuteUpdate(d.db, fmt.Sprintf("USE %s;", d.database))
	if err != nil {
		return err
	}

	err = ExecuteUpdate(d.db, fmt.Sprintf("DROP EVENT IF EXISTS %s_data_expired_event;", d.database))
	if err != nil {
		return err
	}

	logs.Info("DataSave.tableExpired delete data expired event success")

	if enable {
		var buffer bytes.Buffer

		buffer.WriteString(fmt.Sprintf("CREATE EVENT %s_data_expired_event ON SCHEDULE EVERY 1 HOUR ", d.database))
		buffer.WriteString("STARTS CURRENT_TIMESTAMP ON COMPLETION PRESERVE DO BEGIN ")

		for table, _ := range d.tableInfo {
			buffer.WriteString(fmt.Sprintf("DELETE FROM %s.%s WHERE timestamp < DATE_SUB(NOW(), INTERVAL %d DAY);", d.database, table, d.expired))
		}
		buffer.WriteString("END;")

		err = ExecuteUpdate(d.db, buffer.String())
		if err != nil {
			return err
		}

		logs.Info("DataSave.tableExpired create data expired event success")
	}

	return nil
}
