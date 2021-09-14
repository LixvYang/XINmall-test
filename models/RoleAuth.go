// Package models provides 
package models

type RoleAuth struct {
	AuthId int
	RoleId   int
}

func (r *RoleAuth) TableName() string {
	return "role_auth"
}
