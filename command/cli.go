package command

var CommandBuilders []CommandBuilder

type CommandBuilder func() ([]Command, error)

type Command struct {
	Command          []string
	ShortDescription string
	LongDescription  string
	Executor         interface{}
}

type ByCommandLength []Command

func (a ByCommandLength) Len() int           { return len(a) }
func (a ByCommandLength) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByCommandLength) Less(i, j int) bool { return len(a[i].Command) < len(a[j].Command) }

type Execute interface {
	Execute([]string) error
}
