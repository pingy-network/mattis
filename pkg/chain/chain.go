// See https://chainid.network.
package chain

type Chain struct {
	bnm string
	nid int
}

func (c Chain) Id() int {
	return c.nid
}

func (c Chain) Name() string {
	return c.bnm
}
