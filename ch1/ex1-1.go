//Echos the arguments along with the name of the command
package main

import (
	"fmt"
	"os"
)

func main(){
	s,sep := ""," "
	for _,arg:= range os.Args{
		s += sep + arg
	}
	fmt.Println(s)
}