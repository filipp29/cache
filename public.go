package cache

// New Возвращает новый объект cacher
func New() cacher {
	c := cacher{
		list: make(map[string]interface{}),
	}
	return c
}
