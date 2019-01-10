package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"

	"github.com/negaihoshi/daigou/pkg/app"
	"github.com/negaihoshi/daigou/pkg/e"
	"github.com/negaihoshi/daigou/pkg/util"
	"github.com/negaihoshi/daigou/service/auth_service"
	"github.com/negaihoshi/daigou/service/user_service"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
	"github.com/markbates/goth/providers/facebook"
	"github.com/negaihoshi/daigou/pkg/setting"

	"github.com/negaihoshi/daigou/models"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

type User struct {
	Username string `valid:"Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	username := c.PostForm("username")
	password := c.PostForm("password")

	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)

	if !ok {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	authService := auth_service.Auth{Username: username, Password: password}
	isExist, err := authService.Check()
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		return
	}

	if !isExist {
		appG.Response(http.StatusOK, e.ERROR_AUTH, nil)
		return
	}

	token, err := util.GenerateToken()
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_AUTH_TOKEN, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"token": token,
	})
}

func GetGoogle(c *gin.Context) {
	key := ""             // Replace with your SESSION_SECRET or similar
	maxAge := 86400 * 30  // 30 days
	isProd := false       // Set to true when serving over https

	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(maxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = true   // HttpOnly should always be enabled
	store.Options.Secure = isProd

	gothic.Store = store
fmt.Println(setting.AppConfig.Social.FacebookClientKey)
	facebookProvider := facebook.New(setting.AppConfig.Social.FacebookClientKey, setting.AppConfig.Social.FacebookClientSecret, setting.AppConfig.Social.FacebookClientCallbackURL)
	googleProvider := google.New(setting.AppConfig.Social.GoogleClientKey, setting.AppConfig.Social.GoogleClientSecret, setting.AppConfig.Social.GoogleClientCallbackURL)
	// googleProvider.SetPrompt("select_account")

	goth.UseProviders(
		// google.New(os.Getenv("GOOGLE_KEY"), os.Getenv("GOOGLE_SECRET"), "http://localhost:3000/auth/google/callback"),
		googleProvider,
		facebookProvider,
	)

	provider := c.Param("provider")

	q := c.Request.URL.Query()
	q.Add("provider", provider)
	c.Request.URL.RawQuery = q.Encode()
	gothic.BeginAuthHandler(c.Writer, c.Request)
}

func GetGoogleCallback(c *gin.Context) {
	provider := c.Param("provider")

	q := c.Request.URL.Query()
	q.Add("provider", provider)
	c.Request.URL.RawQuery = q.Encode()
	user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	res, err := json.Marshal(user)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	jsonString := string(res)
	html := fmt.Sprintf(`%v`, jsonString)
	c.Data(http.StatusOK, "application/json; charset=utf-8", []byte(html))

	appG := app.Gin{C: c}

	check, _ := models.CheckProvider(user.Provider, user.UserID)

	if !check {
		userService := user_service.User{
			Username:         user.Name,
			Password: "psjh1019",
		}

		if err := userService.Add(); err != nil {
			appG.Response(http.StatusOK, e.ERROR_ADD_ARTICLE_FAIL, nil)
			return
		}

		userServiceProvider := user_service.ThirdPartyProvider{
			Provider: user.Provider,
			Email: user.Email,
			ProviderID: user.UserID,
			Avatar: user.AvatarURL,
			AccessToken: user.AccessToken,
			AccessTokenSecret: user.AccessTokenSecret,
			UserID: userService.ID,
		}

		if err := userServiceProvider.AddProvider(); err != nil {
			appG.Response(http.StatusOK, e.ERROR_ADD_ARTICLE_FAIL, nil)
			return
		}
	}

	// appG := app.Gin{C: c}
	// token, err := util.GenerateToken()
	// if err != nil {
	// 	appG.Response(http.StatusOK, e.ERROR_AUTH_TOKEN, nil)
	// 	return
	// }

	// c.JSON(http.StatusOK, gin.H{"token": token, "status": http.StatusOK})
}
