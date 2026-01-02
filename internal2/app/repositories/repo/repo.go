package repo

//Repository generique interface
type Repository[T any] interface {
	Save(t *T) (*T, error)
}
