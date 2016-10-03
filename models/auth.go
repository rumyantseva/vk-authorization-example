package models

import (
	"time"
)

//go:generate reform

// Auth model.
// reform:auth
type Auth struct {
	ID       int    `reform:"id,pk" json:"id"`
	Email    string `reform:"email" json:"email"`
	Password string `reform:"password" json:"-"`

	VKUserID        int64  `reform:"vk_user_id" json:"vk_user_id"`
	VKToken         string `reform:"vk_token" json:"vk_token"`
	VKTokenLifetime int    `reform:"vk_token_lifetime" json:"vk_token_lifetime"`

	CreatedAt time.Time `reform:"created_at" json:"-"`
	UpdatedAt time.Time `reform:"updated_at" json:"-"`
}

func (a *Auth) BeforeInsert() error {
	now := time.Now()
	if a.CreatedAt.IsZero() {
		a.CreatedAt = now
	}
	if a.UpdatedAt.IsZero() {
		a.UpdatedAt = now
	}

	return nil
}

func (a *Auth) BeforeUpdate() error {
	if a.UpdatedAt.IsZero() {
		a.UpdatedAt = time.Now()
	}

	return nil
}
