module example.com/hello

go 1.16

require (
	example.com/greetings v0.0.0-00010101000000-000000000000
	example.com/more_actions v0.0.0-00010101000000-000000000000
	rsc.io/quote v1.5.2
)

replace example.com/greetings => ../greetings

replace example.com/more_actions => ../more_actions
