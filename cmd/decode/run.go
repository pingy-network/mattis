package decode

import (
	"fmt"

	"github.com/pingy-network/mattis/pkg/chain"
	"github.com/pingy-network/mattis/pkg/envvar"
	"github.com/spf13/cobra"
	"github.com/xh3b4sd/tracer"
)

type run struct {
	flag *flag
}

func (r *run) runE(cmd *cobra.Command, arg []string) error {
	var err error

	{
		err = r.flag.Verify()
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var env envvar.Config
	{
		env = envvar.Create(r.flag.Env)
	}

	var cha chain.Chain
	{
		cha, err = chain.ById(r.flag.Net)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	fmt.Printf("%#v\n", env)
	fmt.Printf("\n")
	fmt.Printf("%#v\n", cha)
	fmt.Printf("\n")
	fmt.Printf("%#v\n", r.flag.Txn)

	return nil
}
