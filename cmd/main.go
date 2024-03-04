package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kimoscloud/user-management-service/internal/controller"
	logging2 "github.com/kimoscloud/user-management-service/internal/core/ports/logging"
	userRepository "github.com/kimoscloud/user-management-service/internal/core/ports/repository"
	"github.com/kimoscloud/user-management-service/internal/core/usecase"
	"github.com/kimoscloud/user-management-service/internal/infrastructure/configuration"
	"github.com/kimoscloud/user-management-service/internal/infrastructure/db"
	"github.com/kimoscloud/user-management-service/internal/infrastructure/logging"
	userRepositoryPostgres "github.com/kimoscloud/user-management-service/internal/infrastructure/repository/postgres"
	"github.com/kimoscloud/user-management-service/internal/infrastructure/server"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	if os.Getenv("ENV") == "dev" {
		err := godotenv.Load(".env")
		if err != nil {
			panic("Error loading .env file")
		}
	}
	// Create a new instance of the Gin router
	instance := gin.New()
	instance.Use(gin.Recovery())
	conn, err := db.NewConnection()
	if err != nil {
		log.Fatalf("failed to new database err=%s\n", err.Error())
	}
	logger, err := logging.NewLogger()
	if err != nil {
		log.Fatalf("failed to new logger err=%s\n", err.Error())
	}
	// Create the UserRepository
	userRepo := userRepositoryPostgres.NewUserRepository(conn)
	//projectRepo := projectRepositoryPostgres.NewProjectRepository(conn)
	//userProjectRepo := userProjectRepositoryPostgres.NewUserProjectRepository(conn)
	//teamProjectRepo := teamProjectRepositoryPostgres.NewTeamProjectRepository(conn)

	initUserController(instance, userRepo, logger)
	//initOrganizationController(
	//	instance,
	//	orgRepo,
	//	userOrgRepo,
	//	roleRepo,
	//	teamRepo,
	//	teamMemberRepo,
	//	userRepo,
	//	logger,
	//)
	//initProjectController(
	//	instance,
	//	projectRepo,
	//	userProjectRepo,
	//	teamProjectRepo,
	//	roleRepo,
	//	userRepo,
	//	logger,
	//)

	httpServer := server.NewHttpServer(
		instance,
		configuration.GetHttpServerConfig(),
	)
	httpServer.Start()
	defer httpServer.Stop()

	// Listen for OS signals to perform a graceful shutdown
	log.Println("listening signals...")
	c := make(chan os.Signal, 1)
	signal.Notify(
		c,
		os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGTERM,
	)
	<-c
	log.Println("graceful shutdown...")
}

// TODO use interfaces here
//func initProjectController(
//	instance *gin.Engine,
//	projectRepo *projectRepositoryPostgres.RepositoryPostgres,
//	userProjectRepo *userProjectRepositoryPostgres.RepositoryPostgres,
//	teamProjectRepo *teamProjectRepositoryPostgres.RepositoryPostgres,
//	roleRepo *roleRepositoryPostgres.RepositoryPostgres,
//	userRepo userRepository.Repository,
//	teamRepo *teamRepositoryPostgres.RepositoryPostgres,
//	logger logging2.Logger,
//) {
//
//}

func initUserController(
	instance *gin.Engine,
	userRepo userRepository.Repository,
	logger logging2.Logger,
) {
	createUserUseCase := usecase.NewCreateUserUseCase(userRepo, logger)
	authenticateUserUseCase := usecase.NewAuthenticateUserUseCase(
		userRepo,
		logger,
	)
	getUserUseCase := usecase.NewGetUserUseCase(userRepo, logger)
	updateUserProfileUseCase := usecase.NewUpdateUserProfileUseCase(
		userRepo,
		logger,
	)

	changePasswordUseCase := usecase.NewChangePasswordUseCase(
		userRepo,
		logger,
	)

	userController := controller.NewUserController(
		instance,
		logger,
		createUserUseCase,
		authenticateUserUseCase,
		getUserUseCase,
		updateUserProfileUseCase,
		changePasswordUseCase,
	)
	userController.InitRouter()
}
