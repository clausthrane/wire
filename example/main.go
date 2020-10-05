package main

import "context"

func main() {
	app, cancel := NewApplication()
	defer cancel()
	service := app.Service()
	println(service.GetValue(context.TODO()))
}
