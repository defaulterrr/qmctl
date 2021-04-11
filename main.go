package main

import (
	"github.com/defaulterrr/qmctl/cmd"
)

var data = `
hosts:
- vm1:
    id: 1
    cpu: 1
    mem: 2048
    cloudinit: 9000
    ip: 192.168.5.2
    gateway: 192.168.5.1
- vm3:
    id: 3
    cpu: 1
    mem: 2048
    cloudinit: 9000
    ip: 192.168.5.2
    gateway: 192.168.5.1
`

func main() {
	// Config := qm.Config{}
	// file, openErr := ioutil.ReadFile("./input.yaml")
	// if openErr != nil {
	// 	log.Fatal("Failed to open the file")
	// }
	// err := yaml.Unmarshal(file, &Config)
	// if err != nil {
	// 	fmt.Println(yaml.FormatError(err, true, true))
	// }
	// PresentConfig := qm.Config{}
	// _ = yaml.Unmarshal([]byte(data), &PresentConfig)

	// // fmt.Println(Config)

	// neededState := state.ObtainState(Config)
	// // fmt.Println(neededState)
	// presentState := state.ObtainState(PresentConfig)
	// // fmt.Println(presentState)

	// state.MergeStates(neededState, presentState)

	// state.ObtainStateFromQM("qq")
	cmd.Execute()
}
