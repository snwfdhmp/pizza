package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/snwfdhmp/pizza/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
