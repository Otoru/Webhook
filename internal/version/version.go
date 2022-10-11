package version

import (
	"fmt"
	"runtime"
)

const version = "v0.1.0"

type Info struct {
	Version string `json:"version,omitempty"`
	Golang  string `json:"go,omitempty"`
}

func Get(short bool) string {
	if short {
		return version
	}

	info := Info{Version: version, Golang: runtime.Version()}
	return fmt.Sprintf("%#v", info)
}
