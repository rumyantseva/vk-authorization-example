package services

import (
	"github.com/rumyantseva/vk-authorization-example/models"
	"gopkg.in/reform.v1"
	"math/rand"
	"time"

	"gopkg.in/hlandau/passlib.v1"
)

type Registration struct {
	q *reform.Querier
}

func NewRegistration(q *reform.Querier) *Registration {
	return &Registration{
		q: q,
	}
}

func (r *Registration) Register(userID int64, email string, token string, tokenLifetime int) (int, string, error) {
	password := r.makePassword(10)
	hash, err := passlib.Hash(password)
	if err != nil {
		return 0, "", err
	}

	// ToDo: check for unique vk-id and email
	auth := &models.Auth{
		Email:           email,
		Password:        hash,
		VKUserID:        userID,
		VKToken:         token,
		VKTokenLifetime: tokenLifetime,
	}
	err = r.q.Save(auth)
	if err != nil {
		return 0, "", err
	}

	return auth.ID, password, nil
}

func (r *Registration) makePassword(n int) string {
	rand.Seed(time.Now().UnixNano())
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(b)
}
