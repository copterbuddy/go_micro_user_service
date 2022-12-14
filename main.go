package main

import (
	"context"
	"fmt"
	"main/handler"
	"main/intercepter"
	"main/repository"
	"main/service"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"gorm.io/driver/postgres"

	// "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	initConfig()

	r := SetupApi()

	r.Run(fmt.Sprintf(":%v", viper.GetString("app.port")))
}

func SetupApi() *gin.Engine {

	db = initUserDatabasePostgreSql()

	userRepo := repository.NewUserRepositoryDB(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	intercepter := intercepter.NewInterceptor()

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	SetCore(r)

	handler.SetupRouter(r, userHandler, intercepter)
	return r
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	// logs.Info(fmt.Sprintf("this is my viper %v", viper.GetString("app.port")))
}

var db *gorm.DB

type SqlLogger struct {
	logger.Interface
}

func (l SqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, _ := fc()
	fmt.Printf("%v\n==========================================\n", sql)
}

func SetCore(r *gin.Engine) {
	//set core
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000/*"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowWildcard:    true,
		// AllowOriginFunc: func(origin string) bool {
		// 	return origin == "https://github.com"
		// },
		MaxAge: 12 * time.Hour,
	}))
}

func initUserDatabasePostgreSql() *gorm.DB {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Bangkok",
		viper.GetString("db.host"),
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.database"),
		viper.GetInt("db.port"))

	dial := postgres.Open(dsn)

	var err error
	db, err = gorm.Open(dial, &gorm.Config{
		Logger: &SqlLogger{},
		DryRun: false, //????????????????????????????????? db ????????? true
	})
	if err != nil {
		panic(err)
	}
	//set timeout
	// db.SetConnMaxLifetime(3 * time.Minute)
	// db.SetMaxOpenConns(10)
	// db.SetMaxIdleConns(10)

	return db
}

// func initUserDatabaseMySql() *gorm.DB {
// 	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true",
// 		viper.GetString("db.username"),
// 		viper.GetString("db.password"),
// 		viper.GetString("db.host"),
// 		viper.GetInt("db.port"),
// 		viper.GetString("db.database"),
// 	)

// 	dial := mysql.Open(dsn)

// 	var err error
// 	db, err = gorm.Open(dial, &gorm.Config{
// 		Logger: &SqlLogger{},
// 		DryRun: false, //????????????????????????????????? db ????????? true
// 	})
// 	if err != nil {
// 		panic(err)
// 	}
// 	//set timeout
// 	// db.SetConnMaxLifetime(3 * time.Minute)
// 	// db.SetMaxOpenConns(10)
// 	// db.SetMaxIdleConns(10)

// 	return db
// }
