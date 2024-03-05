// Code generated by ent, DO NOT EDIT.

package ent

import (
	"Real-Time-Chat/ent/chat"
	"Real-Time-Chat/ent/user"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Chat is the model entity for the Chat schema.
type Chat struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Message holds the value of the "message" field.
	Message string `json:"message,omitempty"`
	// SentAt holds the value of the "sent_at" field.
	SentAt time.Time `json:"sent_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ChatQuery when eager-loading is set.
	Edges                  ChatEdges `json:"edges"`
	user_sent_messages     *int
	user_received_messages *int
	selectValues           sql.SelectValues
}

// ChatEdges holds the relations/edges for other nodes in the graph.
type ChatEdges struct {
	// SentUser holds the value of the sent_user edge.
	SentUser *User `json:"sent_user,omitempty"`
	// ReceivedUser holds the value of the received_user edge.
	ReceivedUser *User `json:"received_user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// SentUserOrErr returns the SentUser value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ChatEdges) SentUserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.SentUser == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.SentUser, nil
	}
	return nil, &NotLoadedError{edge: "sent_user"}
}

// ReceivedUserOrErr returns the ReceivedUser value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ChatEdges) ReceivedUserOrErr() (*User, error) {
	if e.loadedTypes[1] {
		if e.ReceivedUser == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.ReceivedUser, nil
	}
	return nil, &NotLoadedError{edge: "received_user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Chat) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case chat.FieldID:
			values[i] = new(sql.NullInt64)
		case chat.FieldMessage:
			values[i] = new(sql.NullString)
		case chat.FieldSentAt:
			values[i] = new(sql.NullTime)
		case chat.ForeignKeys[0]: // user_sent_messages
			values[i] = new(sql.NullInt64)
		case chat.ForeignKeys[1]: // user_received_messages
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Chat fields.
func (c *Chat) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case chat.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case chat.FieldMessage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field message", values[i])
			} else if value.Valid {
				c.Message = value.String
			}
		case chat.FieldSentAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field sent_at", values[i])
			} else if value.Valid {
				c.SentAt = value.Time
			}
		case chat.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_sent_messages", value)
			} else if value.Valid {
				c.user_sent_messages = new(int)
				*c.user_sent_messages = int(value.Int64)
			}
		case chat.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_received_messages", value)
			} else if value.Valid {
				c.user_received_messages = new(int)
				*c.user_received_messages = int(value.Int64)
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Chat.
// This includes values selected through modifiers, order, etc.
func (c *Chat) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// QuerySentUser queries the "sent_user" edge of the Chat entity.
func (c *Chat) QuerySentUser() *UserQuery {
	return NewChatClient(c.config).QuerySentUser(c)
}

// QueryReceivedUser queries the "received_user" edge of the Chat entity.
func (c *Chat) QueryReceivedUser() *UserQuery {
	return NewChatClient(c.config).QueryReceivedUser(c)
}

// Update returns a builder for updating this Chat.
// Note that you need to call Chat.Unwrap() before calling this method if this Chat
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Chat) Update() *ChatUpdateOne {
	return NewChatClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Chat entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Chat) Unwrap() *Chat {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Chat is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Chat) String() string {
	var builder strings.Builder
	builder.WriteString("Chat(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("message=")
	builder.WriteString(c.Message)
	builder.WriteString(", ")
	builder.WriteString("sent_at=")
	builder.WriteString(c.SentAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Chats is a parsable slice of Chat.
type Chats []*Chat
