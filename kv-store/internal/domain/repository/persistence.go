package repository

import "context"

type PersistencxeRepository interface {
	Append(ctx context.Context, command string, args []string) error
	Replay(ctx context.Context, store KeyValueRepository, args []string) error
	Close() error
}
