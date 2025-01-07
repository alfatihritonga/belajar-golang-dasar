package database

var connection bool

func init() {
	connection = true
}

func ConnectDatabase() bool {
	return connection
}
