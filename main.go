package main

import(
	"github.com/gographics/imagick/imagick"
	"github.com/go-martini/martini"
	"github.com/crowdmob/goamz/aws"
	"github.com/crowdmob/goamz/s3"
	"strings"
	"strconv"
	"log"
	"net/http"
)

func main() {
	m := martini.Classic()

	m.Get("/**", GetResizedImage)

	m.Run()
}

func GetResizedImage(params martini.Params, w http.ResponseWriter, r *http.Request) {
	path := strings.Split(params["_1"], "/")
	sizeParam := path[0]
	sizeArr := strings.Split(sizeParam, "x")
	width, _ := strconv.Atoi(sizeArr[0])
	height, _ := strconv.Atoi(sizeArr[1])

	imgPathArr := path[1:len(path)]
	imgPath := strings.Join(imgPathArr, "/")

	log.Println(imgPath)

	bucket := GetBucket()
	blob, err := GetImage(bucket, imgPath)
	if err != nil {
		WriteNotFound(w)
		return
	}
	image := Resize(uint(width), uint(height), blob)

	w.Write(image)
}

func ParseParam(p string) {

}

func WriteNotFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404"))
}

func GetBucket() *s3.Bucket {
	auth, err := aws.EnvAuth()
	if err != nil {
		panic(err.Error())
	}

	s := s3.New(auth, aws.APNortheast)
	return s.Bucket("kigotest2")
}

func GetImage(bucket *s3.Bucket, path string) ([]byte, error) {
	blob, err := bucket.Get(path)
	return blob, err
}

func Resize(w, h uint, blob []byte) []byte {
	imagick.Initialize()
	defer imagick.Terminate()

	wand := imagick.NewMagickWand()
	defer wand.Destroy()

	_ = wand.ReadImageBlob(blob)
	wand.ResizeImage(w, h, imagick.FILTER_LANCZOS2_SHARP, 1)
	return wand.GetImageBlob()
}