package connectdb

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

/*
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root93@"
	dbname   = "go-api"
)
*/

// Outra maneira, e criar uma constante passando os valores
// e depois usar o printf, mas por algum motivo deu erro
// entao eu comentei e fiz dessa outra maneira, que deu certo

func ConnectDb() (*sql.DB, error) {
	/*psqlInfo := fmt.Sprintf("host=%s port=%d user=%s"+
	"password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)*/

	// LEMBRANDO QUE MEU HOST-> VOU TER QUE MUDAR PARA O NOME AO QUAL
	// EU COLOQUEI LA NO MEU DOCKER-COMPOSE, ISSO EM CASO DEU TA USANDO UM DOCKERFILE

	connStr := "user=postgres dbname=go-api password=root93@ port=5432 host=go_db sslmode=disable"

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to ")

	return db, nil
}
