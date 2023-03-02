package In_memory_cache

import "fmt"

type Cashe struct {
	cashe map[string]interface{}
}

// Set - adds saves value by key
func (c *Cashe) Set(key string, value interface{}) {
	if len(c.cashe) == 0 {
		c.cashe = make(map[string]interface{})
	}
	c.cashe[key] = value
}

// Get - returns the value by key
func (c *Cashe) Get(key string) interface{} {
	value, ok := c.cashe[key]
	if !ok {
		fmt.Println("Value not found")
	}
	return value
}

// Delete - delete the value by key
func (c *Cashe) Delete(key string) {
	delete(c.cashe, key)
}
