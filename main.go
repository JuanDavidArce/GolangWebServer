package main

func main() {
	server := NewServer(":3000")
	server.Handle("GET", "/", HadleRoot)
	server.Handle("POST", "/user", UserPostRequest)
	server.Handle("POST", "/create", PostRequest)
	server.Handle("POST", "/api", server.AddMidleware(HandleHome, CheckAuth(), Logging()))
	server.Listen()
}
