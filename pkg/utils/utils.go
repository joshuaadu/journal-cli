package utils

import (
	//"fmt"
	//"time"
	"github.com/google/uuid"
)

// GenerateID generates a unique ID for the entry (you can use UUID or another method)
func GenerateID() string {
	// Placeholder for now. You can replace this with a UUID generator.
	//return fmt.Sprintf("%d", time.Now().UnixNano())
	//  generates a unique ID using UUID
	return uuid.New().String()
}
