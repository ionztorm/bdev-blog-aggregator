package command

import (
	"fmt"

	"github.com/ionztorm/gator/internal/state"
)

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	handlers map[string]func(*state.State, Command) error
}

func (c *Commands) Run(s *state.State, cmd Command) error {
	handler, exists := c.handlers[cmd.Name]
	if !exists {
		return fmt.Errorf("unknown command: %s", cmd.Name)
	}
	return handler(s, cmd)
}

func (c *Commands) Register(name string, f func(*state.State, Command) error) {
	if c.handlers == nil {
		c.handlers = make(map[string]func(*state.State, Command) error)
	}
	c.handlers[name] = f
}

func newCmdRegister() *Commands {
	return &Commands{
		handlers: make(map[string]func(*state.State, Command) error),
	}
}

var cmdRegistry = newCmdRegister()

func registerGlobal(name string, handler func(*state.State, Command) error) {
	cmdRegistry.Register(name, handler)
}

func GetCmdRegistry() *Commands {
	return cmdRegistry
}
