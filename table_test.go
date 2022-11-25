package elegant

import (
	"database/sql"
	"fmt"
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
	tbl := Table[users](db, options.Table().SetInsertOptions(options.Insert().IgnoreFields("id", "created_at")))

	// err = tbl.BuildSchema().
	// 	AddColumn(columns.NewInteger("id").PrimaryKey().Identity()).
	// 	AddColumn(columns.NewVarchar("first_name", 20).NotNull()).
	// 	AddColumn(columns.NewVarchar("last_name", 20).NotNull()).
	// 	AddColumn(columns.NewTimestamp("created_at").NotNull().Default("now()")).
	// 	AddColumn(columns.NewInteger("age").NotNull().Default("18")).
	// 	Build(true)
	//
	// tbl2 := Table[books](db)
	//
	// err = tbl2.BuildSchema().
	// 	AddColumn(columns.NewText("id").PrimaryKey()).
	// 	Build(true)
	//
	// if err != nil {
	// 	panic(err)
	// }
	// // n := time.Now()
	//
	// u, err := tbl.Insert(users{
	// 	FirstName: "Ali",
	// 	LastName:  "Dehkharghani",
	// })
	// if err != nil {
	// 	panic(err)
	// } else {
	// 	fmt.Println(u.Id)
	// }
	//
	// u, err = tbl.Insert(users{
	// 	FirstName: "Hamed",
	// 	LastName:  "Mehrara",
	// })
	// if err != nil {
	// 	panic(err)
	// } else {
	// 	fmt.Println("user_id", u.Id)
	// }
	//
	// _, err = tbl2.Insert(books{
	// 	Id: "tesssting",
	// })
	// if err != nil {
	// 	panic(err)
	// }

	data, err := tbl.Query(func(builder *QueryBuilder) {
		builder.Where("first_name", "=", "Ali").AndGroup("last_name", "=", "Hamed", func(qb *WhereClause) {
			qb.And("first_name", "=", "H")
		})
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

func TestTbl_mapToUpdateQuery(t *testing.T) {
	db, err := sql.Open("postgres", "user=postgres password=root sslmode=disable database=testapp")
	if err != nil {
		panic(err)
	}

	// ----------------------------------------------------------------------------------
	tbl := Table[users](db, options.Table().SetInsertOptions(options.Insert().IgnoreFields("id", "created_at")))
	user, err := tbl.Insert(users{
		FirstName: "Ali",
		LastName:  "Error",
	})
	if err != nil {
		panic(err)
	}

	user.LastName = "Error - Modified"

	err = tbl.UpdateByID(user.Id, *user, options.Update().IgnoreFields("created_at"))
	if err != nil {
		panic(err)
	}
	// type fields struct {
	// 	name    string
	// 	db      *sql.DB
	// 	options *options.TableOptions
	// }
	// type args struct {
	// 	m             map[string]interface{}
	// 	ignoredFields []string
	// }
	// tests := []struct {
	// 	name            string
	// 	fields          fields
	// 	args            args
	// 	wantColumnNames string
	// 	wantValues      []interface{}
	// }{
	// 	{
	// 		name: "Successful",
	// 		fields: fields{
	// 			name:    "users",
	// 			db:      nil,
	// 			options: nil,
	// 		},
	// 		args: args{
	// 			m: map[string]interface{}{
	// 				"first_name": "Ali",
	// 			},
	// 			ignoredFields: nil,
	// 		},
	// 		wantColumnNames: "first_name=$1",
	// 		wantValues:      []interface{}{"Ali"},
	// 	},
	// }
	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		c := &Tbl[users]{
	// 			name:    tt.fields.name,
	// 			db:      tt.fields.db,
	// 			options: tt.fields.options,
	// 		}
	// 		gotColumnNames, _, gotValues := c.mapToUpdateQuery(tt.args.m, tt.args.ignoredFields...)
	//
	// 		if gotColumnNames != tt.wantColumnNames {
	// 			t.Errorf("mapToUpdateQuery() gotColumnNames = %v, want %v", gotColumnNames, tt.wantColumnNames)
	// 		}
	// 		if !reflect.DeepEqual(gotValues, tt.wantValues) {
	// 			t.Errorf("mapToUpdateQuery() gotValues = %v, want %v", gotValues, tt.wantValues)
	// 		}
	// 	})
	// }
}

func TestTbl_UpdateByID(t *testing.T) {
	type fields struct {
		name    string
		db      *sql.DB
		options *options.TableOptions
	}
	type args[T any] struct {
		id   interface{}
		data T
	}

	type test[T any] struct {
		name    string
		fields  fields
		args    args[T]
		wantErr bool
	}

	u := users{
		FirstName: "Hamed",
		LastName:  "",
		CreatedAt: nil,
	}

	tests := []test[users]{
		// TODO: Add test cases.
		{
			name: "Successful",
			fields: fields{
				name: "users",
			},
			args: args[users]{
				id:   "1",
				data: u,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Tbl[users]{
				name:    tt.fields.name,
				db:      tt.fields.db,
				options: tt.fields.options,
			}
			if err := c.UpdateByID(tt.args.id, tt.args.data, options.Update().IgnoreFields("created_at")); (err != nil) != tt.wantErr {
				t.Errorf("UpdateByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
