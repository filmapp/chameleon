package main

import(
	"github.com/gographics/imagick/imagick"
	"github.com/go-martini/martini"
	"github.com/crowdmob/goamz/s3"
	"log"
	"net/http"
	"github.com/kaiinui/mofu/lib"
)

func main() {
	m := martini.Classic()

	bucket := mofu.GetBucket()
	m.Map(bucket)

	wand := mofu.GetWand()
	m.Map(wand)
	defer imagick.Terminate()
	defer wand.Destroy()

	m.Get("/**", GetResizedImage)

	m.Run()
}

func GetResizedImage(bucket *s3.Bucket, wand *imagick.MagickWand, params martini.Params, w http.ResponseWriter, r *http.Request) {
	width, height, path := mofu.ParsePath(params["_1"])

	log.Println("getting " + path)
	blob, err := bucket.Get(path)
	if err != nil {
		log.Println(err)
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
