package controllers

import (
	"edge/data/request"
	"edge/data/response"
	"edge/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) CreateUser(user request.CreateUsersRequest) (usr models.User) {
	args := m.Called()
	return args.Get(0).(models.User)
}
func (m *MockUserService) FindAll() (user []models.User) {
	args := m.Called()
	return args.Get(0).([]models.User)
}
func (m *MockUserService) GetUserById(userId int) (user models.User) {
	args := m.Called()
	return args.Get(0).(models.User)
}

type UserControllerSuite struct {
	suite.Suite
	mockUserService *MockUserService
	ctx             *gin.Context
	response        *httptest.ResponseRecorder
}

func (suite *UserControllerSuite) SetupTest() {
	suite.mockUserService = &MockUserService{}
	suite.response = httptest.NewRecorder()

	gin.SetMode(gin.TestMode)
	ctx, _ := gin.CreateTestContext(suite.response)
	ctx.Request = &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
	}
	ctx.Request.Header.Set("Content-Type", "application/json")
	suite.ctx = ctx
}

func (suite *UserControllerSuite) TearDownTest() {
	suite.mockUserService.AssertExpectations(suite.T())
}

func (suite *UserControllerSuite) TestUserController_GetByIdUser_Success() {

	suite.ctx.Request.Method = "GET"
	suite.ctx.Params = []gin.Param{{Key: "userId", Value: "1"}}

	user := models.User{
		Id: 1, Name: "test",
	}

	suite.mockUserService.On("GetUserById", mock.Anything).Return(user)
	controller := NewUserController(suite.mockUserService)
	controller.GetUserByIdHandler(suite.ctx)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   user,
	}

	assert.Equal(suite.T(), http.StatusOK, suite.response.Code)
	jsonData, _ := json.Marshal(webResponse)
	assert.Equal(suite.T(), string(suite.response.Body.Bytes()), string(jsonData))

}

func TestUserControllerSuite(t *testing.T) {
	suite.Run(t, new(UserControllerSuite))
}
