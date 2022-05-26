package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct {
}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error ", err)
		os.Exit(1)
	}

	// bs := make([]byte, 99999)

	// resp.Body.Read(bs)

	// fmt.Println("Response body is ", string(bs))

	lw := logWriter{}

	//io.Copy(os.Stdout, resp.Body)
	io.Copy(lw, resp.Body)
}

//In order to implement / use a certain interface in golang (in this case: Writer)
//1. Define a custom type (struct) to be used on a receiver function and apply the interface's method signature
//2. Initialize a new struct that was actually used for a receiver function (logWriter.Write)
//   and pass it out to the actual function that receives a Writer (io.Copy(Writer, Reader))
//logWriter.Write will be called inside io.Copy

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println("Just wrote this many bytes", len(bs))

	//Println returns the actual bytes and the possible error object
	return fmt.Println(string(bs))
}
