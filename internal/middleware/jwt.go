package middleware

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/omerberkcan/banking-transfer/internal/session"
)

type jwtMiddleware struct {
	s *session.Redis
}

type JWT interface {
	AuthControl(next echo.HandlerFunc) echo.HandlerFunc
}

func (j jwtMiddleware) AuthControl(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var userID int
		user, ok := c.Get("user").(*jwt.Token)
		if !ok {
			return c.NoContent(http.StatusUnauthorized)
		}
		if !user.Valid {
			return c.NoContent(http.StatusUnauthorized)
		}

		claims, _ := user.Claims.(jwt.MapClaims)
		tmp, ok := claims["user_id"].(float64)
		if !ok {
			return c.NoContent(http.StatusUnauthorized)
		}
		userID = int(tmp)
		uuid := claims["uuid"].(string)

		token, err := j.s.FindTokenByUserID(userID)
		if err != nil {
			return c.NoContent(http.StatusUnauthorized)
		}

		if token.Uuid.String() != uuid {
			j.s.DeleteTokenByUserID(userID)
			return c.NoContent(http.StatusUnauthorized)

		}

		c.Set("userid", userID)

		return next(c)
	}
}
