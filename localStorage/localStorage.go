package localStorage

import (
	"MyTestingCode/dataTypes"
	"os"
	"path"
	"time"
)

type IRemoteDataSource interface {
	Subscribe(symbol string, sDate, eDate time.Time, onData func(dataTypes.IData)) error
}
type ILocalStorage interface {
	SetTimeRange(sDate, eDate time.Time) error
	LoadData(symbol string) error
}
type ICacheDataConsumer interface {
	// localStorage call this method to query each user to get the last date it need to keep cache for.
	// GetLastDateNeed(tableName string) time.Time
	// GetTables() []string
	// GetColumns(tableName string) []string
	// GetPrereadCount(tableName string) int
}

// HW: Testing connectionpool, pre-read cache, allow multple read , single write, clean up
type localStorage struct {
	consumers []ICacheDataConsumer
}

func (ls *localStorage) SetTimeRange(sDate, eDate time.Time) error {
	return nil
}
func NewLocalStorage(symbol string) ILocalStorage {
	result := &localStorage{}
	return result
}
func (ls *localStorage) getLocation(symbol string) string {
	rootPath := os.Getenv("LOCAL_STORAGE_ROOT")

	if "" == rootPath {

		p, _ := os.Executable()
		exPath := path.Dir(p)
		rootPath = exPath
	}

	firstRune := []rune(symbol)[0]

	return path.Join(rootPath, string(firstRune), symbol)
}
func (ls *localStorage) LoadData(symbol string) error {
	path := ls.getLocation(symbol)

	ls.createIfNotExists(path)

	ls.startConnectionPool()
	return nil
}

func (ls *localStorage) createIfNotExists(path string) {

}
func (ls *localStorage) GetAvailableConnections() {

}
func (ls *localStorage) SetMaxConnection() {

}
func (ls *localStorage) startConnectionPool() {

}
func (ls *localStorage) RegisterCacheConsumer(consumer ICacheDataConsumer) {
	if nil == ls.consumers {
		ls.consumers = make([]ICacheDataConsumer, 0, 1)
	}
	ls.consumers = append(ls.consumers, consumer)

	// setup the columns for table
	// setup the pre-read count for table

}
func (ls *localStorage) cacheMaintainWorker() {
	// check each cache for lowerness
	// read if too low

}

func (ls *localStorage) readFromCache(columnName, tableName string, date time.Time) interface{} {
	// check if current cache has the date
	// if not,  pre-read
	// return the value from cache / or nil if not found
	return nil
}
func (ls *localStorage) ReadValue(columnName, tableName string, date time.Time) interface{} {

	// Read from cache first
	// if not found, read
	return nil

}
func (ls *localStorage) ReadValueNoCache(columnName, tableName string, date time.Time) interface{} {

	// Read from cache first
	//
	return nil

}
