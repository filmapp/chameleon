package main

import(
	"github.com/go-martini/martini"
	"github.com/crowdmob/goamz/aws"
	"github.com/crowdmob/goamz/s3"
	"net/http"
)

func main() {
	m := martini.Classic()

	m.Get("/**", func(res http.ResponseWriter, req *http.Request) {
		auth, err := aws.EnvAuth()
		if err != nil {
			panic(err.Error())
		}

		s := s3.New(auth, aws.USEast)
		bucket := s.Bucket("kigotest")

		data, err := bucket.Get("sk.png")
		if err != nil {
			panic(err.Error())
		}

		res.Write(data)
	})

	m.Run()
}