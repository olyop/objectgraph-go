package inmemorycache

func Clear() {
	state.mu.Lock()
	defer state.mu.Unlock()

	for _, cache := range state.groups {
		cache.Flush()
	}
}
