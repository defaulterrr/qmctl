package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/defaulterrr/qmctl/qm"
	"github.com/defaulterrr/qmctl/state"
	"github.com/goccy/go-yaml"
	"github.com/spf13/cobra"
)

var DeployFile string

var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Applies a deploy described in the yaml file",
	Run: func(command *cobra.Command, args []string) {
		if DeployFile == "" {
			fmt.Println("No deploy file was specified")
			os.Exit(1)
		} else {
			Config := qm.Config{}
			file, openErr := ioutil.ReadFile("./input.yaml")
			if openErr != nil {
				log.Fatal("Failed to open the file")
			}
			err := yaml.Unmarshal(file, &Config)
			if err != nil {
				fmt.Println(yaml.FormatError(err, true, true))
			}
			var QmListBytes bytes.Buffer
			getQmListCmd := exec.Command("qm", "list", "|", "grep", "qmctl-")
			fmt.Println(getQmListCmd)
			// getQmListCmd.Stdout = os.Stdout
			// getQmListCmd.Stderr = os.Stderr
			getQmListCmd.Stdout = &QmListBytes
			getQmListCmd.Start()
			err = getQmListCmd.Wait()
			if err != nil {
				fmt.Println("Failed to obtain qemu running config, is 'qm list' available?")
				if !Testing {
					fmt.Println(err)
					os.Exit(1)
				}
			}
			neededState := state.ObtainState(Config)
			fmt.Println("BYTES" + QmListBytes.String())
			presentState := state.ObtainStateFromQM(QmListBytes.String())
			fmt.Println(neededState)
			fmt.Println(presentState)

			state.MergeStates(neededState, presentState)

		}
	},
}
