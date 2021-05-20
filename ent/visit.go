// Code generated by entc, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/ichandxyx/task/ent/job"
	"github.com/ichandxyx/task/ent/visit"
)

// Visit is the model entity for the Visit schema.
type Visit struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// StoreID holds the value of the "storeID" field.
	StoreID string `json:"storeID,omitempty"`
	// VisitTime holds the value of the "visitTime" field.
	VisitTime time.Time `json:"visitTime,omitempty"`
	// Error holds the value of the "error" field.
	Error *string `json:"error,omitempty"`
	// Perimeter holds the value of the "perimeter" field.
	Perimeter int `json:"perimeter,omitempty"`
	// ImageURLs holds the value of the "imageURLs" field.
	ImageURLs []string `json:"imageURLs,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the VisitQuery when eager-loading is set.
	Edges      VisitEdges `json:"edges"`
	job_visits *int
}

// VisitEdges holds the relations/edges for other nodes in the graph.
type VisitEdges struct {
	// Job holds the value of the job edge.
	Job *Job `json:"job,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// JobOrErr returns the Job value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e VisitEdges) JobOrErr() (*Job, error) {
	if e.loadedTypes[0] {
		if e.Job == nil {
			// The edge job was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: job.Label}
		}
		return e.Job, nil
	}
	return nil, &NotLoadedError{edge: "job"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Visit) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case visit.FieldImageURLs:
			values[i] = new([]byte)
		case visit.FieldID, visit.FieldPerimeter:
			values[i] = new(sql.NullInt64)
		case visit.FieldStoreID, visit.FieldError:
			values[i] = new(sql.NullString)
		case visit.FieldVisitTime:
			values[i] = new(sql.NullTime)
		case visit.ForeignKeys[0]: // job_visits
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Visit", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Visit fields.
func (v *Visit) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case visit.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			v.ID = int(value.Int64)
		case visit.FieldStoreID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field storeID", values[i])
			} else if value.Valid {
				v.StoreID = value.String
			}
		case visit.FieldVisitTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field visitTime", values[i])
			} else if value.Valid {
				v.VisitTime = value.Time
			}
		case visit.FieldError:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field error", values[i])
			} else if value.Valid {
				v.Error = new(string)
				*v.Error = value.String
			}
		case visit.FieldPerimeter:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field perimeter", values[i])
			} else if value.Valid {
				v.Perimeter = int(value.Int64)
			}
		case visit.FieldImageURLs:

			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field imageURLs", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &v.ImageURLs); err != nil {
					return fmt.Errorf("unmarshal field imageURLs: %w", err)
				}
			}
		case visit.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field job_visits", value)
			} else if value.Valid {
				v.job_visits = new(int)
				*v.job_visits = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryJob queries the "job" edge of the Visit entity.
func (v *Visit) QueryJob() *JobQuery {
	return (&VisitClient{config: v.config}).QueryJob(v)
}

// Update returns a builder for updating this Visit.
// Note that you need to call Visit.Unwrap() before calling this method if this Visit
// was returned from a transaction, and the transaction was committed or rolled back.
func (v *Visit) Update() *VisitUpdateOne {
	return (&VisitClient{config: v.config}).UpdateOne(v)
}

// Unwrap unwraps the Visit entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (v *Visit) Unwrap() *Visit {
	tx, ok := v.config.driver.(*txDriver)
	if !ok {
		panic("ent: Visit is not a transactional entity")
	}
	v.config.driver = tx.drv
	return v
}

// String implements the fmt.Stringer.
func (v *Visit) String() string {
	var builder strings.Builder
	builder.WriteString("Visit(")
	builder.WriteString(fmt.Sprintf("id=%v", v.ID))
	builder.WriteString(", storeID=")
	builder.WriteString(v.StoreID)
	builder.WriteString(", visitTime=")
	builder.WriteString(v.VisitTime.Format(time.ANSIC))
	if v := v.Error; v != nil {
		builder.WriteString(", error=")
		builder.WriteString(*v)
	}
	builder.WriteString(", perimeter=")
	builder.WriteString(fmt.Sprintf("%v", v.Perimeter))
	builder.WriteString(", imageURLs=")
	builder.WriteString(fmt.Sprintf("%v", v.ImageURLs))
	builder.WriteByte(')')
	return builder.String()
}

// Visits is a parsable slice of Visit.
type Visits []*Visit

func (v Visits) config(cfg config) {
	for _i := range v {
		v[_i].config = cfg
	}
}
