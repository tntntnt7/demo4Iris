package utils

import (
	"context"
	"time"
)

func GetContext() (ctx context.Context) {
	ctx, _ = context.WithTimeout(context.Background(), time.Second * 10)
	return
}
