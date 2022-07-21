package configuration

type Interface interface {
	database
}

type database interface {
	GetDatabaseConnString() string
	SetDatabaseConnString()

	GetDatabaseUser() string
	SetDatabaseUser(user string)

	GetDatabasePassword() string
	SetDatabasePassword(password string)

	GetDatabaseHost() string
	SetDatabaseHost(host string)

	GetDatabasePort() int
	SetDatabasePort(port int)

	GetDatabaseName() string
	SetDatabaseName(name string)

	GetDatabaseSSLMode() bool
	SetDatabaseSSLMode(useSSL bool)

	GetDatabaseMaxConnectionCount() int
	SetDatabaseMaxConnectionCount(maxConns int)
}
