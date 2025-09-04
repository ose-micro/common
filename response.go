package common

type Response[T any] struct {
	Message string
	Records []T
	Record  T
}
