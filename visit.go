package main

type visit struct {
	id        int
	date      string
	visitor   int
	apartment int
	status    string
}

func newVisit(id int, date string, visitor int, apartment int, status string) visit {
	return visit{id, date, visitor, apartment, status}
}
