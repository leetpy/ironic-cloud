package main

import "time"

type user struct {
	name     string
	password string
}

type fault struct {
	Hadler   user
	CreateAt time.Time
	FinishAt time.Time
	Status   string
}
