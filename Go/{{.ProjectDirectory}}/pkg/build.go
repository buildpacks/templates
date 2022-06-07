package buildpack

import (
	"github.com/buildpacks/libcnb"
)

type Build struct {
}

func (b Build) Build(context libcnb.BuildContext) (libcnb.BuildResult, error) {
	return libcnb.BuildResult{}, nil
}
