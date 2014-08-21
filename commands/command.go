package commands

import (
	"github.com/dulo-tech/thumbnailer/thumbnailer"
)

type ChannelFinished chan bool
type ChannelError chan error

// Commander is an interface for types which execute command line instructions.
type Commander interface {
	SetChannels(*ChannelFinished, *ChannelError)
	Execute(string, string)
}

// Command is used to create thumbs from the command line.
type Command struct {
	opts         *thumbnailer.Options
	chanFinished *ChannelFinished
	chanError    *ChannelError
}

// New creates and returns a new Command instance.
func newCommand(opts *thumbnailer.Options) *Command {
	return &Command{
		opts: opts,
	}
}

// SetChannels is used to set the channels used to coordinate command executors.
func (c *Command) SetChannels(f *ChannelFinished, e *ChannelError) {
	c.chanFinished = f
	c.chanError = e
}
