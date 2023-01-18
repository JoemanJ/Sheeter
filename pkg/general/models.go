package sheeters

import "github.com/golang-jwt/jwt"

const USERSDATA = "data/users.json"

/*
This is just a prototype used for testing.
*/

type User struct{
  Username string
  Password string
  OwnedSheets map[int]string
}

type Claims struct{
  Username string
  OwnedSheets []int
  jwt.StandardClaims
}
