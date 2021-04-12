package qm

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func ObtainQmList() []VM {
	listBuffer := bytes.NewBuffer(make([]byte, 0, 60))
	getQmListCmd := exec.Command("qm", "list")
	getQmListCmd.Stdout = listBuffer
	getQmListCmd.Start()
	err := getQmListCmd.Wait()
	if err != nil {
		fmt.Println("Failed to obtain qemu running config, is 'qm list' available?")
		fmt.Println(err)
		os.Exit(1)

	}
	return ObtainStateFromQM(listBuffer.String())
}
