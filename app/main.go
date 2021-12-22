package main

import (
	"context"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	CustomApiValidator "golang_hexagonal_architecture/adapters/api/utils/validator"
	apiV1 "golang_hexagonal_architecture/adapters/api/v1"
	userAdapterApi "golang_hexagonal_architecture/adapters/api/v1/user"
	userMongoRepo "golang_hexagonal_architecture/adapters/repositories/mongodb/user"
	"golang_hexagonal_architecture/adapters/repositories/mongodb/utils/mongodb"
	"golang_hexagonal_architecture/config"
	userCoreService "golang_hexagonal_architecture/core/user"

	"os"
	"os/signal"
	"time"
)

func initRepoMongoDb(db *mongo.Database) map[string]interface{} {
	userRepo, err := userMongoRepo.NewMongoDBRepository(db)
	if err != nil {
		panic(err)
	}
	return map[string]interface{}{
		"user": userRepo,
	}
}
func ConnectDatabase(dbConfig config.MongoDbConfig) *mongo.Database {
	uri := fmt.Sprintf("%s://", dbConfig.Driver)

	if config.GetConfigs().SystemEnv == config.PRODUCTION {
		uri = fmt.Sprintf("%s+srv://", dbConfig.Driver)
	}

	if dbConfig.Username != "" {
		uri = fmt.Sprintf("%s%v:%v@", uri, dbConfig.Username, dbConfig.Password)
	}

	if config.GetConfigs().SystemEnv == string(config.PRODUCTION) {
		uri = fmt.Sprintf("%s%v/%v?retryWrites=true&w=majority",
			uri,
			dbConfig.Host,
			dbConfig.Name,
		)
	} else {
		uri = fmt.Sprintf("%s%v:%v/?connect=direct",
			uri,
			dbConfig.Host,
			dbConfig.Port,
		)
	}
	db, err := mongodb.Connect(uri, dbConfig.Name)
	if err != nil {
		panic(err)
	}
	return db
}
func CustomValidator() *validator.Validate {
	myValidator := validator.New()
	return myValidator
}

func main() {
	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.JSONFormatter{})
	mongoDb := ConnectDatabase(config.GetConfigs().MasterMongoDb)
	repos := initRepoMongoDb(mongoDb)

	userService := userCoreService.New(repos["user"].(*userMongoRepo.MongoDb))
	userController := userAdapterApi.NewController(userService)

	e := echo.New()
	routerParam := apiV1.Router{
		Echo:           e,
		UserController: userController,
	}
	apiV1.RegisterRouter(routerParam)

	e.Validator = &CustomApiValidator.Validator{Validator: CustomValidator()}
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))
	go func() {
		address := fmt.Sprintf("0.0.0.0:%d", config.GetConfigs().Port)

		if err := e.Start(address); err != nil {
			log.Error("server failed to start")
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Error("server stopped immediately")
	}
}
