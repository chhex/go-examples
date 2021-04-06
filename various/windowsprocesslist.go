package main

import (
	"fmt"
	wapi "github.com/iamacarpet/go-win64api"
)

func main(){
	pr, err := wapi.ProcessList()
	if err != nil {
		fmt.Printf("Error fetching process list... %s\r\n", err.Error())
	}
	for _, p := range pr {
		fmt.Printf("%8d - %-30s - %-30s - %s\r\n", p.Pid, p.Username, p.Executable, p.Fullpath)
	}
}
