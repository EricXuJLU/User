package main

import (
	"context"
	"encoding/json"
	"golang.org/x/oauth2"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	address     = "localhost:8080"
	state       = "xxn" //oauth 2.0 的验证信息
	uriUserInfo = "https://api.github.com/user"
)

var (
	config = oauth2.Config{
		ClientID:     "d4762f4329a6d78159cf",
		ClientSecret: "7569c0e6bd9c23c9a13dc7aeec9ec99410d30562",
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://github.com/login/oauth/authorize",
			TokenURL: "https://github.com/login/oauth/access_token",
		},
		RedirectURL: "http://localhost:8080/oauth/redirect",
		Scopes:      []string{"user"},
	}
	globalToken   *oauth2.Token
	tokenLifeTime = time.Minute * 30
)

func main() {
	mux := createService()
	err := http.ListenAndServe(address, mux)
	if err != nil {
		log.Fatal(err)
	}
}

func createService() (mux *http.ServeMux) {
	mux = http.NewServeMux()

	//访问主页时，重定向到github登录界面
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		url := config.AuthCodeURL(state)
		http.Redirect(writer, request, url, http.StatusFound)
	})

	//获取到code，去交换access_token
	mux.HandleFunc("/oauth/redirect", func(writer http.ResponseWriter, request *http.Request) {
		request.ParseForm()
		stateGet := request.Form.Get("state")
		if stateGet != state { //校验有误
			http.Error(writer, "Invalid State", http.StatusBadRequest)
			return
		}
		code := request.Form.Get("code")
		if code == "" {
			http.Error(writer, "Code Not Find", http.StatusBadRequest)
			return
		}
		token, err := config.Exchange(context.Background(), code)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		token.Expiry = time.Now().Add(tokenLifeTime)
		globalToken = token

		e := json.NewEncoder(writer)
		e.SetIndent("", "  ")
		e.Encode(token)
	})

	//刷新token
	mux.HandleFunc("/refresh", func(writer http.ResponseWriter, request *http.Request) {
		//原有token为空，重新登录
		if globalToken == nil {
			http.Redirect(writer, request, "/", http.StatusFound)
			return
		}
		globalToken.Expiry = time.Now().Add(tokenLifeTime)
		token, err := config.TokenSource(context.Background(), globalToken).Token()
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		//token.Expiry = time.Now().Add(time.Minute*30)
		globalToken = token

		e := json.NewEncoder(writer)
		e.SetIndent("", "")
		e.Encode(token)
	})

	//尝试获取用户身份信息
	mux.HandleFunc("/user", func(writer http.ResponseWriter, request *http.Request) {
		if globalToken == nil {
			http.Redirect(writer, request, "/", http.StatusFound)
			return
		}
		req, _ := http.NewRequest("GET", uriUserInfo, nil)
		req.Header.Add("Authorization", "Bearer "+globalToken.AccessToken)
		client := http.Client{}
		resp, err := client.Do(req)
		defer resp.Body.Close()
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}
		var keyValue = make(map[string]interface{})
		err = json.Unmarshal(data, &keyValue)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}

		e := json.NewEncoder(writer)
		e.SetIndent("", "")
		e.Encode(keyValue)
	})

	return
}
