package house

type House struct {
	WindowType string
	DoorType   string
	Floor      int
}

func (h *House) GetWindowInfo() string {
	return h.WindowType
}

func (h *House) GetDoorInfo() string {
	return h.DoorType
}

func (h *House) GetFoorInfo() int {
	return h.Floor
}
