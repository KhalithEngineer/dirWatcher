// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package dbmodels

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/v4/types"
	"github.com/volatiletech/strmangle"
)

// TaskRun is an object representing the database table.
type TaskRun struct {
	ID               int               `boil:"id" json:"id" toml:"id" yaml:"id"`
	StartTime        null.Time         `boil:"start_time" json:"start_time,omitempty" toml:"start_time" yaml:"start_time,omitempty"`
	EndTime          null.Time         `boil:"end_time" json:"end_time,omitempty" toml:"end_time" yaml:"end_time,omitempty"`
	TotalRuntime     null.String       `boil:"total_runtime" json:"total_runtime,omitempty" toml:"total_runtime" yaml:"total_runtime,omitempty"`
	DirectoryPath    null.String       `boil:"directory_path" json:"directory_path,omitempty" toml:"directory_path" yaml:"directory_path,omitempty"`
	Interval         null.String       `boil:"interval" json:"interval,omitempty" toml:"interval" yaml:"interval,omitempty"`
	MagicString      null.String       `boil:"magic_string" json:"magic_string,omitempty" toml:"magic_string" yaml:"magic_string,omitempty"`
	FilesAdded       types.StringArray `boil:"files_added" json:"files_added,omitempty" toml:"files_added" yaml:"files_added,omitempty"`
	FilesDeleted     types.StringArray `boil:"files_deleted" json:"files_deleted,omitempty" toml:"files_deleted" yaml:"files_deleted,omitempty"`
	TotalOccurrences null.Int          `boil:"total_occurrences" json:"total_occurrences,omitempty" toml:"total_occurrences" yaml:"total_occurrences,omitempty"`
	Status           null.String       `boil:"status" json:"status,omitempty" toml:"status" yaml:"status,omitempty"`

	R *taskRunR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L taskRunL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var TaskRunColumns = struct {
	ID               string
	StartTime        string
	EndTime          string
	TotalRuntime     string
	DirectoryPath    string
	Interval         string
	MagicString      string
	FilesAdded       string
	FilesDeleted     string
	TotalOccurrences string
	Status           string
}{
	ID:               "id",
	StartTime:        "start_time",
	EndTime:          "end_time",
	TotalRuntime:     "total_runtime",
	DirectoryPath:    "directory_path",
	Interval:         "interval",
	MagicString:      "magic_string",
	FilesAdded:       "files_added",
	FilesDeleted:     "files_deleted",
	TotalOccurrences: "total_occurrences",
	Status:           "status",
}

var TaskRunTableColumns = struct {
	ID               string
	StartTime        string
	EndTime          string
	TotalRuntime     string
	DirectoryPath    string
	Interval         string
	MagicString      string
	FilesAdded       string
	FilesDeleted     string
	TotalOccurrences string
	Status           string
}{
	ID:               "task_run.id",
	StartTime:        "task_run.start_time",
	EndTime:          "task_run.end_time",
	TotalRuntime:     "task_run.total_runtime",
	DirectoryPath:    "task_run.directory_path",
	Interval:         "task_run.interval",
	MagicString:      "task_run.magic_string",
	FilesAdded:       "task_run.files_added",
	FilesDeleted:     "task_run.files_deleted",
	TotalOccurrences: "task_run.total_occurrences",
	Status:           "task_run.status",
}

// Generated where

type whereHelpernull_Time struct{ field string }

func (w whereHelpernull_Time) EQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Time) NEQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Time) LT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Time) LTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Time) GT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Time) GTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

func (w whereHelpernull_Time) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Time) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

type whereHelpertypes_StringArray struct{ field string }

func (w whereHelpertypes_StringArray) EQ(x types.StringArray) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpertypes_StringArray) NEQ(x types.StringArray) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpertypes_StringArray) LT(x types.StringArray) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertypes_StringArray) LTE(x types.StringArray) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertypes_StringArray) GT(x types.StringArray) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertypes_StringArray) GTE(x types.StringArray) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

func (w whereHelpertypes_StringArray) IsNull() qm.QueryMod { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpertypes_StringArray) IsNotNull() qm.QueryMod {
	return qmhelper.WhereIsNotNull(w.field)
}

var TaskRunWhere = struct {
	ID               whereHelperint
	StartTime        whereHelpernull_Time
	EndTime          whereHelpernull_Time
	TotalRuntime     whereHelpernull_String
	DirectoryPath    whereHelpernull_String
	Interval         whereHelpernull_String
	MagicString      whereHelpernull_String
	FilesAdded       whereHelpertypes_StringArray
	FilesDeleted     whereHelpertypes_StringArray
	TotalOccurrences whereHelpernull_Int
	Status           whereHelpernull_String
}{
	ID:               whereHelperint{field: "\"task_run\".\"id\""},
	StartTime:        whereHelpernull_Time{field: "\"task_run\".\"start_time\""},
	EndTime:          whereHelpernull_Time{field: "\"task_run\".\"end_time\""},
	TotalRuntime:     whereHelpernull_String{field: "\"task_run\".\"total_runtime\""},
	DirectoryPath:    whereHelpernull_String{field: "\"task_run\".\"directory_path\""},
	Interval:         whereHelpernull_String{field: "\"task_run\".\"interval\""},
	MagicString:      whereHelpernull_String{field: "\"task_run\".\"magic_string\""},
	FilesAdded:       whereHelpertypes_StringArray{field: "\"task_run\".\"files_added\""},
	FilesDeleted:     whereHelpertypes_StringArray{field: "\"task_run\".\"files_deleted\""},
	TotalOccurrences: whereHelpernull_Int{field: "\"task_run\".\"total_occurrences\""},
	Status:           whereHelpernull_String{field: "\"task_run\".\"status\""},
}

// TaskRunRels is where relationship names are stored.
var TaskRunRels = struct {
}{}

// taskRunR is where relationships are stored.
type taskRunR struct {
}

// NewStruct creates a new relationship struct
func (*taskRunR) NewStruct() *taskRunR {
	return &taskRunR{}
}

// taskRunL is where Load methods for each relationship are stored.
type taskRunL struct{}

var (
	taskRunAllColumns            = []string{"id", "start_time", "end_time", "total_runtime", "directory_path", "interval", "magic_string", "files_added", "files_deleted", "total_occurrences", "status"}
	taskRunColumnsWithoutDefault = []string{}
	taskRunColumnsWithDefault    = []string{"id", "start_time", "end_time", "total_runtime", "directory_path", "interval", "magic_string", "files_added", "files_deleted", "total_occurrences", "status"}
	taskRunPrimaryKeyColumns     = []string{"id"}
	taskRunGeneratedColumns      = []string{}
)

type (
	// TaskRunSlice is an alias for a slice of pointers to TaskRun.
	// This should almost always be used instead of []TaskRun.
	TaskRunSlice []*TaskRun
	// TaskRunHook is the signature for custom TaskRun hook methods
	TaskRunHook func(context.Context, boil.ContextExecutor, *TaskRun) error

	taskRunQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	taskRunType                 = reflect.TypeOf(&TaskRun{})
	taskRunMapping              = queries.MakeStructMapping(taskRunType)
	taskRunPrimaryKeyMapping, _ = queries.BindMapping(taskRunType, taskRunMapping, taskRunPrimaryKeyColumns)
	taskRunInsertCacheMut       sync.RWMutex
	taskRunInsertCache          = make(map[string]insertCache)
	taskRunUpdateCacheMut       sync.RWMutex
	taskRunUpdateCache          = make(map[string]updateCache)
	taskRunUpsertCacheMut       sync.RWMutex
	taskRunUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var taskRunAfterSelectMu sync.Mutex
var taskRunAfterSelectHooks []TaskRunHook

var taskRunBeforeInsertMu sync.Mutex
var taskRunBeforeInsertHooks []TaskRunHook
var taskRunAfterInsertMu sync.Mutex
var taskRunAfterInsertHooks []TaskRunHook

var taskRunBeforeUpdateMu sync.Mutex
var taskRunBeforeUpdateHooks []TaskRunHook
var taskRunAfterUpdateMu sync.Mutex
var taskRunAfterUpdateHooks []TaskRunHook

var taskRunBeforeDeleteMu sync.Mutex
var taskRunBeforeDeleteHooks []TaskRunHook
var taskRunAfterDeleteMu sync.Mutex
var taskRunAfterDeleteHooks []TaskRunHook

var taskRunBeforeUpsertMu sync.Mutex
var taskRunBeforeUpsertHooks []TaskRunHook
var taskRunAfterUpsertMu sync.Mutex
var taskRunAfterUpsertHooks []TaskRunHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *TaskRun) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range taskRunAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *TaskRun) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range taskRunBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *TaskRun) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range taskRunAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *TaskRun) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range taskRunBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *TaskRun) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range taskRunAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *TaskRun) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range taskRunBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *TaskRun) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range taskRunAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *TaskRun) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range taskRunBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *TaskRun) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range taskRunAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddTaskRunHook registers your hook function for all future operations.
func AddTaskRunHook(hookPoint boil.HookPoint, taskRunHook TaskRunHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		taskRunAfterSelectMu.Lock()
		taskRunAfterSelectHooks = append(taskRunAfterSelectHooks, taskRunHook)
		taskRunAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		taskRunBeforeInsertMu.Lock()
		taskRunBeforeInsertHooks = append(taskRunBeforeInsertHooks, taskRunHook)
		taskRunBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		taskRunAfterInsertMu.Lock()
		taskRunAfterInsertHooks = append(taskRunAfterInsertHooks, taskRunHook)
		taskRunAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		taskRunBeforeUpdateMu.Lock()
		taskRunBeforeUpdateHooks = append(taskRunBeforeUpdateHooks, taskRunHook)
		taskRunBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		taskRunAfterUpdateMu.Lock()
		taskRunAfterUpdateHooks = append(taskRunAfterUpdateHooks, taskRunHook)
		taskRunAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		taskRunBeforeDeleteMu.Lock()
		taskRunBeforeDeleteHooks = append(taskRunBeforeDeleteHooks, taskRunHook)
		taskRunBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		taskRunAfterDeleteMu.Lock()
		taskRunAfterDeleteHooks = append(taskRunAfterDeleteHooks, taskRunHook)
		taskRunAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		taskRunBeforeUpsertMu.Lock()
		taskRunBeforeUpsertHooks = append(taskRunBeforeUpsertHooks, taskRunHook)
		taskRunBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		taskRunAfterUpsertMu.Lock()
		taskRunAfterUpsertHooks = append(taskRunAfterUpsertHooks, taskRunHook)
		taskRunAfterUpsertMu.Unlock()
	}
}

// OneG returns a single taskRun record from the query using the global executor.
func (q taskRunQuery) OneG(ctx context.Context) (*TaskRun, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single taskRun record from the query.
func (q taskRunQuery) One(ctx context.Context, exec boil.ContextExecutor) (*TaskRun, error) {
	o := &TaskRun{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "dbmodels: failed to execute a one query for task_run")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all TaskRun records from the query using the global executor.
func (q taskRunQuery) AllG(ctx context.Context) (TaskRunSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all TaskRun records from the query.
func (q taskRunQuery) All(ctx context.Context, exec boil.ContextExecutor) (TaskRunSlice, error) {
	var o []*TaskRun

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "dbmodels: failed to assign all query results to TaskRun slice")
	}

	if len(taskRunAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all TaskRun records in the query using the global executor
func (q taskRunQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all TaskRun records in the query.
func (q taskRunQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: failed to count task_run rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table using the global executor.
func (q taskRunQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q taskRunQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "dbmodels: failed to check if task_run exists")
	}

	return count > 0, nil
}

// TaskRuns retrieves all the records using an executor.
func TaskRuns(mods ...qm.QueryMod) taskRunQuery {
	mods = append(mods, qm.From("\"task_run\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"task_run\".*"})
	}

	return taskRunQuery{q}
}

// FindTaskRunG retrieves a single record by ID.
func FindTaskRunG(ctx context.Context, iD int, selectCols ...string) (*TaskRun, error) {
	return FindTaskRun(ctx, boil.GetContextDB(), iD, selectCols...)
}

// FindTaskRun retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTaskRun(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*TaskRun, error) {
	taskRunObj := &TaskRun{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"task_run\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, taskRunObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "dbmodels: unable to select from task_run")
	}

	if err = taskRunObj.doAfterSelectHooks(ctx, exec); err != nil {
		return taskRunObj, err
	}

	return taskRunObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *TaskRun) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *TaskRun) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("dbmodels: no task_run provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(taskRunColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	taskRunInsertCacheMut.RLock()
	cache, cached := taskRunInsertCache[key]
	taskRunInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			taskRunAllColumns,
			taskRunColumnsWithDefault,
			taskRunColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(taskRunType, taskRunMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(taskRunType, taskRunMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"task_run\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"task_run\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "dbmodels: unable to insert into task_run")
	}

	if !cached {
		taskRunInsertCacheMut.Lock()
		taskRunInsertCache[key] = cache
		taskRunInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// UpdateG a single TaskRun record using the global executor.
// See Update for more documentation.
func (o *TaskRun) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the TaskRun.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *TaskRun) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	taskRunUpdateCacheMut.RLock()
	cache, cached := taskRunUpdateCache[key]
	taskRunUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			taskRunAllColumns,
			taskRunPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("dbmodels: unable to update task_run, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"task_run\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, taskRunPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(taskRunType, taskRunMapping, append(wl, taskRunPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to update task_run row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: failed to get rows affected by update for task_run")
	}

	if !cached {
		taskRunUpdateCacheMut.Lock()
		taskRunUpdateCache[key] = cache
		taskRunUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q taskRunQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q taskRunQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to update all for task_run")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to retrieve rows affected for task_run")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o TaskRunSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o TaskRunSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("dbmodels: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), taskRunPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"task_run\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, taskRunPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to update all in taskRun slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to retrieve rows affected all in update all taskRun")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *TaskRun) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns, opts...)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *TaskRun) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("dbmodels: no task_run provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(taskRunColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	taskRunUpsertCacheMut.RLock()
	cache, cached := taskRunUpsertCache[key]
	taskRunUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			taskRunAllColumns,
			taskRunColumnsWithDefault,
			taskRunColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			taskRunAllColumns,
			taskRunPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("dbmodels: unable to upsert task_run, could not build update column list")
		}

		ret := strmangle.SetComplement(taskRunAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(taskRunPrimaryKeyColumns) == 0 {
				return errors.New("dbmodels: unable to upsert task_run, could not build conflict column list")
			}

			conflict = make([]string, len(taskRunPrimaryKeyColumns))
			copy(conflict, taskRunPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"task_run\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(taskRunType, taskRunMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(taskRunType, taskRunMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "dbmodels: unable to upsert task_run")
	}

	if !cached {
		taskRunUpsertCacheMut.Lock()
		taskRunUpsertCache[key] = cache
		taskRunUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// DeleteG deletes a single TaskRun record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *TaskRun) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single TaskRun record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *TaskRun) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("dbmodels: no TaskRun provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), taskRunPrimaryKeyMapping)
	sql := "DELETE FROM \"task_run\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to delete from task_run")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: failed to get rows affected by delete for task_run")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

func (q taskRunQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q taskRunQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("dbmodels: no taskRunQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to delete all from task_run")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: failed to get rows affected by deleteall for task_run")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o TaskRunSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o TaskRunSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(taskRunBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), taskRunPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"task_run\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, taskRunPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to delete all from taskRun slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: failed to get rows affected by deleteall for task_run")
	}

	if len(taskRunAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *TaskRun) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("dbmodels: no TaskRun provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *TaskRun) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindTaskRun(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TaskRunSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("dbmodels: empty TaskRunSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TaskRunSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := TaskRunSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), taskRunPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"task_run\".* FROM \"task_run\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, taskRunPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "dbmodels: unable to reload all in TaskRunSlice")
	}

	*o = slice

	return nil
}

// TaskRunExistsG checks if the TaskRun row exists.
func TaskRunExistsG(ctx context.Context, iD int) (bool, error) {
	return TaskRunExists(ctx, boil.GetContextDB(), iD)
}

// TaskRunExists checks if the TaskRun row exists.
func TaskRunExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"task_run\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "dbmodels: unable to check if task_run exists")
	}

	return exists, nil
}

// Exists checks if the TaskRun row exists.
func (o *TaskRun) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return TaskRunExists(ctx, exec, o.ID)
}
