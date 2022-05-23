package store

import "errors"

type Store interface {
	Get(key string) (string, error)
	Set(key, value string)
	Delete(key string) error
	Count(value string) int
	Begin()
	Commit() error
	Rollback() error
}

type store struct {
	localStore    map[string]string
	currenTxScope *transactionScope
}

var ErrNoTransaction = errors.New("no transaction")
var ErrKeyNotSet = errors.New("key not set")

func NewStore() Store {
	s := &store{
		localStore:    make(map[string]string),
		currenTxScope: nil,
	}

	return s
}

func (s *store) Get(key string) (string, error) {
	return s.findByKey(key)
}

func (s *store) Set(key, value string) {
	if s.currenTxScope != nil {
		s.currenTxScope.localStore[key] = txValue{
			value:   value,
			deleted: false,
		}
	} else {
		s.localStore[key] = value
	}
}

func (s *store) Delete(key string) error {
	if _, err := s.findByKey(key); err != nil {
		return err
	}

	s.delete(key)

	return nil
}

func (s *store) Count(value string) int {
	valueKeys := make(map[string]struct{})
	for k, v := range s.localStore {
		if v == value {
			valueKeys[k] = struct{}{}
		}
	}

	if s.currenTxScope != nil {
		s.currenTxScope.scanValues(valueKeys, value)
	}

	return len(valueKeys)
}

func (s *store) Begin() {
	curTxScope := s.currenTxScope
	newTxScope := &transactionScope{
		localStore: make(map[string]txValue),
		parent:     curTxScope,
	}

	s.currenTxScope = newTxScope
}

func (s *store) Commit() error {
	if s.currenTxScope == nil {
		return ErrNoTransaction
	}

	curTxScope := s.currenTxScope
	parentTxScope := curTxScope.parent

	if parentTxScope != nil {
		for k, v := range curTxScope.localStore {
			parentTxScope.localStore[k] = v
		}
	} else {
		for k, v := range curTxScope.localStore {
			if v.deleted {
				delete(s.localStore, k)
			} else {
				s.localStore[k] = v.value
			}
		}
	}

	s.currenTxScope = parentTxScope

	return nil
}

func (s *store) Rollback() error {
	if s.currenTxScope == nil {
		return ErrNoTransaction
	}

	s.currenTxScope = s.currenTxScope.parent

	return nil
}

func (s *store) findByKey(key string) (string, error) {
	if s.currenTxScope != nil {
		txV, err := s.currenTxScope.findByKey(key)
		if err == nil {
			if txV.deleted {
				return "", ErrKeyNotSet
			}

			return txV.value, nil
		}
	}

	if v, ok := s.localStore[key]; ok {
		return v, nil
	}

	return "", ErrKeyNotSet
}

func (s *store) delete(key string) {
	if s.currenTxScope != nil {
		s.currenTxScope.delete(key)
	} else {
		delete(s.localStore, key)
	}
}
