package cli

import (
	"github.com/defaulterrr/qmctl/internal/qm"
	"github.com/spf13/cobra"
)

var flushCmd = &cobra.Command{
	Use:   "flush",
	Short: "Flush all vms created by qmctl",
	Run: func(cmd *cobra.Command, args []string) {
		s := qm.ObtainQmList()
		for i := range s {
			s[i].Stop()
			s[i].Destroy()
		}
	},
}
