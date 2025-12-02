package envvar

import (
	"github.com/xh3b4sd/tracer"
)

var rpcConfigError = &tracer.Error{
	Description: "This critical error occurs if not a single RPC provider is properly configured.",
}
