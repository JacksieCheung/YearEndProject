module github.com/JacksieCheung/YearEndProject

go 1.14

replace YearEndProject => ./

require (
	YearEndProject v0.0.0-00010101000000-000000000000
	github.com/gin-gonic/gin v1.6.3
	github.com/go-sql-driver/mysql v1.5.0
	github.com/jinzhu/gorm v1.9.16
	github.com/spf13/viper v1.7.1
	go.uber.org/zap v1.16.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)
