package middleware

import "github.com/omerberkcan/banking-transfer/internal/session"

type Middelwares struct {
	JwtMiddleware JWT
}

func New(s *session.Redis) *Middelwares {
	return &Middelwares{JwtMiddleware: jwtMiddleware{s: s}}
}
