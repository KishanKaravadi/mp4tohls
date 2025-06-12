package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/KishanKaravadi/mp4tohls/internal/data"
	"github.com/KishanKaravadi/mp4tohls/internal/process"
)

type config struct {
	port       int
	storageDir string
	db         struct {
		dsn string
	}
}

type application struct {
	cfg       config
	logger    *Logger
	models    data.Models
	processor *process.Processor
}

func main() {
	var cfg config
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.storageDir, "storage", "./vids", "Directory to store application user files")
	//flag.StringVar(&cfg.db.dsn, "db-dsn", "", "PostgreSQL data source name")

	logger := NewLogger(os.Stderr, os.Stdout)

	err := validateStorageDir(cfg.storageDir)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Info(fmt.Sprintf("storage directory '%s' is valid", cfg.storageDir))

	// db, err := openDB(cfg)
	// if err != nil {
	// 	logger.Fatal(err)
	// }
	// defer db.Close()
	// logger.Info("database connection pool established")

	app := &application{
		cfg:       cfg,
		logger:    logger,
		models:    data.NewModels(nil),
		processor: process.NewProcessor(cfg.storageDir),
	}

	app.logger.Info(fmt.Sprintf("Starting server on port %d", app.cfg.port))
	err = app.serve()
	if err != nil {
		app.logger.Fatal(err)
	}

}

func validateStorageDir(dir string) error {
	info, err := os.Stat(dir)

	if os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0766)
		if err != nil {
			return fmt.Errorf("failed to create storage director '%s': %w", dir, err)
		}
	} else if err != nil {
		return fmt.Errorf("failed to access storage directory '%s': %w", dir, err)
	} else if !info.IsDir() {
		return fmt.Errorf("path '%s' exists but is not a directory", dir)
	}

	return nil
}

func (app *application) serve() error {
	err := http.ListenAndServe(fmt.Sprintf(":%d", app.cfg.port), app.routes())
	if err != nil {
		return err
	}
	return nil
}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
