package middleware

type Logger interface {
	Info(...any)
	Error(...any)
	WithFields(fields map[string]any) Logger
}
