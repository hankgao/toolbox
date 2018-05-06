package main

import (
	"github.com/boltdb/bolt"
	"github.com/skycoin/skycoin/src/visor"
)

const (
	dbpath = "~/.suncoin/data.db"
)

func main() {

	c := visor.NewVisorConfig()

	db, err := bolt.Open(dbpath, 0700, nil)
	if err != nil {
		panic(err)
	}

	v, err := visor.NewVisor(c, db)

	for i := 0; i < 100; i++ {

	}

}

type treasureBookT map[string]uint64

func listAllAddresses() []string {
	// read the blocks one by one

	var addresses []string

	return addresses
}
