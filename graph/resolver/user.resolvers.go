package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/kendricko-adrio/gqlgen-todos/database"
	"github.com/kendricko-adrio/gqlgen-todos/graph/generated"
	"github.com/kendricko-adrio/gqlgen-todos/graph/jwt"
	"github.com/kendricko-adrio/gqlgen-todos/graph/model"
	"github.com/kendricko-adrio/gqlgen-todos/middleware"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	//print("test")

	var checkUser model.User

	test := db.Where("email = ?", input.Email).First(&checkUser)

	if test.RowsAffected != 0 {
		return &checkUser, nil
	}

	user := model.User{
		FirstName:  input.FirstName,
		LastName:   input.LastName,
		AuthToken:  input.AuthToken,
		Email:      input.Email,
		IDToken:    input.IDToken,
		PhotoURL:   input.PhotoURL,
		ProviderID: input.ProviderID,
		StatusID:   1,
	}

	//db.Select("first_name", "last_name", "auth_token", "email", "id_token",
	//	"photo_url", "provider", "created_at").Create(&user)
	db.Create(&user)

	return &user, nil
}

func (r *mutationResolver) CreateAccount(ctx context.Context, input model.CreateAccount) (*model.User, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	var checkUser model.User

	test := db.Where("email = ?", input.Email).First(&checkUser)

	if test.RowsAffected != 0 {
		return &checkUser, nil
	}

	otp := model.CreateOTP()
	user := model.User{
		Email:     input.Email,
		CountryID: input.CountryID,
		StatusID:  1,
		Otp:       otp,
	}

	//mailjet.SendEmail(input.Email, otp)

	db.Create(&user)
	fmt.Println(otp)
	return &user, nil
}

func (r *mutationResolver) CreateFullAccount(ctx context.Context, input model.FinalizeAccount) (string, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	fmt.Println("test")
	var findUser model.User
	test := db.Where("otp = ? AND status_id = ?", input.Otp, 1).First(&findUser)

	if test.RowsAffected == 0 {
		return "", nil
	}
	pass, _ := model.HashPassword(input.Password)
	findUser.StatusID = 2
	findUser.UserName = input.UserName
	findUser.Password = pass

	db.Save(&findUser)

	token, _ := jwt.GenerateToken(findUser.UserName)

	//http.Set

	return token, nil
}

func (r *mutationResolver) LoginUser(ctx context.Context, input model.Login) (string, error) {
	isValid := model.ValidateUser(input.Username, input.Password)

	if isValid == false {
		return "", nil
	}
	token, _ := jwt.GenerateToken(input.Username)

	c := http.Cookie{
		Name:     "staem",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HttpOnly: true,
	}

	w := *middleware.ForWriter(ctx)
	http.SetCookie(w, &c)
	//w.Header().Add("Set-Cookie", c.String())
	//http.SetCookie(*w, &c)

	return token, nil
}

func (r *mutationResolver) LogoutUser(ctx context.Context, username string) (*model.User, error) {
	user := middleware.ForContext(ctx)
	if user == nil {
		return &model.User{}, fmt.Errorf("access denied")
	}
	w := *middleware.ForWriter(ctx)

	c := &http.Cookie{
		Name:     "staem",
		Value:    "",
		Expires:  time.Now().Add(time.Hour * -24),
		HttpOnly: true,
	}

	http.SetCookie(w, c)

	return user, nil
}

func (r *queryResolver) GetUsers(ctx context.Context) ([]*model.User, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	var user []*model.User

	db.Preload("Country").Preload("Games").Find(&user).Debug()
	fmt.Print(&user)

	return user, err
}

func (r *queryResolver) GetUser(ctx context.Context, input *model.Email) (*model.User, error) {
	// not email sebenernya, ngetest username
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	var user model.User

	db.Where("user_name = ?", input.Email).Preload("Games").First(&user)

	return &user, nil
}

func (r *queryResolver) GetAuthUser(ctx context.Context) (*model.User, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	user := middleware.ForContext(ctx)
	if user == nil {
		return &model.User{}, fmt.Errorf("access denied")
	}

	var findUser model.User
	db.Where("id = ?", user.ID).Preload("Games").First(&findUser)

	return &findUser, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
