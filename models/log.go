package models

type Log []string // this creates a *new type*

func (l *Log) Add(text string) {
	*l = append(*l, text)
}

func (l Log) LastN(n int) []string {
	if len(l) < n {
		return l
	}
	return l[len(l)-n:]
}

