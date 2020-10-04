package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/beito123/goleveldb/leveldb"
)

func main() {
	path := "./TestWorld_v1.16.40/db/"

	db, err := openDB(path)
	if err != nil {
		panic(err)
	}

	err = bumpKeyValue(db)
	if err != nil {
		panic(err)
	}
}

func openDB(path string) (*leveldb.DB, error) {
	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func bumpKeyValue(db *leveldb.DB) error {
	os.MkdirAll("./db/", os.ModePerm)

	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		key := iter.Key()
		value := iter.Value()
		ioutil.WriteFile("./db/"+fmt.Sprintf("0x%x", key)+".dat", value, os.ModePerm)
		ioutil.WriteFile("./db/"+fmt.Sprintf("0x%x", key)+"_dump.txt", []byte(hex.Dump(value)), os.ModePerm)
	}

	iter.Release()

	err := iter.Error()
	if err != nil {
		return err
	}

	return nil
}
