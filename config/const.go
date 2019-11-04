package config

const (
	ServiceName = "Privy"
	Version     = "0.1.0"
	RequestID   = "request_id"

	// Table Name
	TableNameProduct          = "products"
	TableNameCategory         = "categories"
	TableNameProductImages    = "product_images"
	TableNameCategoryProducts = "category_products"

	envServerAddress = "SERVER_ADDRESS"

	// Headers
	HeaderRequestID = "X-Request-ID"

	// Database EnvVar
	envDatabaseHost        = "DATABASE_HOST"
	envDatabasePort        = "DATABASE_PORT"
	envDatabaseName        = "DATABASE_NAME"
	envDatabaseUsername    = "DATABASE_USERNAME"
	envDatabasePassword    = "DATABASE_PASSWORD"
	envDatabaseEnableLog   = "DATABASE_ENABLE_LOG"
	envDatabaseMaxOpenConn = "DATABASE_MAX_OPEN_CONN"
	envDatabaseMaxIdleConn = "DATABASE_MAX_IDLE_CONN"

	envLoggerFormat = "LOGGER_FORMAT"
	envLogLevel     = "LOG_LEVEL"
)
