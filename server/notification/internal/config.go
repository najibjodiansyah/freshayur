package internal

type Config struct {
	RabbitMQ RabbitMQConfig
}

type RabbitMQConfig struct {
	Host     string
	Port     int
	User     string
	Password string
}

type EmailConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Sending  string
	IsTLS    bool
}
