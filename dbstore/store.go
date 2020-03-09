// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package dbstore

import (
	"database/sql"
	"time"

	"github.com/clevergo/captchas"
)

type item struct {
	ID        string `db:"id"`
	Answer    string `db:"answer"`
	CreatedAt int64  `db:"created_at"`
	ExpiresIn int64  `db:"expires_in"`
}

// Option is a function that receives a pointer of store.
type Option func(*Store)

// Expiration sets expiration.
func Expiration(expiration time.Duration) Option {
	return func(s *Store) {
		s.expiration = expiration
	}
}

// GCInterval sets garbage collection .
func GCInterval(interval time.Duration) Option {
	return func(s *Store) {
		s.gcInterval = interval
	}
}

type Store struct {
	db         *sql.DB
	dialect    Dialect
	tableName  string
	prefix     string
	expiration time.Duration
	gcInterval time.Duration
}

// New returns a db store.
func New(db *sql.DB, dialect Dialect, opts ...Option) *Store {
	s := &Store{
		db:         db,
		dialect:    dialect,
		tableName:  "captchas",
		prefix:     "default",
		expiration: 10 * time.Minute,
		gcInterval: time.Hour,
	}

	for _, f := range opts {
		f(s)
	}

	go s.gc()

	return s
}

func (s *Store) getID(id string) string {
	return s.prefix + ":" + id
}

// Get implements Store.Get.
func (s *Store) Get(id string, clear bool) (string, error) {
	id = s.getID(id)
	row := s.db.QueryRow(s.dialect.QueryRow(s.tableName), id)
	if row == nil {
		return "", captchas.ErrCaptchaIncorrect
	}
	item := item{}
	if err := row.Scan(&item.ID, &item.Answer, &item.ExpiresIn); err != nil {
		if err == sql.ErrNoRows {
			return "", captchas.ErrCaptchaIncorrect
		}
		return "", err
	}
	if time.Now().Unix() > item.ExpiresIn {
		return "", captchas.ErrCaptchaExpired
	}

	if clear {
		_, err := s.db.Exec(s.dialect.Delete(s.tableName), id)
		if err != nil {
			return "", err
		}
	}

	return item.Answer, nil
}

// Set implements Store.Set.
func (s *Store) Set(id, answer string) error {
	id = s.getID(id)
	now := time.Now()
	_, err := s.db.Exec(
		s.dialect.Insert(s.tableName),
		id, answer, now.Unix(), now.Add(s.expiration).Unix(),
	)
	return err
}

func (s *Store) gc() {
	ticker := time.NewTicker(s.gcInterval)
	for {
		select {
		case <-ticker.C:
			s.deleteExpired()
		}
	}
}

func (s *Store) deleteExpired() {
	s.db.Exec(s.dialect.DeleteExpired(s.tableName), time.Now().Unix())
}
