package runtime

import "runtime"

var (
	pre = "MATTIS"
	src = "https://github.com/pingy-network/mattis"
	sha = "n/a"
	tag = "n/a"
)

func Arc() string {
	return runtime.GOARCH
}

func Gos() string {
	return runtime.GOOS
}

func Pre() string {
	return pre
}

func Sha() string {
	return sha
}

func Src() string {
	return src
}

func Tag() string {
	return tag
}

func Ver() string {
	return runtime.Version()
}
