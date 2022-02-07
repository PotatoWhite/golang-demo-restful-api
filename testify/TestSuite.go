package testify

import (
	"demo-restful-api/config/database"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"testing"
)

type TestSuiteEnv struct {
	suite.Suite
	db *gorm.DB
}

func (suite *TestSuiteEnv) SetupSuite() {
	database.Setup()
	suite.db = database.GetDB()
}

func (suite *TestSuiteEnv) TearDownTest() {
}

func (suite *TestSuiteEnv) TearDownSuite() {
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(TestSuiteEnv))
}
