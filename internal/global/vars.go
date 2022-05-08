package global

import (
	"os"

	"github.com/gorilla/sessions"
)

// SessionStore is the session key used by the application
var SessionStore = sessions.NewCookieStore([]byte(os.Getenv("SESSIONS_KEY")))
