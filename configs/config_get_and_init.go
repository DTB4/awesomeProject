package configs

import (
	"awesomeProject/models"
	"log"
	"os"
	"strconv"
	"strings"
)

func NewConfig(path string) {

	fileName := path

	b, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal("fail to load config ", err)
	}
	s := string(b)
	lines := strings.Split(s, "\n")
	for i := range lines {
		if lines[i] == "" {
			break
		}
		params := strings.Split(lines[i], "=")
		err = os.Setenv(params[0], params[1])
		if err != nil {
			log.Fatal("fail to load config ", err)
		}
	}
}

func InitConfig() *models.Config {
	accessLifeTimeMinutes, err := strconv.Atoi(os.Getenv("ACCESS_LIFE_TIME_MINUTES"))
	if err != nil {
		log.Fatal("fail to load config", err)
	}
	refreshLifeTimeMinutes, err := strconv.Atoi(os.Getenv("REFRESH_LIFE_TIME_MINUTES"))
	if err != nil {
		log.Fatal("fail to load config", err)
	}
	parsingDelaySeconds, err := strconv.Atoi(os.Getenv("PARSER_DELAY_SECONDS"))
	if err != nil {
		log.Fatal("fail to load config", err)
	}

	config := models.Config{
		ServerPort: os.Getenv("HTTP_PORT"),
		AuthConfig: models.AuthConfig{
			AccessLifeTimeMinutes:  accessLifeTimeMinutes,
			RefreshLifeTimeMinutes: refreshLifeTimeMinutes,
			AccessSecretString:     os.Getenv("ACCESS_SECRET_STRING"),
			RefreshSecretString:    os.Getenv("REFRESH_SECRET_STRING"),
		},
		DBConfig: models.DBConfig{
			DatabaseUser:              os.Getenv("DATABASE_USER"),
			DatabasePassword:          os.Getenv("DATABASE_PASSWORD"),
			DatabaseProtocolIPAndPort: os.Getenv("DATABASE_PROTOCOL_IP_PORT"),
			DatabaseName:              os.Getenv("DATABASE_NAME"),
		},
		LogsPath: os.Getenv("LOGS_DIRECTORY_PATH"),
		ParserConfig: models.ParserConfig{
			URL:                 os.Getenv("URL_FOR_API_PARSER"),
			FormatString:        os.Getenv("FORMAT_STRING_FOR_API_URL"),
			ParsingDelaySeconds: parsingDelaySeconds,
		},
		CorsHandlerConfig: os.Getenv("CORS_ORIGIN_ADDRESS_&_PORT"),
	}
	return &config
}
