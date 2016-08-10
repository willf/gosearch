package set

type Set map[string]bool

func New(items ...string) Set {
	set := make(Set)
	for _, item := range items {
		set[item] = true
	}
	return set
}

func NewFromSlice(items []string) Set {
	return New(items...)
}

func (set Set) Add(key string) {
	set[key] = true
}

func (set Set) Contains(key string) bool {
	return set[key]
}

func (set Set) Remove(key string) {
	delete(set, key)
}

func (set Set) Union(other Set) (result Set) {
	result = New()
	for key, _ := range set {
		result.Add(key)
	}
	for key, _ := range other {
		result.Add(key)
	}
	return
}

func (set Set) Difference(other Set) (result Set) {
	result = New()
	for key, _ := range set {
		if !other.Contains(key) {
			result.Add(key)
		}
	}
	return
}

func (set Set) Intersection(other Set) (result Set) {
	result = New()
	for key, _ := range set {
		if other.Contains(key) {
			result.Add(key)
		}
	}
	return
}

func (set Set) Copy() (result Set) {
	result = New()
	for key, _ := range set {
		result.Add(key)
	}
	return
}
