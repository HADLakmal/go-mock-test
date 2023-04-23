package domian

//go:generate mockgen -source=addition.go -package=mocks -destination=../mocks/addition.go
type Additional interface {
	Add(a, b int) int
}
