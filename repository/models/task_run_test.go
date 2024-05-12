// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package dbmodels

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testTaskRuns(t *testing.T) {
	t.Parallel()

	query := TaskRuns()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testTaskRunsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TaskRun{}
	if err = randomize.Struct(seed, o, taskRunDBTypes, true, taskRunColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaskRun struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := TaskRuns().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTaskRunsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TaskRun{}
	if err = randomize.Struct(seed, o, taskRunDBTypes, true, taskRunColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaskRun struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := TaskRuns().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := TaskRuns().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTaskRunsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TaskRun{}
	if err = randomize.Struct(seed, o, taskRunDBTypes, true, taskRunColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaskRun struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := TaskRunSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := TaskRuns().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTaskRunsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TaskRun{}
	if err = randomize.Struct(seed, o, taskRunDBTypes, true, taskRunColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaskRun struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := TaskRunExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if TaskRun exists: %s", err)
	}
	if !e {
		t.Errorf("Expected TaskRunExists to return true, but got false.")
	}
}

func testTaskRunsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TaskRun{}
	if err = randomize.Struct(seed, o, taskRunDBTypes, true, taskRunColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaskRun struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	taskRunFound, err := FindTaskRun(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if taskRunFound == nil {
		t.Error("want a record, got nil")
	}
}

func testTaskRunsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TaskRun{}
	if err = randomize.Struct(seed, o, taskRunDBTypes, true, taskRunColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaskRun struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = TaskRuns().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testTaskRunsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TaskRun{}
	if err = randomize.Struct(seed, o, taskRunDBTypes, true, taskRunColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaskRun struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := TaskRuns().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testTaskRunsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	taskRunOne := &TaskRun{}
	taskRunTwo := &TaskRun{}
	if err = randomize.Struct(seed, taskRunOne, taskRunDBTypes, false, taskRunColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaskRun struct: %s", err)
	}
	if err = randomize.Struct(seed, taskRunTwo, taskRunDBTypes, false, taskRunColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaskRun struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = taskRunOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = taskRunTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := TaskRuns().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testTaskRunsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	taskRunOne := &TaskRun{}
	taskRunTwo := &TaskRun{}
	if err = randomize.Struct(seed, taskRunOne, taskRunDBTypes, false, taskRunColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaskRun struct: %s", err)
	}
	if err = randomize.Struct(seed, taskRunTwo, taskRunDBTypes, false, taskRunColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaskRun struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = taskRunOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = taskRunTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := TaskRuns().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func taskRunBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *TaskRun) error {
	*o = TaskRun{}
	return nil
}

func taskRunAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *TaskRun) error {
	*o = TaskRun{}
	return nil
}

func taskRunAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *TaskRun) error {
	*o = TaskRun{}
	return nil
}

func taskRunBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *TaskRun) error {
	*o = TaskRun{}
	return nil
}

func taskRunAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *TaskRun) error {
	*o = TaskRun{}
	return nil
}

func taskRunBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *TaskRun) error {
	*o = TaskRun{}
	return nil
}

func taskRunAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *TaskRun) error {
	*o = TaskRun{}
	return nil
}

func taskRunBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *TaskRun) error {
	*o = TaskRun{}
	return nil
}

func taskRunAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *TaskRun) error {
	*o = TaskRun{}
	return nil
}

func testTaskRunsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &TaskRun{}
	o := &TaskRun{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, taskRunDBTypes, false); err != nil {
		t.Errorf("Unable to randomize TaskRun object: %s", err)
	}

	AddTaskRunHook(boil.BeforeInsertHook, taskRunBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	taskRunBeforeInsertHooks = []TaskRunHook{}

	AddTaskRunHook(boil.AfterInsertHook, taskRunAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	taskRunAfterInsertHooks = []TaskRunHook{}

	AddTaskRunHook(boil.AfterSelectHook, taskRunAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	taskRunAfterSelectHooks = []TaskRunHook{}

	AddTaskRunHook(boil.BeforeUpdateHook, taskRunBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	taskRunBeforeUpdateHooks = []TaskRunHook{}

	AddTaskRunHook(boil.AfterUpdateHook, taskRunAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	taskRunAfterUpdateHooks = []TaskRunHook{}

	AddTaskRunHook(boil.BeforeDeleteHook, taskRunBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	taskRunBeforeDeleteHooks = []TaskRunHook{}

	AddTaskRunHook(boil.AfterDeleteHook, taskRunAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	taskRunAfterDeleteHooks = []TaskRunHook{}

	AddTaskRunHook(boil.BeforeUpsertHook, taskRunBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	taskRunBeforeUpsertHooks = []TaskRunHook{}

	AddTaskRunHook(boil.AfterUpsertHook, taskRunAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	taskRunAfterUpsertHooks = []TaskRunHook{}
}

func testTaskRunsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TaskRun{}
	if err = randomize.Struct(seed, o, taskRunDBTypes, true, taskRunColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaskRun struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := TaskRuns().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTaskRunsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TaskRun{}
	if err = randomize.Struct(seed, o, taskRunDBTypes, true); err != nil {
		t.Errorf("Unable to randomize TaskRun struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(taskRunColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := TaskRuns().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTaskRunsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TaskRun{}
	if err = randomize.Struct(seed, o, taskRunDBTypes, true, taskRunColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaskRun struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testTaskRunsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TaskRun{}
	if err = randomize.Struct(seed, o, taskRunDBTypes, true, taskRunColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaskRun struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := TaskRunSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testTaskRunsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TaskRun{}
	if err = randomize.Struct(seed, o, taskRunDBTypes, true, taskRunColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaskRun struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := TaskRuns().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	taskRunDBTypes = map[string]string{`ID`: `integer`, `StartTime`: `timestamp without time zone`, `EndTime`: `timestamp without time zone`, `TotalRuntime`: `interval`, `DirectoryPath`: `text`, `Interval`: `interval`, `MagicString`: `text`, `FilesAdded`: `ARRAYtext`, `FilesDeleted`: `ARRAYtext`, `TotalOccurrences`: `integer`, `Status`: `text`}
	_              = bytes.MinRead
)

func testTaskRunsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(taskRunPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(taskRunAllColumns) == len(taskRunPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &TaskRun{}
	if err = randomize.Struct(seed, o, taskRunDBTypes, true, taskRunColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaskRun struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := TaskRuns().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, taskRunDBTypes, true, taskRunPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize TaskRun struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testTaskRunsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(taskRunAllColumns) == len(taskRunPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &TaskRun{}
	if err = randomize.Struct(seed, o, taskRunDBTypes, true, taskRunColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaskRun struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := TaskRuns().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, taskRunDBTypes, true, taskRunPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize TaskRun struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(taskRunAllColumns, taskRunPrimaryKeyColumns) {
		fields = taskRunAllColumns
	} else {
		fields = strmangle.SetComplement(
			taskRunAllColumns,
			taskRunPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := TaskRunSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testTaskRunsUpsert(t *testing.T) {
	t.Parallel()

	if len(taskRunAllColumns) == len(taskRunPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := TaskRun{}
	if err = randomize.Struct(seed, &o, taskRunDBTypes, true); err != nil {
		t.Errorf("Unable to randomize TaskRun struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert TaskRun: %s", err)
	}

	count, err := TaskRuns().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, taskRunDBTypes, false, taskRunPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize TaskRun struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert TaskRun: %s", err)
	}

	count, err = TaskRuns().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}