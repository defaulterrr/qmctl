package cmd

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"

	"github.com/defaulterrr/qmctl/state"
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show running config",
	Run: func(cmd *cobra.Command, args []string) {
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
			if !Testing {
				fmt.Println(err)
				os.Exit(1)
			}
		}
		// fmt.Println("BYTES" + QmListBytes.String())
		presentState := state.ObtainStateFromQM(QmListBytes.String())
		for i := range presentState.Hosts {
			fmt.Println(presentState.Hosts[i].Name)
		}

	},
}
