module main

go 1.15

replace crawler/pkg/spider => ../../pkg/spider

require (
	crawler/pkg/spider v0.0.0-00010101000000-000000000000
	golang.org/x/net v0.0.0-20201026091529-146b70c837a4 // indirect
)
