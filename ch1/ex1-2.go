//Echos the arguments along with the name of the command
package main

import (
	"fmt"
	"os"
)

func main(){
	for idx,arg:= range os.Args{
		fmt.Printf(" %d:%s",idx,arg)
	}
	fmt.Println()
}