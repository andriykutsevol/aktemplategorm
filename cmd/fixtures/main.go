package main

import (
	"fmt"
	"github.com/andriykusevol/aktemplategorm/tests/unit/database/fixtures"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Provide a version please")
		return
	}

	version := args[1]
	orm_fixture.LoadFixtures(version)
}
