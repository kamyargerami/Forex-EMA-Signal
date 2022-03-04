module main

go 1.17

replace bed/providers => ./providers

require bed/providers v0.0.0-00010101000000-000000000000

require (
	bed/helpers v0.0.0-00010101000000-000000000000 // indirect
	github.com/joho/godotenv v1.4.0 // indirect
)

replace bed/helpers => ./helpers
