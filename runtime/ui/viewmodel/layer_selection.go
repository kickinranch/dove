package viewmodel

import (
	"github.com/kickinranch/dove/dove/image"
)

type LayerSelection struct {
	Layer                                                      *image.Layer
	BottomTreeStart, BottomTreeStop, TopTreeStart, TopTreeStop int
}
