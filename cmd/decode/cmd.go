package decode

import (
	"github.com/spf13/cobra"
)

const (
	use = "decode"
	sho = "Decode a specific multisig transaction."
	lon = "Decode a specific multisig transaction."
)

func New() *cobra.Command {
	var flg *flag
	{
		flg = &flag{}
	}

	var cmd *cobra.Command
	{
		cmd = &cobra.Command{
			Use:   use,
			Short: sho,
			Long:  lon,
			RunE:  (&run{flag: flg}).runE,
		}
	}

	{
		flg.Init(cmd)
	}

	return cmd
}
