package main

import (
	"RyftFramework/database"
	"RyftFramework/framework"
	"flag"
	"os"
)

func main() {
	migratorFlag := flag.NewFlagSet("migrate", flag.ExitOnError)
	fresh := migratorFlag.Bool("fresh", false, "Drop all table defined in RegisterModel")
	seed := migratorFlag.Bool("seed", false, "Seed the database with data defined in the seeder")

	if os.Args[1] == "migrate" {
		if err := migratorFlag.Parse(os.Args[2:]); err == nil {
			database.RunMigrator(*fresh, *seed)
		}
	} else {
		framework.BootstrapFramework()
	}

}
