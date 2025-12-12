package decoder

import (
	"fmt"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/go-cmp/cmp"
	"github.com/pingy-network/mattis/pkg/contract/erc20"
	"github.com/xh3b4sd/tracer"
)

func Test_Decoder_Decode(t *testing.T) {
	testCases := []struct {
		abi abi.ABI
		byt []byte
		met string
		inp map[string]any
	}{
		// Case 000
		{
			abi: musAbi(erc20.Erc20ContractABI),
			byt: common.FromHex("0xa9059cbb000000000000000000000000fb1456057ee45d465f2af571e9d178560e1bba9b000000000000000000000000000000000000000000000000000000001dcd6500"),
			met: "transfer",
			inp: map[string]any{
				"_to":    common.HexToAddress("0xfB1456057Ee45d465f2AF571e9d178560E1BBa9b"),
				"_value": big.NewInt(500000000),
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var err error

			var dec *Decoder
			{
				dec = New(Config{
					Abi: tc.abi,
				})
			}

			var met string
			var inp map[string]any
			{
				met, inp, err = dec.Decode(tc.byt)
				if err != nil {
					t.Fatal(err)
				}
			}

			var opt []cmp.Option
			{
				opt = []cmp.Option{
					cmp.AllowUnexported(big.Int{}),
				}
			}

			if dif := cmp.Diff(tc.met, met); dif != "" {
				t.Fatalf("-expected +actual:\n%s", dif)
			}

			if dif := cmp.Diff(tc.inp, inp, opt...); dif != "" {
				t.Fatalf("-expected +actual:\n%s", dif)
			}
		})
	}
}

func musAbi(jsn string) abi.ABI {
	a, e := abi.JSON(strings.NewReader(jsn))
	if e != nil {
		tracer.Panic(tracer.Mask(e))
	}

	return a
}
