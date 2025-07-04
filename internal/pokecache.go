package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	data map[string]cacheEntry
	lock sync.RWMutex
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		data: make(map[string]cacheEntry),
		lock: sync.RWMutex{},
	}

	go c.reapLoop(interval)

	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.data[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	el, ok := c.data[key]
	if !ok {
		return []byte{}, false
	}

	return el.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		c.lock.Lock()
		for key, entry := range c.data {
			if time.Since(entry.createdAt) > interval {
				delete(c.data, key)
			}
		}
		c.lock.Unlock()
	}

}

type Pokemon struct {
	Name string
	Xp   int
}

type Pokedex struct {
	pokemons map[string]Pokemon
	lock     sync.RWMutex
}

func CreateNewPokedex() *Pokedex {
	px := &Pokedex{
		pokemons: make(map[string]Pokemon),
		lock:     sync.RWMutex{},
	}

	return px
}

func (px *Pokedex) CapturePokemon(name string, exp int) {

	px.lock.Lock()
	defer px.lock.Unlock()

	px.pokemons[name] = Pokemon{
		Name: name,
		Xp:   exp,
	}
}

func (px Pokedex) GetPokemon(name string) (Pokemon, bool) {
	pokemon, exists := px.pokemons[name]
	return pokemon, exists
}

func (px Pokedex) GetPokedex() ([]Pokemon, bool) {

	if len(px.pokemons) == 0 {
		return nil, false
	}

	pokemons := []Pokemon{}

	for _, pokemon := range px.pokemons {
		pokemons = append(pokemons, pokemon)
	}

	return pokemons, true
}
