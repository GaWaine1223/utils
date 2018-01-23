package idl

import (
	"context"
)

type IController interface {
	GenIdl() interface{}
	Do(context.Context, interface{}) (interface{}, error)
}
