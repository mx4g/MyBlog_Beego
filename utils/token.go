package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (

	//密匙
	KEY string = "JWT-ARY-STARK"
	// 过期时间 1 minutes
	DEFAULT_EXPIRE_SECONDS int = 60

)

type TokenUser struct {
	Id string `json:"id"`
	Name string `json:"json"`
}


// JWT -- json web token
// HEADER PAYLOAD SIGNATURE
// This struct is the PAYLOAD
type MyCustomClaims struct {
	TokenUser
	jwt.StandardClaims
}


// 刷新令牌
func RefreshToken(tokenString string)(string, error) {
	//验证要发过来的令牌是否正确
	//先分析令牌
	token, err := jwt.ParseWithClaims(
		tokenString,
		&MyCustomClaims{},
		func(token *jwt.Token)(interface{}, error) {
			return []byte(KEY), nil
	    })
	claims, ok := token.Claims.(*MyCustomClaims)
	if !ok || !token.Valid {
		return "", err
	}
	mySigningKey := []byte(KEY)
	expireAt  := time.Now().Add(time.Second * time.Duration(DEFAULT_EXPIRE_SECONDS)).Unix()
	newClaims := MyCustomClaims{
		claims.TokenUser,
		jwt.StandardClaims{
			ExpiresAt: expireAt,
			Issuer:    claims.TokenUser.Name,
			IssuedAt:  time.Now().Unix(),
		},
	}
	//使用新声明生成新令牌
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims)
	tokenStr, err := newToken.SignedString(mySigningKey)
	if err != nil {
		fmt.Println("生成新的令牌失败 !! error :", err)
		return  "" , err
	}
	return tokenStr, err
}

//验证令牌
func ValidateToken(tokenString string) error {
	//先分析令牌
	token, err := jwt.ParseWithClaims(
		tokenString,
		&MyCustomClaims{},
		func(token *jwt.Token)(interface{}, error) {
			return []byte(KEY), nil
		})


	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {

		fmt.Printf("%v %v", claims.TokenUser, claims.StandardClaims.ExpiresAt)
		fmt.Println("令牌将于过期", time.Unix(claims.StandardClaims.ExpiresAt, 0))
	} else {
		fmt.Println("验证令牌字符串失败 !!!",err)
		return err
	}
	return nil
}


//生成令牌
func GenerateToken(id,username string) (tokenString string) {

   expiredSeconds := DEFAULT_EXPIRE_SECONDS

	// 创建 Claims
	mySigningKey := []byte(KEY)
	expireAt  := time.Now().Add(time.Second * time.Duration(expiredSeconds)).Unix()
	fmt.Println("令牌将于过期 ", time.Unix(expireAt, 0) )
	user := TokenUser{id, username}
	//将参数传递到此函数
	claims := MyCustomClaims{
		user,
		jwt.StandardClaims{
			//过期时间
			ExpiresAt: expireAt,
			//发行人
			Issuer:    user.Name,
			//发行日期
			IssuedAt:  time.Now().Unix(),
		},
	}
	//生成token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//加密
	tokenStr, err := token.SignedString(mySigningKey)
	if err != nil {
		fmt.Println("生成JSON Web令牌失败 !! error :", err)
	}
	return tokenStr

}

// 将此结果返回到客户端，然后所有以后的请求都应具有头 "Authorization: Bearer <token> "
func getHeaderTokenValue(tokenString string) string {
	//Authorization: Bearer <token>
	return fmt.Sprintf("Bearer %s", tokenString)
}




