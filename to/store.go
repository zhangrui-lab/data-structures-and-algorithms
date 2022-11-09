package to

type Store interface {
	Put(arg, reply *string) error
	Get(arg, reply *string) error
}
