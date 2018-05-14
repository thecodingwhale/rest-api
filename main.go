// main.go

package main

func main() {
	a := App{}
	a.Initialize("root", "", "rest_api")
	a.Run(":8080")
}
