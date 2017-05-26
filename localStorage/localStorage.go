package localStorage

import (
	"MyTestingCode/priceData"
	"database/sql"
	"os"
	"path"
	"time"
	// load sqlite driver
	"log"

	"io"

	_ "github.com/mattn/go-sqlite3"
)

type ILocalStorage interface {
	Register(dataInfo ILocalStorageDataInfo) error
	LoadData(symbol string) error
	io.Closer
}
type ILocalStorageDataInfo interface {
	GetCatalog() string
	GetName() string
	GetValueType() priceData.DataValueType
}

// HW: Testing connectionpool, pre-read cache, allow multple read , single write, clean up
type sqliteLocalStorage struct {
	tables   map[string]map[string]ILocalStorageDataInfo
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
		ls.updateTables()
		ls.startConnectionPool()
	}
	return err
}
func (ls *sqliteLocalStorage) updateTables() (err error) {
	_, err = ls.database.Exec("CREATE TABLE `customers` (`till_id` INTEGER PRIMARY KEY AUTOINCREMENT, `client_id` VARCHAR(64) NULL, `first_name` VARCHAR(255) NOT NULL, `last_name` VARCHAR(255) NOT NULL, `guid` VARCHAR(255) NULL, `dob` DATETIME NULL, `type` VARCHAR(1))")
	if err != nil {
		log.Println(err)
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
func (ls *sqliteLocalStorage) Register(dataInfo ILocalStorageDataInfo) error {
	if nil == ls.tables {
		ls.tables = make(map[string]map[string]ILocalStorageDataInfo)
	}
	if _, Found := ls.tables[dataInfo.GetCatalog()]; !Found {
		ls.tables[dataInfo.GetCatalog()] = make(map[string]ILocalStorageDataInfo)
	}
	ls.tables[dataInfo.GetCatalog()][dataInfo.GetName()] = dataInfo
	// setup the columns for table
	// setup the pre-read count for table
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
