package main

import(
	"github.com/gographics/imagick/imagick"
	"github.com/go-martini/martini"
	"github.com/crowdmob/goamz/s3"
	"strings"
	"strconv"
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
	width, height, path := ParsePath(params["_1"])

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

// @args p [string] image/Bs3iSswCIAE075C.500x500.jpg
func ParsePath(p string) (int, int, string) {
	path := strings.Split(p, "/")
	name := path[len(path) - 1]
	nameArr := strings.Split(name, ".")
	sizeParam := nameArr[1]
	width, height := ParseSizeParam(sizeParam)
	fname := nameArr[0] + "." + nameArr[2]
	path[len(path) - 1] = fname
	imgPath := strings.Join(path, "/")

	return width, height, imgPath
}

// @args p [string] 100x100
func ParseSizeParam(p string) (int, int) {
	sizeArr := strings.Split(p, "x")
	width, _ := strconv.Atoi(sizeArr[0])
	height, _ := strconv.Atoi(sizeArr[1])
	return width, height
}

func WriteNotFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404"))
}
