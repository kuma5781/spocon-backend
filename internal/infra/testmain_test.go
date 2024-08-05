package infra_test

import (
	"context"
	"os"
	"testing"

	"spocon-backend/test"
)

var testDB *test.TestDB

func TestMain(m *testing.M) {

	ctx := context.Background()
	testDB = test.SetupDB(ctx, m)
	defer testDB.TearDown()

	code := m.Run()

	os.Exit(code)
}
