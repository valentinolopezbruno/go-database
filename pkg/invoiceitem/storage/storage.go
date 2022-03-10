package Storage

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
)

var (
	db   *sql.DB
	once sync.Once
)

func newPostgresDB() {
	once.Do(func() { //Singleton
		var err error
		db, err = sql.Open("postgres", "postgres://edteam:edteam@localhost:7530/godb?sslmode=disable")
		if err != nil {
			log.Fatalf("no pudimos abrir la bd", err) //AVISAMOS DE LA CONEXION FALLIDA
		}

		if err = db.Ping(); err != nil {
			log.Fatalf("no pudimos hace rping", err) //AVISAMOS DE LA CONEXION CORRECTA
		}

		fmt.Println("CONECTADO A POSTFRES")
	})
}

func Pool() *sql.DB {
	return db
}
