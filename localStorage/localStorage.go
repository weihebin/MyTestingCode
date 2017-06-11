package localStorage

import (
	"MyTestingCode/constants"
	"MyTestingCode/priceData"
	"context"
	"database/sql"
	"fmt"
	"os"
	"path"
	"time"

	// load sqlite driver
	"log"

	"io"

	_ "github.com/mattn/go-sqlite3"
)

type ILocalStorage interface {
	GetLastDateOfColumn(catalog constants.ColumnCatalog, columnName string) time.Time
	SubscribeColumnData(ctx context.Context, timeRange priceData.TimeRange, catalog constants.ColumnCatalog, columns []string, onData func(*time.Time, priceData.IData)) error
	AddOrUpdateValues(catalog constants.ColumnCatalog, columnsName []string, date *time.Time, values priceData.IData) error
	RegisterColumn(catalog constants.ColumnCatalog, columnName string, valueType constants.DataValueType) error
	LoadData(symbol string) error
	io.Closer
}
type dataInfo struct {
	Catalog   constants.ColumnCatalog
	Name      string
	ValueType constants.DataValueType
}

// HW: Testing connectionpool, pre-read cache, allow multple read , single write, clean up
type sqliteLocalStorage struct {
	tables   map[constants.ColumnCatalog]map[string]dataInfo
	database *sql.DB
}

func NewLocalStorage(symbol string) ILocalStorage {
	result := &sqliteLocalStorage{}
	return result
}
func (ls *sqliteLocalStorage) Close() (err error) {
	if nil != ls.database {
		err = ls.database.Close()
	}
	return err
}
func (ls *sqliteLocalStorage) getLocation(symbol string) string {
	rootPath := os.Getenv("LOCAL_STORAGE_ROOT")

	if "" == rootPath {

		p, _ := os.Executable()
		exPath := path.Dir(p)
		rootPath = exPath
	}

	firstRune := []rune(symbol)[0]

	return path.Join(rootPath, string(firstRune), symbol, symbol+".db")
}
func (ls *sqliteLocalStorage) LoadData(symbol string) (err error) {
	path := ls.getLocation(symbol)

	err = ls.createIfNotExists(path)
	if nil == err {
		err = ls.updateTables()
		if nil == err {
			ls.startConnectionPool()
		}
	}
	return err
}
func getDataType(cDataType constants.DataValueType) string {
	switch cDataType {
	case constants.DataValueType_Decimal:
		return "decimal(20,5)"

	}
	return "text"
}
func (ls *sqliteLocalStorage) updateTables() (err error) {

	for tableName, columns := range ls.tables {
		// create table if not exist
		sql := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %v ( Year INTEGER NOT NULL, Month INTEGER NOT NULL,Day INTEGER NOT NULL, PRIMARY KEY ( Year,Month,Day ) );", tableName)
		_, err = ls.database.Exec(sql)
		if err != nil {
			log.Println(err)
		}
		for columnName, columnInfo := range columns {
			// append column to table if not exists
			sql = fmt.Sprintf("ALTER TABLE %v ADD COLUMN %v %v", tableName, columnName, getDataType(columnInfo.ValueType))
			ls.database.Exec(sql) // err != nil if duplicate
		}
	}

	return err
}
func (ls *sqliteLocalStorage) createIfNotExists(filename string) (err error) {

	err = os.MkdirAll(path.Dir(filename), 0777)
	if nil == err {
		ls.database, err = sql.Open("sqlite3", filename)
		if nil != err {
			log.Println(err)
		}
	} else {
		log.Println(err)
	}

	return err
}
func (ls *sqliteLocalStorage) GetAvailableConnections() {

}
func (ls *sqliteLocalStorage) SetMaxConnection() {

}
func (ls *sqliteLocalStorage) startConnectionPool() {

}
func (ls *sqliteLocalStorage) RegisterColumn(catalog constants.ColumnCatalog, columnName string, valueType constants.DataValueType) error {
	if nil == ls.tables {
		ls.tables = make(map[constants.ColumnCatalog]map[string]dataInfo)
	}
	if _, Found := ls.tables[catalog]; !Found {
		ls.tables[catalog] = make(map[string]dataInfo)
	}
	ls.tables[catalog][columnName] = dataInfo{Catalog: catalog, Name: columnName, ValueType: valueType}
	return nil

}
func (ls *sqliteLocalStorage) GetLastDateOfColumn(catalog constants.ColumnCatalog, columnName string) time.Time {

	sqlStr := fmt.Sprintf("select max(Year * 10000 + Month * 100 + Day) from %v where %v is not null;", catalog, columnName)

	rows, err := ls.database.Query(sqlStr)
	if nil != err {
		log.Println(err)
	}
	var maxDate sql.NullInt64
	if nil != rows && rows.Next() {
		err = rows.Scan(&maxDate)
		if nil != err {
			log.Println(err)
		}
	}
	if maxDate.Valid {
		v, _ := maxDate.Value()
		val := (int)(v.(int64))
		return time.Date(val/10000, (time.Month)((val%10000)/100), val%100, 0, 0, 0, 0, time.UTC)
	} else {
		return time.Time{}
	}
}
func (ls *sqliteLocalStorage) SubscribeColumnData(ctx context.Context, timeRange priceData.TimeRange, catalog constants.ColumnCatalog, columns []string, onData func(*time.Time, priceData.IData)) error {
	return nil
}
func (ls *sqliteLocalStorage) AddOrUpdateValues(catalog constants.ColumnCatalog, columnsName []string, date *time.Time, values priceData.IData) error {
	return nil
}

func (ls *sqliteLocalStorage) cacheMaintainWorker() {
	// check each cache for lowerness
	// read if too low
}
func (ls *sqliteLocalStorage) readFromCache(columnName, tableName string, date time.Time) interface{} {
	// check if current cache has the date
	// if not,  pre-read
	// return the value from cache / or nil if not found
	return nil
}
func (ls *sqliteLocalStorage) ReadValue(columnName, tableName string, date time.Time) interface{} {

	// Read from cache first
	// if not found, read
	return nil

}
func (ls *sqliteLocalStorage) ReadValueNoCache(columnName, tableName string, date time.Time) interface{} {

	// Read from cache first
	//
	return nil

}
