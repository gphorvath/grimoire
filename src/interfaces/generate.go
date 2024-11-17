package interfaces

type Generate interface {
	Generate(prompt string) (string, error)
	GenerateStream(prompt string) (<-chan string, <-chan error)
}
