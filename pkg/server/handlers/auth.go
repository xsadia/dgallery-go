package handlers

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/ravener/discord-oauth2"
	"github.com/xsadia/kgallery/config"
	"golang.org/x/oauth2"
)

type AuthHandler struct {
	OAuthConfig *oauth2.Config
}

func NewAuthHandler() *AuthHandler {
	conf := &oauth2.Config{
		RedirectURL:  config.Ctx.Env["OAUTH_REDIRECT_URL"],
		ClientID:     config.Ctx.Env["DISCORD_CLIENT_ID"],
		ClientSecret: config.Ctx.Env["DISCORD_SECRET"],
		Scopes:       []string{discord.ScopeIdentify, discord.ScopeGuilds},
		Endpoint:     discord.Endpoint,
	}

	return &AuthHandler{
		OAuthConfig: conf,
	}
}

func (a *AuthHandler) Auth(c *fiber.Ctx) error {
	return c.Redirect(a.OAuthConfig.AuthCodeURL("random"), http.StatusTemporaryRedirect)
}

func (a *AuthHandler) Create(c *fiber.Ctx) error {
	fmt.Println(c.FormValue("state"))

	token, err := a.OAuthConfig.Exchange(context.Background(), c.FormValue("code"))

	fmt.Println(token)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	res, err := a.OAuthConfig.Client(context.Background(), token).Get("https://discord.com/api/users/@me/guilds")

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if res.StatusCode != http.StatusOK {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "unable to reach auth provider"})
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	fmt.Println(string(body))

	return c.JSON(fiber.Map{"ok": true})
}
