package main

import (
	"fmt"
	"os"
)

// Program should iterate and find all files like tree in linux
// it should use also goroutines
func main() {

	writeCurrentPath()

	readFilesForDir(0, ".")

}

func readFilesForDir(iteration int, dir string) {
	files, err := os.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		for i := 0; i <= iteration; i++ {
			fmt.Print("\t")
		}
		fmt.Println("|-", file.Name())
		if file.Type().IsDir() {
			readFilesForDir(iteration+1, dir+"/"+file.Name())
		}
	}
}

func writeCurrentPath() {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Current working dir is: %v \n", path)
}
