package main

func main() {
	a := App{}
	// set usr pwd
	a.Initialize("root", "root", "golangdb")

	a.Run(":3000")
}
