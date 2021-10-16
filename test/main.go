package main

import (
	"archiver"
	"fmt"
)

func main() {
	rarFile := "E:/workspace_go/src/qubit/static/upload/craft/20211015/1634267973.rar"
	destPath := "E:/workspace_go/src/qubit/static/upload/craft/20211015/1634267973/"
	r := archiver.NewRar()
	err := r.Unarchive(rarFile, destPath)
	fmt.Print(err)
}
