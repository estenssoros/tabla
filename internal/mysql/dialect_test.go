package mysql

import (
	"fmt"
	"testing"

	"github.com/estenssoros/tabla/internal/gopher"
	"github.com/stretchr/testify/assert"
)

func TestDielectDropIfExists(t *testing.T) {
	s := &gopher.Struct{
		Name: "asdf",
	}
	assert.Equal(t, "DROP TABLE IF EXISTS `asdf`;\n", Dialect{}.DropIfExists(s))
}

var testDialectFieldsTables = []struct {
	fields []*gopher.Field
	err    bool
}{
	{
		[]*gopher.Field{
			&gopher.Field{
				Name: "id",
				Type: "int",
			},
		},
		false,
	},
	{
		[]*gopher.Field{
			&gopher.Field{
				Name:    "id",
				Type:    "int",
				SQLType: "datetime",
			},
			&gopher.Field{
				Name: "id",
				Type: "string",
			},
			&gopher.Field{
				Name: "id",
				Type: "time.Time",
			},
			&gopher.Field{
				Name: "id",
				Type: "bool",
			},
			&gopher.Field{
				Name: "id",
				Type: "float64",
			},
			&gopher.Field{
				Name: "id",
				Type: "uuid.UUID",
			},
		},
		false,
	},
	{
		[]*gopher.Field{
			&gopher.Field{
				Name:     "id",
				Type:     "int",
				SQLType:  "varchar",
				SQLExtra: "30",
			},
		},
		false,
	},
	{
		[]*gopher.Field{
			&gopher.Field{
				Name: "id",
			},
		},
		true,
	},
}

func TestDialectFields(t *testing.T) {
	for i, tt := range testDialectFieldsTables {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			s, err := Dialect{}.Fields(tt.fields)
			if tt.err {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				assert.NotEmpty(t, s)
			}
		})
	}
}

var testDialectCreateTables = []struct {
	in  *gopher.Struct
	err bool
}{
	{
		&gopher.Struct{},
		false,
	},
	{
		&gopher.Struct{
			Fields: []*gopher.Field{
				&gopher.Field{},
			},
		},
		true,
	},
}

func TestDialectCreate(t *testing.T) {
	for i, tt := range testDialectCreateTables {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			s, err := Dialect{}.Create(tt.in)
			if tt.err {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				assert.NotEmpty(t, s)
			}
		})
	}
}
