module github.com/robotic-framework/srv-configuration

go 1.14

replace k8s.io/client-go => k8s.io/client-go v0.18.8

require (
	github.com/eden-framework/client v0.0.0-20201022095936-63753150b326
	github.com/eden-framework/context v0.0.2
	github.com/eden-framework/courier v1.0.4
	github.com/eden-framework/eden-framework v1.1.10-0.20201127094502-ef783952fb49
	github.com/eden-framework/sqlx v0.0.1
	github.com/konsorten/go-windows-terminal-sequences v1.0.3 // indirect
	github.com/sirupsen/logrus v1.7.0
	github.com/spf13/cobra v0.0.5
)
