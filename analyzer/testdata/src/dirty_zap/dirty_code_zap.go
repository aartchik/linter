package test

import "go.uber.org/zap"

func main() {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	zap.ReplaceGlobals(logger)
	password := "secret"
	token := "abc"

	logger.Info("Upper first letter") // want "log message should start with lowercase"

	logger.Info("server started!!!") // want "log message contains special symbols or emoji"

	logger.Info("сервер запущен") // want "log message should contain only English letters"

	logger.Warn("Warning message") // want "log message should start with lowercase"

	logger.Error("error occurred!!!") // want "log message contains special symbols or emoji"


	logger.Info(
		"request completed",
		zap.String("Field", "value"), // want "log field key should start with lowercase"
	)


	logger.Info(
		"request completed",
		zap.String("password", password), // want "log field key contains potential sensitive word" "log field key contains potential sensitive word"
	)

	logger.Info(
		"request completed",
		zap.String("ключ", "value"), // want "log field key should contain only English letters"
	)

	logger.Info(
		"request completed",
		zap.String("field!!!", "value"), // want "log field key contains special symbols or emoji"
	)


	logger.Info(
		"request completed",
		zap.String("field", password), // want "log field key contains potential sensitive word"
	)
	zap.L().Info(
		"request completed",
		zap.String("field", password), // want "log field key contains potential sensitive word"
	)
	logger.Info(
		"request completed",
		zap.String("field", password), // want "log field key contains potential sensitive word"
	)


	logger.Info(
		"request completed",
		zap.String("test", token), // want "log field key contains potential sensitive word"
	)

	logger.Info(
		"request completed",
		zap.String("api_key", token), // want "log field key contains potential sensitive word" "log field key contains potential sensitive word"
	)




	logger.Info(
		"request completed",
		zap.String("user", password), // want "log field key contains potential sensitive word"
	)

	logger.Info(
		"request completed",
		zap.Any("payload", token), // want "log field key contains potential sensitive word"
	)


	

	zap.L().Info("Upper global message") // want "log message should start with lowercase"

	zap.L().Info("global message!!!") // want "log message contains special symbols or emoji"

	zap.L().Info(
		"request completed",
		zap.String("Password", "ete"), // want "log field key should start with lowercase"
	)
	token = ""
	zap.L().Info(
		"request completed",
		zap.String("token", token), // want "log field key contains potential sensitive word" "log field key contains potential sensitive word"
	)
}