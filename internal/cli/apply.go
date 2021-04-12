package cli

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/defaulterrr/qmctl/internal/qm"
	"github.com/goccy/go-yaml"
	"github.com/spf13/cobra"
)

var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Applies a deploy described in the yaml file",
	Run: func(command *cobra.Command, args []string) {
		if DeployFile == "" {
			fmt.Println("No deploy file was specified")
			os.Exit(1)
		} else {
			config := qm.YamlConfig{}
			file, openErr := ioutil.ReadFile("./input.yaml")
			if openErr != nil {
				log.Fatal("Failed to open the file")
			}
			err := yaml.Unmarshal(file, &config)
			if err != nil {
				fmt.Println(yaml.FormatError(err, true, true))
			}
			neededState := qm.ObtainState(config)
			presentState := qm.ObtainQmList()
			fmt.Println(neededState)
			fmt.Println(presentState)

			qm.MergeStates(neededState, presentState)

		}
	},
}
