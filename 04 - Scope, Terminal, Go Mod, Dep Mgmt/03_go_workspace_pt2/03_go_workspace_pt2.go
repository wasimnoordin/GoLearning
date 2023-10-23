/*

Using the code you wrote in the previous hands-on exercise:
	- look at the contents in the folder of your module
		- ls -la
	- build your program
		- any of these
			- go build main.go
			- go build .
			- go build ./…
	- run your executable
		- ./<name of executable>
		- this could vary on your machine, if so, google how to do it on your machine, for example, "on a mac, run
		  executable at the terminal" or "run executable from terminal mac"
*/

/* TERMINAL:

wnoordin@LUK-5CG32158TW MINGW64 ~
$ pwd
/c/Users/wnoordin

wnoordin@LUK-5CG32158TW MINGW64 ~
$ cd Documents/Golang/04\ -\ Scope\,\ Terminal\,\ Go\ Mod\,\ Dep\ Mgmt/03_go_workspace_pt2/

wnoordin@LUK-5CG32158TW MINGW64 ~/Documents/Golang/04 - Scope, Terminal, Go Mod, Dep Mgmt/03_go_workspace_pt2 (main)
$ go build main.go

wnoordin@LUK-5CG32158TW MINGW64 ~/Documents/Golang/04 - Scope, Terminal, Go Mod, Dep Mgmt/03_go_workspace_pt2 (main)
$ ls -la
total 1886
drwxr-xr-x 1 wnoordin 1049089       0 Oct 20 12:27 ./
drwxr-xr-x 1 wnoordin 1049089       0 Oct 20 12:23 ../
-rw-r--r-- 1 wnoordin 1049089     470 Oct 20 12:19 03_go_workspace_pt2.go
-rwxr-xr-x 1 wnoordin 1049089 1923584 Oct 20 12:27 main.exe*
-rw-r--r-- 1 wnoordin 1049089      81 Oct 20 12:22 main.go

wnoordin@LUK-5CG32158TW MINGW64 ~/Documents/Golang/04 - Scope, Terminal, Go Mod, Dep Mgmt/03_go_workspace_pt2 (main)
$ ./main.exe
Hello Gophers

wnoordin@LUK-5CG32158TW MINGW64 ~/Documents/Golang/04 - Scope, Terminal, Go Mod, Dep Mgmt/03_go_workspace_pt2 (main)

*/