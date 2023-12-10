package exceptionhandler

import "log"

func HandleError(err error) {
	if err != nil {
		log.Printf("Error %v", err)
	}
}
