package transaction

import (
	"context"

	"gorm.io/gorm"
)

type transactionManager struct {
	db *gorm.DB
}

func NewTransactionManager(db *gorm.DB) TransactionManager {
	return &transactionManager{db: db}
}

func (tm *transactionManager) WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	return tm.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// Inject transaction ke dalam context
		txCtx := context.WithValue(ctx, "tx", tx)
		return fn(txCtx)
	})
}

func (tm *transactionManager) GetDB() *gorm.DB {
	return tm.db
}

// Helper function untuk mendapatkan DB dari context
func GetTxFromContext(ctx context.Context) *gorm.DB {
	if tx, ok := ctx.Value("tx").(*gorm.DB); ok {
		return tx
	}
	return nil
}
