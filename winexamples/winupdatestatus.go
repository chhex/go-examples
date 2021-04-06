package main

import (
	"fmt"
	wapi "github.com/iamacarpet/go-win64api"
	"time"
)

func main() {
	ret, err := wapi.UpdatesPending()
	if err != nil {
		fmt.Printf("Error fetching data... %s\r\n", err.Error())
	}

	fmt.Printf("Number of Updates Available: %d\n", ret.NumUpdates)
	fmt.Printf("Updates Pending:             %t\n\n", ret.UpdatesReq)
	fmt.Printf("%25s | %25s | %s\n", "EVENT DATE", "STATUS", "UPDATE NAME")
	for _, v := range ret.UpdateHistory {
		fmt.Printf("%25s | %25s | %s\n", v.EventDate.Format(time.RFC822), v.Status, v.UpdateName)
	}
}
