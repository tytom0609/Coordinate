// main.go

package main

func main() {
	a := App{}
	// You need to set your Username and Password here
	a.Initialize()

	a.Run(":8080")
}
