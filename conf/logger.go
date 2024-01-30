package conf

type Logger struct {
	Level        string
	Prefix       string
	Director     string
	FilePath     string
	ShowLine     bool
	LogInConsole bool
	MaxSize      int
	MaxBackups   int
	MaxAge       int
}
