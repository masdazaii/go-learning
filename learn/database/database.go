package database

var Connection string

func init() {
	Connection = "MySql"
}

func TestConnection() string {
	return Connection
}