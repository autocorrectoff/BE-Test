package utils

import "log"

// HandleErrors -> just logging the error at the moment
func HandleError(err error) {
	if err != nil {
		log.Println("An error occurred", err)
	}
}
