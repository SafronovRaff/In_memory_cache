# In memory cache


```
package main

import (
	in "github.com/SafronovRaff/In_memory_cache"
  "time"
)

func main() {
	//NewCash - creates cache storage and cache storage time
	cashe := in.NewCash(30 * time.Second)

	//Set - adds the value of saving by key along with the current time
	cashe.Set("userId-1", "3242565469")

	// Get - returns the value by key
	p := cashe.Get("userId-1")
	fmt.Println(p)

	// Delete - delete the value by key
	cashe.Delete("userId-1")
  }
  ```
