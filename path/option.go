package path

type Option func(*options)

// 可选参数
type options struct {
	file string
	skip int
}

func WithFile(file string) Option {
	return func(o *options) {
		o.file = file
	}
}

func WithSkip(skip int) Option {
	return func(o *options) {
		o.skip = skip
	}
}
