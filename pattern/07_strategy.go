package pattern

import "fmt"

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

// Общий интерфейс для подключения к БД
type IDBconnection interface {
	Connect()
}

// Структура подключения к MySql
type MySqlConnection struct {
	connectionString string
}

// Реализация метода по подключению к БД MySql
func (con MySqlConnection) Connect() {
	fmt.Println("Connect sql: " + con.connectionString)
}

// Cтруктура подключения к Oracle
type OracleConnection struct {
	connectionString string
}

// Реализация метода по подключению к БД Oracle
func (con OracleConnection) Connect() {
	fmt.Println("Connect sql: " + con.connectionString)
}

// Структура с полем конкретной БД
type DBConnection struct {
	db IDBconnection
}

// Функция по подключению к конкретной БД
func (con DBConnection) DBConnect() {
	con.db.Connect()
}

func main() {
	mySqlConnection := MySqlConnection{"Connection to MySql"}
	connectionSql := DBConnection{mySqlConnection}
	connectionSql.DBConnect()

	oracleConnection := OracleConnection{"Connection to oracle"}
	connectionOracle := DBConnection{oracleConnection}
	connectionOracle.DBConnect()
}
