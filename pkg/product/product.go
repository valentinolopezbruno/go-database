package product

import "time"

//Modelo de producto
type Model struct {
	ID           uint
	Name         string
	Observations string
	Price        string
	CreatedAt    time.Time
	UpdateAt     time.Time
}

// Models es un Slice de Model
type Models []Model

type Storage interface {
	Migrate() error
	//Create(*Model) error
	//Update(*Model) error
	//GetAll() (Models, error)
	//GetByID(uint) (*Model, error)
	//Delete(uint) error
}
