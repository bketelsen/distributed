package stringsvc

// Args is the input struct
import "strings"

type Args struct {
	Name string
}

// Result is the output struct
type Result struct {
	Name string
}

// Upper is the type that implements the RPC methods that are exported
type Upper string

func (u *Upper) Uppercase(args *Args, result *Result) error {
	result.Name = strings.ToUpper(args.Name)
	return nil
}
