package model

import (
	"fmt"
	"github.com/kendricko-adrio/gqlgen-todos/database"
	"golang.org/x/crypto/bcrypt"
	//"fmt"
	"math/rand"
	"time"
)

type NewUser struct {
	FirstName  string `json:"FirstName"`
	LastName   string `json:"LastName"`
	AuthToken  string `json:"AuthToken"`
	Email      string `json:"Email"`
	IDToken    string `json:"IdToken"`
	PhotoURL   string `json:"PhotoUrl"`
	ProviderID int    `json:"ProviderId"`
}


type User struct {
	ID         int        `json:"ID" gorm:"primaryKey"`
	UserName   string     `json:"userName"`
	Password   string     `json:"password"`
	FirstName  string     `json:"FirstName"`
	LastName   string     `json:"LastName"`
	AuthToken  string     `json:"AuthToken"`
	Email      string     `json:"Email"`
	IDToken    string     `json:"IdToken"`
	PhotoURL   string     `json:"PhotoUrl"`
	Provider   *Provider  `json:"Provider"`
	ProviderID int        `json:"ProviderId"`
	CountryID  int        `json:"CountryId"`
	Country    *Country   `json:"Country" gorm:"foreignKey:CountryID"`
	Status     *Status    `json:"status gorm:"foreignKey:StatusID"`
	StatusID   int        `json:"statusId"`
	Otp        string     `json:"otp"`
	Games      []*Game    `json:"games" gorm:"many2many:users_games"`
	CreatedAt  time.Time  `json:"CreatedAt"`
	UpdatedAt  time.Time  `json:"UpdatedAt"`
	DeletedAt  *time.Time `json:"DeletedAt"`
}

//func main(){
//
//	//fmt.Print(rand.Intn(100))
//	//for i := 0; i<100; i++ {
//	//	fmt.Println(rand.Intn(10))
//	//}
//	fmt.Println(CreateOTP())
//}

func CreateOTP() string{
	charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var s string
	r:= rand.New(rand.NewSource(time.Now().Unix()))
	for i := 0; i<5; i++ {
		s += string(charset[r.Intn(36)])
	}

	return s
}

//HashPassword hashes given password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//CheckPassword hash compares raw password with it's hashed values
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GetUserIdByUsername(username string) (int, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	var user User
	test := db.Where("user_name = ?", username).First(&user)
	if test.RowsAffected == 0{
		return 0, err
	}

	return user.ID,nil

}

func ValidateUser(username, password string) bool{
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	var user User
	test :=db.Where("user_name = ?", username).First(&user)
	if test.RowsAffected == 0{
		return false
	}else if CheckPasswordHash(password, user.Password) == false{
		fmt.Println(CheckPasswordHash(password, user.Password))
		return false
	}

	return true
}