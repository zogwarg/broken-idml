package main

import (
	"archive/zip"
	"encoding/xml"
	"fmt"
	"github.com/sqweek/dialog"
	"io"
	"io/ioutil"
	"os"
	"path"
)

func IsValid(input io.Reader) error {
	decoder := xml.NewDecoder(input)
	for {
		err := decoder.Decode(new(interface{}))
		if err != nil {
			return err
		}
	}
}

func main() {
	ok := dialog.Message("%s", "Do you want to select a broken IDML file?").Title("Yes or No ?").YesNo()
	if ok {
		filename, err := dialog.File().Filter("Broken IDML source file", "idml").Load()

		// Exit if error
		if err != nil || filename == "" {
			os.Exit(1)
		}

		dir, err := os.MkdirTemp("", "example")

		// Exit if error
		if err != nil || dir == "" {
			os.Exit(1)
		}

		defer os.RemoveAll(dir) // clean up

		input, err := ioutil.ReadFile(filename)
		ioutil.WriteFile(path.Join(dir, "source.idml"), input, 0644)
		os.Chdir(dir)
		zipListing, err := zip.OpenReader("source.idml")

		for _, file := range zipListing.File {
			if file.FileHeader.Name[len(file.FileHeader.Name)-4:] != ".xml" {
				continue
			}
			f, _ := file.Open()
			valid := IsValid(f)

			if valid != io.EOF {
				fmt.Println(file.FileHeader.Name)
				fmt.Println(valid)
			}

		}
	}
}
