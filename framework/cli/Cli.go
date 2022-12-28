package cli

type Command struct {
	Command      string
	Title        string
	Description  string
	Args         []CommandArgument
	ExampleUsage map[string]string
	Handler      func(arg ...string)
}

type CommandArgument struct {
	Name        string
	Description string
	Required    bool
	Flag        bool
}

type Registry struct {
	Commands map[string]Command
}

func MakeRegistry() Registry {
	return Registry{
		Commands: make(map[string]Command),
	}
}
