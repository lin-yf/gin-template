package jwt

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWT struct {
	SigningKey []byte
}

func NewJWT() *JWT {
	return &JWT{
		[]byte("owl"),
	}
}

type MyClaims struct {
	UserId   string
	Username string
	jwt.StandardClaims
}

// 定义过期时间,7天后过期
// 	expirationTime := time.Now().Add(7 * 24 * time.Hour)
// claims := &MyClaims{
//   UserId:   id,
//   Username: name,
//   StandardClaims: jwt.StandardClaims{
//     ExpiresAt: expirationTime.Unix(), // 过期时间
//     IssuedAt:  time.Now().Unix(),     // 发布时间
//     Subject:   "token",               // 主题
//     Issuer:    "owl",                 // 发布者
//   },
// }

// 定义生成token的方法
func (j *JWT) CreateToken(claims *MyClaims) (string, error) {
	// 注意单词别写错了
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// 定义解析token的方法
func (j *JWT) ParseToken(tokenString string) (*jwt.Token, *MyClaims, error) {
	claims := &MyClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			// ValidationErrorMalformed是一个uint常量，表示token不可用
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return token, nil, fmt.Errorf("token不可用")
				// ValidationErrorExpired表示Token过期
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return token, nil, fmt.Errorf("token过期")
				// ValidationErrorNotValidYet表示无效token
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return token, nil, fmt.Errorf("无效的token")
			} else {
				return token, nil, fmt.Errorf("token不可用")
			}
		}
		// 将token中的claims信息解析出来并断言成用户自定义的有效载荷结构
		if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
			return token, claims, nil
		}
	}
	return token, claims, err
}

func GenerateToken(username string) (string, error) {
	// 构造SignKey: 签名和解签名需要使用一个值
	j := NewJWT()
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &MyClaims{
		UserId:   "1",
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(), // 过期时间
			IssuedAt:  time.Now().Unix(),     // 发布时间
			Subject:   "token",               // 主题
			Issuer:    "owl",                 // 发布者
		},
	}
	// 根据claims生成token对象
	token, err := j.CreateToken(claims)
	return token, err
}
