package main

import (
  "github.com/joho/godotenv"
  "github.com/julienschmidt/httprouter"
  log "github.com/sirupsen/logrus"
  "net/http"
  "os"
  "time"

  _ "github.com/go-sql-driver/mysql"
  routes "go.web/routes"
  "go.web/utils/database"
  "go.web/utils/redis"
)

// Database config
type Database struct {
  Hostname string
  Username string
  Password string
  Dialect  string
  Database string
  Port     string
}

// RedisConfig .
type RedisConfig struct {
  Hostname string
  Password string
  DB       int
}

// AppConfig .
type AppConfig struct {
  Database    Database
  RedisConfig RedisConfig
}

// SetupLogger setup logrus log
func SetupLogger() {
  format := new(log.TextFormatter)
  format.TimestampFormat = time.RFC3339
  format.FullTimestamp = true
  log.SetFormatter(format)
}

// SetupEnvironment .
func SetupEnvironment(appConfig *AppConfig) {
  appConfig.Database.Hostname = os.Getenv("DB_HOSTNAME")
  appConfig.Database.Username = os.Getenv("DB_USERNAME")
  appConfig.Database.Password = os.Getenv("DB_PASSWORD")
  appConfig.Database.Dialect = os.Getenv("DB_DIALECT")
  appConfig.Database.Database = os.Getenv("DB_DATABASE")
  appConfig.Database.Port = os.Getenv("DB_PORT")

  appConfig.RedisConfig.Hostname = os.Getenv("REDIS_HOSTNAME")
  appConfig.RedisConfig.Password = os.Getenv("REDIS_PASSWORD")
}

func main() {
  SetupLogger()

  switch ev := os.Getenv("GO_ENV"); ev {
  case "development":
    err := godotenv.Load()
    if err != nil {
      log.Warning(err)
    }
  }

  var appConfig AppConfig

  SetupEnvironment(&appConfig)

  log.Info("Starting server...")

  log.Info("Connecting to database...")
  err := database.SetupDatabases(appConfig.Database.Dialect, appConfig.Database.Username, appConfig.Database.Password, appConfig.Database.Hostname, appConfig.Database.Port, appConfig.Database.Database)
  if err != nil {
    log.Fatal(err)
  }

  err = database.GetHandle().Ping()
  if err != nil {
    log.Fatal(err)
  }
  log.Info("Success connecting to database!")

  log.Info("Connecting to redis...")
  err = redis.SetupRedis(appConfig.RedisConfig.Hostname, appConfig.RedisConfig.Password, 0)
  if err != nil {
    log.Fatal(err)
  }
  log.Info("Success connecting to redis!")

  router := httprouter.New()

  // start router
  router.GET("/ping", routes.Ping)

  log.Info("Start http server listening...")
  log.Fatal(http.ListenAndServe(os.Getenv("SERVER_LISTENING_HOST"), router))
}
