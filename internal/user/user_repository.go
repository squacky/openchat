package user

import (
	"context"
	"time"

	"github.com/squacky/openchat/internal/user/domain"

	"github.com/pkg/errors"
	"github.com/squacky/openchat/internal/user/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository interface {
	CreateUser(ctx context.Context, user *domain.User) error
	GetUserByID(ctx context.Context, id string) (*domain.User, error)
}

type UserRepository struct {
	collection *mongo.Collection
}

func (r *UserRepository) CreateUser(ctx context.Context, user *domain.User) error {
	opts := options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After)
	now := time.Now()
	res := r.collection.FindOneAndUpdate(ctx, bson.M{"email": user.Email}, bson.M{
		"$set": bson.M{
			"full_name":  user.FullName,
			"avatar":     user.Avatar,
			"updated_on": now,
		},
		"$setOnInsert": bson.M{
			"email":      user.Email,
			"phone":      user.Phone,
			"status":     "active",
			"created_on": now,
		},
	}, opts)
	if err := res.Err(); err != nil {
		return errors.Wrap(err, "failed to persist user")
	}
	return nil
}

func (r *UserRepository) GetUserByID(ctx context.Context, id string) (*domain.User, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.Wrap(err, "un-supported user id format provided")
	}
	res := r.collection.FindOne(ctx, bson.M{"_id": _id})
	if err := res.Err(); err != nil || errors.Is(err, mongo.ErrNoDocuments) {
		return nil, errors.Wrap(err, "user not found")
	}
	var user models.User
	if err := res.Decode(&user); err != nil {
		return nil, errors.Wrap(err, "failed to decode the user")
	}
	return user.Domain(), nil
}

func NewUserRepository(db *mongo.Database) *UserRepository {
	return &UserRepository{
		collection: db.Collection("users"),
	}
}
