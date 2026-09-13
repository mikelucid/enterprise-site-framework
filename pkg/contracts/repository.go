package contracts

type Repository[T any] interface {
	Create(T) error
	FindByID(id string) (T, error)
	Update(T) error
	Delete(id string) error
}
