package main

import(
	"fmt"
	"os"
	"bufio"
)

func main(){
	//map of line of text to its overall count
	counts := make(map[string]int)
	//map of line of text to map of file-names wherever its found
	file_counts := make(map[string]map[string]bool)
	files := os.Args[1:]
	for _,file := range files{
		f,err := os.Open(file)
		if err !=nil {
			fmt.Fprint(os.Stderr,"%v\n",err)
		}
		input := bufio.NewScanner(f)
		for input.Scan(){
			if file_counts[input.Text()] == nil {
				file_counts[input.Text()] = make(map[string]bool)
			}
			counts[input.Text()]++
			file_counts[input.Text()][file]=true
		}
	}
	for text,count := range counts{
		if count>1{
			files := file_counts[text]
			fmt.Printf("COUNT:%d\nLINE:%s\nFILENAMES:\n",count,text)
			for filename,_:= range files{
				fmt.Printf("%s\n",filename)
			}
			fmt.Println("--------------")
		}
	}
}