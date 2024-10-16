package utils

type ConnectionManager struct {
	Connections *ConcurrentMap[*ConcurrentConn, *ConcurrentSet[string]]
}

func NewConnectionManager() *ConnectionManager {
	return &ConnectionManager{
		Connections: NewConcurrentMap[*ConcurrentConn, *ConcurrentSet[string]](),
	}
}

func (c *ConnectionManager) ContainsKey(key *ConcurrentConn, target string) bool {
	set, exists := c.Connections.Load(key)
	if !exists {
		return false
	}

	contained := false
	set.Range(func(item string) bool {
		if item == target {
			contained = true
			return false
		}
		return true
	})

	return contained
}
