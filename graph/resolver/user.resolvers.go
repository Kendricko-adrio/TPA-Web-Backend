package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/kendricko-adrio/gqlgen-todos/database"
	"github.com/kendricko-adrio/gqlgen-todos/graph/jwt"
	"github.com/kendricko-adrio/gqlgen-todos/graph/mailjet"
	"github.com/kendricko-adrio/gqlgen-todos/graph/model"
	"github.com/kendricko-adrio/gqlgen-todos/middleware"
	"gorm.io/gorm/clause"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	close, err := db.DB()
	defer close.Close()
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
	close, err := db.DB()
	defer close.Close()
	var checkUser model.User

	test := db.Where("email = ?", input.Email).First(&checkUser)

	if test.RowsAffected != 0 {
		return &checkUser, nil
	}

	otp := model.CreateOTP()
	user := model.User{
		Email:      input.Email,
		CountryID:  input.CountryID,
		StatusID:   1,
		Otp:        otp,
		ProviderID: 2,
	}

	mailjet.SendEmail(input.Email, otp)

	db.Create(&user)
	fmt.Println(otp)
	return &user, nil
}

func (r *mutationResolver) CreateFullAccount(ctx context.Context, input model.FinalizeAccount) (string, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	close, err := db.DB()
	defer close.Close()
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
	findUser.CustomURL = input.UserName

	db.Save(&findUser)

	token, _ := jwt.GenerateToken(findUser.UserName)

	//http.Set

	return token, nil
}

func (r *mutationResolver) LoginUser(ctx context.Context, input model.Login) (*model.User, error) {
	user, err := model.ValidateUser(input.Username, input.Password)
	if err != nil {
		return nil, err
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

	return &user, nil
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
	fmt.Println("test")
	http.SetCookie(w, c)

	return user, nil
}

func (r *mutationResolver) UpdateGeneral(ctx context.Context, profileName string, realName string, customURL string, countryID int, summary string, hideAward bool) (*model.User, error) {
	//panic(fmt.Errorf("not implemented"))
	user := middleware.ForContext(ctx)
	if user == nil {
		return &model.User{}, fmt.Errorf("access denied")
	}
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	close, err := db.DB()
	defer close.Close()
	var userFind model.User
	test := db.Find(user).First(&userFind)

	if test.RowsAffected == 0 {
		panic(test.Error)
	}

	userFind.UserName = profileName
	userFind.RealName = realName
	userFind.CustomURL = customURL
	userFind.CountryID = countryID
	userFind.HideAward = hideAward
	userFind.Summary = summary
	db.Save(&userFind)
	token, _ := jwt.GenerateToken(profileName)

	c := http.Cookie{
		Name:     "staem",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HttpOnly: true,
	}

	w := *middleware.ForWriter(ctx)
	http.SetCookie(w, &c)
	return &userFind, nil
}

func (r *mutationResolver) UpdateImage(ctx context.Context, imageURL string) (*model.User, error) {
	user := middleware.ForContext(ctx)
	if user == nil {
		return &model.User{}, fmt.Errorf("access denied")
	}
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	close, err := db.DB()
	defer close.Close()
	var userFind model.User
	test := db.Find(user).First(&userFind)

	if test.RowsAffected == 0 {
		panic(test.Error)
	}

	user.PhotoURL = imageURL

	db.Save(user)
	return user, nil
}

func (r *mutationResolver) AddGameToWishlist(ctx context.Context, gameID int) (*model.User, error) {
	user := middleware.ForContext(ctx)
	if user == nil {
		return &model.User{}, fmt.Errorf("access denied")
	}
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}
	close, err := db.DB()
	defer close.Close()
	db.First(&user)

	game := &model.Game{ID: gameID}

	db.First(game)

	user.Wishlist = append(user.Wishlist, game)

	db.Save(user)

	promo := model.Promo{}

	test := db.Preload(clause.Associations).Where("game_id = ?", gameID).First(&promo)

	if test.RowsAffected != 0 {
		//var userPromo
		//db.First()
		fmt.Println("masuk email", user.Email, promo.Game.Name)
		//mailjet.OnSale(user.Email, promo.Game.Name)
	}

	return user, nil
}

func (r *mutationResolver) DeleteWishlist(ctx context.Context, gameID int) (*model.User, error) {
	//panic(fmt.Errorf("not implemented"))
	user := middleware.ForContext(ctx)
	if user == nil {
		return &model.User{}, fmt.Errorf("access denied")
	}
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	close, err := db.DB()
	defer close.Close()
	db.Exec("DELETE FROM games_wishlist WHERE game_id = ? AND user_user_id = ?", gameID, user.UserID)

	return user, nil
}

func (r *mutationResolver) SetCurrBadge(ctx context.Context, badgeID int) (*model.User, error) {
	user := middleware.ForContext(ctx)
	if user == nil {
		return &model.User{}, fmt.Errorf("access denied")
	}
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}
	close, err := db.DB()
	defer close.Close()
	db.First(&user)

	user.CurrBadgeID = badgeID
	db.Save(&user)

	return user, err
}

func (r *mutationResolver) SetCurrMiniBg(ctx context.Context, imageURL *string) (*model.User, error) {
	user := middleware.ForContext(ctx)
	if user == nil {
		return &model.User{}, fmt.Errorf("access denied")
	}
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	close, err := db.DB()
	defer close.Close()
	db.First(&user)

	user.CurrMiniBgImg = *imageURL

	db.Save(&user)
	return user, nil
}

func (r *mutationResolver) SetCurrTheme(ctx context.Context, themeID *int) (*model.User, error) {
	user := middleware.ForContext(ctx)
	if user == nil {
		return &model.User{}, fmt.Errorf("access denied")
	}
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	close, err := db.DB()
	defer close.Close()

	db.Find(&user)

	user.CurrThemeID = *themeID
	db.Save(&user)

	return user, nil
}

func (r *mutationResolver) SetCurrProfileBackground(ctx context.Context, backgroundIt int) (*model.User, error) {
	user := middleware.ForContext(ctx)
	if user == nil {
		return &model.User{}, fmt.Errorf("access denied")
	}
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	close, err := db.DB()
	defer close.Close()

	db.First(&user)
	user.CurrProfileBackgroundID = backgroundIt
	db.Save(&user)
	return user, nil
}

func (r *queryResolver) GetUsers(ctx context.Context) ([]*model.User, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	close, err := db.DB()
	defer close.Close()
	var user []*model.User

	db.Preload("Country").Preload(clause.Associations).Find(&user).Debug()
	fmt.Print(&user)

	return user, err
}

func (r *queryResolver) GetUser(ctx context.Context, input *model.Email) (*model.User, error) {
	// not email sebenernya, ngetest username
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	close, err := db.DB()
	defer close.Close()
	var user model.User

	db.Where("user_name = ?", input.Email).Preload(clause.Associations).First(&user)

	return &user, nil
}

func (r *queryResolver) GetAuthUser(ctx context.Context) (*model.User, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	close, err := db.DB()
	defer close.Close()
	user := middleware.ForContext(ctx)
	if user == nil {
		return &model.User{}, fmt.Errorf("access denied")
	}

	var findUser model.User
	db.Where("user_id = ?", user.UserID).Preload(clause.Associations).First(&findUser)

	return &findUser, err
}

func (r *queryResolver) GetUserByLink(ctx context.Context, customURL string) (*model.User, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	close, err := db.DB()
	defer close.Close()
	var user model.User

	db.Where("custom_url = ?", customURL).Preload("CurrTheme").Preload("Friends.OwnBadge").Preload("Friends.CurrBadge").Preload("Games").Preload("OwnBadge").First(&user)

	return &user, nil
}

func (r *queryResolver) GetAdmin(ctx context.Context, username string, password string) (*model.User, error) {
	user, err := model.ValidateUser(username, password)
	if err != nil {
		return nil, err
	}

	if user.Role != "admin" {
		panic("user is not admin")
	}

	token, _ := jwt.GenerateToken(username)

	c := http.Cookie{
		Name:     "staem",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HttpOnly: true,
	}

	w := *middleware.ForWriter(ctx)
	http.SetCookie(w, &c)

	return &user, nil
}

func (r *queryResolver) GetAllUserPaginated(ctx context.Context, page int) ([]*model.User, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}
	close, err := db.DB()
	defer close.Close()
	var users []*model.User
	db.Limit(10).Offset((page - 1) * 10).Find(&users)
	return users, nil
}

func (r *queryResolver) GetTotalUser(ctx context.Context) (int, error) {
	db, err := database.Connect()
	if err != nil {
		return 0, err
	}
	var user []*model.User
	test := db.Find(&user)
	close, err := db.DB()
	defer close.Close()
	return int(test.RowsAffected), nil
}

func (r *queryResolver) GetAllBadge(ctx context.Context) ([]*model.Badge, error) {
	panic(fmt.Errorf("not implemented"))
}
