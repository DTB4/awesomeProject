package services

import (
	"awesomeProject/configs"
	"awesomeProject/models"
	"awesomeProject/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type TestCaseTokenService1 struct {
	TestName     string
	BearerString string
	Want         string
}
type TestCaseValidateToken struct {
	TestName     string
	AccessToken  string
	WantError    bool
	WantErrorMsg string
	WantID       int
}

type TokenServiceTestSuite struct {
	suite.Suite
	cfg          *models.Config
	TokenService *TokenService
}

func TestTokenServiceTestSuite(t *testing.T) {
	suite.Run(t, new(TokenServiceTestSuite))
}

func (suite *TokenServiceTestSuite) SetupSuite() {
	configs.NewConfig("../configs/test_config.env")
	suite.cfg = configs.InitConfig()
	mockTokenRepository := repository.NewMockTokenRepository()
	suite.TokenService = NewTokenService(&suite.cfg.AuthConfig, mockTokenRepository)
}

func (suite *TokenServiceTestSuite) TearDownSuite() {

}
func (suite *TokenServiceTestSuite) SetupTest() {

}
func (suite *TokenServiceTestSuite) TearDownTest() {

}
func (suite *TokenServiceTestSuite) Test() {

}

func (suite *TokenServiceTestSuite) TestTokenServiceGetTokenFromBearerString() {

	testCases := []TestCaseTokenService1{
		{
			TestName:     "Success",
			BearerString: "Bearer token",
			Want:         "token",
		},
		{
			TestName:     "Empty token",
			BearerString: "Bearer",
			Want:         "",
		},
		{
			TestName:     "More than one bearer",
			BearerString: "Bearer token Bearer token",
			Want:         "",
		},
	}

	for _, testCase := range testCases {
		suite.T().Run(testCase.TestName, func(t *testing.T) {
			got := suite.TokenService.ParseFromBearerString(testCase.BearerString)

			assert.Equal(t, testCase.Want, got)
		})

	}
}

const userID = 1

func (suite *TokenServiceTestSuite) TestTokenService_ValidateAccessToken() {

	accessToken, refreshToken, _ := suite.TokenService.generateToken(userID, suite.cfg.AuthConfig.AccessLifeTimeMinutes, suite.cfg.AuthConfig.AccessSecretString, "")
	expAccessToken, _, _ := suite.TokenService.generateToken(userID, -1, suite.cfg.AuthConfig.AccessSecretString, "")

	testCases := []TestCaseValidateToken{
		{
			TestName:     "Success",
			AccessToken:  accessToken,
			WantError:    false,
			WantErrorMsg: "",
			WantID:       userID,
		},
		{
			TestName:     "Invalid token",
			AccessToken:  accessToken + "c",
			WantError:    true,
			WantErrorMsg: "invalid",
			WantID:       0,
		},
		{
			TestName:     "Token with wrong signature",
			AccessToken:  refreshToken,
			WantError:    true,
			WantErrorMsg: "invalid",
			WantID:       0,
		},
		{
			TestName:     "Expired token",
			AccessToken:  expAccessToken,
			WantError:    true,
			WantErrorMsg: "expired",
			WantID:       0,
		},
	}

	for _, testCase := range testCases {
		suite.T().Run(testCase.TestName, func(t *testing.T) {
			gotClaims, gotErr := suite.TokenService.validateToken(testCase.AccessToken, suite.cfg.AuthConfig.AccessSecretString)
			if testCase.WantError {
				if assert.Error(t, gotErr) {
					assert.Contains(t, gotErr.Error(), testCase.WantErrorMsg)
				}
			} else {
				if assert.NoError(t, gotErr) {
					assert.Equal(t, testCase.WantID, gotClaims.ID)
				}
			}
		})
	}
}
