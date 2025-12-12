package decoder

import (
	"bytes"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/xh3b4sd/tracer"
)

type Config struct {
	Abi abi.ABI
}

type Decoder struct {
	abi abi.ABI
}

func New(c Config) *Decoder {
	if len(c.Abi.Methods) == 0 {
		tracer.Panic(fmt.Errorf("%T.Abi must not be empty", c))
	}

	return &Decoder{
		abi: c.Abi,
	}
}

// Decode takes the raw calldata bytes and returns the decoded transaction
// intent.
func (d *Decoder) Decode(byt []byte) (string, map[string]any, error) {
	var err error

	var met abi.Method
	{
		met, err = d.method(byt)
		if err != nil {
			return "", nil, tracer.Mask(err)
		}
	}

	var inp map[string]any
	{
		inp, err = d.inputs(met, byt)
		if err != nil {
			return "", nil, tracer.Mask(err)
		}
	}

	return met.Name, inp, nil
}

func (d *Decoder) method(byt []byte) (abi.Method, error) {
	if len(byt) < 4 {
		return abi.Method{}, tracer.Mask(decodeError,
			tracer.Context{Key: "reason", Value: "calldata too short"},
			tracer.Context{Key: "length", Value: len(byt)},
		)
	}

	var sel [4]byte
	{
		sel = [4]byte(byt[0:4])
	}

	for _, x := range d.abi.Methods {
		if bytes.Equal(x.ID, sel[:]) {
			return x, nil
		}
	}

	return abi.Method{}, tracer.Mask(decodeError,
		tracer.Context{Key: "reason", Value: "invalid method selector"},
		tracer.Context{Key: "selector", Value: string(sel[:])},
	)
}

func (d *Decoder) inputs(met abi.Method, byt []byte) (map[string]any, error) {
	var err error

	var unp []any
	{
		unp, err = met.Inputs.Unpack(byt[4:])
		if err != nil {
			return nil, tracer.Mask(err, tracer.Context{Key: "method", Value: met.Name})
		}
	}

	var inp map[string]any
	{
		inp = make(map[string]any, len(met.Inputs))
	}

	for i, x := range met.Inputs {
		var key string
		{
			key = x.Name
		}

		// Input names can be empty for some ABIs.

		if key == "" {
			key = fmt.Sprintf("[%d]", i)
		}

		{
			inp[key] = unp[i]
		}
	}

	return inp, nil
}
