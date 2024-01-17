package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/doug-martin/goqu/v9"
	"github.com/jmoiron/sqlx"
	"github.com/pavelzhirnov/student-service/internal/model"
	"github.com/pavelzhirnov/student-service/pkg/logging"
)

const teacherTName = "teacher"

type TeacherRepositoryImpl struct {
	db *sqlx.DB
}

func NewTeacherRepository(db *sqlx.DB) *TeacherRepositoryImpl {
	return &TeacherRepositoryImpl{
		db: db,
	}
}

func (t *TeacherRepositoryImpl) Get(ctx context.Context, teacherID string) (*model.Teacher, error) {
	query, _, err := goqu.From(teacherTName).Where(
		goqu.I("id").Eq(teacherID),
	).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("не удалось сформировать запрос get Teacher: %w", err)
	}

	teacher := &model.Teacher{}
	if err = t.db.GetContext(ctx, teacher, query); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrEntityNotFound
		}
		return nil, fmt.Errorf("не удалось выполнить запрос get Teacher: %w", err)
	}

	return teacher, nil
}

func (t *TeacherRepositoryImpl) Create(ctx context.Context, teachers []*model.Teacher) ([]*model.Teacher, error) {
	tx, err := t.db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("error transaction: %w", err)
	}

	defer func() {
		if err != nil {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				logging.GetLogger(ctx).WithContext(ctx).WithField("err", rollbackErr).Error("не удалось выполнить Rollback teacher")
			}
			return
		}
	}()

	if len(teachers) == 0 {
		return nil, nil
	}
	rows := make([]goqu.Record, 0, len(teachers))
	for _, teacher := range teachers {
		rows = append(rows, goqu.Record{
			"position_type": teacher.PositionType,
			"full_name":     teacher.FullName,
			"student_id":    teacher.StudentID,
		})
	}

	query, _, err := goqu.Insert(teacherTName).Rows(rows).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("не удалось сформировать запрос create teachers: %w", err)
	}

	if _, err = tx.ExecContext(ctx, query); err != nil {
		return nil, fmt.Errorf("не удалось выполнить запрос create teachers: %w", err)
	}

	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("не удалось выполнить commit: %w", err)
	}

	teachersList, err := t.List(ctx, &TeacherListFilter{
		IDList:    []string{teachers[0].StudentID},
		FieldName: "student_id",
	})
	if err != nil {
		return nil, fmt.Errorf("не удалось получить созданных teachers: %w", err)
	}

	return teachersList, nil
}

func (t *TeacherRepositoryImpl) Update(ctx context.Context, teacher *model.Teacher) (*model.Teacher, error) {
	tx, err := t.db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("error transaction: %w", err)
	}

	defer func() {
		if err != nil {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				logging.GetLogger(ctx).WithContext(ctx).WithField("err", rollbackErr).Error("не удалось выполнить Rollback teacher")
			}
			return
		}
	}()

	query, _, err := goqu.Update(teacherTName).Set(
		teacher,
	).Where(
		goqu.I("id").Eq(teacher.ID),
	).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("не удалось создать Update teacher: %w", err)
	}
	if _, err = tx.ExecContext(ctx, query); err != nil {
		return nil, fmt.Errorf("не удалось выполнить запрос по обновлению teacher: %v", err)
	}

	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("не удалось выполнить commit: %w", err)
	}

	teacher, err = t.Get(ctx, teacher.ID)
	if err != nil {
		return nil, fmt.Errorf("не удалось получить teacher после обновления: %w", err)
	}

	return teacher, nil
}

func (t *TeacherRepositoryImpl) UpdateTeachers(ctx context.Context, studentID string, teachers []*model.Teacher) ([]*model.Teacher, error) {
	err := t.DeleteByStudentID(ctx, studentID)
	if err != nil {
		return nil, fmt.Errorf("не удалось удалить вложенные объекты: %w", err)
	}

	if len(teachers) > 0 {
		teachers, err = t.Create(ctx, teachers)
		if err != nil {
			return nil, fmt.Errorf("не удалось создатиь вложенные объекты: %w", err)
		}
	}

	return teachers, nil
}

func (t *TeacherRepositoryImpl) DeleteByStudentID(ctx context.Context, studentID string) error {
	tx, err := t.db.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("error transaction: %w", err)
	}

	defer func() {
		if err != nil {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				logging.GetLogger(ctx).WithContext(ctx).WithField("err", rollbackErr).Error("не удалось выполнить Rollback teacher")
			}
			return
		}
	}()

	query, _, err := goqu.Delete(teacherTName).Where(
		goqu.I("student_id").Eq(studentID),
	).ToSQL()
	if err != nil {
		return fmt.Errorf("не удалось сформирвоать запрос для удаления вложенных объектов: %w", err)
	}

	if _, err = tx.ExecContext(ctx, query); err != nil {
		return fmt.Errorf("не удалось удалить вложенные объекты: %w", err)
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("не удалось выполнить commit: %w", err)
	}
	return nil
}

type TeacherListFilter struct {
	IDList    []string
	FieldName string
}

func (t *TeacherListFilter) toDataSet() *goqu.SelectDataset {
	selectDataset := goqu.From(teacherTName)
	if t.IDList == nil {
		return selectDataset
	}

	selectDataset = selectDataset.Where(
		goqu.I(t.FieldName).In(t.IDList),
	)

	return selectDataset
}

func (t *TeacherRepositoryImpl) List(ctx context.Context, filter *TeacherListFilter) ([]*model.Teacher, error) {
	teacherList := make([]*model.Teacher, 0)

	query, _, err := filter.toDataSet().ToSQL()
	if err != nil {
		return nil, fmt.Errorf("не удалось сформировать запрос list teacher: %w", err)
	}

	if err = t.db.SelectContext(ctx, &teacherList, query); err != nil {
		return nil, fmt.Errorf("не удалось выполнить запрос list teacher: %w", err)
	}

	return teacherList, nil
}
