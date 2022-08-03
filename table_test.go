package elegant

import (
	"database/sql"
	"fmt"
	"github.com/aliforever/go-elegant/columns"
	"github.com/aliforever/go-elegant/options"
	_ "github.com/lib/pq"
	"testing"
)

type users struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (users) TableName() string {
	return "users"
}

func TestNewCreateTable(t *testing.T) {
	db, err := sql.Open("postgres", "user=postgres password=root sslmode=disable database=testapp")
	if err != nil {
		panic(err)
	}

	// ----------------------------------------------------------------------------------
	tbl := Table[users](db)

	err = tbl.DropTable()

	err = tbl.BuildSchema().
		AddColumn(columns.NewInteger("id").PrimaryKey().Identity()).
		AddColumn(columns.NewVarchar("first_name", 20).NotNull()).
		AddColumn(columns.NewVarchar("last_name", 20).NotNull()).
		Build()

	err = tbl.Insert(users{
		FirstName: "Ali",
		LastName:  "Dehkharghani",
	}, options.NewInsert().IgnoreFields("id"))
	if err != nil {
		panic(err)
	}

	err = tbl.Insert(users{
		FirstName: "Hamed",
		LastName:  "Mehrara",
	}, options.NewInsert().IgnoreFields("id"))
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
