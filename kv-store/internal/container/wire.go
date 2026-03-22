//go:build wireinject

// *build wireinject

package container

import (
	"kv-store/internal/adapter/handler"
	"kv-store/internal/adapter/protocol"
	"kv-store/internal/infrastructure/persistence"
	"kv-store/internal/infrastructure/storage"
	"kv-store/internal/usecase"

	"github.com/google/wire"
)

func InitializeContainer(opt persistence.AOFProviderOption) (*Container, error) {
	wire.Build(
		storage.NewStore,
		persistence.NewAOFProvider,
		protocol.NewParser,
		usecase.NewStats,
		usecase.NewCommandHadler,
		handler.NewTCPHandler,
		NewContainer,
	)
	return nil, nil
}
