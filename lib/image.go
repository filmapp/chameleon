package mofu

import(
	"github.com/gographics/imagick/imagick"
)

func GetWand() *imagick.MagickWand {
	imagick.Initialize()
	wand := imagick.NewMagickWand()
	return wand
}

func Resize(wand *imagick.MagickWand, w, h uint, blob []byte) []byte {
	_ = wand.ReadImageBlob(blob)
	wand.ResizeImage(w, h, imagick.FILTER_LANCZOS2_SHARP, 1)
	return wand.GetImageBlob()
}