// Code generated by entc, DO NOT EDIT.

package job

import (
	"time"
)

const (
	// Label holds the string label denoting the job type in the database.
	Label = "job"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldCreatedAt holds the string denoting the createdat field in the database.
	FieldCreatedAt = "created_at"
	// EdgeVisits holds the string denoting the visits edge name in mutations.
	EdgeVisits = "visits"
	// Table holds the table name of the job in the database.
	Table = "jobs"
	// VisitsTable is the table the holds the visits relation/edge.
	VisitsTable = "visits"
	// VisitsInverseTable is the table name for the Visit entity.
	// It exists in this package in order to avoid circular dependency with the "visit" package.
	VisitsInverseTable = "visits"
	// VisitsColumn is the table column denoting the visits relation/edge.
	VisitsColumn = "job_visits"
)

// Columns holds all SQL columns for job fields.
var Columns = []string{
	FieldID,
	FieldStatus,
	FieldCreatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus string
	// DefaultCreatedAt holds the default value on creation for the "createdAt" field.
	DefaultCreatedAt func() time.Time
)