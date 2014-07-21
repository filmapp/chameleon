package mofu

import(
	"github.com/crowdmob/goamz/aws"
	"github.com/crowdmob/goamz/s3"
)

func GetBucket() *s3.Bucket {
	auth, err := aws.EnvAuth()
	if err != nil {
		panic(err.Error())
	}

	s := s3.New(auth, aws.USEast)
	return s.Bucket("filmapp-development")
}
