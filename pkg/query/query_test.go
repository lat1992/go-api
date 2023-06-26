package query

import (
	"reflect"
	"testing"
)

func TestSql_Build(t *testing.T) {
	tests := []struct {
		name      string
		exec      func() *Sql
		wantQuery string
		wantArgs  []any
	}{
		{
			name: "nominal select",
			exec: func() *Sql {
				return Select("id", "name").From("table").Where("column1").Equal(1).And("column2").Like("value2")
			},
			wantQuery: "SELECT id,name FROM table WHERE column1 = $1 AND column2 LIKE $2",
			wantArgs:  []any{1, "value2"},
		},
		{
			name: "nominal insert",
			exec: func() *Sql {
				return Insert("table").Columns("id", "name").Values(2, "name2")
			},
			wantQuery: "INSERT INTO table (id,name) VALUES ($1,$2)",
			wantArgs:  []any{2, "name2"},
		},
		{
			name: "nominal update",
			exec: func() *Sql {
				return Update("table").Set([]ColumnValue{
					{"id", 1},
					{"name", "name1"},
				}).Where("column1").Equal(3).And("column2").Like("value2")
			},
			wantQuery: "UPDATE table SET id = $1,name = $2 WHERE column1 = $3 AND column2 LIKE $4",
			wantArgs:  []any{1, "name1", 3, "value2"},
		},
		{
			name: "nominal delete",
			exec: func() *Sql {
				return Delete("table").Where("column1").Equal(1).And("column2").Like("value2")
			},
			wantQuery: "DELETE FROM table WHERE column1 = $1 AND column2 LIKE $2",
			wantArgs:  []any{1, "value2"},
		},
		{
			name: "nominal select with join",
			exec: func() *Sql {
				return Select("id", "name").From("table").Inner("table2", "id", "id").Where("column1").Equal(1).And("column2").Like("value2")
			},
			wantQuery: "SELECT id,name FROM table INNER JOIN table2 ON table.id = table2.id WHERE column1 = $1 AND column2 LIKE $2",
			wantArgs:  []any{1, "value2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.exec()
			if got.Build() != tt.wantQuery {
				t.Errorf("Sql.Build() = %v, want %v", got.Build(), tt.wantQuery)
			}
			if !reflect.DeepEqual(got.args, tt.wantArgs) {
				t.Errorf("Sql.args = %v, want %v", got.Args(), tt.wantArgs)
			}
		})
	}
}
