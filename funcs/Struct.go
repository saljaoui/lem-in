package funcs

type Ant struct {
	id       int
	path     []string
	position int
}

type Path struct {
	id    int
	rooms []string
	ants  int
}
type AntGraph struct {
	connections map[string][]string
	Ants  int
	StartRoom string
	EndRoom string
}