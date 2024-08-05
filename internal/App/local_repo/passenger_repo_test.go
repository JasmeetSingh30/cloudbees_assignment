package local_repo

import (
	"testing"

	customErrors "cloudbees_assignmnet.com/ass1/internal/App/custom_errors"
	"cloudbees_assignmnet.com/ass1/internal/App/model"
	"github.com/stretchr/testify/suite"
)

type PassengerRepositorySuite struct {
	suite.Suite
	repo *PassengerRepository
}

func (suite *PassengerRepositorySuite) SetupTest() {
	suite.repo = NewPassengerRepository()
}

func (suite *PassengerRepositorySuite) TestGetByID_ExistingPassenger() {
	user := model.NewUser("John", "Singh", "Jasmeet@cloudbees.com")
	suite.repo.passengers[1] = user

	foundUser, err := suite.repo.GetByID(1)

	suite.Equal(user, foundUser)
	suite.NoError(err)
}

func (suite *PassengerRepositorySuite) TestGetByID_NonExistingPassenger() {
	foundUser, err := suite.repo.GetByID(2)

	suite.Nil(foundUser)
	suite.Equal(customErrors.NotFound("user %d", 2), err)
}

func (suite *PassengerRepositorySuite) TestGetByData_ExistingPassenger() {
	user := model.NewUser("Jasmeet", "Singh", "Jasmeet@cloudbees.com")
	suite.repo.passengers[1] = user

	foundUser, err := suite.repo.GetByData("Jasmeet", "Singh", "Jasmeet@cloudbees.com")

	suite.Equal(user, foundUser)
	suite.NoError(err)
}

func (suite *PassengerRepositorySuite) TestGetByData_NonExistingPassenger() {
	foundUser, err := suite.repo.GetByData("Jasmeet", "Singh", "Jasmeet@cloudbees.com")

	suite.Nil(foundUser)
	suite.Equal(customErrors.NotFound("user Jasmeet Singh"), err)
}

func (suite *PassengerRepositorySuite) TestFindOrCreate_NewPassenger() {
	newUser := suite.repo.FindOrCreate("Jasmeet", "Singh", "Jasmeet@cloudbees.com")

	suite.NotNil(newUser)
	suite.Equal("Jasmeet", newUser.FirstName)
	suite.Equal("Singh", newUser.LastName)
	suite.Equal("Jasmeet@cloudbees.com", newUser.Email)
	suite.Equal(uint(1), newUser.ID)
}

func (suite *PassengerRepositorySuite) TestFindOrCreate_ExistingPassenger() {
	existingUser := model.NewUser("Jasmeet", "Singh", "jasmeet@cloudbees.com")
	suite.repo.passengers[1] = existingUser

	foundUser := suite.repo.FindOrCreate("Jasmeet", "Singh", "jasmeet@cloudbees.com")

	suite.Equal(existingUser, foundUser)
}


func TestPassengerRepositorySuite(t *testing.T) {
	suite.Run(t, new(PassengerRepositorySuite))
}
