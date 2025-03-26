// Package condigo
// Code in this file is borrowed from [https://gist.github.com/thanhminhmr]
// Source: https://gist.github.com/thanhminhmr/05593d24fd3ee533e5e1970ea82e3952
package condigo

func If[T any](condition bool, ifTrue Supplier[T], ifFalse Supplier[T]) T {
	if condition {
		return ifTrue.Get()
	}
	return ifFalse.Get()
}

func Func[T any](supplier func() T) Supplier[T] {
	return FuncSupplier[T](supplier)
}

func Value[T any](value T) Supplier[T] {
	return ValueSupplier[T]{Value: value}
}

type Supplier[T any] interface {
	Get() T
}

type FuncSupplier[T any] func() T

func (s FuncSupplier[T]) Get() T {
	return s()
}

type ValueSupplier[T any] struct {
	Value T
}

func (s ValueSupplier[T]) Get() T {
	return s.Value
}
