package imports

type Inject interface {
	Asset(string)([]byte, error)
}
