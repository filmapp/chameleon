package mofu

import(
	"net/http"
	"github.com/go-martini/martini"

	"github.com/crowdmob/goamz/s3"
)

func StartServer() {
	m := martini.Classic()

	bucket := GetBucket()
	m.Map(bucket)

	s := NewImageService()
	m.Map(s)
	defer s.Destroy()

	m.Get("/**", RenderResizedImage)

	m.Run()
}

func RenderResizedImage(bucket *s3.Bucket, s *ImageService, params martini.Params, w http.ResponseWriter, r *http.Request) {
	width, height, path := ParsePath(params["_1"])

	blob, err := bucket.Get(path)
	if err != nil {
		RenderNotFound(w)
		return
	}
	image := s.Resize(uint(width), uint(height), blob)

	w.Write(image)
}

func RenderNotFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404"))
}
