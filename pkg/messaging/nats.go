package messaging

import (
	"time"

	"github.com/nats-io/nats.go"
)

type Config struct {
	URL           string `mapstructure:"url"`
	SubjectPrefix string `mapstructure:"subject_prefix"`
	Workers       int    `mapstructure:"workers"`
}

type Client struct {
	conn *nats.Conn
	cfg  Config
}

func NewClient(cfg Config) (*Client, error) {
	conn, err := nats.Connect(cfg.URL, nats.MaxReconnects(5), nats.ReconnectWait(time.Second))
	if err != nil {
		return nil, err
	}
	return &Client{conn: conn, cfg: cfg}, nil
}

func (c *Client) Publish(subject string, data []byte) error {
	return c.conn.Publish(c.prefixed(subject), data)
}

func (c *Client) Subscribe(subject string, handler nats.MsgHandler) (*nats.Subscription, error) {
	return c.conn.Subscribe(c.prefixed(subject), handler)
}

func (c *Client) Request(subject string, data []byte, timeout time.Duration) (*nats.Msg, error) {
	return c.conn.Request(c.prefixed(subject), data, timeout)
}

func (c *Client) Close() { c.conn.Close() }

func (c *Client) prefixed(subject string) string {
	if c.cfg.SubjectPrefix == "" {
		return subject
	}
	return c.cfg.SubjectPrefix + "." + subject
}
