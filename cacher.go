package cache

type cacher struct {
	list map[string]interface{}
}

/*-----------------------------------------------------------------------*/

// Get возвращает значение по ключу. Если значение отсутствует возвращает nil
func (c *cacher) Get(
	key string,
) interface{} {
	if value, ok := c.list[key]; ok {
		return value
	}
	return nil
}

/*-----------------------------------------------------------------------*/

// Set устанавливает значение valuer для ключа key
func (c *cacher) Set(
	key string,
	value interface{},
) {
	c.list[key] = value
}

/*-----------------------------------------------------------------------*/

// Delete удаляет элемент с ключом key
func (c *cacher) Delete(
	key string,
) {
	delete(c.list, key)
}

/*-----------------------------------------------------------------------*/
