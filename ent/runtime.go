// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/ichandxyx/task/ent/job"
	"github.com/ichandxyx/task/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	jobFields := schema.Job{}.Fields()
	_ = jobFields
	// jobDescStatus is the schema descriptor for status field.
	jobDescStatus := jobFields[0].Descriptor()
	// job.DefaultStatus holds the default value on creation for the status field.
	job.DefaultStatus = jobDescStatus.Default.(string)
	// jobDescCreatedAt is the schema descriptor for createdAt field.
	jobDescCreatedAt := jobFields[1].Descriptor()
	// job.DefaultCreatedAt holds the default value on creation for the createdAt field.
	job.DefaultCreatedAt = jobDescCreatedAt.Default.(func() time.Time)
}
