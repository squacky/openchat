package mongodb

import (
	"context"

	"github.com/pkg/errors"
	"github.com/squacky/openchat/config"
	"github.com/squacky/openchat/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/mongo/driver/connstring"
)

type Config struct {
	URI string
}

func NewConfig(c *config.Config) *Config {
	return &Config{URI: c.MongoDB.URI}
}

func Connect(ctx context.Context, cfg *Config) (*mongo.Database, error) {
	cs, err := connstring.ParseAndValidate(cfg.URI)
	if err != nil {
		return nil, err
	}

	c, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.URI))
	if err != nil {
		return nil, errors.Wrap(err, "failed to establish db connection")
	}

	if err := c.Ping(ctx, nil); err != nil {
		return nil, errors.Wrap(err, "failed to ping database")
	}

	logger.Info("mongodb connected")

	return c.Database(cs.Database), nil
}
