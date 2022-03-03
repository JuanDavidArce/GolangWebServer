package main

func main() {
	server := NewServer(":3000")
	server.Handle("/", HadleRoot)
	server.Handle("/api", server.AddMidleware(HandleHome, CheckAuth()))
	server.Listen()
}
