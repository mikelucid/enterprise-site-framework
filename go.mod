module github.com/mikelucid/enterprise-site-framework

go 1.21

require (
	github.com/golang-jwt/jwt/v5 v5.2.1
	github.com/google/uuid v1.6.0
	github.com/nats-io/nats.go v1.36.0
	github.com/spf13/cobra v1.8.1
	github.com/spf13/viper v1.19.0
	github.com/stretchr/testify v1.9.0
	github.com/gin-gonic/gin v1.10.0
	go.uber.org/zap v1.27.0
	gorm.io/driver/postgres v1.5.9
	gorm.io/gorm v1.25.12
	go.opentelemetry.io/otel v1.28.0
	go.opentelemetry.io/otel/exporters/jaeger v1.17.0
	go.opentelemetry.io/otel/sdk v1.28.0
	github.com/go-playground/validator/v10 v10.22.0
	github.com/redis/go-redis/v9 v9.6.1
)
