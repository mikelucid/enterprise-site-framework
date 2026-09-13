package facades

import (
	"time"

	"github.com/nats-io/nats.go"
	"github.com/mikelucid/enterprise-site-framework/bootstrap"
	"github.com/mikelucid/enterprise-site-framework/pkg/ai"
	"github.com/mikelucid/enterprise-site-framework/pkg/auth"
	"github.com/mikelucid/enterprise-site-framework/pkg/cache"
	"github.com/mikelucid/enterprise-site-framework/pkg/messaging"
)

func AI() (*ai.Engine, error) { return bootstrap.ResolveAs[*ai.Engine](bootstrap.GlobalContainer(), "ai") }
func Auth() (*auth.Manager, error) {
	return bootstrap.ResolveAs[*auth.Manager](bootstrap.GlobalContainer(), "auth")
}
func Cache() (*cache.InMemory, error) {
	return bootstrap.ResolveAs[*cache.InMemory](bootstrap.GlobalContainer(), "cache")
}

type CharacterFacade struct{}
type SiteFacade struct{}
type RoyaltyFacade struct{}
type PaymentFacade struct{}
type QueueFacade struct{}

func Character() CharacterFacade { return CharacterFacade{} }
func Site() SiteFacade           { return SiteFacade{} }
func Royalty() RoyaltyFacade     { return RoyaltyFacade{} }
func Payment() PaymentFacade     { return PaymentFacade{} }
func Queue() QueueFacade         { return QueueFacade{} }

func (CharacterFacade) Create(data map[string]any) map[string]any { return data }
func (CharacterFacade) Find(id string) map[string]any             { return map[string]any{"id": id} }
func (CharacterFacade) Update(id string, data map[string]any) map[string]any {
	data["id"] = id
	return data
}
func (CharacterFacade) Delete(string) bool { return true }

func (SiteFacade) Create(data map[string]any) map[string]any { return data }
func (SiteFacade) Deploy(id string) string                    { return "deployed:" + id }
func (SiteFacade) Scale(id string) string                     { return "scaled:" + id }
func (SiteFacade) GetByID(id string) map[string]any           { return map[string]any{"id": id} }

func (RoyaltyFacade) Calculate(amount float64, rate float64) float64 { return amount * rate }
func (RoyaltyFacade) Distribute(total float64, recipients int) float64 {
	if recipients <= 0 {
		return 0
	}
	return total / float64(recipients)
}

func (PaymentFacade) Collect(id string) string { return "collected:" + id }
func (PaymentFacade) Refund(id string) string  { return "refunded:" + id }
func (PaymentFacade) Status(id string) string  { return "status:" + id }

func (QueueFacade) Dispatch(subject string, payload []byte) error {
	client, err := bootstrap.ResolveAs[*messaging.Client](bootstrap.GlobalContainer(), "queue")
	if err != nil {
		return err
	}
	return client.Publish(subject, payload)
}
func (QueueFacade) Listen(subject string, handler func([]byte) error) error {
	client, err := bootstrap.ResolveAs[*messaging.Client](bootstrap.GlobalContainer(), "queue")
	if err != nil {
		return err
	}
	_, err = client.Subscribe(subject, func(msg *nats.Msg) {
		_ = handler(msg.Data)
	})
	return err
}

func Recommend(input string) ([]string, error) {
	e, err := AI()
	if err != nil {
		return nil, err
	}
	return e.Recommend(input), nil
}
func SearchHarmonics(query string) ([]string, error) {
	e, err := AI()
	if err != nil {
		return nil, err
	}
	return e.SearchHarmonics(query), nil
}
func Resonate(signal string) (string, error) {
	e, err := AI()
	if err != nil {
		return "", err
	}
	return e.Resonate(signal), nil
}

func User(token string) (string, error) {
	m, err := Auth()
	if err != nil {
		return "", err
	}
	claims, err := m.ValidateToken(token)
	if err != nil {
		return "", err
	}
	return claims.Subject, nil
}
func HasPermission(roles []string, permission string) bool {
	return auth.NewRBAC().HasPermission(roles, permission)
}
func GenerateToken(subject string, roles []string) (string, error) {
	m, err := Auth()
	if err != nil {
		return "", err
	}
	return m.GenerateToken(subject, roles)
}

func Get(key string) (string, bool, error) {
	c, err := Cache()
	if err != nil {
		return "", false, err
	}
	v, ok := c.Get(key)
	return v, ok, nil
}
func Set(key, value string, ttl time.Duration) error {
	c, err := Cache()
	if err != nil {
		return err
	}
	c.Set(key, value, ttl)
	return nil
}
func Delete(key string) error {
	c, err := Cache()
	if err != nil {
		return err
	}
	c.Delete(key)
	return nil
}
func Remember(key string, ttl time.Duration, loader func() (string, error)) (string, error) {
	v, ok, err := Get(key)
	if err != nil {
		return "", err
	}
	if ok {
		return v, nil
	}
	v, err = loader()
	if err != nil {
		return "", err
	}
	return v, Set(key, v, ttl)
}
