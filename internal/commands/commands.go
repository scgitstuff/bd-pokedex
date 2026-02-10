package commands

var commands map[string]CliCommand

// to allign command names in the help function
var helpNameWidth = 0

// I want singleton behavior
// could not do this at declaration, circular dependency in Help
func GetCommandsMap() map[string]CliCommand {
	if commands == nil {
		commands = map[string]CliCommand{
			"exit": {
				Description: "Exit the Pokedex",
				Callback:    commandExit,
			},
			"help": {
				Description: "Displays a help message",
				Callback:    commandHelp,
			},
			"map": {
				Description: "Displays 20 Pokemon world locations at a time",
				Callback:    commandMap,
			},
			"mapb": {
				Description: "Displays previous 20 Pokemon world locations at a time",
				Callback:    commandMapb,
			},
			"explore": {
				Description: "List Pokemon in a given area",
				Callback:    commandExplore,
			},
			"catch": {
				Description: "Catch a Pokemon",
				Callback:    commandCatch,
			},
			"inspect": {
				Description: "Catch a Pokemon",
				Callback:    comamndInspect,
			},
			"pokedex": {
				Description: "Catch a Pokemon",
				Callback:    commandPokedex,
			},
		}

		for k := range commands {
			if helpNameWidth < len(k) {
				helpNameWidth = len(k)
			}
		}
	}

	return commands
}
