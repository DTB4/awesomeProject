package handlers

import (
	"awesomeProject/configs"
	"awesomeProject/midleware"
	"awesomeProject/models"
	"awesomeProject/repository"
	"awesomeProject/services"
	"context"
	"encoding/json"
	"github.com/DTB4/logger/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type Request struct {
	Method    string
	Url       string
	AuthToken string
}

type ExpectedResponse struct {
	StatusCode int
	BodyPart   string
}

type TestCaseHandler struct {
	TestName    string
	Request     Request
	HandlerFunc http.HandlerFunc
	Want        ExpectedResponse
}

type AuthTestSuite struct {
	suite.Suite
	accessToken  string
	refreshToken string
	cfg          *models.Config
	userHandler  *UserHandler
	myLogger     *logger.Logger
	recorder     httptest.ResponseRecorder
	tokenService *services.TokenService
	authHandler  *midleware.AuthHandler
}

func TestTokenServiceTestSuite(t *testing.T) {
	suite.Run(t, new(AuthTestSuite))
}

func (suite *AuthTestSuite) SetupSuite() {
	configs.NewConfig("../configs/test_config.env")
	suite.cfg = configs.InitConfig()
	suite.myLogger = logger.NewLogger(suite.cfg.LogsPath)

	mockTokenRepository := repository.NewMockTokenRepository()
	mockUserService := services.NewMockUserService()

	suite.tokenService = services.NewTokenService(&suite.cfg.AuthConfig, mockTokenRepository)
	suite.authHandler = midleware.NewAuthHandler(suite.tokenService, suite.myLogger)
	suite.userHandler = NewUserHandler(mockUserService, suite.tokenService, suite.myLogger)

	suite.accessToken, suite.refreshToken, _ = suite.tokenService.GeneratePairOfTokens(1)
}

func (suite *AuthTestSuite) TearDownSuite() {

}
func (suite *AuthTestSuite) SetupTest() {

}
func (suite *AuthTestSuite) TearDownTest() {

}
func (suite *AuthTestSuite) Test() {

}

func AssertUserProfileResponse(t *testing.T, recorder *httptest.ResponseRecorder) {
	t.Helper()

	var response models.User
	err := json.Unmarshal([]byte(recorder.Body.String()), &response)

	if assert.NoError(t, err) {
		assert.Equal(t, models.User{
			ID:        1,
			Email:     "email",
			FirstName: "FirstName",
			LastName:  "LastName",
		}, response)
	}
}

func (suite *AuthTestSuite) TestUserHandlerGetProfile() {
	t := suite.T()
	handlerFunc := suite.authHandler.AccessTokenCheck(suite.userHandler.ShowUserProfile)
	cases := []TestCaseHandler{
		{
			TestName: "Successfully get user profile",
			Request: Request{
				Method:    http.MethodGet,
				Url:       "/profile",
				AuthToken: suite.accessToken,
			},
			HandlerFunc: handlerFunc,
			Want: ExpectedResponse{
				StatusCode: 200,
				BodyPart:   "email",
			},
		},
		{
			TestName: "Unauthorized getting user profile",
			Request: Request{
				Method:    http.MethodGet,
				Url:       "/profile",
				AuthToken: "",
			},
			HandlerFunc: handlerFunc,
			Want: ExpectedResponse{
				StatusCode: 401,
				BodyPart:   "invalid",
			},
		},
	}

	for _, test := range cases {
		t.Run(test.TestName, func(t *testing.T) {
			request := httptest.NewRequest(test.Request.Method, test.Request.Url, strings.NewReader(""))

			if test.Request.AuthToken != "" {
				request.Header.Set("Authorization", "Bearer "+test.Request.AuthToken)
			}
			request = request.WithContext(context.WithValue(request.Context(), "CurrentUser", models.ActiveUserData{ID: 1}))
			recorder := httptest.NewRecorder()

			test.HandlerFunc(recorder, request)

			assert.Contains(t, recorder.Body.String(), test.Want.BodyPart)
			if assert.Equal(t, test.Want.StatusCode, recorder.Code) {
				if recorder.Code == http.StatusOK {
					AssertUserProfileResponse(t, recorder)
				}
			}
		})
	}
}
