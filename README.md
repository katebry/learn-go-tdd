# learn-go-tdd

Repo created whilst following [Learn Go With Tests](https://quii.gitbook.io/learn-go-with-tests/) 

Useful packages:
- [goimports](https://pkg.go.dev/golang.org/x/tools/cmd/goimports) 
- [errcheck](https://pkg.go.dev/github.com/kisielk/errcheck/internal/errcheck)

Commands:
- `go test` to run _test files
- `go test -v` runs _test files and gives you a breakdown of each test run
- `go test -bench=.` run benchmark tests in _test files
- `go test -cover` shows your test coverage

Test 1: `basics/hello.go` - a function that returns a greeting

Test 2: `integers/adder.go` - a function that adds two numbers together

Test 3: `for/iteration.go` - a function that repeats a string 

Test 4: `arrays/sum.go` - a function that adds the values of an array

Test 5: `structs/shapes.go` - a function that calculates the perimeter of a rectangle

Test 6: `pointers/wallet.go` - a function that allows us to deposit money into a wallet

Test 7: `maps/dictionary.go` - a function that returns the definition when a word is searched for

Further Reading:

- [slices](https://go.dev/blog/slices-intro)

- [table driven tests](https://github.com/golang/go/wiki/TableDrivenTests)

- [parametric polymorphism](https://en.wikipedia.org/wiki/Parametric_polymorphism)

- [pointers](https://gobyexample.com/pointers)

- [comparable types](https://go.dev/ref/spec#Comparison_operators)