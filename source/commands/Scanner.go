package commands

import "fmt"

func NewScannerCmd() *Command {
	return &Command{
		Use:     "vext scan <target>",
		Example: "vext scan 192.168.1.1",
		Short:   "Scan a target for vulnerabilities",

		Run: func(cmd *Command, args []string) {
			fmt.Println("Executing scan command...")
			fmt.Printf("Example: %s\n", cmd.Example)

			if len(args) > 0 {
				fmt.Printf("Target: %s\n", args[0])
			} else {
				fmt.Println("No target specified")
			}
		},
	}
}

var Scanner = NewScannerCmd()
