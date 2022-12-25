package command

import "errors"

func (container *Registry) AddCommand(callable string, command Command) {
	if _, ok := container.Commands[callable]; ok {
		panic("Command " + callable + " already exists")
	}

	container.Commands[callable] = command
}

func (container *Registry) GetCommand(callable string) (Command, error) {
	if _, ok := container.Commands[callable]; ok {
		return container.Commands[callable], nil
	}
	return Command{}, errors.New("command" + callable + "not found")
}

func (container *Registry) GetCommands() map[string]Command {
	return container.Commands
}
