package chain

import (
	"github.com/xh3b4sd/tracer"
)

var chainIdError = &tracer.Error{
	Description: "This runtime error occurs if a blockchain name cannot be mapped to a network ID.",
}
