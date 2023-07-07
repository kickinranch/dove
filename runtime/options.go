package runtime

import (
	"github.com/spf13/viper"

	"github.com/kickinranch/dove/dove"
)

type Options struct {
	Ci           bool
	Image        string
	Source       dove.ImageSource
	IgnoreErrors bool
	ExportFile   string
	CiConfig     *viper.Viper
	BuildArgs    []string
}
