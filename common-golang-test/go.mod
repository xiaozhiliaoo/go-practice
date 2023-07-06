module common-golang-test

go 1.18

require (
	github.com/xiaozhiliaoo/common-golang/common-log v0.0.0
	github.com/xiaozhiliaoo/common-golang/common-mysql v0.0.0
)

require (
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/jmoiron/sqlx v1.3.5 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.24.0 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.2.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace (
	github.com/xiaozhiliaoo/common-golang/common-log => ../../common-golang/common-log
	github.com/xiaozhiliaoo/common-golang/common-mysql => ../../common-golang/common-mysql
)
