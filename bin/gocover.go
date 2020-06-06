package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Sereal/Sereal/Go/sereal"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	db := os.Args[1]
	fmt.Println("Dumping the database: %v\n", db)
	var decoder sereal.Decoder
	decoder.PerlCompat = false
	dat, err := ioutil.ReadFile(db)
	check(err)
	var m interface{}
	err = decoder.Unmarshal(dat, &m)
	check(err)
	spew.Dump(m)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
