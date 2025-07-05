package main
import (
	"net/http"
	"os"
	"alist/server"
)
func init() {
	server.Init()
}
var Handler http.Handler = http.DefaultServeMux
func main() {
	// 用于本地运行
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	http.ListenAndServe(":"+port, Handler)
}
