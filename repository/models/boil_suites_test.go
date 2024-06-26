// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package dbmodels

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Configs", testConfigs)
	t.Run("TaskRuns", testTaskRuns)
}

func TestDelete(t *testing.T) {
	t.Run("Configs", testConfigsDelete)
	t.Run("TaskRuns", testTaskRunsDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Configs", testConfigsQueryDeleteAll)
	t.Run("TaskRuns", testTaskRunsQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Configs", testConfigsSliceDeleteAll)
	t.Run("TaskRuns", testTaskRunsSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Configs", testConfigsExists)
	t.Run("TaskRuns", testTaskRunsExists)
}

func TestFind(t *testing.T) {
	t.Run("Configs", testConfigsFind)
	t.Run("TaskRuns", testTaskRunsFind)
}

func TestBind(t *testing.T) {
	t.Run("Configs", testConfigsBind)
	t.Run("TaskRuns", testTaskRunsBind)
}

func TestOne(t *testing.T) {
	t.Run("Configs", testConfigsOne)
	t.Run("TaskRuns", testTaskRunsOne)
}

func TestAll(t *testing.T) {
	t.Run("Configs", testConfigsAll)
	t.Run("TaskRuns", testTaskRunsAll)
}

func TestCount(t *testing.T) {
	t.Run("Configs", testConfigsCount)
	t.Run("TaskRuns", testTaskRunsCount)
}

func TestHooks(t *testing.T) {
	t.Run("Configs", testConfigsHooks)
	t.Run("TaskRuns", testTaskRunsHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Configs", testConfigsInsert)
	t.Run("Configs", testConfigsInsertWhitelist)
	t.Run("TaskRuns", testTaskRunsInsert)
	t.Run("TaskRuns", testTaskRunsInsertWhitelist)
}

func TestReload(t *testing.T) {
	t.Run("Configs", testConfigsReload)
	t.Run("TaskRuns", testTaskRunsReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Configs", testConfigsReloadAll)
	t.Run("TaskRuns", testTaskRunsReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Configs", testConfigsSelect)
	t.Run("TaskRuns", testTaskRunsSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Configs", testConfigsUpdate)
	t.Run("TaskRuns", testTaskRunsUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Configs", testConfigsSliceUpdateAll)
	t.Run("TaskRuns", testTaskRunsSliceUpdateAll)
}
