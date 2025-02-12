//go:build !linux && !darwin
// +build !linux,!darwin

package podman

import (
	"fmt"

	"github.com/kickinranch/dove/dove/image"
)

type resolver struct{}

func NewResolverFromEngine() *resolver {
	return &resolver{}
}

func (r *resolver) Build(args []string) (*image.Image, error) {
	return nil, fmt.Errorf("unsupported platform")
}

func (r *resolver) Fetch(id string) (*image.Image, error) {
	return nil, fmt.Errorf("unsupported platform")
}
