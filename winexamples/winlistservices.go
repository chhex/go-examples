package main

import (
	"fmt"

	wapi "github.com/iamacarpet/go-win64api"
)

func main(){
	svc, err := wapi.GetServices()
	if err != nil {
		fmt.Printf("%s\r\n", err.Error())
	}

	for _, v := range svc {
		fmt.Printf("%-50s - %-75s - Status: %-20s - Accept Stop: %-5t, Running Pid: %d\r\n", v.SCName, v.DisplayName, v.StatusText, v.AcceptStop, v.RunningPid)
	}
}