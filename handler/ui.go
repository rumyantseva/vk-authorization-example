package handler

import (
	"errors"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"
	"github.com/rumyantseva/go-vk/vk"
	"github.com/rumyantseva/vk-authorization-example/services"
	"net/http"
)

func (h *Handler) Hello(c echo.Context) error {
	data := struct {
		ClientID    int
		RedirectURI string
		Version     string
	}{
		h.VK.ClientID,
		h.VK.OAuthRedirectURI,
		h.VK.Version,
	}

	return c.Render(http.StatusOK, "hello", data)
}

func (h *Handler) Verify(c echo.Context) error {
	code := c.QueryParam("code")

	client := vk.NewOAuthClient(nil, h.VK.ClientID, h.VK.ClientSecret)
	aToken, _, err := client.OAuth.AccessToken(h.VK.OAuthRedirectURI, code)
	if err != nil {
		return err
	}

	if aToken.Email == nil {
		return errors.New("E-mail mustn't be empty.")
	}

	regService := services.NewRegistration(h.DB.Querier)
	id, pass, err := regService.Register(aToken.UserID, *aToken.Email, aToken.AccessToken, aToken.ExpiresIn)
	if err != nil {
		return err
	}

	c.Logger().Debugf("ID: %d, pass: %s", id, pass)

	data := struct {
		Login int
		Pass  string
	}{
		id,
		pass,
	}

	return c.Render(http.StatusOK, "verify", data)
}
