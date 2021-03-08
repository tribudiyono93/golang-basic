package main

import (
	"fmt"
	"github.com/tribudiyono93/golang-basic/util"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {

	f, err := ioutil.TempFile("", "sample")
	util.CheckErr(err)

	fmt.Println("Temp file name:", f.Name())

	defer os.Remove(f.Name())

	_, err = f.Write([]byte{1, 2, 3, 4})
	util.CheckErr(err)

	dname, err := ioutil.TempDir("", "sampledir")
	util.CheckErr(err)
	fmt.Println("Temp dir name:", dname)

	defer os.RemoveAll(dname)

	fname := filepath.Join(dname, "file1")
	err = ioutil.WriteFile(fname, []byte{1, 2}, 0666)
	util.CheckErr(err)
}