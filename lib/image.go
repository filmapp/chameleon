package mofu

import(
	"github.com/gographics/imagick/imagick"
)

func GetWand() *imagick.MagickWand {
	imagick.Initialize()
	wand := imagick.NewMagickWand()
	return wand
}