package clipboard

type Clipboard interface {
	Read() string
	Write(string)
}
