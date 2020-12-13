package main

import (
    "fmt"
    "time"
    "os"
    jwt "github.com/dgrijalva/jwt-go"
)

var secret string = os.Getenv("TOKEN_SECRET")
var mySigningKey = []byte(secret)

func GenerateJWT() (string, error){
    token := jwt.New(jwt.SigningMethodHS256)
    claims := token.Claims.(jwt.MapClaims)

    claims["authorized"] = true
    claims["user"] = "Sean Young"
    claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

    tokenString, err := token.SignedString(mySigningKey)

    if err != nil {
        fmt.Errorf("Something went wrong: %s", err.Error())
        return "", err
    }

    return tokenString, nil
}

func main(){
    fmt.Println("JWT Client")

    tokenString, err := GenerateJWT()
    if( err != nil){
        fmt.Println("Error generating token string")
    }
    fmt.Println(tokenString)
}