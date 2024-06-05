package models

import (
	"time"
	"github.com/sev-2/raiden"
)

type Books struct {
	raiden.ModelBase
	Id int64 `json:"id,omitempty" column:"name:id;type:bigint;primaryKey;autoIncrement;nullable:false"`
	CreatedAt time.Time `json:"created_at,omitempty" column:"name:created_at;type:timestampz;nullable:false;default:now()"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" column:"name:updated_at;type:timestampz;nullable;default:now()"`
	Title *string `json:"title,omitempty" column:"name:title;type:text;nullable"`
	Writer *string `json:"writer,omitempty" column:"name:writer;type:text;nullable"`
	PublishedAt *time.Time `json:"published_at,omitempty" column:"name:published_at;type:date;nullable"`
	Page *int64 `json:"page,omitempty" column:"name:page;type:bigint;nullable"`
	Description *string `json:"description,omitempty" column:"name:description;type:text;nullable"`

	// Table information
	Metadata string `json:"-" schema:"public" rlsEnable:"true" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"anon" write:"anon" readUsing:"true" writeCheck:"true" writeUsing:"true"`

	// Relations
	Categories []*Categories `json:"categories,omitempty" join:"joinType:manyToMany;through:book_categories;sourcePrimaryKey:id;sourceForeignKey:book_id;targetPrimaryKey:id;targetForeign:book_id"`
	BookCategories []*BookCategories `json:"book_categories,omitempty" join:"joinType:hasMany;primaryKey:id;foreignKey:book_id"`
}
