// Code generated by ent, DO NOT EDIT.

package ent

import (
	"BazarMessenger/ent/profileandsetting"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// ProfileAndSetting is the model entity for the ProfileAndSetting schema.
type ProfileAndSetting struct {
	config
	// ID of the ent.
	ID           int `json:"id,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ProfileAndSetting) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case profileandsetting.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ProfileAndSetting fields.
func (pas *ProfileAndSetting) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case profileandsetting.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pas.ID = int(value.Int64)
		default:
			pas.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ProfileAndSetting.
// This includes values selected through modifiers, order, etc.
func (pas *ProfileAndSetting) Value(name string) (ent.Value, error) {
	return pas.selectValues.Get(name)
}

// Update returns a builder for updating this ProfileAndSetting.
// Note that you need to call ProfileAndSetting.Unwrap() before calling this method if this ProfileAndSetting
// was returned from a transaction, and the transaction was committed or rolled back.
func (pas *ProfileAndSetting) Update() *ProfileAndSettingUpdateOne {
	return NewProfileAndSettingClient(pas.config).UpdateOne(pas)
}

// Unwrap unwraps the ProfileAndSetting entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pas *ProfileAndSetting) Unwrap() *ProfileAndSetting {
	_tx, ok := pas.config.driver.(*txDriver)
	if !ok {
		panic("ent: ProfileAndSetting is not a transactional entity")
	}
	pas.config.driver = _tx.drv
	return pas
}

// String implements the fmt.Stringer.
func (pas *ProfileAndSetting) String() string {
	var builder strings.Builder
	builder.WriteString("ProfileAndSetting(")
	builder.WriteString(fmt.Sprintf("id=%v", pas.ID))
	builder.WriteByte(')')
	return builder.String()
}

// ProfileAndSettings is a parsable slice of ProfileAndSetting.
type ProfileAndSettings []*ProfileAndSetting