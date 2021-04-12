package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

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
			neededState := state.ObtainState(Config)
			// fmt.Println("BYTES" + QmListBytes.String())
			presentState := state.ObtainQmList()
			fmt.Println(neededState)
			fmt.Println(presentState)

			state.MergeStates(neededState, presentState)

		}
	},
}
