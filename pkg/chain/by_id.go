package chain

import "github.com/xh3b4sd/tracer"

func ById(nid int) (Chain, error) {
	switch nid {
	case arbitrumOne.nid:
		return arbitrumOne, nil
	case baseMainnet.nid:
		return baseMainnet, nil
	case ethereumMainnet.nid:
		return ethereumMainnet, nil
	case opMainnet.nid:
		return opMainnet, nil
	default:
		return Chain{}, tracer.Mask(chainIdError, tracer.Context{Key: "id", Value: nid})
	}
}
