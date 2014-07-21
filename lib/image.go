package mofu

import(
	"github.com/gographics/imagick/imagick"
)

type ImageService struct {
	Wand *imagick.MagickWand
}

func NewImageService() *ImageService {
	imagick.Initialize()
	wand := imagick.NewMagickWand()
	return &ImageService{Wand: wand}
}

func (s *ImageService) Resize(w, h uint, blob []byte) []byte {
	wand := s.Wand
	_ = wand.ReadImageBlob(blob)
	wand.ResizeImage(w, h, imagick.FILTER_LANCZOS2_SHARP, 1)
	return wand.GetImageBlob()
}

func (s *ImageService) Destroy() {
	imagick.Terminate()
	s.Wand.Destroy()
}