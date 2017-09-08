# Sqlq Simple Query Builder

Simple Query Builder

## Getting Started

### Prerequisites

- go - https://golang.org/

```
https://golang.org/doc/install
```

### Installing

- Get this library using go get

```
go get github.com/valuppo/sqlq
```

### How To Use
- Select Query Builder <br />
  There are 2 ways two select columns: Select(Column, ...) or Columns(Column, ...)

```
sql := sqlq.Select("id","name").Columns("email","address").From("users").Where("id", "=", "1").WhereOr("name", "LIKE", "%sqlq%")
        .OrderBy("created_at","DESC").Limit(7).Sql()
fmt.Println(sql) //SELECT id, name, email, address FROM users WHERE id = '1' OR name LIKE '%sqlq' ORDER BY created_at DESC LIMIT 7
```

- Insert Query Builder

```
sql := sqlq.Insert().Into("users").Columns("name", "email").Values("sqlq", "sqlq@valuppo.com").Sql()
fmt.Println(sql) //INSERT INTO users (name,email) VALUES ('sqlq', 'sqlq@valuppo.com')
```

- Update Query Builder <br />
  There are 2 ways two set columns and values: Set(Column, Value) or SetMultiple(Columns, Values) <br />
  Where in set multiple you can pass slice of columns and values into it

```
sql := sqlq.Update("users").Set("name", "sqlq").SetMultiple([]string{"email","password"}, []string{"sqlq@valuppo.com", "sqlqpass")
        .Where("id", "=", "5").WhereOr("name", "LIKE", "%sqlq%").Sql()
fmt.Println(sql) //UPDATE users SET name = 'sqlq', email = 'sqlq@valuppo.com', password = 'sqlqpass' WHERE id = '5' OR name LIKE '%sqlq%'
```

- Delete Query Builder

```
sql := sqlq.Delete().From("users").Where("id", "=", "3").WhereOr("created_at", "<", "CURDATE()").Sql()
fmt.Println(sql) //DELETE FROM users WHERE id = '3' OR created_at < 'CURDATE()'
```

## Built With

* [Golang](https://golang.org/) - Programming Language

## Authors

* **Antony Gunawan** - *Initial work* - [valuppo](https://github.com/valuppo)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
