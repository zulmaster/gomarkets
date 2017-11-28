package connector

// Connector - универсальный интерфейс к биржам
type Connector interface {
	Connect() error
	Close()
}
