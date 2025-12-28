package commands

type Command struct {
	Use     string
	Short   string
	Example string

	Run func(cmd *Command, args []string)
}

func (c *Command) Execute(args []string) {
	c.Run(c, args)
}
