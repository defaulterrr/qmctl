package qm

import (
	"os/exec"
	"strings"

	"github.com/defaulterrr/qmctl/internal/iptables"
)

type VM struct {
	// UID of a vm
	ID string `yaml:"id"`
	// Amount of CPU cores
	CPU string `yaml:"cpu"`
	// Amount of memory in MB
	Mem string `yaml:"mem"`
	// Cloud-Init vm to be cloned from
	Cloudinit string `yaml:"cloudinit"`
	// IP to be set using cloud-init
	IP string `yaml:"ip"`
	// Auto-generated name
	Name string
	// IP gateway
	Gateway string `yaml:"gateway"`
	// Bridge interface to be used as a gateway for private subnet
	Bridge string `yaml:"bridge"`
	// ForwardedPort port to forward vm ssh
	ForwardedPort string `yaml:"forwardedPort"`
}

func (v VM) Clone() {
	command := exec.Command("qm", "clone", v.Cloudinit, v.ID, "--name", v.Name)
	command.Run()
}

func (v VM) Stop() {
	command := exec.Command("qm", "stop", v.ID)
	command.Run()
}

func (v VM) Start() {
	command := exec.Command("qm", "start", v.ID)
	command.Run()
}

func (v VM) Destroy() {
	command := exec.Command("qm", "destroy", v.ID)
	command.Run()
}

func (v VM) Set() {
	setCommand := exec.Command("qm", "set", v.ID, "--cores", v.CPU)
	setCommand.Run()

	setCommand = exec.Command("qm", "set", v.ID, "--memory", v.Mem)
	setCommand.Run()

	setCommand = exec.Command("qm", "set", v.ID, "--net0", "virtio,bridge="+v.Bridge)
	setCommand.Run()

	setCommand = exec.Command("qm", "set", v.ID, "--ipconfig0", "ip="+v.IP+",gw="+v.Gateway)
	setCommand.Run()

	lines := strings.Split(v.IP, "/")
	iptables.AddSSHForwarding(lines[0], "22", v.ForwardedPort)
}
