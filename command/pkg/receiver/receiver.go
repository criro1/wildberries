// Package receiver ...
package receiver

import (
	"github.com/criro1/wildberries/command/pkg/api/v1"
)

// Receiver ...
type Receiver interface {
	Copy() (str string, err error)
	Paste() (str string, err error)
}

type receiver struct {
}

// Copy return the string about suceess of copy
func (r *receiver) Copy() (str string, err error) {
	str = v1.CopyS
	return
}

// Paste return the string about suceess of paste
func (r *receiver) Paste() (str string, err error) {
	str = v1.PasteS
	return
}

// NewReceiver ...
func NewReceiver() Receiver {
	return &receiver{}
}
