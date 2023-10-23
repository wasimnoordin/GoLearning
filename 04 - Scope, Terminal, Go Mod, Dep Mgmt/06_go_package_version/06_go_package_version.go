/*
Using the code from the previous hands-on exercise:

	- use a function from the package found at github.com/GoesToEleven/puppy but make your code depend on v1.2.0
		-- go get github.com/GoesToEleven/puppy@v1.2.0
	- inspect your go.mod file
*/

/*
PS C:\Users\wnoordin\Documents\Golang> go get github.com/GoesToEleven/puppy@v1.2.0
go: downloading github.com/GoesToEleven/puppy v1.2.0
go: downgraded github.com/GoesToEleven/puppy v1.3.0 => v1.2.0
*/

/* To revert to latest version:
PS C:\Users\wnoordin\Documents\Golang> go get github.com/GoesToEleven/puppy@latest
go: upgraded github.com/GoesToEleven/puppy v1.2.0 => v1.3.0
*/