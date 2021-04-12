package state

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func ObtainQmList() State {

	var QmListBytes bytes.Buffer
	getQmListCmd := exec.Command("qm", "list")
	// fmt.Println(getQmListCmd)
	// getQmListCmd.Stdout = os.Stdout
	// getQmListCmd.Stderr = os.Stderr
	getQmListCmd.Stdout = &QmListBytes
	getQmListCmd.Start()
	err := getQmListCmd.Wait()
	if err != nil {
		fmt.Println("Failed to obtain qemu running config, is 'qm list' available?")
		fmt.Println(err)
		os.Exit(1)

	}
	// fmt.Println("BYTES" + QmListBytes.String())
	presentState := ObtainStateFromQM(QmListBytes.String())
	// fmt.Println(presentState)

	return presentState
}
