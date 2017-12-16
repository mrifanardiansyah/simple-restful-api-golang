package main

func main() {
	var app = App{}

	app.Initialize("root", "sunday", "rest_api_example")
	app.Run(":8080")
}
