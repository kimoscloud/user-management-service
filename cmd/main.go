package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kimoscloud/user-management-service/internal/controller"
	logging2 "github.com/kimoscloud/user-management-service/internal/core/ports/logging"
	"github.com/kimoscloud/user-management-service/internal/core/usecase/organization"
	"github.com/kimoscloud/user-management-service/internal/core/usecase/user"
	"github.com/kimoscloud/user-management-service/internal/infrastructure/configuration"
	"github.com/kimoscloud/user-management-service/internal/infrastructure/db"
	"github.com/kimoscloud/user-management-service/internal/infrastructure/logging"
	organizationRepository "github.com/kimoscloud/user-management-service/internal/infrastructure/repository/postgres/organization"
	roleRepository "github.com/kimoscloud/user-management-service/internal/infrastructure/repository/postgres/organization/role"
	teamRepository "github.com/kimoscloud/user-management-service/internal/infrastructure/repository/postgres/organization/team"
	teamMemberRepository "github.com/kimoscloud/user-management-service/internal/infrastructure/repository/postgres/organization/team-member"
	userOrganizationRepository "github.com/kimoscloud/user-management-service/internal/infrastructure/repository/postgres/organization/user-organization"
	projectRepository "github.com/kimoscloud/user-management-service/internal/infrastructure/repository/postgres/project"
	teamProjectRepository "github.com/kimoscloud/user-management-service/internal/infrastructure/repository/postgres/project/team-project"
	userProjectRepository "github.com/kimoscloud/user-management-service/internal/infrastructure/repository/postgres/project/user-project"
	user2 "github.com/kimoscloud/user-management-service/internal/infrastructure/repository/postgres/user"
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
	userRepo := user2.NewUserRepository(conn)
	orgRepo := organizationRepository.NewOrganizationRepository(conn)
	userOrgRepo := userOrganizationRepository.NewUserOrganizationRepository(conn)
	roleRepo := roleRepository.NewRoleRepository(conn)
	teamRepo := teamRepository.NewTeamRepository(conn)
	teamMemberRepo := teamMemberRepository.NewTeamMemberRepository(conn)
	projectRepo := projectRepository.NewProjectRepository(conn)
	userProjectRepo := userProjectRepository.NewUserProjectRepository(conn)
	teamProjectRepo := teamProjectRepository.NewTeamProjectRepository(conn)

	initUserController(instance, userRepo, logger)
	initOrganizationController(
		instance,
		orgRepo,
		userOrgRepo,
		roleRepo,
		teamRepo,
		teamMemberRepo,
		logger,
	)
	initProjectController(
		instance,
		projectRepo,
		userProjectRepo,
		teamProjectRepo,
		roleRepo,
		userRepo,
		teamRepo,
		logger,
	)

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
func initProjectController(
	instance *gin.Engine,
	projectRepo *projectRepository.RepositoryPostgres,
	userProjectRepo *userProjectRepository.RepositoryPostgres,
	teamProjectRepo *teamProjectRepository.RepositoryPostgres,
	roleRepo *roleRepository.RepositoryPostgres,
	userRepo *user2.RepositoryPostgres,
	teamRepo *teamRepository.RepositoryPostgres,
	logger logging2.Logger,
) {

}

// TODO use interfaces here
func initOrganizationController(
	instance *gin.Engine,
	orgRepo *organizationRepository.RepositoryPostgres,
	userOrgRepo *userOrganizationRepository.RepositoryPostgres,
	roleRepo *roleRepository.RepositoryPostgres,
	teamRepo *teamRepository.RepositoryPostgres,
	teamMemberRepo *teamMemberRepository.RepositoryPostgres,
	logger logging2.Logger,
) {
	createOrganizationUseCase := organization.NewCreateOrganizationUseCase(
		orgRepo,
		userOrgRepo,
		roleRepo,
		logger,
	)
	getOrgByUserIdAndOrgIdUseCase := organization.NewGetOrganizationByOrgIdAndUserIdUseCase(
		orgRepo,
		logger,
	)
	getOrganizationsByUserIdUseCase := organization.NewGetOrganizationsByUserUseCase(
		orgRepo,
		logger,
	)
	createOrganizationUserUseCase := organization.NewCreateOrganizationMemberUseCase(
		orgRepo,
		userOrgRepo,
		roleRepo,
		logger,
	)
	organizationController := controller.NewOrganizationController(
		instance,
		logger,
		createOrganizationUseCase,
		getOrgByUserIdAndOrgIdUseCase,
		getOrganizationsByUserIdUseCase,
		createOrganizationUserUseCase,
	)
	organizationController.InitRouter()
}

// TODO use interfaces here`
func initUserController(
	instance *gin.Engine,
	userRepo *user2.RepositoryPostgres,
	logger logging2.Logger,
) {
	createUserUseCase := user.NewCreateUserUseCase(userRepo, logger)
	authenticateUserUseCase := user.NewAuthenticateUserUseCase(
		userRepo,
		logger,
	)
	getUserUseCase := user.NewGetUserUseCase(userRepo, logger)
	updateUserProfileUseCase := user.NewUpdateUserProfileUseCase(
		userRepo,
		logger,
	)

	changePasswordUseCase := user.NewChangePasswordUseCase(
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
