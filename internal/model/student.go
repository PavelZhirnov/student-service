package model

type Student struct {
	ID       string     `db:"id" goqu:"skipinsert,skipupdate"`
	FullName string     `db:"full_name"`
	Age      int64      `db:"age"`
	Salary   int64      `db:"salary"`
	Teachers []*Teacher `db:"-"`
}
