package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ClickHouse/clickhouse-go"
	"github.com/sirupsen/logrus"
)

type ClickHouseHook struct {
	db        *sql.DB
	entries   []logrus.Entry
	batchSize int
}

// NewClickHouseHook establishes a connection to ClickHouse using the provided DSN.
func NewClickHouseHook(dsn string, batchSize int) (*ClickHouseHook, error) {
	db, err := sql.Open("clickhouse", dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			log.Fatalf("[%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		} else {
			log.Fatal(err)
		}
	}
	return &ClickHouseHook{db: db, batchSize: batchSize}, nil
}

// Fire is triggered by Logrus to log entries to ClickHouse.
func (hook *ClickHouseHook) Fire(entry *logrus.Entry) error {
	hook.entries = append(hook.entries, *entry)
	if len(hook.entries) >= hook.batchSize {
		if err := hook.flush(); err != nil {
			return err
		}
	}
	return nil
}

// flush sends the collected log entries to ClickHouse in a batch.
func (hook *ClickHouseHook) flush() error {
	tx, err := hook.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("INSERT INTO tiered_logs (event_time, level, message) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, entry := range hook.entries {
		if _, err := stmt.Exec(entry.Time, entry.Level.String(), entry.Message); err != nil {
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	// Clear the entries after flushing
	hook.entries = nil
	return nil
}

// Levels returns the logging levels for which the hook is triggered.
func (hook *ClickHouseHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func main() {
	// ClickHouse DSN (replace with your credentials and host)
	dsn := "tcp://localhost:9000?database=default&username=default&password=&debug=true"

	// Create ClickHouse hook with a batch size of 5
	hook, err := NewClickHouseHook(dsn, 5)
	if err != nil {
		log.Fatalf("failed to connect to ClickHouse: %v", err)
	}
	defer hook.db.Close()

	// Set up logrus
	logger := logrus.New()
	logger.Out = os.Stdout
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logger.AddHook(hook)

	// Log some entries
	for i := 0; i < 10; i++ {
		logger.WithFields(logrus.Fields{
			"iteration": i,
		}).Info("This is an info log entry")

		time.Sleep(time.Second)
	}

	// Flush any remaining log entries before exiting
	if err := hook.flush(); err != nil {
		log.Fatalf("failed to flush logs to ClickHouse: %v", err)
	}

	fmt.Println("Logs sent to ClickHouse.")
}
