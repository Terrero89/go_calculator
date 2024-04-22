package iomanager

// IOManager interfaces that implenent the methods in cmd manager and filemanager
type IOManager interface {
	ReadLines() ([]string, error)
	WriteResult(data interface{}) error
}
