package qm

import (
	"fmt"

	"github.com/defaulterrr/qmctl/internal/iptables"
)

func MergeStates(needed, running []VM) {
	var (
		found  bool
		config VM
	)
	iptables.FlushSSHForwarding()
	for id := range needed {
		config = needed[id]
		found = false
		for id := range running {
			if config.ID == running[id].ID && config.Name == running[id].Name {
				found = true
				break
			}
		}

		if found {
			fmt.Println("Found existing instance: " + config.Name)
			config.Stop()
			config.Set()
			config.Start()
		} else {
			fmt.Println("Didn't find existing instance: " + config.Name)
			config.Clone()
			config.Set()
			config.Start()
		}
	}
}
