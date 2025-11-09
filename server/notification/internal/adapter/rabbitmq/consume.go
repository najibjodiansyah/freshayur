package rabbitmq

import (
	"github.com/najibjodiansyah/freshayur/notification/internal"
	"github.com/najibjodiansyah/freshayur/notification/internal/core"
)

type Consumer interface {
	ConsumeMessage(qName string, req core.NotificationEntity) error
}

type consume struct {
	cfg internal.Config
}

func NewConcumer(cfg internal.Config) Consumer {
	return &consume{
		cfg: cfg,
	}
}

func (c *consume) ConsumeMessage(qName string, req core.NotificationEntity) error {
	return nil
}
