package handler

import (
	"github.com/gavv/httpexpect/v2"
	"github.com/gin-gonic/gin"
	"github.com/simonkimi/minebangumi/api"
	"github.com/simonkimi/minebangumi/internal/app/config"
	"github.com/simonkimi/minebangumi/internal/pkg/middleware"
	"github.com/simonkimi/minebangumi/pkg/hash"
	"github.com/simonkimi/minebangumi/pkg/testutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func initMockConfig(token string) config.Config {
	conf := config.NewMockConfig()
	conf.SetString(config.UserApiToken, token)
	return conf
}

func TestJwtAuthMiddleware_Ok(t *testing.T) {
	token := hash.GenerateRandomKey(40)
	conf := initMockConfig(token)
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.Use(TokenAuthMiddleware(conf))
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"is_login": IsLogin(c),
		})
	})

	server := httptest.NewServer(r)
	defer server.Close()

	httpexpect.Default(t, server.URL).
		GET("/test").
		WithHeader("Authorization", "Token "+token).
		Expect().
		Status(200).
		JSON().Object().
		HasValue("is_login", true)
}

func TestJwtAuthMiddleware_NoToken(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	token := hash.GenerateRandomKey(40)
	conf := initMockConfig(token)
	r.Use(TokenAuthMiddleware(conf))
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"is_login": IsLogin(c),
		})
	})

	server := httptest.NewServer(r)
	defer server.Close()

	httpexpect.Default(t, server.URL).
		GET("/test").
		Expect().
		Status(200).
		JSON().Object().
		HasValue("is_login", false)
}

func TestJwtAuthMiddleware_TokenError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	token := hash.GenerateRandomKey(40)
	conf := initMockConfig(token)
	r.Use(TokenAuthMiddleware(conf))
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"is_login": IsLogin(c),
		})
	})

	server := httptest.NewServer(r)
	defer server.Close()

	httpexpect.Default(t, server.URL).
		GET("/test").
		WithHeader("Authorization", "invalid-token").
		Expect().
		Status(200).
		JSON().Object().
		HasValue("is_login", false)
}

func TestRequireAuthMiddleware_Ok(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	token := hash.GenerateRandomKey(40)
	conf := initMockConfig(token)
	group := r.Group("/api")
	group.Use(TokenAuthMiddleware(conf))
	group.Use(RequireAuthMiddleware())
	group.GET("/test", func(c *gin.Context) {
		api.OkResponse(c, &gin.H{
			"key": "value",
		})
	})

	server := httptest.NewServer(r)
	defer server.Close()

	httpexpect.Default(t, server.URL).
		GET("/api/test").
		WithHeader("Authorization", "Token "+token).
		Expect().
		Status(http.StatusOK).
		JSON().Object().
		HasValue("code", api.APIStatusEnumSuccess).
		Value("data").Object().HasValue("key", "value")
}

func TestRequireAuthMiddleware_Error(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	token := hash.GenerateRandomKey(40)
	conf := initMockConfig(token)
	group := r.Group("/api")
	group.Use(TokenAuthMiddleware(conf))
	group.Use(middleware.ResponseWrapperMiddleware())
	group.Use(RequireAuthMiddleware())
	group.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": "ok",
		})
	})

	server := httptest.NewServer(r)
	defer server.Close()

	testutil.NewDebugHttp(t, server.URL).
		GET("/api/test").
		Expect().
		Status(http.StatusOK).
		JSON().Object().
		HasValue("code", api.APIStatusEnumUnauthorized).
		HasValue("message", "Unauthorized")
}
