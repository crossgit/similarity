package main

import (
	"log"

	"github.com/tidwall/buntdb"
)

func main() {
	var err error
	db, err = buntdb.Open("data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set("myKey", "myvalue", nil)
		return err
	})

	db.View(func(tx *buntdb.Tx) error {
		val, err := tx.Get("myKey1")
		if err != nil {
			log.Println(err)
			return err
		}
		log.Printf("value is %s \n", val)
		return nil
	})
}

var db *buntdb.DB

// func init() {
// 	var err error
// 	db, err = buntdb.Open("data.db")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()

// 	db.Update(func(tx *buntdb.Tx) error {
// 		_, _, err := tx.Set("myKey", "myvalue", nil)
// 		return err
// 	})
// }
