module example.com/mainFolder

go 1.23.1

replace example.com/examplepackage => ../examplepackage

replace example.com/custompackage => ../custompackage

replace example.com/arrays => ../arrays

require (
	example.com/arrays v0.0.0-00010101000000-000000000000
	example.com/custompackage v0.0.0-00010101000000-000000000000
	example.com/examplepackage v0.0.0-00010101000000-000000000000
)
