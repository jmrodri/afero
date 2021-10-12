package main

import (
	"fmt"

	"github.com/spf13/afero"
)

func main() {
	fs := afero.NewMemMapFs()

	_, err := fs.Create("/tmp/test.sh")
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fs.Chmod("/tmp/test.sh", 0755)

	fi, err := fs.Stat("/tmp/test.sh")
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Printf("Fileinfo.name: %v\n", fi.Name())
	fmt.Printf("Fileinfo.size: %v\n", fi.Size())
	fmt.Printf("Fileinfo.mode: %v\n", fi.Mode())

	if fi.Mode()&0111 != 0 {
		fmt.Println("File is executable")
	} else {
		fmt.Println("File is not executable")
	}

	// f, err := fs.FS.Create(filepath.Join(filepath.Dir(pluginFilePath), filename))
	//         Expect(err).To(BeNil())
	//         Expect(f).ToNot(BeNil())
	//         // ensure that the file exists
	//         _, err = fs.FS.Stat(pluginFilePath)
	//         Expect(err).To(BeNil())
}
