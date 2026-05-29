package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"pipe-api/config"
	"github.com/google/uuid"
)

type YandexTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

type YandexUserInfo struct {
	ID           string   `json:"id"`
	Login        string   `json:"login"`
	DefaultEmail string   `json:"default_email"`
	Emails       []string `json:"emails"`
}

func GetYandexAuthURL() (string, string) {
	state := uuid.New().String()
	authURL := fmt.Sprintf(
		"https://oauth.yandex.ru/authorize?response_type=code&client_id=%s&redirect_uri=%s&state=%s",
		config.AppConfig.YandexClientID,
		url.QueryEscape(config.AppConfig.YandexRedirectURI),
		state,
	)
	return authURL, state
}

func ExchangeYandexCode(code string) (*YandexTokenResponse, error) {
	data := url.Values{
		"grant_type":    {"authorization_code"},
		"code":          {code},
		"client_id":     {config.AppConfig.YandexClientID},
		"client_secret": {config.AppConfig.YandexClientSecret},
		"redirect_uri":  {config.AppConfig.YandexRedirectURI},
	}
	resp, err := http.PostForm("https://oauth.yandex.ru/token", data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	var tokenResp YandexTokenResponse
	if err := json.Unmarshal(body, &tokenResp); err != nil {
		return nil, err
	}
	return &tokenResp, nil
}

func GetYandexUserInfo(accessToken string) (*YandexUserInfo, error) {
	req, _ := http.NewRequest("GET", "https://login.yandex.ru/info?format=json", nil)
	req.Header.Set("Authorization", "OAuth "+accessToken)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	var userInfo YandexUserInfo
	if err := json.Unmarshal(body, &userInfo); err != nil {
		return nil, err
	}
	return &userInfo, nil
}