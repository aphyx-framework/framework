package cli

type Command struct {
	Command     string
	Title       string
	Description string
	Args        map[string]string
	Handler     func(arg ...string)
}

type Registry struct {
	Commands map[string]Command
}

func MakeRegistry() Registry {
	return Registry{
		Commands: make(map[string]Command),
	}
}
