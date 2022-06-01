package model

type UserContextKey string

// Context key for user information, used on middleware and every handler
// to get the user information on context
const (
	ContextKeyUserInformation UserContextKey = "UserContextInformation"
)

// This user context used on Auth Middleware for passing context
// user information to handler
type UserContext struct {
	Id    int
	Name  string
	Email string
}
