module chapter2

go 1.24.2

replace numerics => ../modules/numerics

replace greetings => ../modules/greetings

require (
	greetings v0.0.0-00010101000000-000000000000
	numerics v0.0.0-00010101000000-000000000000
)
