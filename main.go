package main

func main() {
	host, port, threads := parser()
	r := setupRouter(threads)
	r.Run(address(host, port)) // Use string conversion for port
}
