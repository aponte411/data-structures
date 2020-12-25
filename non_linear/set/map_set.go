package map_set

type MapSet map[interface{}]bool

func (s *MapSet) Add(key interface{}) {
	(*s)[key] = true
}

func (s *MapSet) Remove(key interface{}) {
	delete((*s), key)
}

func (s *MapSet) Find(key interface{}) bool {
	return (*s)[key]
}
