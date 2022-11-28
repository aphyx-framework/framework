package main

import (
	"RyftFramework/migration"
	"flag"
	"os"
)

func main() {
	migratorFlag := flag.NewFlagSet("migrate", flag.ExitOnError)
	fresh := migratorFlag.Bool("fresh", false, "Drop all table defined in RegisterModel")
	seed := migratorFlag.Bool("seed", false, "Seed the migration with data defined in the seeder")

	if len(os.Args) < 2 {
		// If no argument is passed, start the server
		BootstrapFramework()
	}

	switch os.Args[1] {
	case "migrate":
		err := migratorFlag.Parse(os.Args[2:])
		if err != nil {
			panic(err)
		}
		migration.RunMigrator(*fresh, *seed)
	default:
		panic("Unknown command")
	}

}
