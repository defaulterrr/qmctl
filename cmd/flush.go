package cmd

import (
	"fmt"

	"github.com/defaulterrr/qmctl/state"
	"github.com/spf13/cobra"
)

var flushCmd = &cobra.Command{
	Use:   "flush",
	Short: "Flush all vms created by qmctl",
	Run: func(cmd *cobra.Command, args []string) {
		presentState := state.ObtainQmList()
		fmt.Println(presentState)
		presentState.Flush()
	},
}
