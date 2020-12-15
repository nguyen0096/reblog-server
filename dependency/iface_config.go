package dependency

// IConfig server configurations
type IConfig interface {
	IDatabaseConfig
}

// IDatabaseConfig contains getters for database configurations
type IDatabaseConfig interface {
	GetHostname() string
	GetPort() int
	GetUser() string
	GetPassword() string
	GetDatabase() string
}
