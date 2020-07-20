// Package ctrlc ...
package ctrlc

import (
	"github.com/criro1/wildberries/command/pkg/receiver"
)

// CtrlC ...
type CtrlC interface {
	Execute() (str string, err error)
}

type ctrlC struct {
	receiver receiver.Receiver
}

// Executed the required method
func (c *ctrlC) Execute() (str string, err error) {
	str, err = c.receiver.Copy()
	return
}

// NewCtrlC ...
func NewCtrlC(rec receiver.Receiver) CtrlC {
	return &ctrlC{
		receiver: rec,
	}
}