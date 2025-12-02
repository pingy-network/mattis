package envvar

import "github.com/xh3b4sd/tracer"

func (c Config) verify() error {
	if c.RpcAlchemy == "" {
		return tracer.Mask(rpcConfigError,
			tracer.Context{Key: "reason", Value: "provider api key must not be empty"},
			tracer.Context{Key: "providers", Value: "alchemy"},
		)
	}

	return nil
}
