package orm_fixture

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-testfixtures/testfixtures/v3"
	"os"
	"path"
	"runtime"
)

// LoadFixtures connects to the test database, and loads pre-defined fixtures
// into it.
func LoadFixtures(folder string) {
	// It reads from the .env file included from the Makefile

	// Connecto to the test database
	dbDsn := os.Getenv("DATABASE_DSN")
	db, err := sql.Open("mysql", dbDsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Get the template parent directory
	_, filename, _, _ := runtime.Caller(0)
	//dir := path.Join(path.Dir(filename), "..")
	dir := path.Dir(filename)
	templateDir := fmt.Sprintf("%s/%s", dir, folder)

	//fmt.Println("dbDsn: ", dbDsn)
	//fmt.Println("templateDir: ", templateDir)

	// Initialize fixtures
	fixtures, _ := testfixtures.New(
		// In order to prevent you from accidentally wiping the wrong database,
		// this package will refuse to load fixtures if the database name (or database filename for SQLite)
		// doesn't contains "test". If you want to disable this check, use:
		testfixtures.DangerousSkipTestDatabaseCheck(),

		// If you want to disable cleanup, you can also do like below.
		// This is usually not recommended, and should be used mostly for debugging.
		// It skips the automatic cleanup of the database tables before loading the fixtures.
		// Use it if you need to add additional fixture data to an already loaded dataset for specific test scenarios.
		//testfixtures.DangerousSkipCleanupFixtureTables(),

		// For PostgreSQL and MySQL/MariaDB, this package also resets all sequences to a high number
		// to prevent duplicated primary keys while running the tests.
		// The default is 10000, but you can change that with:
		// testfixtures.ResetSequencesTo(10000),

		// Or, if you want to skip the reset of sequences entirely:
		// testfixtures.SkipResetSequences(),

		testfixtures.Database(db),
		testfixtures.Dialect("mysql"),
		testfixtures.Directory(templateDir),
	)

	// // Initialize fixtures
	// fixtures, _ := testfixtures.New(
	// 	testfixtures.DangerousSkipTestDatabaseCheck(),
	// 	testfixtures.DangerousSkipCleanupFixtureTables(),
	// 	testfixtures.Database(db),
	// 	testfixtures.Dialect("mysql"),
	// 	testfixtures.Directory(templateDir),
	// )

	// Load fixtures
	err = fixtures.Load()
	if err != nil {
		fmt.Println(err)
	}
}
