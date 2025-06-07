module chapter1

go 1.24.2

replace greetings => ../modules/greetings
replace numerics => ../modules/numerics

require (
	greetings v0.0.0-00010101000000-000000000000
	numerics v0.0.0-00010101000000-000000000000
)

