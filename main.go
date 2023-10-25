package main

const port = ":8080"

func main() {
	httpServer := NewAPIserver(port)
	httpServer.Run()
}
