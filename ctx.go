package utils

import (
	"context"
	"time"
)

func CtxWithTimeout(seconds int) context.Context {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*time.Duration(seconds))
	return ctx
}
