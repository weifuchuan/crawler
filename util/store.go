package util

import (
	"log"
	"os"
)

//SaveAsFile for save data by file.
func SaveAsFile(path string, data []byte) {
	f, err := os.Create(path)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	n, err := f.Write(data)
	if err != nil {
		log.Fatalf("f.Write error(%v)", err)
		return
	}
	log.Printf("f.Write (%d)", n)
}
