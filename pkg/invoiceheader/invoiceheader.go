package invoiceheader

import "time"

//Modelo de invoiceheader
type Model struct {
	ID        uint
	Name      string
	CreatedAt time.Time
	UpdateAt  time.Time
}
