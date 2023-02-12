package mappers

import (
	"ginskill/src/dbs"
	"gorm.io/gorm"
)

type SqlMapper struct {
	Sql  string
	Args []interface{}
	db   *gorm.DB
}

func (this *SqlMapper) setDB(db *gorm.DB) {
	this.db = db
}

func NewSqlMapper(sql string, args []interface{}) *SqlMapper {
	return &SqlMapper{Sql: sql, Args: args}
}
func Mapper(sql string, args []interface{}, err error) *SqlMapper {
	if err != nil {
		panic(err.Error())
	}
	return NewSqlMapper(sql, args)
}

// select
func (this *SqlMapper) Query() *gorm.DB {
	if this.db != nil {
		this.db.Raw(this.Sql, this.Args)
	}
	return dbs.Orm.Raw(this.Sql, this.Args)
}

// update delete insert
func (this *SqlMapper) Exec() *gorm.DB {
	if this.db != nil {
		this.db.Raw(this.Sql, this.Args)
	}
	return dbs.Orm.Exec(this.Sql, this.Args)
}

type SqlMappers []*SqlMapper

func Mappers(SqlMappers ...*SqlMapper) (list SqlMappers) {
	list = SqlMappers
	return
}
func (this SqlMappers) apply(tx *gorm.DB) {
	for _, sql := range this {
		sql.setDB(tx)
	}
}
func (this SqlMappers) Exec(f func() error) error {
	return dbs.Orm.Transaction(func(tx *gorm.DB) error {
		this.apply(tx)
		return f()
	})
}
