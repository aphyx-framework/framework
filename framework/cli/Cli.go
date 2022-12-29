package cli

type Command struct {
	Command      string
	Title        string
	Description  string
	Args         []CommandArgument
	ExampleUsage map[string]string
	Handler      func(c CommandArgumentValue)
}

type CommandArgument struct {
	Name        string
	Description string
	Required    bool
	Flag        bool
}

type CommandArgumentValue struct {
	Store map[string]string
}

type Registry struct {
	Commands map[string]Command
}

func MakeRegistry() Registry {
	return Registry{
		Commands: make(map[string]Command),
	}
}
