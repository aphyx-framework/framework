package commands

import (
	"fmt"
	"github.com/TwiN/go-color"
	"github.com/aphyx-framework/framework/framework/cli"
	"github.com/gosuri/uitable"
	"go.uber.org/fx"
	"os"
	"strings"
)

var FrameworkCommands = fx.Options(
	fx.Invoke(TestCommand),
	fx.Invoke(HelpCommand),
)

func RunCommand(registry cli.Registry) {
	command := os.Args[1]
	args := os.Args[2:]

	commandResult, err := registry.GetCommand(command)

	if err != nil {
		panic(err)
	}

	for _, arg := range args {
		if arg == "--help" {
			printHelp(commandResult)
		}

	}

	commandResult.Handler(args...)
	os.Exit(0)
}

func printHelp(command cli.Command) {
	table := uitable.New()
	table.MaxColWidth = 80
	table.Wrap = true // wrap columns

	joinedArgs := "[No arguments needed]"

	if len(command.Args) > 0 {
		args := make([]string, 0, len(command.Args))
		for k := range command.Args {
			args = append(args, k)
		}

		// Join the slice with commas
		joinedArgs = strings.Join(args, ", ")
	}

	table.AddRow(color.GreenBackground+color.Black+command.Command, color.Reset+"   "+command.Title)
	table.AddRow("Description:", "   "+command.Description)
	table.AddRow("Arguments:", "   "+joinedArgs)
	table.AddRow("Usages:", "   "+"Type in `"+command.Command+" --help` for usage example")
	table.AddRow("") // blank

	fmt.Println(table)
}

//func RunCliApplication(enableNopLogger bool) {
//	nop := fx.Options()
//
//	if enableNopLogger == false {
//		nop = fx.Options(fx.NopLogger)
//	}
//
//	fx.New(
//		nop,
//		fx.Provide(configuration.NewConfiguration),
//		fx.Provide(logging.NewLogger),
//		fx.Provide(database.NewDbConnection),
//		fx.Invoke(runCliCommand),
//	).Run()
//}
//
//func runCliCommand(logger logging.ApplicationLogger, config configuration.Configuration, db *gorm.DB) {
//	if os.Args[1] == "repl" {
//		repl(logger, config, db)
//		os.Exit(0)
//	} else {
//		err := parseCommand(os.Args[1], os.Args[2:], logger, config, db)
//		if err != nil {
//			logger.ErrorLogger.Fatalln(err)
//		}
//		os.Exit(0)
//	}
//}
//
//func repl(logger logging.ApplicationLogger, config configuration.Configuration, db *gorm.DB) {
//	logger.InfoLogger.Println("RyFT Framework REPL")
//	logger.InfoLogger.Println("Type 'exit' to exit the REPL")
//
//	for {
//		fmt.Print("[Ryft] >> ")
//		scanner := bufio.NewScanner(os.Stdin)
//		input := ""
//
//		for scanner.Scan() {
//			input = scanner.Text()
//			break
//		}
//
//		if input == "exit" {
//			break
//		}
//
//		args := strings.Split(input, " ")
//		err := parseCommand(args[0], args[1:], logger, config, db)
//
//		if err != nil {
//			logger.ErrorLogger.Println(err)
//		}
//	}
//
//	logger.InfoLogger.Println("Exiting REPL")
//	os.Exit(0)
//
//}
//func parseCommand(
//	command string,
//	args []string,
//	logger logging.ApplicationLogger,
//	config configuration.Configuration,
//	db *gorm.DB,
//) error {
//	switch command {
//	case "migrate":
//		fresh := args[0] == "fresh"
//		seed := args[1] == "seed"
//		migration.RunMigrator(fresh, seed, logger, db)
//	case "make":
//		generator.Generator(args[0], args[1:], logger)
//	case "createkey":
//		var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
//		b := make([]rune, 32)
//		for i := range b {
//			b[i] = letterRunes[rand.Intn(len(letterRunes))]
//		}
//		logger.InfoLogger.Println("Generated key: ", string(b))
//		logger.InfoLogger.Println("Please add this key to your config.toml file in the 'security' > key section")
//	default:
//		return errors.New("invalid cli")
//	}
//	return nil
//}
