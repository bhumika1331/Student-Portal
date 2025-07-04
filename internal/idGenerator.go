package internal

func SeqID() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
