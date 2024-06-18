package database

type Scanner interface {
	Scan(dest ...interface{}) error
}
