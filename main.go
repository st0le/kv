package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"
)

func getStorageDirectory() string {
	usr, _ := user.Current()
	path := filepath.Join(usr.HomeDir, ".kv")
	return path
}

func ensureDirectory() string {
	path := getStorageDirectory()
	os.MkdirAll(path, os.ModePerm)
	return path
}

func list() {
	path := ensureDirectory()
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}

func set(key, value string) {
	keyFile := filepath.Join(ensureDirectory(), key)

	f, err := os.Create(keyFile)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	if value == "" {
		io.Copy(f, os.Stdin)
	} else {
		f.WriteString(value)
	}
}

func get(key string) {
	keyFile := filepath.Join(ensureDirectory(), key)
	f, err := os.Open(keyFile)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	io.Copy(os.Stdout, f)
}

func main() {

	var args = os.Args

	fi, _ := os.Stdin.Stat()

	if (fi.Mode() & os.ModeCharDevice) == 0 {
		// if there is a piped stdin, assume its a set command.
		if len(args) < 2 {
			log.Fatal("specify a key")
		} else {
			set(args[1], "")
		}

	} else {

		switch len(args) {
		case 1:
			// list
			list()
		case 2:
			// get
			get(args[1])
		case 3:
			// set
			set(args[1], args[2])
		}
	}

}
