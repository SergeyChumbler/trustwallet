package store

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Store_GetAbsentKey(t *testing.T) {
	s := NewStore()

	_, err := s.Get("x")
	assert.ErrorIs(t, err, ErrKeyNotSet)
}

func Test_Store_SetValue(t *testing.T) {
	s := NewStore()
	s.Set("x", "123")

	v, err := s.Get("x")
	assert.NoError(t, err)
	assert.Equal(t, v, "123")
}

func Test_Store_DeleteAbsentKey(t *testing.T) {
	s := NewStore()

	err := s.Delete("x")
	assert.ErrorIs(t, err, ErrKeyNotSet)
}

func Test_Store_DeleteKey(t *testing.T) {
	s := NewStore()

	s.Set("x", "123")
	err := s.Delete("x")
	assert.NoError(t, err)

	_, err = s.Get("x")
	assert.ErrorIs(t, err, ErrKeyNotSet)
}

func Test_Store_GetZeroCount(t *testing.T) {
	s := NewStore()

	count := s.Count("123")
	assert.Equal(t, 0, count)
}

func Test_Store_GetCount(t *testing.T) {
	s := NewStore()
	s.Set("x", "123")

	count := s.Count("123")
	assert.Equal(t, 1, count)
}

func Test_Store_GetMultipleCount(t *testing.T) {
	s := NewStore()
	s.Set("x", "123")
	s.Set("y", "123")
	s.Set("z", "123")

	count := s.Count("123")
	assert.Equal(t, 3, count)
}

func Test_Store_MultiSetAndGet(t *testing.T) {
	s := NewStore()
	s.Set("x", "123")
	s.Set("y", "456")

	v, err := s.Get("x")
	assert.NoError(t, err)
	assert.Equal(t, v, "123")

	v, err = s.Get("y")
	assert.NoError(t, err)
	assert.Equal(t, v, "456")
}

func Test_Store_NoTransactionCommit(t *testing.T) {
	s := NewStore()
	s.Set("x", "123")

	err := s.Commit()
	assert.ErrorIs(t, err, ErrNoTransaction)
}

func Test_Store_NoTransactionRollback(t *testing.T) {
	s := NewStore()
	s.Set("x", "123")

	err := s.Rollback()
	assert.ErrorIs(t, err, ErrNoTransaction)
}

func Test_Store_DeleteNotExistedKey(t *testing.T) {
	s := NewStore()
	err := s.Delete("x")

	assert.ErrorIs(t, err, ErrKeyNotSet)
}

func Test_Store_DeleteInTransaction(t *testing.T) {
	s := NewStore()
	s.Set("x", "123")
	s.Begin()

	_ = s.Delete("x")
	_, err := s.Get("x")

	//assert.Equal(t, v, "213")
	assert.ErrorIs(t, err, ErrKeyNotSet)
}

func Test_Store_TransactionCommit(t *testing.T) {
	s := NewStore()
	s.Set("x", "123")
	s.Begin()

	v, err := s.Get("x")
	assert.NoError(t, err)
	assert.Equal(t, v, "123")

	s.Set("x", "456")
	v, err = s.Get("x")
	assert.NoError(t, err)
	assert.Equal(t, v, "456")

	s.Set("y", "789")
	v, err = s.Get("y")
	assert.NoError(t, err)
	assert.Equal(t, v, "789")

	err = s.Commit()
	assert.NoError(t, err)

	v, err = s.Get("x")
	assert.NoError(t, err)
	assert.Equal(t, v, "456")

	v, err = s.Get("y")
	assert.NoError(t, err)
	assert.Equal(t, v, "789")
}

func Test_Store_TransactionRollback(t *testing.T) {
	s := NewStore()
	s.Set("x", "123")
	s.Begin()

	s.Set("x", "456")
	v, err := s.Get("x")
	assert.NoError(t, err)
	assert.Equal(t, v, "456")

	err = s.Rollback()
	assert.NoError(t, err)

	v, err = s.Get("x")
	assert.NoError(t, err)
	assert.Equal(t, v, "123")
}

func Test_Store_NestedTransactionOverrideValue(t *testing.T) {
	s := NewStore()
	s.Set("x", "123")

	s.Begin()
	s.Set("x", "456")

	s.Begin()
	s.Set("x", "789")

	_ = s.Commit()

	v, err := s.Get("x")
	assert.NoError(t, err)
	assert.Equal(t, v, "789")

	err = s.Commit()

	v, err = s.Get("x")
	assert.NoError(t, err)
	assert.Equal(t, v, "789")

	err = s.Commit()
	assert.ErrorIs(t, err, ErrNoTransaction)
}

func Test_Store_DeleteFromNestedTransactionCommit(t *testing.T) {
	s := NewStore()
	s.Set("x", "123")

	s.Begin()
	s.Set("x", "456")

	s.Begin()
	s.Set("x", "789")
	err := s.Delete("x")
	err = s.Commit()

	_, err = s.Get("x")
	assert.ErrorIs(t, err, ErrKeyNotSet)

	err = s.Commit()

	_, err = s.Get("x")
	assert.ErrorIs(t, err, ErrKeyNotSet)
}

func Test_Store_DeleteFromNestedTransactionRollback(t *testing.T) {
	s := NewStore()
	s.Set("x", "123")

	s.Begin()
	s.Set("x", "456")

	s.Begin()
	s.Set("x", "789")
	err := s.Delete("x")
	err = s.Commit()

	_, err = s.Get("x")
	assert.ErrorIs(t, err, ErrKeyNotSet)

	err = s.Rollback()

	v, err := s.Get("x")
	assert.NoError(t, err)
	assert.Equal(t, v, "123")
}

func Test_Store_ComplexFlowTest(t *testing.T) {
	s := NewStore()
	s.Set("x", "123")
	s.Set("y", "456")
	s.Begin()
	s.Set("x", "789")
	s.Set("y", "123")
	s.Begin()
	s.Set("z", "333")
	s.Set("t", "444")
	_ = s.Delete("y")
	s.Begin()
	s.Set("y", "555")
	_ = s.Rollback()
	s.Set("x", "666")
	_ = s.Commit()
	s.Set("y", "777")
	_ = s.Delete("t")
	_ = s.Commit()

	x, err := s.Get("x")
	assert.NoError(t, err)
	y, err := s.Get("y")
	assert.NoError(t, err)
	z, err := s.Get("z")
	assert.NoError(t, err)
	_, err = s.Get("t")
	assert.ErrorIs(t, err, ErrKeyNotSet)

	assert.Equal(t, x, "666")
	assert.Equal(t, y, "777")
	assert.Equal(t, z, "333")

	err = s.Commit()
	assert.ErrorIs(t, err, ErrNoTransaction)
}

func Test_Store_CountAfterDeleteInTransaction(t *testing.T) {
	s := NewStore()
	s.Set("x", "123")
	s.Set("y", "123")
	s.Set("z", "123")
	s.Begin()
	_ = s.Delete("y")

	count := s.Count("123")
	assert.Equal(t, 2, count)
}

func Test_Store_CountInTransactionsWithRollbacks(t *testing.T) {
	s := NewStore()
	s.Set("x", "123")
	s.Begin()
	s.Set("y", "123")
	s.Set("t", "123")

	count := s.Count("123")
	assert.Equal(t, 3, count)

	s.Begin()
	s.Set("x", "345")
	s.Set("z", "123")
	s.Set("d", "123")
	s.Set("t", "444")

	count = s.Count("123")
	assert.Equal(t, 3, count)

	_ = s.Rollback()

	count = s.Count("123")
	assert.Equal(t, 3, count)

	s.Set("s", "123")
	_ = s.Commit()
	s.Set("y", "777")
	_ = s.Delete("t")
	count = s.Count("123")
	assert.Equal(t, 2, count)

	_ = s.Commit()

	count = s.Count("123")
	assert.Equal(t, 2, count)
}
