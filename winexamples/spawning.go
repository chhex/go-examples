package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
)


func handleError_(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err.Error())
	}
}

func main() {

	javaCmd := exec.Command("java", "-version")
	javaOut, errorOut := javaCmd.CombinedOutput()
	handleError_(errorOut, "Failed on Java")
	fmt.Println("> java -version")
	fmt.Println(string(javaOut))

	grepCmd := exec.Command("grep", "hello")

	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepBytes, _ := ioutil.ReadAll(grepOut)
	grepCmd.Wait()

	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, errorOut := lsCmd.Output()
	handleError_(errorOut, "Failed on ls")
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}