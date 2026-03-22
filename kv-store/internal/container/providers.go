package container

import (
	"kv-store/internal/adapter/handler"
	"kv-store/internal/adapter/protocol"
	"kv-store/internal/domain/repository"
	"kv-store/internal/usecase"
)

type Container struct {
	Store          repository.KeyValueRepository
	Persistence    repository.PersistencxeRepository
	CommandHandler *usecase.CommandHandler
	TCPHandler     *handler.TCPHandler
	Parser         *protocol.Parser
}

func NewContainer(
	store repository.KeyValueRepository,
	persist repository.PersistencxeRepository,
	cmdHandler *usecase.CommandHandler,
	tcpHandler *handler.TCPHandler,
	parser *protocol.Parser,
) *Container {
	return &Container{
		Store:          store,
		Persistence:    persist,
		CommandHandler: cmdHandler,
		TCPHandler:     tcpHandler,
		Parser:         parser,
	}
}

func (c *Container) Close() error {
	c.Store.StopCleanup()

	if c.Persistence != nil {
		return c.Persistence.Close()
	}

	return nil
}
