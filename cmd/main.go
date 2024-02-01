package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kimoscloud/user-management-service/internal/controller"
	logging2 "github.com/kimoscloud/user-management-service/internal/core/ports/logging"
	organizationRepository "github.com/kimoscloud/user-management-service/internal/core/ports/repository/organization"
	userOrganizationRepository "github.com/kimoscloud/user-management-service/internal/core/ports/repository/organization/user-organization"
	userRepository "github.com/kimoscloud/user-management-service/internal/core/ports/repository/user"
	"github.com/kimoscloud/user-management-service/internal/core/usecase/organization"
	"github.com/kimoscloud/user-management-service/internal/core/usecase/user"
	"github.com/kimoscloud/user-management-service/internal/infrastructure/configuration"
	"github.com/kimoscloud/user-management-service/internal/infrastructure/db"
	"github.com/kimoscloud/user-management-service/internal/infrastructure/logging"
	organizationRepositoryPostgres "github.com/kimoscloud/user-management-service/internal/infrastructure/repository/postgres/organization"
	roleRepositoryPostgres "github.com/kimoscloud/user-management-service/internal/infrastructure/repository/postgres/organization/role"
	teamRepositoryPostgres "github.com/kimoscloud/user-management-service/internal/infrastructure/repository/postgres/organization/team"
	teamMemberRepositoryPostgres "github.com/kimoscloud/user-management-service/internal/infrastructure/repository/postgres/organization/team-member"
	userOrganizationRepositoryPostgres "github.com/kimoscloud/user-management-service/internal/infrastructure/repository/postgres/organization/user-organization"
	projectRepositoryPostgres "github.com/kimoscloud/user-management-service/internal/infrastructure/repository/postgres/project"
	teamProjectRepositoryPostgres "github.com/kimoscloud/user-management-service/internal/infrastructure/repository/postgres/project/team-project"
	userProjectRepositoryPostgres "github.com/kimoscloud/user-management-service/internal/infrastructure/repository/postgres/project/user-project"
	userRepositoryPostgres "github.com/kimoscloud/user-management-service/internal/infrastructure/repository/postgres/user"
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
	orgRepo := organizationRepositoryPostgres.NewOrganizationRepository(conn)
	userOrgRepo := userOrganizationRepositoryPostgres.NewUserOrganizationRepository(conn)
	roleRepo := roleRepositoryPostgres.NewRoleRepository(conn)
	teamRepo := teamRepositoryPostgres.NewTeamRepository(conn)
	teamMemberRepo := teamMemberRepositoryPostgres.NewTeamMemberRepository(conn)
	projectRepo := projectRepositoryPostgres.NewProjectRepository(conn)
	userProjectRepo := userProjectRepositoryPostgres.NewUserProjectRepository(conn)
	teamProjectRepo := teamProjectRepositoryPostgres.NewTeamProjectRepository(conn)

	initUserController(instance, userRepo, logger)
	initOrganizationController(
		instance,
		orgRepo,
		userOrgRepo,
		roleRepo,
		teamRepo,
		teamMemberRepo,
		userRepo,
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
	projectRepo *projectRepositoryPostgres.RepositoryPostgres,
	userProjectRepo *userProjectRepositoryPostgres.RepositoryPostgres,
	teamProjectRepo *teamProjectRepositoryPostgres.RepositoryPostgres,
	roleRepo *roleRepositoryPostgres.RepositoryPostgres,
	userRepo userRepository.Repository,
	teamRepo *teamRepositoryPostgres.RepositoryPostgres,
	logger logging2.Logger,
) {

}

// TODO use interfaces here
func initOrganizationController(
	instance *gin.Engine,
	orgRepo organizationRepository.Repository,
	userOrgRepo userOrganizationRepository.Repository,
	roleRepo *roleRepositoryPostgres.RepositoryPostgres,
	teamRepo *teamRepositoryPostgres.RepositoryPostgres,
	teamMemberRepo *teamMemberRepositoryPostgres.RepositoryPostgres,
	userRepo userRepository.Repository,
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

	checkIfUserHasEnoughPermissionsUseCase := organization.NewCheckUserHasPermissionsToMakeAction(
		userOrgRepo,
		logger)

	createOrganizationUserUseCase := organization.NewCreateOrganizationMemberUseCase(
		orgRepo,
		userOrgRepo,
		roleRepo,
		userRepo,
		checkIfUserHasEnoughPermissionsUseCase,
		logger,
	)
	removeOrganizationUserUseCase := organization.NewRemoveOrganizationMemberUseCase(
		orgRepo,
		userOrgRepo,
		logger,
	)

	createTeamUsecase := organization.NewCreateTeamUseCase(
		userOrgRepo,
		teamRepo,
		teamMemberRepo,
		checkIfUserHasEnoughPermissionsUseCase,
		logger)

	addMemberToTeamUseCase := organization.NewAddTeamMembersUseCase(
		userOrgRepo,
		teamRepo,
		teamMemberRepo,
		checkIfUserHasEnoughPermissionsUseCase,
		logger)

	organizationController := controller.NewOrganizationController(
		instance,
		logger,
		createOrganizationUseCase,
		getOrgByUserIdAndOrgIdUseCase,
		getOrganizationsByUserIdUseCase,
		createOrganizationUserUseCase,
		removeOrganizationUserUseCase,
		createTeamUsecase,
		addMemberToTeamUseCase,
	)
	organizationController.InitRouter()
}

// TODO use interfaces here`
func initUserController(
	instance *gin.Engine,
	userRepo userRepository.Repository,
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
