package envvar

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/pingy-network/mattis/pkg/runtime"
	"github.com/xh3b4sd/tracer"
)

// Create loads the environment config of this process given a local file path.
// For instance ".env" will load the files ".env" and ".env.secret".
func Create(pat string) Config {
	var err error
	var env Config

	{
		err = godotenv.Load(pat, pat+".secret")
		if err != nil {
			tracer.Panic(tracer.Mask(err))
		}
	}

	{
		err = envconfig.Process(runtime.Pre(), &env)
		if err != nil {
			tracer.Panic(tracer.Mask(err))
		}
	}

	{
		err = env.verify()
		if err != nil {
			tracer.Panic(tracer.Mask(err))
		}
	}

	return env
}
