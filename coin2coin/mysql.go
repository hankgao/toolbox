package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbPassword = "Temp#1234"
	dbUser     = "robot"
	dbAddress  = "120.55.60.140:3306"
	dbName     = "coinsell"
)

//
type Account struct {
	User     string
	Password string
}

// Table is a table
type Table interface {
	Load() error //

}

// CoinPairTable represents coin2coin table
type CoinPairTable struct {
	Pairs []CoinPair
}

// OrderTable represents order table
type OrderTable struct {
}

// Database represents the databas we are talking to
type Database struct {
	CurrentAccount Account
	Name           string
	Tables         []Table

	Handle        *sql.DB
	stmtCoinPairs *sql.Stmt
	stmtOrders    *sql.Stmt
}

var db Database

// Init does some preparation stuff to use the database
func (db *Database) Init() {

	db.CurrentAccount = Account{
		User:     dbUser,
		Password: dbPassword,
	}

	conn := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		db.CurrentAccount.User, db.CurrentAccount.Password, dbAddress, dbName)

	h, err := sql.Open("mysql", conn)
	if err != nil {
		panic(err)
	}

	db.Handle = h

	// prepare statements
	db.stmtCoinPairs, err = db.Handle.Prepare(
		"SELECT from_coin, to_coin, exchange_rate, start_date, end_date FROM coin2coin WHERE ? >= start_date && ? <= end_date")
	if err != nil {
		panic(err)
	}

	db.stmtOrders, err = db.Handle.Prepare(
		`INSERT INTO c2corders (coin_from, coin_to, depositing_c_address, receiving_c_address, i_name, i_mobile, i_m_address, date_placed) 
		VALUES(?, ?, ?, ?, ?, ?, ?, ?)`)

	if err != nil {
		panic(err)
	}

}

// Close closes all prepred statements
func (db *Database) Close() {
	db.stmtCoinPairs.Close()
}

// LoadCoinPairs read coin2coin table and load data into memory
func LoadCoinPairs() []CoinPair {
	today := time.Now().Format("2006-01-02 15:04:05")
	rows, err := db.stmtCoinPairs.Query(today, today)
	if err != nil {
		panic(err)
	}

	res := []CoinPair{}

	pair := CoinPair{}
	for rows.Next() {
		// from_coin, to_coin, exchange_rate, start_date, end_date
		err := rows.Scan(&pair.FromCoin, &pair.ToCoin, &pair.ExhangeRate, &pair.ValidFrom, &pair.ValidUntil)
		if err != nil {
			panic(err)
		}

		res = append(res, pair)
	}

	return res
}
