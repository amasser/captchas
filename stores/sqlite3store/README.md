# SQLite3 Store for Captchas

```shell
$ go get github.com/clevergo/captchas/stores/sqlite3store
```

```go
import (
	"github.com/clevergo/captchas/stores/dbstore"
	"github.com/clevergo/captchas/stores/sqlite3store"
	_ "github.com/mattn/go-sqlite3"
)
```

```go
store := sqlite3store.New(
	dbstore.Expiration(10*time.Minute), // captcha expiration, optional.
	dbstore.GCInterval(time.Minute), // garbage collection interval to delete expired captcha, optional.
	dbstore.TableName("captchas"), // table name, optional.
	dbstore.Category("default"), // category, optional.
)
```