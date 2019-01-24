package accountRoutes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ttnmapper/ttnmapper-api-v2/internal/userHandler"

	jwt "github.com/dgrijalva/jwt-go"
)

const (
	LOGIN_TICKET_MINUTES int = 10
)

var SECRET_KEY = []byte("AllYourBase")

type MyCustomClaims struct {
	LoginTicket string `json:"loginTicket"`
	jwt.StandardClaims
}

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

	result := make(map[string]string)
	w.Header().Set("Content-Type", "application/json")

	loginTicket, ok := r.Form["ticket"]
	if !ok {
		fmt.Printf("No ticket supplied\n")
		result["error"] = "No ticket supplied in request"
		jsonStr, err := json.Marshal(result)
		if err != nil {
			w.Write(jsonStr)
		}
		return
	}

	// The loginTicket is the entire jwt token. Verify it
	token, err := jwt.ParseWithClaims(loginTicket[0], &MyCustomClaims{}, keyFunction)
	if err != nil {
		fmt.Printf(err.Error()) // eg. token is expired by 12m54s
		result["error"] = err.Error()
		jsonStr, err := json.Marshal(result)
		if err != nil {
			w.Write(jsonStr)
		}
		return
	}

	claims := token.Claims.(*MyCustomClaims)
	if claims.StandardClaims.VerifyExpiresAt(time.Now().Unix(), true) {
		fmt.Printf("loginTicket time valid")
		state := userHandler.CheckTicketState(claims.LoginTicket)
		fmt.Printf(string(state))

		result["login_state"] = string(state)
		jsonStr, err := json.Marshal(result)
		if err != nil {
			w.Write(jsonStr)
		}
		return
	} else {
		fmt.Printf("loginTicket time invalid")
		jsonStr, _ := json.Marshal(result)
		w.Write([]byte(jsonStr))
		return
	}

}

/*
 * This function is called when a user returns to the site, with an existing
 * token in the local store.
 */
func VerifyToken(w http.ResponseWriter, r *http.Request) {

}

func CheckStatus(w http.ResponseWriter, r *http.Request) {

}

func keyFunction(token *jwt.Token) (interface{}, error) {
	return SECRET_KEY, nil
}
