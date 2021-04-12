package qm

import (
	"os/exec"
	"fmt"
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
}

func (mach VM) Clone() {
	command := exec.Command("qm", "clone", mach.Cloudinit, mach.ID, "--name", mach.Name)
	command.Start()
	command.Wait()
}

func (mach VM) Stop() {
	command := exec.Command("qm", "stop", mach.ID)
	command.Start()
	command.Wait()
}

func (mach VM) Start() {
	command := exec.Command("qm", "start", mach.ID)
	command.Start()
	command.Wait()
}

func (mach VM) Destroy() {
	// command := exec.Command("qm", "destroy", mach.ID)
	// command.Run()
	// command.Wait()
}

func (mach VM) Set() {
	setCommand := exec.Command("qm", "set", mach.ID, "--cores", mach.CPU)
	setCommand.Start()
	err := setCommand.Wait()
	fmt.Println(err)

	setCommand = exec.Command("qm", "set", mach.ID, "--memory", mach.Mem)
	setCommand.Start()
	err = setCommand.Wait()
	fmt.Println(err)


	setCommand = exec.Command("qm", "set", mach.ID, "--ipconfig0", "ip="+mach.IP+",gw="+mach.Gateway)
	fmt.Println(*setCommand)
	setCommand.Start()
	err = setCommand.Wait()
	fmt.Println(err)

}
