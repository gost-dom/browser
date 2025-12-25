package js

type Option struct {
	// instanceMember indicates that the member is on
	instanceMember bool
}

type PropertyOption func(*Option)

func LegacyUnforgable() PropertyOption {
	return func(o *Option) { o.instanceMember = true }
}

func InitOpts(opts ...PropertyOption) Option {
	var o Option
	for _, opt := range opts {
		opt(&o)
	}
	return o
}
