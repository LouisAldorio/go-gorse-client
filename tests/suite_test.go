package tests

import (
	"os"
	"testing"

	"github.com/LouisAldorio/go-gorse-client/gorseclient"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/suite"
)

type SuiteTest struct {
	suite.Suite
	api *gorseclient.API
}

func (s *SuiteTest) SetupSuite() {
	if err := godotenv.Load("../.env"); err != nil {
		panic(err)
	}
}

func (s *SuiteTest) SetupTest() {
	s.api = gorseclient.NewAPI(
		os.Getenv("BASE_URL"),
		os.Getenv("TOKEN"),
		gorseclient.Option{
			DebugMode: gorseclient.DEBUG_INFO,
		},
	)
}

func (s *SuiteTest) TearDownTest() {
	s.api = nil
}

func TestRun(t *testing.T) {
	suite.Run(t, new(SuiteTest))
}
