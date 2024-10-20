package models

type User struct {
	UserID string `json:"user_id,omitempty" bson:"user_id"`
	Name   string `json:"name,omitempty" bson:"name,omitempty"`
	Email  string `json:"email,omitempty" bson:"email,omitempty"`
}
