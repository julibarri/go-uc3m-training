package conjunto

type Set struct {
	store map[int]bool
}

func (set *Set) constructor() {
	if (set.store == nil) { set.store = make(map[int]bool)}
}

func (set *Set) Añadir(value int) {
	set.constructor()
	set.store[value] = true;
}

func (set *Set) Existe(value int) bool {
	set.constructor()
	return set.store[value]
}

func (set *Set) Borrar(value int) {
	set.constructor()
	set.store[value] = false;
}
