package postgressql

import "github.com/jmoiron/sqlx"

// const (
// 	CustomersTable      = "customers"
// 	TransactionTable    = "transactions"
// 	AccountsTable       = "accounts"
// 	HistoryTable        = "history"
// 	ExpectedTransaction = "expected_transactions"
// 	ReportView          = "history_report"
// 	CustomerReportView  = "customer_report"
// )

type userBalanceStorage struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *userBalanceStorage {
	return &userBalanceStorage{
		db: db,
	}
}