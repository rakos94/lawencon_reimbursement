package config

// DB
const (
	host    = "103.30.180.34"
	user    = "postgres"
	pass    = "bootlawen123"
	dbname  = "reimbursement"
	port    = 9595
	sslmode = "disable"
)

// APPS
const (
	AppPort        = "8090"
	HostCredential = "149.28.141.16:8091"
	GRPCPort       = "8190"
	// HostCredential = "192.168.15.245:1111"
)

// JWT
var (
	UserIdExample        = "20022012"
	SecretKeyExample     = "mLmHu8f1IxFo4dWurBG3jEf1Ex0wDZvvwND6eFmcaX"
	SigningMethodExample = "HS512"
	JwtExpiredHour       = 12
)
