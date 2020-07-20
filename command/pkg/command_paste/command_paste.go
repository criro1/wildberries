// Package ctrlv ...
package ctrlv

import (
	"github.com/criro1/wildberries/command/pkg/receiver"
)

// CtrlV ...
type CtrlV interface {
	Execute() (str string, err error)
}

type ctrlV struct {
	receiver receiver.Receiver
}

// Executed the required method
func (c *ctrlV) Execute() (str string, err error) {
	str, err = c.receiver.Paste()
	return
}

// NewCtrlV ...
func NewCtrlV(rec receiver.Receiver) CtrlV {
	return &ctrlV{
		receiver: rec,
	}
}