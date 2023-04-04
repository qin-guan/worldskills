// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/qin-guan/WorldSkills/Kazan2019/ent/employee"
)

// Employee is the model entity for the Employee schema.
type Employee struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// FirstName holds the value of the "FirstName" field.
	FirstName string `json:"FirstName,omitempty"`
	// LastName holds the value of the "LastName" field.
	LastName string `json:"LastName,omitempty"`
	// Phone holds the value of the "Phone" field.
	Phone string `json:"Phone,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the EmployeeQuery when eager-loading is set.
	Edges EmployeeEdges `json:"edges"`
}

// EmployeeEdges holds the relations/edges for other nodes in the graph.
type EmployeeEdges struct {
	// Asset holds the value of the Asset edge.
	Asset []*Asset `json:"Asset,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// AssetOrErr returns the Asset value or an error if the edge
// was not loaded in eager-loading.
func (e EmployeeEdges) AssetOrErr() ([]*Asset, error) {
	if e.loadedTypes[0] {
		return e.Asset, nil
	}
	return nil, &NotLoadedError{edge: "Asset"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Employee) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case employee.FieldID:
			values[i] = new(sql.NullInt64)
		case employee.FieldFirstName, employee.FieldLastName, employee.FieldPhone:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Employee", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Employee fields.
func (e *Employee) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case employee.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			e.ID = int(value.Int64)
		case employee.FieldFirstName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field FirstName", values[i])
			} else if value.Valid {
				e.FirstName = value.String
			}
		case employee.FieldLastName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field LastName", values[i])
			} else if value.Valid {
				e.LastName = value.String
			}
		case employee.FieldPhone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Phone", values[i])
			} else if value.Valid {
				e.Phone = value.String
			}
		}
	}
	return nil
}

// QueryAsset queries the "Asset" edge of the Employee entity.
func (e *Employee) QueryAsset() *AssetQuery {
	return NewEmployeeClient(e.config).QueryAsset(e)
}

// Update returns a builder for updating this Employee.
// Note that you need to call Employee.Unwrap() before calling this method if this Employee
// was returned from a transaction, and the transaction was committed or rolled back.
func (e *Employee) Update() *EmployeeUpdateOne {
	return NewEmployeeClient(e.config).UpdateOne(e)
}

// Unwrap unwraps the Employee entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (e *Employee) Unwrap() *Employee {
	_tx, ok := e.config.driver.(*txDriver)
	if !ok {
		panic("ent: Employee is not a transactional entity")
	}
	e.config.driver = _tx.drv
	return e
}

// String implements the fmt.Stringer.
func (e *Employee) String() string {
	var builder strings.Builder
	builder.WriteString("Employee(")
	builder.WriteString(fmt.Sprintf("id=%v, ", e.ID))
	builder.WriteString("FirstName=")
	builder.WriteString(e.FirstName)
	builder.WriteString(", ")
	builder.WriteString("LastName=")
	builder.WriteString(e.LastName)
	builder.WriteString(", ")
	builder.WriteString("Phone=")
	builder.WriteString(e.Phone)
	builder.WriteByte(')')
	return builder.String()
}

// Employees is a parsable slice of Employee.
type Employees []*Employee