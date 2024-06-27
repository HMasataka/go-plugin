package shared

type Plugin interface {
	Greet(string) (string, error)
}
