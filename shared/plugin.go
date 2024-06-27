package shared

type Arg struct {
	P1 string
	P2 string
}

type Plugin interface {
	Greet(string, string) (string, error)
}
