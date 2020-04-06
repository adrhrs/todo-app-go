package main

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Activity struct {
	ID       int    `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	Priority int    `json:"priority" db:"priority"`
}

type Category struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	Desc string `json:"desc" db:"desc"`
}

type User struct {
	ID       int    `json:"id" db:"id"`
	Username string `json:"user_name" db:"user_name"`
	Password string `json:"password" db:"password"`
	Fullname string `json:"full_name" db:"full_name"`
}
