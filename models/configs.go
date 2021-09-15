package models

type Config struct {
	ServerPort   string
	AuthConfig   AuthConfig
	DBConfig     DBConfig
	LogsPath     string
	ParserConfig ParserConfig
}

type AuthConfig struct {
	AccessLifeTimeMinutes  int
	RefreshLifeTimeMinutes int
	AccessSecretString     string
	RefreshSecretString    string
}

type DBConfig struct {
	DatabaseUser              string
	DatabasePassword          string
	DatabaseProtocolIPAndPort string
	DatabaseName              string
}

type ParserConfig struct {
	URL                 string
	FormatString        string
	ParsingDelaySeconds int
}
