module tweets

go 1.21

require github.com/google/uuid v1.3.0

require gorm.io/gorm v1.25.2

require github.com/joho/godotenv v1.5.1 // indirect

require (
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.4.2 // indirect
	golang.org/x/crypto v0.11.0 // indirect
	gorm.io/driver/postgres v1.5.2
)

require (
	github.com/golang/protobuf v1.5.3
	github.com/grpc-ecosystem/go-grpc-middleware/v2 v2.0.0-rc.5
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	golang.org/x/net v0.12.0 // indirect
	golang.org/x/sys v0.10.0 // indirect
	golang.org/x/text v0.11.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230726155614-23370e0ffb3e // indirect
	google.golang.org/grpc v1.57.0
	google.golang.org/protobuf v1.31.0
)
