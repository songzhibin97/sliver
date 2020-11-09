package interfaces

type BaseServer interface {
	ListenAndServe() error
}