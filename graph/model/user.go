package model

import (
	"fmt"
	"github.com/kendricko-adrio/gqlgen-todos/database"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm/clause"

	//"fmt"
	"math/rand"
	"time"
)

type User struct {
	UserID                  int                  `json:"userID" gorm:"primaryKey"`
	Role                    string               `json:"role"`
	UserName                string               `json:"userName"`
	Password                string               `json:"password"`
	FirstName               string               `json:"FirstName"`
	LastName                string               `json:"LastName"`
	RealName                string               `json:"realName"`
	CustomURL               string               `json:"customUrl"`
	HideAward               bool                 `json:"hideAward"`
	Money                   int                  `json:"money"`
	Level                   int                  `json:"level"`
	Summary                 string               `json:"summary"`
	AuthToken               string               `json:"AuthToken"`
	Email                   string               `json:"Email"`
	IDToken                 string               `json:"IdToken"`
	PhotoURL                string               `json:"PhotoUrl"`
	Provider                *Provider            `json:"Provider" gorm:"foreignKey:ProviderID"`
	ProviderID              int                  `json:"ProviderId" `
	CountryID               int                  `json:"CountryId"`
	Country                 *Country             `json:"Country" gorm:"foreignKey:CountryID"`
	Status                  *Status              `json:"status" gorm:"foreignKey:StatusID"`
	StatusID                int                  `json:"statusId"`
	Otp                     string               `json:"otp"`
	Friends                 []*User              `json:"friends" gorm:"many2many:friends"`
	OwnBadge                []*Badge             `json:"ownBadge" gorm:"many2many:users_badge"`
	CurrBadge               *Badge               `json:"currBadge" gorm:"foreignKey:CurrBadgeID"`
	CurrBadgeID             int                  `json:"currBadgeId"`
	Games                   []*Game              `json:"games" gorm:"many2many:users_games"`
	Wishlist                []*Game              `json:"wishlist" gorm:"many2many:games_wishlist"`
	OwnMiniBg               []*MiniBackground    `json:"ownMiniBg" gorm:"many2many:users_mini_bg"`
	CurrMiniBgImg           string               `json:"currMiniBgImg"`
	OwnTheme                []*Theme             `json:"ownTheme" gorm:"many2many:users_theme"`
	CurrTheme               *Theme               `json:"currTheme"  gorm:"foreignKey:"CurrThemeID"`
	CurrThemeID             int                  `json:"currThemeId"`
	OwnProfileBackground    []*ProfileBackground `json:"ownProfileBackground" gorm:"many2many:users_profile_bg"`
	CurrProfileBackground   *ProfileBackground   `json:"currProfileBackground" gorm:"foreignKey:CurrProfileBackgroundID"`
	CurrProfileBackgroundID int                  `json:"currProfileBackgroundId"`
	Post                    []*Post              `json:"post" gorm:"foreignKey:UserID"`
	LikeDetail              []*LikeDetail        `json:"likeDetail" gorm:"foreignKey:UserID"`
	CommandDetail           []*CommandDetail     `json:"commandDetail" gorm:"foreignKey:UserID"`
	Items                   []*Item              `json:"items" gorm:"many2many:users_item"`
	SellItem                []*ItemTransaction   `json:"sellItem" gorm:"foreignKey:UserSellerID"`
	BuyItem                 []*ItemTransaction   `json:"buyItem" gorm:"foreignKey:UserBuyerID"`
	OwnFrame                []*AvatarFrame       `json:"ownFrame" gorm:"many2many:users_frame"`
	CurrFrame               string               `json:"currFrame"`
	Point                   int                  `json:"point"`
	MyItem                  []*MyItem            `json:"myItem" gorm:"foreignKey:UserID"`
	OwnAvatar               []*Avatar            `json:"ownAvatar" gorm:"many2many:users_avatar"`
	CurrAvatar              string               `json:"currAvatar"`
	CreatedAt               time.Time            `json:"CreatedAt"`
	UpdatedAt               time.Time            `json:"UpdatedAt"`
	DeletedAt               *time.Time           `json:"DeletedAt"`
}

type NewUser struct {
	FirstName  string `json:"FirstName"`
	LastName   string `json:"LastName"`
	AuthToken  string `json:"AuthToken"`
	Email      string `json:"Email"`
	IDToken    string `json:"IdToken"`
	PhotoURL   string `json:"PhotoUrl"`
	ProviderID int    `json:"ProviderId"`
}

//func main(){
//
//	//fmt.Print(rand.Intn(100))
//	//for i := 0; i<100; i++ {
//	//	fmt.Println(rand.Intn(10))
//	//}
//	fmt.Println(CreateOTP())
//}

func CreateOTP() string {
	charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var s string
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for i := 0; i < 5; i++ {
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
	close, err := db.DB()
	defer close.Close()
	var user User
	test := db.Where("user_name = ?", username).First(&user)
	if test.RowsAffected == 0 {
		return 0, err
	}

	return user.UserID, nil
}

func ValidateUser(username, password string) (User, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	close, err := db.DB()
	defer close.Close()
	var user User
	test := db.Where("user_name = ?", username).Preload(clause.Associations).First(&user)
	if test.RowsAffected == 0 {
		return user, test.Error
	} else if CheckPasswordHash(password, user.Password) == false {
		fmt.Println(CheckPasswordHash(password, user.Password))
		return user, test.Error
	}

	return user, nil
}
