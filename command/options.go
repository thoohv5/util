package command

import (
	"bytes"
	"io"
)

// Option is command option.
type Option func(*options)

type options struct {
	args   []string
	inArgs []string
	env    []string
	reader func(io.Reader)
	result bytes.Buffer
}

// WithArgs with command file.
func WithArgs(args ...string) Option {
	return func(o *options) {
		o.args = append(o.args, args...)
	}
}

// WithInArgs with command file.
func WithInArgs(inArgs ...string) Option {
	return func(o *options) {
		o.inArgs = inArgs
	}
}

func WithEnv(env ...string) Option {
	return func(o *options) {
		o.env = append(o.env, env...)
	}
}

func WithWriter(reader func(io.Reader)) Option {
	return func(o *options) {
		o.reader = reader
	}
}

func WithResult(result bytes.Buffer) Option {
	return func(o *options) {
		o.result = result
	}
}
