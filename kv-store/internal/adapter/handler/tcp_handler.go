package handler

import (
	"bufio"
	"context"
	"io"
	"kv-store/internal/adapter/protocol"
	"kv-store/internal/domain/command"
	"kv-store/internal/usecase"
	"log"
	"net"
	"time"
)

type TCPHandler struct {
	commandHandler *usecase.CommandHandler
	parser         *protocol.Parser
}

func NewTCPHandler(commandHandler *usecase.CommandHandler, parser *protocol.Parser) *TCPHandler {
	return &TCPHandler{
		commandHandler: commandHandler,
		parser:         parser,
	}
}

func (h *TCPHandler) HandleConnection(conn net.Conn) {
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		if ctx.Err() != nil {
			return
		}

		line := scanner.Text()
		if line == "" {
			continue
		}

		cmd, err := h.parser.ParseCommand(line)
		if err != nil {
			response := h.parser.FormatError(err.Error())
			h.WriteResponse(conn, response)
			continue
		}

		if cmd.Type == command.QUIT {
			h.WriteResponse(conn, h.parser.FormatOK())
			return
		}

		response := h.commandHandler.ExecuteCommand(ctx, cmd)
		h.WriteResponse(conn, response)
	}

	if err := scanner.Err(); err != nil && err != io.EOF {
		log.Printf("ERROR FROM CONNECTION")
	}
}

func (h *TCPHandler) WriteResponse(conn net.Conn, response string) {
	conn.Write([]byte(response + "\n"))
}
