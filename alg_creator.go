package middleware

import (
	"github.com/gbrlsnchs/jwt/v3"
)

var cachedAlgs map[string]jwt.Algorithm // чтобы не создавать постоянно
var isFirstCall bool

func сreateAlg(client Client) (jwt.Algorithm, error) {
	if isFirstCall == false {
		isFirstCall = true
		cachedAlgs = make(map[string]jwt.Algorithm)
	}

	if cachedAlgs[client.Code] != nil {
		return  cachedAlgs[client.Code], nil
	}

	switch client.Alg {
	case "HS256":
		cachedAlgs[client.Code] = jwt.NewHS256([]byte(client.Secret))
		break
	case "HS384":
		cachedAlgs[client.Code] = jwt.NewHS384([]byte(client.Secret))
		break
	case "HS512":
		cachedAlgs[client.Code] = jwt.NewHS512([]byte(client.Secret))
		break
	default:
		return nil, &Error{Message:"Unknown alg"}
	}

	return cachedAlgs[client.Code], nil
}