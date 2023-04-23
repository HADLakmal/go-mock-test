package domian

//go:generate mockgen -source=divide.go -package=mocks -destination=../mocks/divide.go
type Divide interface {
	Divide(a, b int) int
}
