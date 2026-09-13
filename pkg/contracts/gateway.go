package contracts

type Gateway interface {
	Handle(input any) (any, error)
}
