package fileutil

import (
	"os"
	"io"
	"io/ioutil"
	"log"
	"fmt"
)

func CopyAllHooks(srcDir, destDir string){
	files, err := ioutil.ReadDir(srcDir)
	if err != nil {
		log.Fatalf("error on read dirctory", err)
		os.Exit(1)
	}

	for i := 0; i < len(files); i++ {
		filename := files[i].Name()
		srcPath := fmt.Sprintf("%s/%s", srcDir, filename)
		destPath := fmt.Sprintf("%s/%s", destDir, filename)
		copy(srcPath, destPath)
		os.Chmod(destPath, os.FileMode(0755))
	}
}

func copy(srcName, destName string){
	src, err := os.Open(srcName)
	if err != nil {
		panic(err)
	}
	defer src.Close()

	dst, err := os.Create(destName)
	if err != nil {
		panic(err)
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if  err != nil {
		panic(err)
	}
}
