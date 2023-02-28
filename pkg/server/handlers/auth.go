package handlers

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/ravener/discord-oauth2"
	"github.com/xsadia/kgallery/config"
	"github.com/xsadia/kgallery/pkg/repository"
	"github.com/xsadia/kgallery/pkg/storage"
	"golang.org/x/oauth2"
)

type AuthHandler struct {
	OAuthConfig *oauth2.Config
	Storage     *storage.Storage
}

func NewAuthHandler(storage *storage.Storage) *AuthHandler {
	conf := &oauth2.Config{
		RedirectURL:  config.Ctx.Env["OAUTH_REDIRECT_URL"],
		ClientID:     config.Ctx.Env["DISCORD_CLIENT_ID"],
		ClientSecret: config.Ctx.Env["DISCORD_SECRET"],
		Scopes:       []string{discord.ScopeIdentify, discord.ScopeGuilds},
		Endpoint:     discord.Endpoint,
	}

	return &AuthHandler{
		OAuthConfig: conf,
		Storage:     storage,
	}
}

func (a *AuthHandler) Auth(c *fiber.Ctx) error {
	return c.Redirect(a.OAuthConfig.AuthCodeURL("random"), http.StatusTemporaryRedirect)
}

func (a *AuthHandler) Create(c *fiber.Ctx) error {
	token, err := a.OAuthConfig.Exchange(context.Background(), c.FormValue("code"))

	if err != nil {
		log.Printf("[Error]: Unable to exchange code for OAuth token")
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	res, err := a.OAuthConfig.Client(context.Background(), token).Get("https://discord.com/api/users/@me/guilds")

	if err != nil {
		log.Printf("[Error]: Unable to request user's guilds %s", err.Error())
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if res.StatusCode != http.StatusOK {
		log.Printf("[Error]: Unable to request user's guilds")
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "unable to reach auth provider"})
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Printf("[Error]: Unable to read request body %s", err.Error())
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	var gm []map[string]string

	json.Unmarshal(body, &gm)

	if !isUserOnGuild(gm) {
		log.Printf("[Error]: user is not on discriminatory guild")
		return c.Status(http.StatusForbidden).JSON(fiber.Map{"error": "user unauthorized"})
	}

	userRes, err := a.OAuthConfig.Client(context.Background(), token).Get("https://discord.com/api/users/@me")

	if err != nil {
		log.Printf("[Error]: Unable to request user data %s", err.Error())
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if res.StatusCode != http.StatusOK {
		log.Printf("[Error]: Unable to request user data")
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "unable to reach auth provider"})
	}

	defer userRes.Body.Close()

	body, err = ioutil.ReadAll(userRes.Body)

	var um map[string]string

	json.Unmarshal(body, &um)

	u := repository.NewUser(um)

	err = u.Create(a.Storage.SQL)

	if err != nil {
		log.Printf("[Error]: error while creating user '%s'", err.Error())
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "unable create user"})
	}

	return c.Status(http.StatusCreated).JSON(u)
}

func isUserOnGuild(guildArray []map[string]string) bool {
	for _, v := range guildArray {
		if v["name"] == config.Ctx.Env["DISCRIMINATORY_GUILD"] {
			return true
		}
	}

	return false
}
