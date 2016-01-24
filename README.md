euler1
===
compares various approaches to summing multiples of 3 and 5 below 1000; includes unit tests, benchmarks, and code coverage

usage
===
       > go test github.com/wrightmikea/euler1/main -cover -coverprofile=cover.out -v -bench=. -benchtime=5s
	   > go tool cover -func=cover.out
