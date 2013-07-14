/*
新浪微博登录
新浪微博SDK golang实现（目前只完成OAuth认证以及发送文字微博以及带图片的微博）
*/

package sina

import (
	"encoding/json"
	"net/http"
	"net/url"
)

/*
Sina API author
*/
type SaeTOAuth struct {
	clientId     string
	clientSecret string
}

type AccessToken struct {
	Access_Token string `json:access_token`
	Remind_In    string `json:remind_in`
	Expires_In   int    `json:expires_in`
	Uid          string `json:uid`
}

const (
	SINA_OAUTH_API_URL = "https://api.weibo.com/oauth2"
)

func NewSaeTOAuth(clientID, clientSecret string) *SaeTOAuth {
	return &SaeTOAuth{clientID, clientSecret}
}

func (s *SaeTOAuth) GetAuthorizeURL(redirect_url, response_type, state, display string) string {
	if response_type == "" {
		response_type = "code"
	}
	if display == "" {
		display = "default"
	}
	v := url.Values{}
	v.Add("client_id", s.clientId)
	v.Add("redirect_uri", redirect_url)
	v.Add("response_type", response_type)
	if state != "" {
		v.Add("state", state)
	}
	v.Add("display", display)

	params := v.Encode()

	return SINA_OAUTH_API_URL + "/authorize?" + params
}

func (s *SaeTOAuth) GetAccessToken(grant_type string, keys map[string]string) (AccessToken, error) {
	v := url.Values{}
	v.Add("client_id", s.clientId)
	v.Add("client_secret", s.clientSecret)
	if grant_type == "code" {
		v.Add("grant_type", "authorization_code")
		v.Add("code", keys["code"])
		v.Add("redirect_uri", keys["redirect_uri"])
	} else if grant_type == "token" {
		v.Add("grant_type", "refresh_token")
		v.Add("refresh_token", keys["refresh_token"])
	} else if grant_type == "password" {
		v.Add("grant_type", "password")
		v.Add("username", keys["username"])
		v.Add("password", keys["password"])
	}

	accessTokenUrl := SINA_OAUTH_API_URL + "/access_token"

	response, err := http.PostForm(accessTokenUrl, v)

	if err != nil {
		return AccessToken{}, err
	}
	defer response.Body.Close()
	jsonMap := AccessToken{}
	json.NewDecoder(response.Body).Decode(&jsonMap)
	return jsonMap, nil

}
