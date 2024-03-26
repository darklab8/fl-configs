package lower_map

import "strings"

type KeyLoweredMap[K ~string, V any] struct {
	data map[K]V
}

type LoweredMapOpt[K ~string, V any] func(m *KeyLoweredMap[K, V])

func WithData[K ~string, V any](data map[K]V) LoweredMapOpt[K, V] {
	return func(m *KeyLoweredMap[K, V]) {
		for key, value := range data {
			m.data[K(strings.ToLower(string(key)))] = value
		}
	}
}

func NewKeyLoweredMap[K ~string, V any](opts ...LoweredMapOpt[K, V]) *KeyLoweredMap[K, V] {
	m := &KeyLoweredMap[K, V]{
		data: make(map[K]V),
	}
	for _, opt := range opts {
		opt(m)
	}

	return m
}

func (d *KeyLoweredMap[K, V]) MapSet(key K, value V) {
	d.data[K(strings.ToLower(strings.ReplaceAll(string(key), " ", "")))] = value
}

func (d *KeyLoweredMap[K, V]) MapGet(key K) V {
	value, ok := d.data[K(strings.ToLower(string(key)))]
	if !ok {
		panic("unable to access key=" + key)
	}
	return value
}

func (d *KeyLoweredMap[K, V]) MapGetValue(key K) (V, bool) {
	value, ok := d.data[K(strings.ToLower(string(key)))]
	return value, ok
}

func (d *KeyLoweredMap[K, V]) Foreach(callback func(key K, value V)) {
	for key, value := range d.data {
		callback(key, value)
	}
}
