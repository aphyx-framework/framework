package command

type Command struct {
	Name        string
	Description string
	Args        map[string]string
	Handler     func()
}

type Registry struct {
	Commands map[string]Command
}

func MakeRegistry() Registry {
	return Registry{
		Commands: make(map[string]Command),
	}
}
