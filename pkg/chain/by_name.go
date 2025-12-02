package chain

import "github.com/xh3b4sd/tracer"

func ByName(bnm string) (Chain, error) {
	switch bnm {
	case arbitrumOne.bnm:
		return arbitrumOne, nil
	case baseMainnet.bnm:
		return baseMainnet, nil
	case ethereumMainnet.bnm:
		return ethereumMainnet, nil
	case opMainnet.bnm:
		return opMainnet, nil
	default:
		return Chain{}, tracer.Mask(chainIdError, tracer.Context{Key: "name", Value: bnm})
	}
}
