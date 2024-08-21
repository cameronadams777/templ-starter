package structs

type AppContext struct {
	Key   string
	Value interface{}
}

type SessionContext struct {
	CSRFToken string
	UserID    uint
}

func (s SessionContext) IsAuthenticated() bool {
  return s.UserID != 0
}
