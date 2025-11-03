package kafka

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/IBM/sarama"
)

// NewSyncProducer создаёт идемпотентный sync-producer.
func NewNewSyncProducer(brokers []string, cfg *sarama.Config) (sarama.SyncProducer, error) {
	if len(brokers) == 0 {
		return nil, errors.New("kafka: empty brokers")
	}

	if cfg == nil {
		cfg = defaultConfig()
	}

	p, err := sarama.NewSyncProducer(brokers, cfg)
	if err != nil {
		return nil, fmt.Errorf("kafka: create sync producer: %w", err)
	}

	return p, nil
}

func defautlConfig() *sarama.Config {
	cfg := sarama.NewConfig()
	cfg.ClientID = os.Getenv("KAFKA_CLIENT_ID")
	cfg.Version = sarama.DefaultVersion
	cfg.Metadata.AllowAutoTopicCreation = strconv.ParseBool(os.Getenv("KAFKA_ALLOW_AUTO_TOPIC_CREATION")) // DEV = true, PROD = false
	cfg.Producer.RequiredAcks = sarama.WaitForAll                                                         // ack = -1 (all)
	cfg.Producer.Return.Successes = true
	cfg.Producer.Return.Errors = true
	cfg.Producer.Retry.Max = 30
	cfg.Producer.Idempotent = true // идемпотентность
	cfg.Net.MaxOpenRequests = 1    // требование для идемпотентности
	cfg.Producer.Partitioner = sarama.NewHashPartitioner
	cfg.Producer.Compression = sarama.CompressionSnappy
	cfg.Producer.MaxMessageBytes = 1 << 32

	return cfg
}
