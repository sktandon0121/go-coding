package main

func revStringByte(s []byte) {
	if len(s) == 1 || len(s) == 0 {
		return
	}
	t := s[0]
	s[0] = s[len(s)-1]
	s[len(s)-1] = t
	revStringByte(s[1 : len(s)-1])
}
