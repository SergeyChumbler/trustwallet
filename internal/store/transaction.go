package store

type txValue struct {
	value   string
	deleted bool
}

type transactionScope struct {
	localStore map[string]txValue
	parent     *transactionScope
}

func (ts *transactionScope) scanValues(target map[string]struct{}, value string) {
	if ts.parent != nil {
		ts.parent.scanValues(target, value)
	}

	for k, v := range ts.localStore {
		if v.value == value && !v.deleted {
			target[k] = struct{}{}
		} else if _, ok := target[k]; ok {
			delete(target, k)
		}
	}
}

func (ts *transactionScope) findByKey(key string) (txValue, error) {
	if v, ok := ts.localStore[key]; ok {
		return v, nil
	}

	if ts.parent != nil {
		return ts.parent.findByKey(key)
	}

	return txValue{}, ErrKeyNotSet
}

func (ts *transactionScope) delete(key string) {
	ts.localStore[key] = txValue{
		deleted: true,
	}
}
