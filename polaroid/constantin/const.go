package constantin

type Constante struct {
	Const chan Act
}

type Act struct {
	Data  interface{}
	Types int
	Name  string
	Err   errors
}

func constante(c chan interface{}) {
	var g map[string]interface{}
	g = make(map[string]interface{})
	for val, ok := range c {
		if !ok {
			return
		}
		println(val)
	}

}

func NewConstante() *Constante {
	return &Constante{
		Const: make(chan Act),
	}
}

func (c *Constante) Start() {
	go constante(c.Const)
}
