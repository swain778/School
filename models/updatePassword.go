package models

type UpdatePassword struct {
	ID              uint
	UserID          uint
	Password        string
	ConformPassword string
}
