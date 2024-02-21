// +build tools
// Code generated by ent, DO NOT EDIT.

package ent

import (
	"Real-Time-Chat/ent/chat"
	"Real-Time-Chat/ent/schema"
	"Real-Time-Chat/ent/user"
	"Real-Time-Chat/ent/userrelation"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	chatFields := schema.Chat{}.Fields()
	_ = chatFields
	// chatDescSentAt is the schema descriptor for sent_at field.
	chatDescSentAt := chatFields[3].Descriptor()
	// chat.DefaultSentAt holds the default value on creation for the sent_at field.
	chat.DefaultSentAt = chatDescSentAt.Default.(func() time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[3].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	userrelationFields := schema.UserRelation{}.Fields()
	_ = userrelationFields
	// userrelationDescCreatedAt is the schema descriptor for created_at field.
	userrelationDescCreatedAt := userrelationFields[2].Descriptor()
	// userrelation.DefaultCreatedAt holds the default value on creation for the created_at field.
	userrelation.DefaultCreatedAt = userrelationDescCreatedAt.Default.(func() time.Time)
}
