package mysql

import (
	"fmt"
	"testing"

	"github.com/estenssoros/tabla/internal/gopher"
	"github.com/stretchr/testify/assert"
)

func TestDielectDropIfExists(t *testing.T) {
	s := &gopher.GoStruct{
		Name: "asdf",
	}
	assert.Equal(t, "DROP TABLE IF EXISTS `asdf`;\n", Dialect{}.DropIfExists(s))
}

var testDialectFieldsTables = []struct {
	fields []*gopher.GoField
	err    bool
}{
	{
		[]*gopher.GoField{
			&gopher.GoField{
				Name: "id",
				Type: "int",
			},
		},
		false,
	},
	{
		[]*gopher.GoField{
			&gopher.GoField{
				Name:    "id",
				Type:    "int",
				SQLType: "datetime",
			},
			&gopher.GoField{
				Name: "id",
				Type: "string",
			},
			&gopher.GoField{
				Name: "id",
				Type: "time.Time",
			},
			&gopher.GoField{
				Name: "id",
				Type: "bool",
			},
			&gopher.GoField{
				Name: "id",
				Type: "float64",
			},
			&gopher.GoField{
				Name: "id",
				Type: "uuid.UUID",
			},
		},
		false,
	},
	{
		[]*gopher.GoField{
			&gopher.GoField{
				Name:     "id",
				Type:     "int",
				SQLType:  "varchar",
				SQLExtra: "30",
			},
		},
		false,
	},
	{
		[]*gopher.GoField{
			&gopher.GoField{
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
	in  *gopher.GoStruct
	err bool
}{
	{
		&gopher.GoStruct{},
		false,
	},
	{
		&gopher.GoStruct{
			Fields: []*gopher.GoField{
				&gopher.GoField{},
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
