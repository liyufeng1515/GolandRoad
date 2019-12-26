package main

import (
	"fmt"
	"io"
	"os"
)

func main()  {
	file_reader, error := os.Open("/Users/dwyane/Downloads/loadrunner-11.zip")
	if error != nil{
		fmt.Print("Open error:",error)
		return
	}
	defer file_reader.Close()
	file_writer,error := os.Create("/Users/dwyane/Downloads/loadrunner-11-copy.zip")
	if error != nil{
		fmt.Print("Create error:",error)
		return
	}
	buf := make([]byte,4096)//MMU,最小单位是4K
	for{
		n,error :=  file_reader.Read(buf)
		if n==0 || error == io.EOF{
			fmt.Print("读完文件.",n)
			break
		}
		file_writer.Write(buf[:n])
	}
	file_writer.Close()
}


