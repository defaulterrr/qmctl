package state

import "fmt"

func MergeStates(neededState, presentState State) {
	neededVMs := neededState.Hosts
	presentVMs := presentState.Hosts

	var isFound bool
	for _, value := range neededVMs {
		isFound = false
		for _, value2 := range presentVMs {
			if value.ID == value2.ID && value.Name == value2.Name {
				isFound = true
				break
			}
		}

		if isFound {
			fmt.Println("Found existing instance: " + value.Name)
			value.Stop()
			value.Set()
			value.Start()
		} else {
			fmt.Println("Didn't find existing instance: " + value.Name)
			value.Clone()
			value.Set()
			value.Start()
		}
	}
}
