module dbms

go 1.18

require databaselocal v0.0.0-00010101000000-000000000000

require github.com/go-sql-driver/mysql v1.6.0 // indirect

replace databaselocal => ./databaselocal
