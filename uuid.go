package gocelery

import uuid "github.com/satori/go.uuid"

// generateUUID generates a v4 uuid and returns it as a string
func generateUUID() string {
	id, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}
	return id.String()
}
