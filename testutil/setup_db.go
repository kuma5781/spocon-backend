package testutil

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type TestDB struct {
	DB       *sql.DB
	TearDown func()
}

func SetupDB(ctx context.Context, m *testing.M) *TestDB {
	const mysqlImage = "mysql:8.0"
	const databaseName = "spocon"
	const port = "3306"

	req := testcontainers.ContainerRequest{
		Image:        mysqlImage,
		ExposedPorts: []string{port},
		Env: map[string]string{
			"MYSQL_DATABASE":             databaseName,
			"MYSQL_ALLOW_EMPTY_PASSWORD": "yes",
		},
		WaitingFor: wait.ForListeningPort(port),
	}

	mysqlContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not start mysql: %s", err)
		os.Exit(1)
	}

	host, err := mysqlContainer.Host(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not get mysql host: %s", err)
		os.Exit(1)
	}
	mappedPort, err := mysqlContainer.MappedPort(ctx, port)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not get mysql port: %s", err)
		os.Exit(1)
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", "root", "", host, mappedPort.Port(), databaseName)
	tdb, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not open db: %s", err)
		os.Exit(1)
	}

	if err := tdb.Ping(); err != nil {
		fmt.Fprintf(os.Stderr, "Could not ping db: %s", err)
		os.Exit(1)
	}

	if err := migrateUp(dsn); err != nil {
		fmt.Fprintf(os.Stderr, "Could not migrate up: %s", err)
		os.Exit(1)
	}

	teardown := func() {
		if err := mysqlContainer.Terminate(ctx); err != nil {
			fmt.Fprintf(os.Stderr, "Could not stop mysql: %s", err)
			os.Exit(1)
		}
	}

	return &TestDB{DB: tdb, TearDown: teardown}
}

func migrateUp(dsn string) error {
	m, err := migrate.New(
		"file://../../database/migrations",
		"mysql://"+dsn,
	)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}
