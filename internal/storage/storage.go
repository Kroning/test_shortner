package storage

import (
	"context"
	"errors"
)

var (
	LinkNotExists error = errors.New("storage: Link exists")
)

type Storage interface {
	LoadLinksByAlias(ctx context.Context, alias string) ([][]string, error)
	CheckLinkExistance(ctx context.Context, alias string) (string, error)
	InsertLink(ctx context.Context, alias string, url string) error
}
