package waterbottle

type WaterBottle struct {
	material *string
	content  *string
	cap      *string
	label    *string
}

func (wb *WaterBottle) SetMaterial(material *string) {
	wb.material = material
}

func (wb *WaterBottle) SetContent(content *string) {
	wb.content = content
}

func (wb *WaterBottle) SetCap(cap *string) {
	wb.cap = cap
}

func (wb *WaterBottle) SetLabel(label *string) {
	wb.label = label
}

func (wb *WaterBottle) GetMaterial() *string {
	return wb.material
}

func (wb *WaterBottle) GetContent() *string {
	return wb.content
}

func (wb *WaterBottle) GetCap() *string {
	return wb.cap
}

func (wb *WaterBottle) GetLabel() *string {
	return wb.label
}
