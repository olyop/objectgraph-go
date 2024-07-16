package database

type Scanner interface {
	Scan(dest ...any) error
}
