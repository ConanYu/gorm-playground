module gorm.io/playground

go 1.16

require (
	github.com/mattn/go-sqlite3 v1.14.16 // indirect
	golang.org/x/crypto v0.5.0 // indirect
	golang.org/x/tools v0.5.0 // indirect
	gorm.io/datatypes v1.1.0 // indirect
	gorm.io/driver/mysql v1.4.5
	gorm.io/driver/postgres v1.4.6
	gorm.io/driver/sqlite v1.4.4
	gorm.io/driver/sqlserver v1.4.2
	gorm.io/gen v0.3.19
	gorm.io/gorm v1.24.2
	gorm.io/hints v1.1.1 // indirect
	gorm.io/plugin/dbresolver v1.4.0 // indirect
)

replace gorm.io/gorm => ./gorm
