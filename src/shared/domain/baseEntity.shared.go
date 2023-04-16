package shared_domain

type BaseEntities[T any] struct {
	Id    string
	Props T
}
