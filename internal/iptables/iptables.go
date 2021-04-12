package iptables

import (
	"fmt"
	"os"
	"os/exec"
)

func FlushSSHForwarding() {
	FlushCommand := exec.Command("iptables", "--flush")
	if err := FlushCommand.Run(); err != nil {
		fmt.Println("Failed to flush rules")
		fmt.Println(err.Error())
		os.Exit(2)
	}
}

func AddSSHForwarding(destinationAddress, destinationPort, forwardedPort string) {
	fmt.Println(destinationAddress, destinationPort, forwardedPort)

	PreroutingCommand := exec.Command("iptables", "-A", "PREROUTING", "-t", "nat", "-i", "vmbr0", "-p", "tcp", "--dport", forwardedPort, "-j", "DNAT", "--to-destination", destinationAddress+":"+destinationPort)
	ForwardCommand := exec.Command("iptables", "-A", "FORWARD", "-p", "tcp", "-d", destinationAddress, "--dport", destinationPort, "-j", "ACCEPT")
	PostroutingCommand := exec.Command("iptables", "-A", "POSTROUTING", "-t", "nat", "-s", destinationAddress, "-o", "vmbr0", "-j", "MASQUERADE")

	err := PreroutingCommand.Run()
	if err != nil {
		fmt.Println("Failed to add PREROUTING rule-chain rule")
		fmt.Println(err.Error())
		os.Exit(2)
	}

	err = ForwardCommand.Run()
	if err != nil {
		fmt.Println("Failed to add FORWARD rule-chain rule")
		fmt.Println(err.Error())
		os.Exit(2)
	}

	err = PostroutingCommand.Run()
	if err != nil {
		fmt.Println("Failed to add POSTROUTING rule-chain rule")
		fmt.Println(err.Error())
		os.Exit(2)
	}

	// 	iptables -A PREROUTING -t nat -i vmbr0 -p tcp --dport 1000 -j DNAT --to-destination $DESTINATION:$DESTINATIONPORT

	// iptables -A FORWARD -p tcp -d $DESTINATION --dport $DESTINATIONPORT -j ACCEPT

	// iptables -A POSTROUTING -t nat -s $DESTINATION -o vmbr0 -j MASQUERADE

}

func RemoveSSHForwarding(destinationAddress, destinationPort, forwardedPort string) {
	PreroutingCommand := exec.Command("iptables", "-D", "PREROUTING", "-t", "nat", "-i", "vmbr0", "-p", "tcp", "--dport", forwardedPort, "-j", "DNAT", "--to-destination", destinationAddress+":"+destinationPort)
	ForwardCommand := exec.Command("iptables", "-D", "FORWARD", "-p", "tcp", "-d", destinationAddress, "--dport", destinationPort, "-j", "ACCEPT")
	PostroutingCommand := exec.Command("iptables", "-D", "POSTROUTING", "-t", "nat", "-s", destinationAddress, "-o", "vmbr0", "-j", "MASQUERADE")

	err := PreroutingCommand.Run()
	if err != nil {
		fmt.Println("Failed to delete PREROUTING rule-chain rule")
		fmt.Println(err.Error())
		os.Exit(2)
	}

	err = ForwardCommand.Run()
	if err != nil {
		fmt.Println("Failed to delete FORWARD rule-chain rule")
		fmt.Println(err.Error())
		os.Exit(2)
	}

	err = PostroutingCommand.Run()
	if err != nil {
		fmt.Println("Failed to delete POSTROUTING rule-chain rule")
		fmt.Println(err.Error())
		os.Exit(2)
	}
}
