package usecase

import (
	"context"
	"fmt"
	"kv-store/internal/adapter/protocol"
	"kv-store/internal/domain/command"
	"kv-store/internal/domain/repository"
	"strconv"
	"strings"
)

type CommandHandler struct {
	store   repository.KeyValueRepository
	persist repository.PersistencxeRepository
	parser  *protocol.Parser
	stats   *Stats
}

func NewCommandHadler(store repository.KeyValueRepository, persist repository.PersistencxeRepository, stats *Stats, parser *protocol.Parser) *CommandHandler {
	return &CommandHandler{
		store:   store,
		persist: persist,
		parser:  parser,
		stats:   stats,
	}
}

func (h *CommandHandler) ExecuteCommand(ctx context.Context, cmd *protocol.Command) string {
	h.stats.IncrementCommads()

	switch cmd.Type {
	case command.SET:
		return h.handleSet(ctx, cmd.Args)
	case command.GET:
		return h.handleGet(ctx, cmd.Args)
	case command.DEL:
		return h.handleDel(ctx, cmd.Args)
	case command.EXPIRE:
		return h.handleExpire(ctx, cmd.Args)
	case command.TTL:
		return h.handleTTL(ctx, cmd.Args)
	case command.PERSIST:
		return h.handlePersists(ctx, cmd.Args)
	case command.KEYS:
		return h.handleKeys(ctx, cmd.Args)
	case command.PING:
		return h.handlePing(ctx, cmd.Args)
	case command.INFO:
		return h.handleInfo(ctx, cmd.Args)
	default:
		return h.parser.FormatError(fmt.Sprintf("Unknown command: %v", cmd.Type))
	}
}

func (h *CommandHandler) handleSet(ctx context.Context, args []string) string {
	if len(args) < 2 {
		return h.parser.FormatError("SET require at least 2 args!")
	}

	key := args[0]
	value := strings.Join(args[1:], " ")
	h.store.Set(ctx, key, value)

	if h.persist != nil {
		h.persist.Append(ctx, command.SET.String(), args)
	}

	return h.parser.FormatOK()
}

func (h *CommandHandler) handleGet(ctx context.Context, args []string) string {
	if len(args) < 1 {
		return h.parser.FormatError("GET require at least 1 args!")
	}

	value, found := h.store.Get(ctx, args[0])
	if found {
		return value
	}

	return h.parser.FormatNil()
}

func (h *CommandHandler) handleDel(ctx context.Context, args []string) string {
	if len(args) < 1 {
		return h.parser.FormatError("DEL require at least 1 args!")
	}

	count := 0
	for _, keys := range args {
		count += h.store.Del(ctx, keys)
	}

	if h.persist != nil && count > 0 {
		for _, keys := range args {
			h.persist.Append(ctx, command.DEL.String(), []string{keys})
		}
	}

	return h.parser.FormatResponse(count)
}

func (h *CommandHandler) handleExists(ctx context.Context, args []string) string {
	if len(args) < 1 {
		return h.parser.FormatError("EXISTS require at least 1 args!")
	}

	count := 0
	for _, keys := range args {
		if h.store.Exists(ctx, keys) {
			count++
		}
	}

	return h.parser.FormatResponse(count)
}

func (h *CommandHandler) handleExpire(ctx context.Context, args []string) string {
	if len(args) < 1 {
		return h.parser.FormatError("EXPIRE require at least 1 args!")
	}

	key := args[0]
	seconds, err := strconv.Atoi(args[1])
	if err != nil {
		return h.parser.FormatError("invalid seconds value")
	}

	success := h.store.Expire(ctx, key, seconds)
	if success {
		if h.persist != nil {
			h.persist.Append(ctx, command.EXPIRE.String(), []string{key})
		}
		return h.parser.FormatOK()
	}

	return h.parser.FormatResponse(success)
}

func (h *CommandHandler) handlePersists(ctx context.Context, args []string) string {
	if len(args) < 1 {
		return h.parser.FormatError("PERSISTS require at least 1 args!")
	}

	ok := h.store.Persist(ctx, args[0])
	if ok {
		if h.persist != nil {
			h.persist.Append(ctx, command.PERSIST.String(), []string{args[0]})
		}
		return h.parser.FormatOK()
	}

	return h.parser.FormatResponse(ok)
}

func (h *CommandHandler) handleTTL(ctx context.Context, args []string) string {
	if len(args) < 1 {
		return h.parser.FormatError("TTL require at least 1 args!")
	}

	ttl := h.store.TTL(ctx, args[0])
	return h.parser.FormatResponse(ttl)
}

func (h *CommandHandler) handleKeys(ctx context.Context, args []string) string {
	pattern := "*"
	if len(args) > 0 {
		pattern = args[0]
	}

	keys := h.store.Keys(ctx, pattern)
	if len(keys) == 0 {
		return ""
	}

	return strings.Join(keys, " ")
}

func (h *CommandHandler) handlePing(ctx context.Context, args []string) string {
	message := "PONG"
	if len(args) > 0 {
		message = strings.Join(args, " ")
	}
	return message
}

func (h *CommandHandler) handleInfo(ctx context.Context, args []string) string {
	return h.stats.GetInfo(ctx)
}
