# Make sure you have at least go version 1.18
go version

# Validate units test pass
go test

# Starting all fuzzing tests
go test --fuzz=Fuzz
