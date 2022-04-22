package history

type GoshCommand struct {
	text   string
	result int // This should be some sort of more unique error type..
}

func NewCommand(text string, result int) GoshCommand {
	return GoshCommand{text, result}
}

type GoshHistory struct {
	commands []GoshCommand
}

func NewHistory() GoshHistory {
	commands := make([]GoshCommand, 0)
	return GoshHistory{commands}
}

func (g *GoshHistory) AddToHistory(c GoshCommand) {
	g.commands = append(g.commands, c)
}
