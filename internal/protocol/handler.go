package protocol

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"own-redis/internal/storage"
)

type Handler struct {
	Store *storage.Store
}

func NewHandler(s *storage.Store) *Handler {
	return &Handler{Store: s}
}

func (h *Handler) HandleProcess(input string) string {
	input = strings.TrimSpace(input)
	if input == "" {
		return ""
	}

	args := strings.Fields(input)
	if len(args) == 0 {
		return ""
	}

	command := strings.ToUpper(args[0])
	switch command {
	case "PING":
		return "PONG\n"
	case "GET":
		return h.handleGet(args)
	case "SET":
		return h.handleSet(args)
	default:
		return fmt.Sprintf("(error) ERR unknown command '%s'\n", command)
	}
}

func (h *Handler) handleGet(args []string) string {
	if len(args) == 0 {
		return "(error) ERR wrong number of arguments for 'GET' command\n"
	}

	key := args[0]
	val, exists := h.Store.Get(key)
	if !exists {
		return "(nil)\n"
	}

	return val + "\n"
}

func (h *Handler) handleSet(args []string) string {
	if len(args) < 2 {
		return "(error) ERR wrong number of arguments of 'SET' command"
	}

	key := args[0]
	var valueArgs []string
	var ttl time.Duration = 0

	if len(args) >= 4 && strings.ToUpper(args[len(args)-2]) == "PX" {
		msStr := args[len(args)-1]
		ms, err := strconv.ParseInt(msStr, 10, 64)
		if err != nil || ms <= 0 {
			return "(error) ERR invalid expire time in 'SET' command"
		}

		ttl = time.Duration(ms) * time.Millisecond
		valueArgs = args[1 : len(args)-2]
	} else {
		valueArgs = args[1:]
	}

	if len(valueArgs) == 0 {
		return "(error) ERR wrong number of arguments for 'SET' command"
	}

	finalValue := strings.Join(valueArgs, " ")
	h.Store.Set(key, finalValue, ttl)
	return "OK\n"
}
