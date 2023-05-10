package divide

import "dinero.go/types"

type Divider[T any] interface {
	Divide(amount T, factor T, calculator types.Calculator[T]) (T, error)
}