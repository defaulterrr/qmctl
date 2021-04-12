package state

import (
	"regexp"
	"strings"

	"github.com/defaulterrr/qmctl/qm"
)

type State struct {
	Hosts []qm.VM
}

func NewState() State {
	newState := State{}
	newState.Hosts = make([]qm.VM, 0, 10)
	return newState
}
func ObtainState(config qm.Config) State {
	newState := NewState()

	for index := range config.Hosts {
		for key := range config.Hosts[index] {
			vm := config.Hosts[index][key]
			vm.Name = "qmctl-" + key + "-from-" + vm.Cloudinit
			newState.Hosts = append(newState.Hosts, vm)
		}
	}
	return newState
}

func ObtainStateFromQM(qmlist string) State {
	// var list = `
	// VMID NAME                 STATUS     MEM(MB)    BOOTDISK(GB) PID
	// 100 alpha001             running    2048               8.00 5858
	// 101 WINXP002             running    1024              32.00 6177
	// 102 Win2K                running    2048              32.00 113893
	// 105 axe002               running    16384            100.00 279613`
	state := NewState()
	// fmt.Println(qmlist)
	qmlist = strings.TrimSuffix(qmlist, "\n")
	lines := strings.Split(qmlist, "\n")
	for i := range lines {
		line := lines[i]

		strings.TrimSpace(line)
		space := regexp.MustCompile(`\s+`)
		line = space.ReplaceAllString(line, " ")
		line = strings.TrimSpace(line)
		// fmt.Println(line)
		params := strings.Split(line, " ")
		// fmt.Println(params)
		id := params[0]
		name := params[1]
		mem := params[3]
		vm := qm.VM{
			ID:   id,
			Name: name,
			Mem:  mem,
		}
		state.Hosts = append(state.Hosts, vm)
	}
	return state
}
