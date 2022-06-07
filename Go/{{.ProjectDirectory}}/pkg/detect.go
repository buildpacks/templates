package buildpack

import (
	"github.com/buildpacks/libcnb"
)

type Detect struct {
}

func (d Detect) Detect(context libcnb.DetectContext) (libcnb.DetectResult, error) {
	return libcnb.DetectResult{}, nil
}
