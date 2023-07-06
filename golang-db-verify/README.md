# GoLang DB Verify

验证数据库连接，主要验证数据库连接的配置**SetMaxIdleConns**，**SetConnMaxLifetime**，**SetConnMaxIdleTime**，*
*SetConnMaxLifetime**和MySQL变量**wait_timeout**，**interactive_timeout**，**max_connections**，**connect_timeout**之间的关系。

## MySQL配置

```
SHOW GLOBAL VARIABLES LIKE "%timeout%";
set global wait_timeout=30;
set global interactive_timeout=30;
```

wait_timeout默认28800s，interactive_timeout默认28800s，max_connections默认151个，connect_timeout默认10s

## 数据库驱动的统计

```
MaxOpenConnections:10 OpenConnections:1 InUse:0 Idle:1 WaitCount:0 WaitDuration:0s MaxIdleClosed:0 MaxIdleTimeClosed:0 MaxLifetimeClosed:0
```

MaxConn >= OpenConn = InUseConn + IdleConn

## 验证方式

http://localhost:8000/    （普通查询）

http://localhost:8000/sleep  （占用连接，到达最大连接数）

http://localhost:8000/db_stats （连接池统计）

## 建议

db.SetConnMaxLifetime()
是必需的，以确保在MySQL服务器、操作系统或其他中间件关闭连接之前，驱动程序能安全地关闭连接。由于一些中间件在5分钟前关闭空闲连接，我们建议超时时间短于5分钟。这个设置也有助于负载平衡和改变系统变量。

强烈建议使用db.SetMaxOpenConns()来限制应用程序使用的连接数。没有推荐的限制数，因为它取决于应用程序和MySQL服务器。

db.SetMaxIdleConns()推荐与db.SetMaxOpenConns()设置相同。当它小于SetMaxOpenConns()
时，连接被打开和关闭的频率会比你预期的高得多。闲置的连接可以通过db.SetConnMaxLifetime()来关闭。如果你想更快地关闭空闲连接，你可以从Go
1.15开始使用db.SetConnMaxIdleTime()。

SetConnMaxLifetime设置一个连接可以被重复使用的最大时间。

SetConnMaxIdleTime设置一个连接可以空闲的最大时间。


## 参考

- https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_wait_timeout
- https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_interactive_timeout

