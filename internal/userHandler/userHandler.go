package userHandler

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"time"

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
	postData := url.Values{"grant_type": {"authorization_code"}, "code": {usercode}}
	req, err := http.NewRequest("POST", "https://account.thethingsnetwork.org/users/token", bytes.NewReader(postData))
	req.Header.Add("User-Agent", "myClient")
	resp, err := client.Do(req)

	for i := range make([]int, 20) {
		//fmt.Printf("%d", i)
		time.Sleep(1 * time.Second)
		if i > 10 {
			fmt.Printf("Changin value\n")
			currentTasks[taskId].status = Status_Updating
		}
	}
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
