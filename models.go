package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/kiruiaaron/rssagg/internal/database"
)


type CreateUserRow struct {
	ID        uuid.UUID   `json:"id"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	Name      string      `json:"name"`
}


func databaseUserToUser(dbUser database.CreateUserRow) CreateUserRow{
	return CreateUserRow{
		ID: dbUser.ID,
		CreatedAt : dbUser.CreatedAt ,
		UpdatedAt: dbUser.UpdatedAt,
		Name: dbUser.Name,
	}
}