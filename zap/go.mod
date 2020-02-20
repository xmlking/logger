module github.com/xmlking/logger/zap
go 1.13

replace github.com/xmlking/logger => ../

require (
	github.com/xmlking/logger v2.1.1-0.20200215215730-b3fc8be24e26
	github.com/pkg/errors v0.9.1
	go.uber.org/zap v1.13.0
)
