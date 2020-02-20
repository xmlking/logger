module github.com/xmlking/logger/zap

go 1.13

replace github.com/xmlking/logger v0.1.0 => ../

require (
	github.com/pkg/errors v0.9.1
	github.com/xmlking/logger v0.1.0
	go.uber.org/zap v1.13.0
)
