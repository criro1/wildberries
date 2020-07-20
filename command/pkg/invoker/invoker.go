// Package invoker ...
package invoker

import (
	"errors"
	"github.com/criro1/wildberries/command/api/v1"
)

type command interface {
	Execute() (str string, err error)
}

// Invoker ...
type Invoker interface {
	Add(com command) (err error)
	Execute() (str string, err error)
	Len() (int)
}

type invoker struct {
	commands []command
}

// Len return amount of commands into slice
func (i *invoker) Len() (int) {
	return len(i.commands)
}

// Add new elem into slice of commands
func (i *invoker) Add(com command) (err error) {
	if len(i.commands) < 0 {
		err = errors.New(v1.BadAmount)
		return
	}
	i.commands = append(i.commands, com)
	return
}

// Execute all methods of invoker
func (i *invoker) Execute() (str string, err error) {
	if len(i.commands) < 0 {
		err = errors.New(v1.BadAmount)
		return
	}
	res := ""
	for _, c := range(i.commands) {
		str, err = c.Execute()
		if err != nil {
			return
		}
		res += str + "\n"
	}
	str = res
	return
}

// NewIvoker ...
func NewIvoker() Invoker {
	return &invoker{}
}
