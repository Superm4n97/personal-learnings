package main

import (
	"fmt"
	"os"
)

func main() {
	//Opening file
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("Error occur while opening the file")
		return
	}
	defer file.Close()

	/*file still open now*/

	//get the file size
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("Error occur while size calculation")
		return
	}

	fmt.Println(fileInfo.Name())
	fmt.Println(fileInfo.Mode())
	fmt.Println(fileInfo.Size())
	fmt.Println(fileInfo.IsDir())
	fmt.Println(fileInfo.ModTime())

	//read the file
	bs := make([]byte, fileInfo.Size())
	bytesRead, err := file.Read(bs)
	if err != nil {
		panic(err)
	}

	if int64(bytesRead) == fileInfo.Size() {
		fmt.Println("All data is read")
	}

	str := string(bs)

	fmt.Println(str)
}
