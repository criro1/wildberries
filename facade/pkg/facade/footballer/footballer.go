// description of package footballer
package footballer

import (
	"fmt"
)

const (
	kickBall = "%d - %s kicks the ball"
	skipAndTouchBall = "%d - %s skips, but touchs and rolls the ball"
	withoutTouch = "%d - %s skips and don't touch the ball"
)

type InterfaceFootballer interface {
	Choose(i, qty int) (string, error)
}

type footballer struct {
	name string
}

func (f *footballer) Choose(i, qty int) (string, error) {
	if (i < qty - 2) {
		return f.skipWithoutTouch(i)
	} else if (i == qty - 2) {
		return f.skipAndTouch(i)
	} else if (i == qty - 1) {
		return f.kick(i)
	} else {
		return "", nil
	}
}

func (f *footballer) kick(i int) (string, error) {
	return fmt.Sprintf(kickBall, i + 1, f.name), nil
}

func (f *footballer) skipAndTouch(i int) (string, error) {
	return fmt.Sprintf(skipAndTouchBall, i + 1, f.name), nil
}

func (f *footballer) skipWithoutTouch(i int) (string, error) {
	return fmt.Sprintf(withoutTouch, i + 1, f.name), nil
}

func NewFootballer(name string) InterfaceFootballer {
	return &footballer{
		name: name,
	}
}

