package controller

import (
	"context"
	"store-bpel/account_service/internal/util"
	"store-bpel/account_service/schema"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func (c *accountServiceController) SignIn(ctx context.Context, request *schema.SignInRequest) (*schema.SignInResponseData, error) {
	account, err := c.repository.GetAccount(ctx, request.Username)
	if err != nil {
		return nil, err
	}
	err = util.CheckPasswordBcrypt([]byte(account.Password), []byte(request.Password))
	if err != nil {
		return nil, err
	}

	jwtToken, err := c.generateJwtToken(request.Username, account.UserRole)
	if err != nil {
		return nil, err
	}

	return &schema.SignInResponseData{
		UserId: account.Username,
		Role:   account.UserRole,
		Token:  jwtToken,
	}, nil
}

func (c *accountServiceController) generateJwtToken(username string, role int) (string, error) {
	var (
		token  = jwt.New(jwt.SigningMethodHS256)
		claims = token.Claims.(jwt.MapClaims)
	)

	// assign username, role and expiration time
	claims["username"] = username
	claims["userrole"] = role
	claims["expired"] = time.Now().UTC().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString([]byte("secretkey"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
