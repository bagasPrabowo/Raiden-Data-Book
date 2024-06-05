package models

import (
	"time"
	"github.com/sev-2/raiden"
)

type BookCategories struct {
	raiden.ModelBase
	Id int64 `json:"id,omitempty" column:"name:id;type:bigint;primaryKey;autoIncrement;nullable:false"`
	CreatedAt time.Time `json:"created_at,omitempty" column:"name:created_at;type:timestampz;nullable:false;default:now()"`
	BookId *int64 `json:"book_id,omitempty" column:"name:book_id;type:bigint;nullable"`
	CategoryId *int64 `json:"category_id,omitempty" column:"name:category_id;type:bigint;nullable"`

	// Table information
	Metadata string `json:"-" schema:"public" rlsEnable:"true" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	Books *Books `json:"books,omitempty" join:"joinType:hasOne;primaryKey:id;foreignKey:book_id"`
	Categories *Categories `json:"categories,omitempty" join:"joinType:hasOne;primaryKey:id;foreignKey:category_id"`
}
