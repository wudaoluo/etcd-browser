package regret

import (
	"github.com/ThreeKing2018/goutil/golog"
	"github.com/syndtr/goleveldb/leveldb"
)

var db *leveldb.DB

func init() {
	var err error
	db, err = leveldb.OpenFile("./leveldb", nil)
	if err != nil {
		panic(err)
	}
}

func Close() {
	db.Close()
}

func Get(key []byte) ([]byte, error) {
	return db.Get(key, nil)
}

func Put(key, value []byte) error {
	return db.Put(key, value, nil)
}

func Delete(key []byte) error {
	return db.Delete(key, nil)
}

func Has(key []byte) (bool, error) {
	return db.Has(key, nil)
}

func GetString(key string) (string, error) {
	value, err := db.Get([]byte(key), nil)
	return string(value), err
}

func PutString(key, value string) {
	err :=  db.Put([]byte(key), []byte(value), nil)
	if err != nil {
		golog.Error(err)
	}
}

func DeleteString(key string) {
	err := db.Delete([]byte(key), nil)
	if err != nil {
		golog.Error(err)
	}
}

func HasString(key string) bool {
	ok, err := db.Has([]byte(key), nil)
	if err != nil {
		golog.Error(err)
		return false
	}

	return ok
}

