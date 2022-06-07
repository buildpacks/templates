package main

import (
	"github.com/buildpacks/libcnb"

	bp "{{ .ModuleName }}/pkg/buildpack"
)

func main() {
	libcnb.Main(
		bp.Detect{},
		bp.Build{},
	)
}
