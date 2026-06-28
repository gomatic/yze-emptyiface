package a

// Handler uses the empty interface and must be flagged (and fixed to any).
type Handler interface{} // want `prefer any`

// accept takes an empty interface parameter and must be flagged.
func accept(x interface{}) { _ = x } // want `prefer any`

// Real has methods and must not be flagged.
type Real interface {
	Do()
}
