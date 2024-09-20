module github.com/nptri/mainfolder

go 1.23.1

replace github.com/nptri/GoLearn/custompackage => ../custompackage

replace github.com/nptri/GoLearn/example => ../example

require (
	github.com/nptri/GoLearn/custompackage v0.0.0
	github.com/nptri/GoLearn/example v0.0.0
)
