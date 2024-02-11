package model

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"

	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
)

var conn *redis.Client

func init() {
	conn = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
}

func NewSession(w http.ResponseWriter, cookieKey, redisValue string) {
	b := make([]byte, 64)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		panic("ランダムな文字作成時にエラーが発生しました。")
	}
	newRedisKey := base64.URLEncoding.EncodeToString(b)

	ctx := context.Background()
	if err := conn.Set(ctx, newRedisKey, redisValue, 0).Err(); err != nil {
		panic("Session登録時にエラーが発生：" + err.Error())
	}
	http.SetCookie(w, &http.Cookie{
		Name:     cookieKey,
		Value:    newRedisKey,
		Path:     "/",
		Domain:   "localhost",
		Secure:   false,
		HttpOnly: false,
	})
}

func GetSession(r *http.Request, cookieKey string) interface{} {
	cookie, err := r.Cookie(cookieKey)
	if err != nil {
		fmt.Println("Cookieが見つかりません。")
		return nil
	}
	redisKey := cookie.Value

	ctx := context.Background()
	redisValue, err := conn.Get(ctx, redisKey).Result()
	switch {
	case err == redis.Nil:
		fmt.Println("SessionKeyが登録されていません。")
		return nil
	case err != nil:
		fmt.Println("Session取得時にエラー発生：" + err.Error())
		return nil
	}
	return redisValue
}

func DeleteSession(w http.ResponseWriter, r *http.Request, cookieKey string) {
	cookie, err := r.Cookie(cookieKey)
	if err != nil {
		fmt.Println("Cookieが見つかりません。")
		return
	}
	redisKey := cookie.Value

	ctx := context.Background()
	conn.Del(ctx, redisKey)
	http.SetCookie(w, &http.Cookie{
		Name:   cookieKey,
		Value:  "",
		Path:   "/",
		Domain: "localhost",
		MaxAge: -1,
	})
}
