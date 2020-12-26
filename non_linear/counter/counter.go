package counter

type Counter map[interface{}]int

func (c *Counter) Add(key interface{}) {
	(*c)[key] += 1
}

func (c *Counter) Find(key interface{}) bool {
	_, ok := (*c)[key]
	return ok
}

func (c *Counter) Get(key interface{}) (int, bool) {
	val, ok := (*c)[key]
	return val, ok
}
