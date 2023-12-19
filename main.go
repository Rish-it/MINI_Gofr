// main.go
package main

import (
	"fmt"
	"sync"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
)

var (
	users      = make(map[int]User)
	userID     = 1
	userIDLock sync.Mutex
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {

	s := g.Server()


	
	s.BindMiddlewareDefault(func(r *ghttp.Request) {
		fmt.Printf("[%s] %s %s\n", gtime.Now().Format("Y-m-d H:i:s"), r.Method, r.URL.Path)
		r.Middleware.Next()
	})

	// API routes
	s.Group("/api").Bind([]ghttp.GroupItem{
		{"ALL", "*", middleware},
		{"GET", "/users", listUsers},
		{"GET", "/user/:id", getUser},
		{"POST", "/user", createUser},
	})

	s.SetPort(8080)
	s.Run()
}

func middleware(r *ghttp.Request) {
	// Your custom middleware logic here
	fmt.Println("Executing custom middleware")
	r.Middleware.Next()
}

func listUsers(r *ghttp.Request) {
	var userList []User
	for _, user := range users {
		userList = append(userList, user)
	}
	r.Response.WriteJson(userList)
}

func getUser(r *ghttp.Request) {
	id := r.GetInt("id")
	user, ok := users[id]
	if !ok {
		r.Response.WriteStatus(404)
		r.Response.WriteJson(g.Map{"error": "User not found"})
		return
	}
	r.Response.WriteJson(user)
}

func createUser(r *ghttp.Request) {
	userIDLock.Lock()
	defer userIDLock.Unlock()

	name := r.Get("name")
	email := r.Get("email")

	if name == "" || email == "" {
		r.Response.WriteStatus(400)
		r.Response.WriteJson(g.Map{"error": "Name and email are required"})
		return
	}

	
	userID++      // unique user ID
	user := User{
		ID:    userID,
		Name:  name.(string),
		Email: email.(string),
	}

	
	users[userID] = user   // storing user in the in-memory map

	r.Response.WriteJson(g.Map{"message": "User created successfully", "user": user})
}
