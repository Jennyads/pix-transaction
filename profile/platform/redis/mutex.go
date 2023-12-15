package redis

type Mutex interface {
	Lock() error
	Unlock() error
}
