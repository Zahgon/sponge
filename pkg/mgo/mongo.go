// Package mgo is a library wrapped on go.mongodb.org/mongo-driver/mongo, with added features paging queries, etc.
package mgo

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type Database = mongo.Database

var ErrNoDocuments = mongo.ErrNoDocuments

const (
	// DBDriverName mongodb driver
	DBDriverName = "mongodb"
)

// Init connecting to mongo
func Init(dsn string, opts ...*options.ClientOptions) (*mongo.Database, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Init2 connecting to mongo using uri
func Init2(uri string, dbName string, opts ...*options.ClientOptions) (*mongo.Database, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Close mongodb
func Close(db *mongo.Database) error { _ = "STUB: not implemented"; return nil }

// WithOption set option for mongodb
func WithOption() *options.ClientOptions { _ = "STUB: not implemented"; return nil }

type customLogger struct {
	zapLogger *zap.Logger
}

func (l *customLogger) Info(_ int, msg string, kvs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (l *customLogger) Error(err error, msg string, kvs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

// NewCustomLogger create a custom logger for mongodb, debug level is used by default.
// example: WithOption().SetLoggerOptions(NewCustomLogger(logger.Get(), true))
func NewCustomLogger(l *zap.Logger, isDebugLevel bool) *options.LoggerOptions {
	_ = "STUB: not implemented"
	return nil
}

// Create a client with our logger options.
