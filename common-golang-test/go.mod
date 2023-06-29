module common-golang-test

go 1.18

require github.com/xiaozhiliaoo/common-golang/common-mysql v0.0.0

require (
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/jmoiron/sqlx v1.3.5 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/xiaozhiliaoo/common-golang/common-mysql => ../../common-golang/common-mysql
