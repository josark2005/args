package args

type Wrap struct {
	Args map[any]any
}

func NewWarp() *Wrap {
	return &Wrap{
		Args: make(map[any]any),
	}
}

func (a *Wrap) Set(key any, value any) *Wrap {
	a.Args[key] = value
	return a
}

func (a *Wrap) Get(key any) any {
	if v, s := a.Args[key]; !s {
		return nil
	} else {
		return v
	}
}

func (a *Wrap) Value(key any) any {
	return a.Get(key)
}

func (a *Wrap) Remove(key any) {
	if _, s := a.Args[key]; s {
	} else {
		delete(a.Args, key)
	}
}
