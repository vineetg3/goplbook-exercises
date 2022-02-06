//Echos the arguments along with the name of the command
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main(){
	s,sep := ""," "
	start := time.Now()
	for _,arg:= range os.Args{
		s += sep + arg
	}
	t :=time.Now().Sub(start)
	fmt.Println(s)
	fmt.Printf("Using + sign time: %v\n",t)
	s= ""
	start = time.Now()
	s = strings.Join(os.Args," ")
	t=time.Now().Sub(start)
	fmt.Println(s)
	fmt.Printf("Using join time: %v\n",t)

}