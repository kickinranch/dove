package runtime

import (
	"github.com/kickinranch/dove/dove"
	"github.com/spf13/viper"
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
