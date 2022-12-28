package cli

import "errors"

func (container *Registry) AddCommand(callable Command) {
	if _, ok := container.Commands[callable.Command]; ok {
		panic("Command " + callable.Command + " already exists")
	}

	container.Commands[callable.Command] = callable
}

func (container *Registry) GetCommand(callable string) (Command, error) {
	if _, ok := container.Commands[callable]; ok {
		return container.Commands[callable], nil
	}
	return Command{}, errors.New("Command " + callable + " not found")
}

func (container *Registry) GetCommands() map[string]Command {
	return container.Commands
}

func (command Command) GetArgument(name string) (CommandArgument, error) {
	for _, arg := range command.Args {
		if arg.Name == name {
			return arg, nil
		}
	}
	return CommandArgument{}, errors.New("Argument " + name + " not found")
}
