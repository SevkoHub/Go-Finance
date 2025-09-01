// audit_log.go
package models

import "time"

// AuditLog struct represents a log entry for actions performed in the application.
// It contains fields to store the ID, action type, user ID, timestamp, and any additional details.
type AuditLog struct {
	ID        int       `json:"id"`        // Unique identifier for the audit log entry
	Action    string    `json:"action"`    // The action that was performed (e.g., "create", "update", "delete")
	UserID    int       `json:"user_id"`   // The ID of the user who performed the action
	Timestamp time.Time `json:"timestamp"`  // The time when the action was performed
	Details   string    `json:"details"`    // Additional details about the action
}

// NewAuditLog creates a new instance of AuditLog with the current timestamp.
// This function can be used to log a new action in the system.
func NewAuditLog(action string, userID int, details string) AuditLog {
	return AuditLog{
		Action:    action,
		UserID:    userID,
		Timestamp: time.Now(), // Set the current time as the timestamp
		Details:   details,
	}
}