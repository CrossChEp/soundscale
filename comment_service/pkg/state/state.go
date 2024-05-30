package state

type State interface {
	Get(entityId string) error
}
