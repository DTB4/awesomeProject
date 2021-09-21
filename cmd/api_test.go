package main

import (
	"awesomeProject/configs"
	"awesomeProject/dbconstructor"
	"awesomeProject/handlers"
	"awesomeProject/midleware"
	"awesomeProject/models"
	"awesomeProject/repository"
	"awesomeProject/services"
	"database/sql"
	"github.com/DTB4/logger/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const userID = 1

type ApiTestSuite struct {
	suite.Suite
	srv          *httptest.Server
	tokenService *services.TokenService
	accessToken  string
	cfg          *models.Config
	myLogger     *logger.Logger
	db           *sql.DB
}

type TestCaseHandler struct {
	TestName   string
	StatusCode int
	BodyPart   string
	Method     string
	Url        string
	AuthToken  string
}

func TestTokenServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ApiTestSuite))
}

func (suite *ApiTestSuite) SetupSuite() {
	configs.NewConfig("../configs/test_config.env")
	suite.cfg = configs.InitConfig()
	suite.myLogger = logger.NewLogger(suite.cfg.LogsPath)
	suite.db = dbconstructor.NewDB(&suite.cfg.DBConfig, suite.myLogger)
	userRepository := repository.NewUserRepository(suite.db)
	tokenRepository := repository.NewTokenRepository(suite.db)
	userService := services.NewUserService(userRepository)
	suite.tokenService = services.NewTokenService(&suite.cfg.AuthConfig, tokenRepository)
	userHandler := handlers.NewUserHandler(userService, suite.tokenService, suite.myLogger)
	middleware := midleware.NewAuthHandler(suite.tokenService, suite.myLogger)

	mux := http.NewServeMux()
	mux.HandleFunc("/profile", middleware.AccessTokenCheck(userHandler.ShowUserProfile))
	suite.srv = httptest.NewServer(mux)

	SQLUpSequence(suite.myLogger, suite.db)

	accessToken, _, err := suite.tokenService.GeneratePairOfTokens(userID)
	if err != nil {
		log.Fatal(err)
	}
	suite.accessToken = accessToken
}

func (suite *ApiTestSuite) TearDownSuite() {
	SQLDownSequence(suite.myLogger, suite.db)
}
func (suite *ApiTestSuite) SetupTest() {

}
func (suite *ApiTestSuite) TearDownTest() {

}
func (suite *ApiTestSuite) Test() {

}

func (suite *ApiTestSuite) TestWalkApiGetProfile() {
	t := suite.T()
	cases := []TestCaseHandler{
		{
			TestName:   "Successfully get user profile",
			Method:     http.MethodGet,
			Url:        suite.srv.URL + "/profile",
			AuthToken:  suite.accessToken,
			StatusCode: 200,
			BodyPart:   "test-1@example.com",
		},
		{
			TestName:   "Unauthorized request",
			Method:     http.MethodGet,
			Url:        suite.srv.URL + "/profile",
			AuthToken:  "",
			StatusCode: 401,
			BodyPart:   "invalid",
		},
		{
			TestName:   "Wrong address",
			Method:     http.MethodGet,
			Url:        suite.srv.URL,
			AuthToken:  suite.accessToken,
			StatusCode: 404,
			BodyPart:   "not found",
		},
		{
			TestName:   "Wrong method",
			Method:     http.MethodPost,
			Url:        suite.srv.URL + "/profile",
			AuthToken:  suite.accessToken,
			StatusCode: 405,
			BodyPart:   "Only POST is Allowed",
		},
	}

	for _, test := range cases {
		t.Run(test.TestName, func(t *testing.T) {

			req, _ := http.NewRequest(test.Method, test.Url, http.NoBody)
			req.Header.Set("Authorization", "Bearer "+test.AuthToken)

			client := http.Client{}
			resp, err := client.Do(req)

			if assert.NoError(t, err) {
				assert.Equal(t, test.StatusCode, resp.StatusCode)
			}
		})
	}
}

func SQLUpSequence(myLogger *logger.Logger, db *sql.DB) {
	file, err := ioutil.ReadFile("sql_suite/api_test_suite.sql")
	if err != nil {
		myLogger.FatalLog("Error while file opening", err)
	}
	stringSQL := string(file)
	sqlCommands := strings.Split(stringSQL, ";")
	tx, err := db.Begin()
	if err != nil {
		myLogger.FatalLog("Fail to start transaction", err)
	}
	myLogger.InfoLog("Transaction find ", len(sqlCommands), " commands")
	for i, command := range sqlCommands {
		if command != "" {
			_, err := tx.Exec(command)
			if err != nil {
				myLogger.ErrorLog("Error in transaction", err)
			}
		}
		myLogger.InfoLog("In process command ", i+1, " from", len(sqlCommands))
	}
	err = tx.Commit()
	if err != nil {
		myLogger.FatalLog("Error during Commit", err)
	}
	myLogger.InfoLog("Setup Suite complete")
}

func SQLDownSequence(myLogger *logger.Logger, db *sql.DB) {
	file, err := ioutil.ReadFile("sql_suite/api_test_suite_drop.sql")
	if err != nil {
		myLogger.FatalLog("Error while file opening", err)
	}
	stringSQL := string(file)
	sqlCommands := strings.Split(stringSQL, ";")
	tx, err := db.Begin()
	if err != nil {
		myLogger.FatalLog("Fail to start transaction", err)
	}
	myLogger.InfoLog("Transaction find ", len(sqlCommands), " commands")
	for i, command := range sqlCommands {
		if command != "" {
			_, err := tx.Exec(command)
			if err != nil {
				myLogger.ErrorLog("Error in transaction", err)
			}
		}
		myLogger.InfoLog("In process command ", i+1, " from", len(sqlCommands))
	}
	err = tx.Commit()
	if err != nil {
		myLogger.FatalLog("Error during Commit", err)
	}
	myLogger.InfoLog("Setup DROP complete")
}
