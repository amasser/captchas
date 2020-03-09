// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package mysqlstore

import (
	"database/sql"
	"fmt"

	"github.com/clevergo/captchas/dbstore"
)

type Store struct {
	*dbstore.Store
}

func New(db *sql.DB, opts ...dbstore.Option) *Store {
	return &Store{dbstore.New(db, &dialect{}, opts...)}
}

type dialect struct {
}

func (d *dialect) Insert(table string) string {
	return fmt.Sprintf(`INSERT INTO %s(id, answer, created_at, expires_in) VALUES(?, ?, ?, ?)`, table)
}

func (d *dialect) QueryRow(table string) string {
	return fmt.Sprintf(`SELECT id, answer, expires_in FROM %s WHERE id = ?`, table)
}

func (d *dialect) Delete(table string) string {
	return fmt.Sprintf(`DELETE FROM %s WHERE id=?`, table)
}

func (d *dialect) DeleteExpired(table string) string {
	return fmt.Sprintf(`DELETE FROM %s WHERE expires_in<?`, table)
}
