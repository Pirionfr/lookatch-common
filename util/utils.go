package util

import (
	"encoding/hex"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"hash/fnv"
	"sync"
)

func ValidateToken(token string, tokenSig string) (*jwt.Token, error) {
	jwtToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Error("Unexpected signing method: %v", t.Header["alg"])
			return nil, fmt.Errorf("invalid token")
		}
		return []byte(tokenSig), nil
	})
	if err == nil && jwtToken.Valid {
		return jwtToken, nil
	}
	return nil, err
}

func FNV1a(s string) uint64 {
	h := fnv.New64a()
	h.Write([]byte(s))
	return h.Sum64()
}

func IntInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func StringInGenSlice(a string, list []interface{}) bool {
	for _, b := range list {
		if b.(string) == a {
			return true
		}
	}
	return false
}

func WithWaitGroup(wg *sync.WaitGroup, f func()) {
	wg.Add(1)
	defer wg.Done()
	f()
}

func IsStandalone(v *viper.Viper) bool {
	if v.Get("auth") != nil {
		return false
	}
	return true
}

func EscapeCtrl(ctrl []byte) (esc []byte) {
	u := []byte(`\u0000`)
	for i, ch := range ctrl {
		if ch <= 31 {
			if esc == nil {
				esc = append(make([]byte, 0, len(ctrl)+len(u)), ctrl[:i]...)
			}
			esc = append(esc, u...)
			hex.Encode(esc[len(esc)-2:], ctrl[i:i+1])
			continue
		}
		if esc != nil {
			esc = append(esc, ch)
		}
	}
	if esc == nil {
		return ctrl
	}
	return esc
}
