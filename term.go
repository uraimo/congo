package congo

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Term struct representing a console manager
type Term struct {
	cmds map[string]func([]string)
	done chan bool
	_prompt string
	unknownHandler func(string)
}


// NewTerm creates a new initialized Term struct
func NewTerm(prompt string) *Term {
	t:= &Term{make(map[string]func([]string), 1), make(chan bool),prompt,nil}
	return t
}

// AddCommand adds a function handler for this command
func (t *Term) AddCommand(name string, funct func([]string)) {
	t.cmds[name] = funct
}

// RemoveCommand removes a command/handler entry
func (t *Term) RemoveCommand(name string) {
	delete(t.cmds, name)
}

// AddUnknownHandler adds a function handler for unknown command
func (t *Term) AddUnknownHandler(funct func(string)) {
	t.unknownHandler = funct
}

// Prompt set the prompt
func (t *Term) Prompt(prompt string) {
	t._prompt=prompt
}

// Display the prompt and listen for new commands
// Add them with .AddCommand(command name, function implementing it)
func (t *Term) listen() {
	scanner := bufio.NewScanner(os.Stdin)
	running := true
	fmt.Print(t._prompt + " ")
	for running && scanner.Scan() {
		tokens := strings.Split(strings.TrimSpace(scanner.Text()), " ")
		command := tokens[0]
		if command == "quit" || command == "exit" {
			running = false
		} else {
			f, found := t.cmds[command]
			switch {
			case found && len(tokens) > 1:
				f(tokens[1:])
			case found:
				f([]string{})
			default:
				if t.unknownHandler!=nil && command!="" {
					t.unknownHandler(command)
				}
			}
			fmt.Print(t._prompt + " ")
		}
	}
	t.done <- true
}

// ListenUntilExit starts reading command until "exit" or "quit"
func (t *Term) ListenUntilExit() {
	go t.listen()
	<-t.done
}
