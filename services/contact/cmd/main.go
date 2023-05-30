package main

import (
	"flag"
	"log"
	"os"
	"time"

	"advanced.microservices/pkg/jsonlog"
	"advanced.microservices/pkg/store/postgres"
	"advanced.microservices/services/contact/internal/domain"
	"advanced.microservices/services/contact/internal/repository"
	"advanced.microservices/services/contact/internal/useCase"
)

type config struct {
	port int
	env  string
	db   struct {
		dsn          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  string
	}
}

type repositories struct {
	contactRepository domain.ContactRepository
	groupRepository   domain.GroupRepository
}

type useCases struct {
	contactUseCase domain.ContactUseCase
	groupUseCase   domain.GroupUseCase
}

type application struct {
	config       config
	logger       *jsonlog.Logger
	repositories *repositories
	useCases     *useCases
	// mailer mailer.Mailer
	// wg sync.WaitGroup
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.StringVar(&cfg.db.dsn, "db-dsn", "", "PostgreSQL DSN")
	flag.IntVar(&cfg.db.maxOpenConns, "db-max-open-conns", 25, "PostgreSQL max open connections")
	flag.IntVar(&cfg.db.maxIdleConns, "db-max-idle-conns", 25, "PostgreSQL max idle connections")
	flag.StringVar(&cfg.db.maxIdleTime, "db-max-idle-time", "15m", "PostgreSQL max connection idle time")
	flag.Parse()

	logger := jsonlog.New(os.Stdout, jsonlog.LevelInfo)

	db, err := postgres.OpenDB(cfg.db.dsn, cfg.db.maxOpenConns, cfg.db.maxIdleConns, cfg.db.maxIdleTime)
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer db.Close()

	repositories := &repositories{
		contactRepository: repository.NewContactRepository(db),
		groupRepository:   repository.NewGroupRepository(db),
	}

	useCases := &useCases{
		contactUseCase: useCase.NewContactUsecase(repositories.contactRepository, time.Minute),
		groupUseCase:   useCase.NewGroupUsecase(repositories.groupRepository, time.Minute),
	}

	app := &application{
		config:       cfg,
		logger:       logger,
		repositories: repositories,
		useCases:     useCases,
		// mailer: mailer.New(cfg.smtp.host, cfg.smtp.port, cfg.smtp.username, cfg.smtp.password, cfg.smtp.sender),
	}

	log.Printf("Success, launched on %d", app.config.port)
}
