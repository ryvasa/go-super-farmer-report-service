package transaction

import (
	"context"

	"gorm.io/gorm"
)


type TransactionManager interface {
    WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error
    GetDB() *gorm.DB
}
