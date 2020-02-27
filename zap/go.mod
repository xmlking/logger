module github.com/xmlking/logger/zap

go 1.14

replace github.com/xmlking/logger => ../

require (
	github.com/pkg/errors v0.9.1
	github.com/xmlking/logger v0.1.2
	go.uber.org/zap v1.13.0
)
