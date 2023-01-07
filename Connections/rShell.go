package Connections

type ReverseShell struct {
	Name string
}

func (r *ReverseShell) Connect() error {
	return nil
}

func (r *ReverseShell) Close() error {
	return nil
}
