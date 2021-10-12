package main

import (
	"fmt"

	"github.com/spf13/afero"
)

func main() {
	fs := afero.NewMemMapFs()

	// create a file
	_, err := fs.Create("/tmp/test.sh")
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	// make it execuable
	fs.Chmod("/tmp/test.sh", 0755)

	// stat the file
	fi, err := fs.Stat("/tmp/test.sh")
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Printf("Fileinfo.name: %v\n", fi.Name())
	fmt.Printf("Fileinfo.size: %v\n", fi.Size())
	fmt.Printf("Fileinfo.mode: %v\n", fi.Mode())

	// see if file is executable
	if fi.Mode()&0111 != 0 {
		fmt.Println("File is executable")
	} else {
		fmt.Println("File is not executable")
	}
}
