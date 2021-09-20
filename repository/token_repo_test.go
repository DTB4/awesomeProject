package repository

import (
	"awesomeProject/configs"
	"awesomeProject/dbconstructor"
	"awesomeProject/models"
	"database/sql"
	"github.com/DTB4/logger/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"strings"
	"testing"
)

type TestCaseTokenRepoGetByUID struct {
	TestName     string
	UIDString    string
	WantUserID   int
	WantError    bool
	WantErrorMsg string
}
type TestCaseTokenRepoUpdate struct {
	TestName         string
	UserID           int
	UIDString        string
	WantRowsAffected int
	WantError        bool
	WantErrorMsg     string
}
type TestCaseTokenRepoNullUID struct {
	TestName         string
	UserID           int
	WantRowsAffected int
	WantError        bool
	WantErrorMsg     string
}

type TokenServiceTestSuite struct {
	suite.Suite
	cfg             *models.Config
	tokenRepository TokenRepositoryI
	myLogger        *logger.Logger
	db              *sql.DB
}

func TestTokenServiceTestSuite(t *testing.T) {
	suite.Run(t, new(TokenServiceTestSuite))
}

func (suite *TokenServiceTestSuite) SetupSuite() {
	configs.NewConfig("../configs/test_config.env")
	suite.cfg = configs.InitConfig()
	suite.myLogger = logger.NewLogger(suite.cfg.LogsPath)
	suite.db = dbconstructor.NewDB(&suite.cfg.DBConfig, suite.myLogger)
	suite.tokenRepository = NewTokenRepository(suite.db)

	file, err := ioutil.ReadFile("sql_suite/token_repo_test_suite.sql")
	if err != nil {
		suite.myLogger.FatalLog("Error while file opening", err)
	}
	stringSQL := string(file)
	sqlCommands := strings.Split(stringSQL, ";")
	tx, err := suite.db.Begin()
	if err != nil {
		suite.myLogger.FatalLog("Fail to start transaction", err)
	}
	suite.myLogger.InfoLog("Transaction find ", len(sqlCommands), " commands")
	for i, command := range sqlCommands {
		if command != "" {
			_, err := tx.Exec(command)
			if err != nil {
				suite.myLogger.ErrorLog("Error in transaction", err)
			}
		}
		suite.myLogger.InfoLog("In process command ", i+1, " from", len(sqlCommands))
	}
	err = tx.Commit()
	if err != nil {
		suite.myLogger.FatalLog("Error during Commit", err)
	}
	suite.myLogger.InfoLog("Setup Suite complete")
}

func (suite *TokenServiceTestSuite) TearDownSuite() {
	file, err := ioutil.ReadFile("sql_suite/token_repo_test_suite_drop.sql")
	if err != nil {
		suite.myLogger.FatalLog("Error while file opening", err)
	}
	stringSQL := string(file)
	sqlCommands := strings.Split(stringSQL, ";")
	tx, err := suite.db.Begin()
	if err != nil {
		suite.myLogger.FatalLog("Fail to start transaction", err)
	}
	suite.myLogger.InfoLog("Transaction find ", len(sqlCommands), " commands")
	for i, command := range sqlCommands {
		if command != "" {
			_, err := tx.Exec(command)
			if err != nil {
				suite.myLogger.ErrorLog("Error in transaction", err)
			}
		}
		suite.myLogger.InfoLog("In process command ", i+1, " from", len(sqlCommands))
	}
	err = tx.Commit()
	if err != nil {
		suite.myLogger.FatalLog("Error during Commit", err)
	}
	suite.myLogger.InfoLog("Setup DROP complete")
}
func (suite *TokenServiceTestSuite) SetupTest() {

}
func (suite *TokenServiceTestSuite) TearDownTest() {

}
func (suite *TokenServiceTestSuite) Test() {

}

func (suite *TokenServiceTestSuite) TestTokenRepoGetByUID() {

	testCases := []TestCaseTokenRepoGetByUID{
		{
			TestName:   "Success",
			UIDString:  "uid_of_user_1",
			WantUserID: 1,
			WantError:  false,
		},
		{
			TestName:   "UID not found",
			UIDString:  "uid_of_user_5",
			WantUserID: 0,
			WantError:  false,
		},
	}

	for _, testCase := range testCases {
		suite.T().Run(testCase.TestName, func(t *testing.T) {
			got, gotErr := suite.tokenRepository.GetByUID(testCase.UIDString)
			if testCase.WantError {
				if assert.Error(t, gotErr) {
					assert.Contains(t, gotErr.Error(), testCase.WantErrorMsg)
				}
			} else {
				if assert.NoError(t, gotErr) {
					assert.Equal(t, testCase.WantUserID, got)
				}
			}
		})

	}
}

func (suite *TokenServiceTestSuite) TestTokenRepoUpdate() {

	testCases := []TestCaseTokenRepoUpdate{
		{
			TestName:         "Success",
			UserID:           2,
			UIDString:        "new_uid_of_user_2",
			WantRowsAffected: 1,
			WantError:        false,
		},
		{
			TestName:         "UserID not found",
			UserID:           5,
			UIDString:        "uid_of_user_5",
			WantRowsAffected: 0,
			WantError:        false,
		},
	}

	for _, testCase := range testCases {
		suite.T().Run(testCase.TestName, func(t *testing.T) {
			got, gotErr := suite.tokenRepository.Update(testCase.UserID, testCase.UIDString)
			if testCase.WantError {
				if assert.Error(t, gotErr) {
					assert.Contains(t, gotErr.Error(), testCase.WantErrorMsg)
				}
			} else {
				if assert.NoError(t, gotErr) {
					assert.Equal(t, testCase.WantRowsAffected, got)
				}
			}
		})

	}
}

func (suite *TokenServiceTestSuite) TestTokenRepoNullUID() {

	testCases := []TestCaseTokenRepoNullUID{
		{
			TestName:         "Success",
			UserID:           3,
			WantRowsAffected: 1,
			WantError:        false,
		},
		{
			TestName:         "UserID not found",
			UserID:           5,
			WantRowsAffected: 0,
			WantError:        false,
		},
	}
	for _, testCase := range testCases {
		suite.T().Run(testCase.TestName, func(t *testing.T) {
			got, gotErr := suite.tokenRepository.NullUID(testCase.UserID)
			if testCase.WantError {
				if assert.Error(t, gotErr) {
					assert.Contains(t, gotErr.Error(), testCase.WantErrorMsg)
				}
			} else {
				if assert.NoError(t, gotErr) {
					assert.Equal(t, testCase.WantRowsAffected, got)
				}
			}
		})

	}
}
