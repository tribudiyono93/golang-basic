package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func check2(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	err := os.Mkdir("subdir", 0755)
	check2(err)

	defer os.RemoveAll("subdir")

	createEmptyFile := func(name string) {
		d := []byte("")
		check2(ioutil.WriteFile(name, d, 0644))
	}

	createEmptyFile("subdir/file1")

	err = os.MkdirAll("subdir/parent/child", 0755)
	check2(err)

	createEmptyFile("subdir/parent/file2")
	createEmptyFile("subdir/parent/file3")
	createEmptyFile("subdir/parent/child/file4")

	c, err := ioutil.ReadDir("subdir/parent")
	check2(err)

	fmt.Println("Listing subdir/parent")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	err = os.Chdir("subdir/parent/child")
	check2(err)

	c, err = ioutil.ReadDir(".")
	check2(err)

	fmt.Println("Listing subdir/parent/child")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	err = os.Chdir("../../..")
	check2(err)

	fmt.Println("Visiting subdir")
	err = filepath.Walk("subdir", visit)
}

func visit(p string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	fmt.Println(" ", p, info.IsDir())
	return nil
}