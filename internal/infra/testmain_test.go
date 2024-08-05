package infra_test

import (
	"context"
	"os"
	"testing"

	"spocon-backend/testutil"
)

var testDB *testutil.TestDB

func TestMain(m *testing.M) {

	ctx := context.Background()
	testDB = testutil.SetupDB(ctx, m)
	defer testDB.TearDown()

	code := m.Run()

	os.Exit(code)
}
