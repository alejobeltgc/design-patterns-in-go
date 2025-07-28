package concretecomponent

type HouseBlend struct {
	Description string
}

func NewHouseBlend() *HouseBlend {
	return &HouseBlend{
		Description: "House Blend",
	}
}

func (h *HouseBlend) Cost() float64 {
	return .89
}

func (h *HouseBlend) GetDescription() string {
	return h.Description
}
