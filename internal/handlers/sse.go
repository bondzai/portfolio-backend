package handlers

type (
	SseHandler interface{}
	sseHandler struct{}
)

func NewSseHandler() SseHandler {
	return &sseHandler{}
}
