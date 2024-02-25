package kms

import (
	"context"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/golang-jwt/jwt/v5"
	"github.com/matelang/jwt-go-aws-kms/v2/jwtkms"
)

const keyID = "447c86fa-884e-4f85-a278-19077bb7c9b7"

func TestKMS() {
	awsCfg, err := config.LoadDefaultConfig(context.Background(),
		config.WithRegion("ap-southeast-1"))
	if err != nil {
		panic(err)
	}
	type Claims struct {
		Username string `json:"username"`
		jwt.RegisteredClaims
	}
	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &Claims{
		Username: "JOHN CENA",
		RegisteredClaims: jwt.RegisteredClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	jwtToken := jwt.NewWithClaims(jwtkms.SigningMethodECDSA256, claims)

	kmsConfig := jwtkms.NewKMSConfig(kms.NewFromConfig(awsCfg), keyID, false)

	// parse JWT object => string token
	str, err := jwtToken.SignedString(kmsConfig.WithContext(context.Background()))
	if err != nil {
		log.Fatalf("can not sign JWT %s", err)
	}

	log.Printf("Signed JWT %s\n", str)

	// parse string token => JWT object
	claimsDecode := Claims{}
	_, err = jwt.ParseWithClaims(str, &claimsDecode, func(token *jwt.Token) (interface{}, error) {
		return kmsConfig, nil
	})
	if err != nil {
		log.Fatalf("can not parse/verify token %s", err)
	}

	log.Printf("Parsed and validated token with claims %v", claimsDecode.Username)
}
