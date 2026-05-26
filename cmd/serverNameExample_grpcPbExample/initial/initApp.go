// Package initial is the package that starts the service to initialize the service, including
// the initialization configuration, service configuration, connecting to the database, and
// resource release needed when shutting down the service.
package initial

//"github.com/go-dev-frame/sponge/internal/database"

var (
	version    string
	configFile string
)

// InitApp initial app configuration
func InitApp() { _ = "STUB: not implemented"; return }

// initializing log

//logger.WithFileName(cfg.Logger.LogFileConfig.Filename),
//logger.WithFileMaxSize(cfg.Logger.LogFileConfig.MaxSize),
//logger.WithFileMaxBackups(cfg.Logger.LogFileConfig.MaxBackups),
//logger.WithFileMaxAge(cfg.Logger.LogFileConfig.MaxAge),
//logger.WithFileIsCompression(cfg.Logger.LogFileConfig.IsCompression),

// initializing tracing

// initializing the print system and process resources

// invalid if it is windows, the default threshold for cpu and memory is 0.8, you can modify them

// initializing database
//database.InitDB()
//logger.Infof("[%s] was initialized", cfg.Database.Driver)
//database.InitCache(cfg.App.CacheType)
//if cfg.App.CacheType != "" {
//	logger.Infof("[%s] was initialized", cfg.App.CacheType)
//}

func initConfig() { _ = "STUB: not implemented"; return }

// get configuration from local configuration file
func getConfigFromLocal() { _ = "STUB: not implemented"; return }
