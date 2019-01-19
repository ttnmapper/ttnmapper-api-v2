package accountRoutes

import (
	"fmt"
	"net/http"
	"time"

	"ttnmapper/api-v2/internal/userHandler"

	jwt "github.com/dgrijalva/jwt-go"
)

const (
	LOGIN_TICKET_MINUTES int = 10
)

var SECRET_KEY = []byte("AllYourBase")

/*
 * @brief 	Called when a user logs in to dispatch the required tasks
 *
 * When the user logs in we need to update all their details from TTNMapper.
 * Since this might take a while, send a loginTicket to the user (just a
 * random UUID). The webpage can use this to lookup the status of the login.
 * Usually this should be fast, but just for in case.
 *
 * The loginTicket is provided in a JWT token with a lifetime of
 * LOGIN_TICKET_MINUTES
 */
func LoginUser(w http.ResponseWriter, r *http.Request) {

	//Get the code from the request
	r.ParseForm()

	loginCode, ok := r.Form["code"]
	if ok {
		fmt.Printf("Received code %s\n", loginCode)
	} else {
		fmt.Printf("No code received\n")
		return
	}

	// Check the code has a sensible length - it is normally about 43 characters
	if len(loginCode) > 100 {
		return
	}

	// Create new login request, with a number to give the user
	loginTicket := userHandler.DispatchUserLogin(loginCode[0])

	type MyCustomClaims struct {
		Foo string `json:"loginTicket"`
		jwt.StandardClaims
	}

	// Create the Claims
	claims := MyCustomClaims{
		loginTicket,
		jwt.StandardClaims{
			Issuer:    "ttnmapper.org",
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(10)).Unix(),
		},
	}

	// Create the token and sign with the secret
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(SECRET_KEY)

	// jData, err := json.Marshal(response)
	if err != nil {
		// handle error
		fmt.Printf("Error: %s", err.Error())
	}
	//w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(tokenString))

}

/*
 *	@brief Check the status of the provided loginTicket
 *
 *
 */
func CheckLoginStatus(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	loginTicket, ok := r.Form["ticket"]
	if !ok {
		fmt.Printf("No ticket supplied\n")
		return
	}

	state := userHandler.CheckTicketState(loginTicket[0])
	fmt.Printf(string(state))
}

/*
 * This function is called when a user returns to the site, with an existing
 * token in the local store.
 */
func VerifyToken(w http.ResponseWriter, r *http.Request) {

}

func CheckStatus(w http.ResponseWriter, r *http.Request) {

}
