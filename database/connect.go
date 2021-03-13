package database
import(

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)


func Connect()(*gorm.DB, error){

	dbUrl := "postgres://rygwkpxrxcijxh:ecb82283e1aecae02b2e2d5b704f56bbbe6ea8ad1cb9aa7187ac60640565ca14@ec2-54-164-22-242.compute-1.amazonaws.com:5432/d900abdh65ie89"
	if dbUrl != "" {
		return gorm.Open(postgres.Open(dbUrl), &gorm.Config{
			PrepareStmt: true,
			Logger:      logger.New(
				log.New(os.Stdout, "\r\n", log.LstdFlags),
				logger.Config{
					SlowThreshold: time.Second,
					Colorful:      true,
					LogLevel:      logger.Info,
				},
			),
		})
	} else {
		dsn := "host=127.0.0.1 user=postgres password=asdfasw22 dbname=staem port=5432 sslmode=disable TimeZone=Asia/Shanghai"
		return gorm.Open(postgres.Open(dsn), &gorm.Config{})
	}


	//dsn := "host=127.0.0.1 user=postgres password=asdfasw22 dbname=staem port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	//return gorm.Open(postgres.Open(dsn), &gorm.Config{})
	//return gorm.Open("postgres","host="+DB_HOST+" port="+DB_PORT+" user="+DB_USER+" dbname="+DB_NAME+" password="+DB_PASS+" sslmode=disable")
}