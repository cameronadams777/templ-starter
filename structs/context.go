package structs

type AppContext struct {
	Key   string
	Value interface{}
}

type SessionContext struct {
	CSRFToken string
	UserID    string
}

func (s SessionContext) IsAuthenticated() bool {
	return s.UserID != ""
}
