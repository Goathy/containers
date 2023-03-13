package list

type list struct{}

func New[V any]() *list {
	return &list{}
}
