package state

import (
	"fmt"
	"sync"

	"github.com/defaulterrr/qmctl/qm"
)

func MergeStates(neededState, presentState State) {
	neededVMs := neededState.Hosts
	presentVMs := presentState.Hosts

	var isFound bool
	for _, value := range neededVMs {
		wg := sync.WaitGroup{}
		wg.Add(1)
		go func(localValue qm.VM) {
			defer wg.Done()
			isFound = false
			for _, value2 := range presentVMs {
				if localValue.ID == value2.ID && localValue.Name == value2.Name {
					isFound = true
					break
				}
			}

			if isFound {
				fmt.Println("Found existing instance: " + localValue.Name)
				localValue.Stop()
				localValue.Set()
				localValue.Start()
			} else {
				fmt.Println("Didn't find existing instance: " + localValue.Name)
				localValue.Clone()
				localValue.Set()
				localValue.Start()
			}
		}(value)
		wg.Wait()
	}
}
