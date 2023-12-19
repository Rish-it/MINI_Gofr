package handler // handler/user.go

import (
	"sync"

	"github.com/Rish-it/MINI_Gofr/model"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var (
	users      = make(map[int]model.User)
	userID     = 1
	userIDLock sync.Mutex
)

func ListUsers(r *ghttp.Request) {
	var userList []model.User
	for _, user := range users {
		userList = append(userList, user)
	}
	r.Response.WriteJson(userList)
}

func GetUser(r *ghttp.Request) {
	id := r.GetInt("id")
	user, ok := users[id]
	if !ok {
		r.Response.WriteStatus(404)
		r.Response.WriteJson(g.Map{"error": "User not found"})
		return
	}
	r.Response.WriteJson(user)
}

func CreateUser(r *ghttp.Request) {
	userIDLock.Lock()
	defer userIDLock.Unlock()

	name := r.Get("name")
	email := r.Get("email")

	// Basic input validation
	if name == "" || email == "" {
		r.Response.WriteStatus(400)
		r.Response.WriteJson(g.Map{"error": "Name and email are required"})
		return
	}

	// Generating a unique user ID
	userID++
	user := model.User{
		ID:    userID,
		Name:  name.(string),
		Email: email.(string),
	}

	// Saving the user in the in-memory map
	users[userID] = user

	r.Response.WriteJson(g.Map{"message": "User created successfully", "user": user})
}
