package Storage

import (
	"database/sql"
	"fmt"
)

const ( //CONSTANTE PARA CARGAR DATOS EN LA BD
	psqlCreateProduct = `CREATE TABBLE IF NOT EXIST products (
		id serial NOT NULL,
		name VERCHAR(25) NOT NULL,
		observations VARCHAR(100),
		price INT NOT NULL,
		create_at TIMESTAMP NOT NULL DEFAULT now(),
		update_at TIMESTAMP,
		CONSTRAINT products_id_pk PRIMARY KEY (id) 
)`
)

// PsqlProduct  es para trabajar con postgras - product
type PsqlProduct struct {
	db *sql.DB
}

// NewPsqlProduct retorna un nuevo puntero de PsqlProduct
func NewPsqlProduct(db *sql.DB) *PsqlProduct {
	return &PsqlProduct{db}
}

// Migrate implementa la interfar product.Storage
func (p *PsqlProduct) Migrate() error {
	stmt, err := p.db.Prepare(psqlCreateProduct)
	if err != nil { //controlamos el error
		return err
	}
	defer stmt.Close() //cerramos stmt

	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	fmt.Println("MIGRACION DE PRODUCTO EJECUTADA CORRECTAMENTE")
	return nil
}
