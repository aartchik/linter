package zap

type Logger struct{}
type Field struct{}

func NewDevelopment(options ...any) (*Logger, error) { return &Logger{}, nil }
func NewProduction(options ...any) (*Logger, error)  { return &Logger{}, nil }

func ReplaceGlobals(logger *Logger) {}

func L() *Logger { return &Logger{} }

func (l *Logger) Sync() error { return nil }

func (l *Logger) Debug(msg string, fields ...Field) {}
func (l *Logger) Info(msg string, fields ...Field)  {}
func (l *Logger) Warn(msg string, fields ...Field)  {}
func (l *Logger) Error(msg string, fields ...Field) {}

func String(key string, value string) Field { return Field{} }
func Int(key string, value int) Field       { return Field{} }
func Bool(key string, value bool) Field     { return Field{} }
func Any(key string, value any) Field       { return Field{} }
func Error(err error) Field                 { return Field{} }