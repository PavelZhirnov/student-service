package model

type Teacher struct {
	ID           string `db:"id" goqu:"skipinsert,skipupdate"`
	PositionType int64  `db:"position_type"`
	FullName     string `db:"full_name"`
	StudentID    string `db:"student_id" goqu:"skipupdate"`
}
