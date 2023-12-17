// main.go
package main

import (
	"github.com/Rish-it/MINI_Gofr/handler"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func main() {
	s := g.Server()
	s.Group("/api").Bind([]ghttp.GroupItem{
		{"ALL", "*", middleware},
		{"GET", "/users", handler.ListUsers},
		{"GET", "/user/:id", handler.GetUser},
		{"POST", "/user", handler.CreateUser},
	})

	s.SetPort(8080)
	s.Run()
}

func middleware(r *ghttp.Request) {
	// You can implement middleware logic here
	r.Middleware.Next()
}
