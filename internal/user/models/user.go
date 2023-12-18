package models

import (
	"time"

	"github.com/squacky/openchat/internal/user/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	FullName  string             `bson:"full_name"`
	Email     string             `bson:"email"`
	Avatar    string             `bson:"avatar"`
	Phone     string             `bson:"phone"`
	Status    string             `bson:"status"`
	CreatedOn time.Time          `bson:"created_on"`
	UpdatedOn time.Time          `bson:"updated_on"`
}

func (u *User) Domain() *domain.User {
	return &domain.User{
		ID:       u.ID.Hex(),
		FullName: u.FullName,
		Email:    u.Email,
		Phone:    u.Phone,
	}
}
