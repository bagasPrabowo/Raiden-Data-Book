package models

import (
	"time"
	"github.com/sev-2/raiden"
)

type Categories struct {
	raiden.ModelBase
	Id int64 `json:"id,omitempty" column:"name:id;type:bigint;primaryKey;autoIncrement;nullable:false"`
	Name *string `json:"name,omitempty" column:"name:name;type:varchar;nullable"`
	CreatedAt *time.Time `json:"created_at,omitempty" column:"name:created_at;type:timestampz;nullable;default:now()"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" column:"name:updated_at;type:timestampz;nullable;default:now()"`

	// Table information
	Metadata string `json:"-" schema:"public" rlsEnable:"true" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	Books []*Books `json:"books,omitempty" join:"joinType:manyToMany;through:book_categories;sourcePrimaryKey:id;sourceForeignKey:category_id;targetPrimaryKey:id;targetForeign:category_id"`
	BookCategories []*BookCategories `json:"book_categories,omitempty" join:"joinType:hasMany;primaryKey:id;foreignKey:category_id"`
}
