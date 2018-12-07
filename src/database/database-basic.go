package main

import (
	"database/sql/driver"
	"sync"
)

/**
We always see the following code when we use third-party drivers:
import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

Here, the underscore (also known as a 'blank') _ can be quite confusing for many beginners,
but this is a great feature in Go. We already know that this underscore identifier is used for discarding values
from function returns, and also that you must use all packages that you've imported in your code in Go.
So when the blank is used with import,
it means that you need to execute the init() function of that package without directly using it,
which is a perfect fit for the use-case of registering database drivers.

import _ "xxx/xxx/xxx" 语法表示执行xxx包的init函数，但不直接使用这个包，特别适合用于注册数据库驱动
 */

type DB struct {
	driver   driver.Driver
	dsn      string
	mu       sync.Mutex // protects freeConn and closed
	freeConn []driver.Conn
	closed   bool
}

func main() {
}
