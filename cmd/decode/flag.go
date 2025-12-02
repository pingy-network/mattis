package decode

import (
	"github.com/spf13/cobra"
	"github.com/xh3b4sd/tracer"
)

type flag struct {
	Env string
	Net int
	Txn string
}

func (f *flag) Init(cmd *cobra.Command) {
	cmd.Flags().StringVar(&f.Env, "env", ".env", "the environment file to load")
	cmd.Flags().IntVar(&f.Net, "net", 1, "the network ID to inspect")
	cmd.Flags().StringVar(&f.Txn, "txn", "", "the transaction hash to decode")
}

func (f *flag) Verify() error {
	if f.Env == "" {
		return tracer.Mask(cmdFlagError, tracer.Context{Key: "reason", Value: "--env must not be empty"})
	}
	if f.Net == 0 {
		return tracer.Mask(cmdFlagError, tracer.Context{Key: "reason", Value: "--net must not be empty"})
	}
	if f.Txn == "" {
		return tracer.Mask(cmdFlagError, tracer.Context{Key: "reason", Value: "--txn must not be empty"})
	}

	return nil
}
