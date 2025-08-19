package cache // Package cache объявляем пакет cache, чтобы его можно было импортировать в проекте

// Cache - структура для хранения данных в памяти
type Cache struct {
	data map[string]interface{} // поле data — это карта (map), где ключи типа string, а значения любого типа
}

// New создает новый инстанс кеша
func New() *Cache {
	return &Cache{ // возвращаем указатель на новый Cache
		data: make(map[string]interface{}), // инициализируем пустую map
	}
}

// Set записывает значение value по ключу key
func (c *Cache) Set(key string, value interface{}) {
	c.data[key] = value // кладем значение в map по ключу
}

// Get возвращает значение по ключу.
// Если ключа нет — вернёт nil.
func (c *Cache) Get(key string) interface{} {
	val, ok := c.data[key] // пытаемся получить значение из map
	if !ok {               // если ключа нет
		return nil // возвращаем nil
	}
	return val // если ключ есть — возвращаем значение
}

// Delete удаляет значение по ключу
func (c *Cache) Delete(key string) {
	delete(c.data, key) // встроенная функция delete удаляет пару ключ-значение из map
}
