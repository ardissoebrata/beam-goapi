# Arisan Yuk! API

## Set SQL Mode (MySQL)

Fix "Incorrect datetime value: '0000-00-00'" error

1. Edit <code>/etc/mysql/mysql.conf.d/mysqld.cnf</code> and add line:
      ```
      sql_mode="ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION"
      ```

      (remove 'NO_ZERO_DATE' from the list)

2. Restart MySQL service
    ```bash
    sudo service mysql restart
    ```

## Migrate

```
$ go run cmd/migrate/main.go
```

## Seed

```
$ go run cmd/seed/main.go
```

## Run

```
$ go run cmd/server/main.go
```

## Live Reload Go

Create <code>.air.toml</code> in your project root directory (you can copy from <code>.air.toml.sample</code>)

Install **Air** (https://github.com/cosmtrek/air)

```
$ go install github.com/cosmtrek/air@latest
```

If you use Laragon, the binary will be installed on
<code>\<path to laragon>/usr/go/bin</code>

Add that path to <code>PATH</code> environment

run **Air**

```
$ air
```