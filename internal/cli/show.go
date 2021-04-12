package cli

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"

	"github.com/defaulterrr/qmctl/internal/qm"
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show running config",
	Run: func(cmd *cobra.Command, args []string) {
		var QmListBytes bytes.Buffer
		getQmListCmd := exec.Command("qm", "list")
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
		s := qm.ObtainStateFromQM(QmListBytes.String())
		for i := range s {
			fmt.Println(s[i].Name)
		}
	},
}
