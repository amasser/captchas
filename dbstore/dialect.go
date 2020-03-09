// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package dbstore

type Dialect interface {
	Insert(table string) string
	Delete(table string) string
	QueryRow(table string) string
	DeleteExpired(table string) string
}
