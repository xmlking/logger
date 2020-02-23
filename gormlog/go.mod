module github.com/xmlking/logger/gormlog

go 1.13

replace github.com/xmlking/logger v0.1.0 => ../

replace github.com/xmlking/logger/zerolog v0.1.0 => ../zerolog

require (
	github.com/xmlking/logger v0.1.0
	github.com/xmlking/logger/zerolog v0.1.0
)
