package sortmap

// sort a map's keys in descending order of its values.

import "sort"

type sortedMap struct {
	M map[string]int
	S []string
}

func (sm *sortedMap) Len() int {
	return len(sm.M)
}

func (sm *sortedMap) Less(i, j int) bool {
	return sm.M[sm.S[i]] < sm.M[sm.S[j]]
}

func (sm *sortedMap) Swap(i, j int) {
	sm.S[i], sm.S[j] = sm.S[j], sm.S[i]
}

func SortedKeys(m map[string]int) []string {
	sm := new(sortedMap)
	sm.M = m
	sm.S = make([]string, len(m))
	i := 0
	for key, _ := range m {
		sm.S[i] = key
		i++
	}
	sort.Sort(sm)
	return sm.S
}
