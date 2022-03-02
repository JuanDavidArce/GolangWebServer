package main

func main() {
	server := NewServer(":3000")
	server.Handle("/", HadleRoot)
	server.Handle("/api", HandleHome)
	server.Listen()
}
