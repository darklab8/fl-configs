package lower_map

import "strings"

type KeyLoweredMap[K ~string, V any] struct {
	data map[K]V
}

func NewKeyLoweredMap[K ~string, V any]() *KeyLoweredMap[K, V] {
	return &KeyLoweredMap[K, V]{
		data: make(map[K]V),
	}
}

func (d *KeyLoweredMap[K, V]) MapSet(key K, value V) {
	d.data[K(strings.ToLower(string(key)))] = value
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
