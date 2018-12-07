package main

import (
	"os/exec"
	"log"
)

func main() {
	cmd := exec.Command("sleep", "5")
	//cmd.SysProcAttr = &syscall.SysProcAttr{}
	//cmd.SysProcAttr.Credential = &syscall.Credential{
	//	Uid: 0,
	//	Gid: 0,
	//}

	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
	err = cmd.Wait()
	log.Printf("Command finished with error: %v", err)
}
