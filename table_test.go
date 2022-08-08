package elegant

import (
	"database/sql"
	"fmt"
	"github.com/aliforever/go-elegant/columns"
	"github.com/aliforever/go-elegant/options"
	_ "github.com/lib/pq"
	"testing"
	"time"
)

type users struct {
	Id        int64      `json:"id"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	CreatedAt *time.Time `json:"created_at"`
}

func (users) TableName() string {
	return "users"
}

type books struct {
	Id string `json:"id"`
}

func (books) TableName() string {
	return "books"
}

func TestNewCreateTable(t *testing.T) {
	db, err := sql.Open("postgres", "user=postgres password=root sslmode=disable database=testapp")
	if err != nil {
		panic(err)
	}

	// ----------------------------------------------------------------------------------
	tbl := Table[users](db)

	err = tbl.BuildSchema().
		AddColumn(columns.NewInteger("id").PrimaryKey().Identity()).
		AddColumn(columns.NewVarchar("first_name", 20).NotNull()).
		AddColumn(columns.NewVarchar("last_name", 20).NotNull()).
		AddColumn(columns.NewTimestamp("created_at").NotNull().Default("now()")).
		AddColumn(columns.NewInteger("age").NotNull().Default("18")).
		Build(true)

	tbl2 := Table[books](db)

	err = tbl2.BuildSchema().
		AddColumn(columns.NewText("id").PrimaryKey()).
		Build(true)

	if err != nil {
		panic(err)
	}
	// n := time.Now()

	u, err := tbl.Insert(users{
		FirstName: "Ali",
		LastName:  "Dehkharghani",
	}, options.NewInsert().IgnoreFields("id", "created_at"))
	if err != nil {
		panic(err)
	} else {
		fmt.Println(u.Id)
	}

	u, err = tbl.Insert(users{
		FirstName: "Hamed",
		LastName:  "Mehrara",
	}, options.NewInsert().IgnoreFields("id", "created_at"))
	if err != nil {
		panic(err)
	} else {
		fmt.Println("user_id", u.Id)
	}

	_, err = tbl2.Insert(books{
		Id: "tesssting",
	}, options.NewInsert().IgnoreFields("created_at"))
	if err != nil {
		panic(err)
	}

	data, err := tbl.Query(func(builder *Builder) {
		builder.Where("first_name", "=", "Ali")
	}).FindOne()
	if err != nil {
		panic(err)
	}

	all, err := tbl.Query(nil).Find()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
	fmt.Println(all)
	// ----------------------------------------------------------------------------------
}
