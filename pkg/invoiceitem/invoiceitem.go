package invoiceitem

import "time"

// Modelo de invoiceitem
type Model struct {
	ID              uint
	invoiceheaderID uint
	ProductID       uint
	CreatedAt       time.Time
	UpdateAt        time.Time
}
