package main

import(
	"github.com/kaiinui/mofu/lib"

	"net/http"
	"github.com/go-martini/martini"

	"github.com/gographics/imagick/imagick"
	"github.com/crowdmob/goamz/s3"
)

func main() {
	m := martini.Classic()

	bucket := mofu.GetBucket()
	m.Map(bucket)

	wand := mofu.GetWand()
	m.Map(wand)
	defer mofu.DestroyWand(wand)

	m.Get("/**", RenderResizedImage)

	m.Run()
}

func RenderResizedImage(bucket *s3.Bucket, wand *imagick.MagickWand, params martini.Params, w http.ResponseWriter, r *http.Request) {
	width, height, path := mofu.ParsePath(params["_1"])

	blob, err := bucket.Get(path)
	if err != nil {
		WriteNotFound(w)
		return
	}
	image := mofu.Resize(wand, uint(width), uint(height), blob)

	w.Write(image)
}

func WriteNotFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404"))
}
