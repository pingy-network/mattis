package envvar

type Config struct {
	HttpHost string `split_words:"true" required:"true"`
	HttpPort string `split_words:"true" required:"true"`

	LogLevel string `split_words:"true" required:"true"`

	RpcAlchemy string `split_words:"true" required:"false"`

	RunServer bool `split_words:"true" required:"true"`
	RunWorker bool `split_words:"true" required:"true"`
}
