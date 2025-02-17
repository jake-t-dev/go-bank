package main

func main() {
	server := NewAPIServer(":3333")
	server.Run()
}
