// Code generated by ent, DO NOT EDIT.

package db

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/dexidp/dex/storage/ent/db/connector"
)

// Connector is the model entity for the Connector schema.
type Connector struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Type holds the value of the "type" field.
	Type string `json:"type,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// ResourceVersion holds the value of the "resource_version" field.
	ResourceVersion string `json:"resource_version,omitempty"`
	// Config holds the value of the "config" field.
	Config []byte `json:"config,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Connector) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case connector.FieldConfig:
			values[i] = new([]byte)
		case connector.FieldID, connector.FieldType, connector.FieldName, connector.FieldResourceVersion:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Connector", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Connector fields.
func (c *Connector) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case connector.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				c.ID = value.String
			}
		case connector.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				c.Type = value.String
			}
		case connector.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		case connector.FieldResourceVersion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field resource_version", values[i])
			} else if value.Valid {
				c.ResourceVersion = value.String
			}
		case connector.FieldConfig:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field config", values[i])
			} else if value != nil {
				c.Config = *value
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Connector.
// Note that you need to call Connector.Unwrap() before calling this method if this Connector
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Connector) Update() *ConnectorUpdateOne {
	return (&ConnectorClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Connector entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Connector) Unwrap() *Connector {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("db: Connector is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Connector) String() string {
	var builder strings.Builder
	builder.WriteString("Connector(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("type=")
	builder.WriteString(c.Type)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(c.Name)
	builder.WriteString(", ")
	builder.WriteString("resource_version=")
	builder.WriteString(c.ResourceVersion)
	builder.WriteString(", ")
	builder.WriteString("config=")
	builder.WriteString(fmt.Sprintf("%v", c.Config))
	builder.WriteByte(')')
	return builder.String()
}

// Connectors is a parsable slice of Connector.
type Connectors []*Connector

func (c Connectors) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
