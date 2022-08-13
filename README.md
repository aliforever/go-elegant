# Elegant
Working with SQL made easier

## Install
```go get github.com/aliforever/go-elegant```

## Use:

- Initialize db connection:
```go
db, err := sql.Open("postgres", "user=postgres password=root sslmode=disable database=testapp")
if err != nil {
panic(err)
}
```

- Define your model
```go
type users struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (users) TableName() string {
	return "users"
}
```

- Create the instance for User (users table):
```go
tbl := Table[users](db)
```

- If you want to pass default options for the table, pass it as the second argument:

    As an example, to always ignore `id` field when inserting data (when Id is of type autoincrement):
```go
tbl := Table[users](db, options.Table().SetInsertOptions(options.Insert().IgnoreFields("id")))
```

- Drop the table:
```go
err := tbl.DropTable()
if err != nil {
	panic(err)
}
```

- Create the table:
```go
err := tbl.BuildSchema().
    AddColumn(columns.NewInteger("id").PrimaryKey().Identity()).
    AddColumn(columns.NewVarchar("first_name", 20).NotNull()).
    AddColumn(columns.NewVarchar("last_name", 20).NotNull()).
    Build()
if err != nil {
    panic(err)
}
```

- Insert data to the table:
```go
err = tbl.Insert(users{
    FirstName: "Ali",
    LastName:  "Error",
}, options.Insert().IgnoreFields("id"))
if err != nil {
    panic(err)
}
```
- Read one row from the table:
```go
data, err := tbl.Query(func(builder *Builder) {
    builder.Where("first_name", "=", "Ali")
}).FindOne()
if err != nil {
    panic(err)
}
fmt.Println(data)
```

- Read all rows from the table:
```go
all, err := tbl.Query(nil).Find()
if err != nil {
    panic(err)
}
fmt.Println(all)
```

## TODO:
- Implement OrderBy, GroupBy, Having
- Implement Delete, Update