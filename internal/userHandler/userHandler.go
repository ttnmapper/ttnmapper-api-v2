package userHandler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/segmentio/ksuid"
)

// Status of tasks
const (
	Status_Started      int = 0
	Status_User_Created int = 1 //Logging in
	Status_Updating     int = 2 //Fetching details
	Status_Done         int = 3 //Fetching details
)

type task struct {
	status  int
	message string
}

// Map between job number and task details.
var currentTasks = make(map[string]*task)

/*
 * @brief		Dispatch all the calls neccesary when a users logs in.
 *
 * A login requires many requests to update user info, before we can give the
 * user a token. Return a 'ticket code' so that the status of the login can be
 * checked periodically
 *
 * @return		int		The code that can be used to lookup the status of the login.
 */
func DispatchUserLogin(usercode string) string {
	// Create a new task
	newTask := task{status: Status_Started, message: ""}
	taskId := ksuid.New().String()
	currentTasks[taskId] = &newTask

	// Spin of goroutine to fetch data
	go UserLoginAndUpdate(usercode, taskId)

	return taskId
}

func UserLoginAndUpdate(usercode string, taskId string) {

	// Send code to TTN Mapper servers
	fmt.Printf("Starting user login\n")
	client := &http.Client{}

	// Create BODY of request
	bodyData := url.Values{}
	bodyData.Set("grant_type", "authorization_code")
	bodyData.Set("code", usercode)

	req, _ := http.NewRequest("POST", "https://account.thethingsnetwork.org/users/token", strings.NewReader(bodyData.Encode()))

	// Add header fields
	req.Header.Add("Authorization", "Basic "+os.Getenv("TTN_AUTH_CODE"))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(bodyData.Encode())))

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error sending post request")
	}
	fmt.Printf("Response was\n")
	fmt.Printf(resp.Status)
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf(string(bodyBytes))

	fmt.Printf("Routine done\n")
	currentTasks[taskId].status = Status_Done

	// Read the response token

	// Dispatch new task to update all user details

	// Login the user, providing them with a token
}

/*
 * @brief		Called periodically to see the state of the login request
 *
 * @return 		int		Returns a number from
 */
func CheckTicketState(ticketCode string) int {
	// Check if the key is in the map:
	if val, ok := currentTasks[ticketCode]; ok {
		return val.status
	} else {
		return -1
	}
}
