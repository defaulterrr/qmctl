package qm

import (
	"regexp"
	"strings"
)

func ObtainState(config YamlConfig) []VM {
	newState := make([]VM, 0, 10)

	for i := range config.Hosts {
		for key := range config.Hosts[i] {
			vm := config.Hosts[i][key]
			vm.Name = "qmctl-" + key + "-from-" + vm.Cloudinit
			newState = append(newState, vm)
		}
	}
	return newState
}

func ObtainStateFromQM(qmlist string) []VM {
	state := make([]VM, 0, 10)
	repRegx := regexp.MustCompile(`\s+`)

	qmlist = strings.TrimSuffix(qmlist, "\n")
	lines := strings.Split(qmlist, "\n")
	lines = lines[1:]

	for i := range lines {
		line := lines[i]

		line = strings.TrimSpace(line)
		line = repRegx.ReplaceAllString(line, " ")
		line = strings.TrimSpace(line)

		params := strings.Split(line, " ")

		// Checking that it's our vm
		if strings.HasPrefix(params[1], "qmctl-") {
			state = append(state, VM{
				ID:   params[0],
				Name: params[1],
				Mem:  params[3],
			})
		}
	}
	return state
}
