package main

import (
	"fmt"
	"io"
	"os"
	"net/http"
)

func main(){
	for _, url := range os.Args[1:]{
		resp,err := http.Get(url)
		if err!= nil {
			fmt.Fprintf(os.Stderr,"Error: %v",err)
			continue
		}
		_,error := io.Copy(os.Stdout,resp.Body)
		if error != nil {
			fmt.Fprintf(os.Stderr,"Error %v",error)
			continue
		}

	}
}