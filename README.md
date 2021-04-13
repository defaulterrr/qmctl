# QMCTL 
qmctl is a tool designed to create complex deploys of virtual machines under proxmox (currently 6.3.1)

qmctl is based on CLI tools, mostly on qm and iptables commands

## Usage:
  qmctl [command]

### Available Commands:
```bash
- apply       Applies a deploy described in the yaml file
- flush       Flush all vms created by qmctl
- help        Help about any command
- show        Show running config
```
    

### Flags:
```bash
-h, --help      help for qmctl

Use "qmctl [command] --help" for more information about a command.
```

Example YAML is available in input.yaml
