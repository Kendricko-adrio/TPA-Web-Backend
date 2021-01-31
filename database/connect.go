package database
import(

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const DB_NAME="staem"
const DB_HOST="127.0.0.1"
const DB_PORT="5432"				//port on installation
const DB_USER="postgres"			//default is postgres
const DB_PASS="asdfasw22"	//password on installation

func Connect()(*gorm.DB, error){
	dsn := "host=127.0.0.1 user=postgres password=asdfasw22 dbname=staem port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
	//return gorm.Open("postgres","host="+DB_HOST+" port="+DB_PORT+" user="+DB_USER+" dbname="+DB_NAME+" password="+DB_PASS+" sslmode=disable")
}