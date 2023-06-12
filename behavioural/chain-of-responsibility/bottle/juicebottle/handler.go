package juicebottle

type JuiceBottle struct {
	material *string
	content  *string
	cap      *string
	label    *string
}

func (jb *JuiceBottle) SetMaterial(material *string) {
	jb.material = material
}

func (jb *JuiceBottle) SetContent(content *string) {
	jb.content = content
}

func (jb *JuiceBottle) SetCap(cap *string) {
	jb.cap = cap
}

func (jb *JuiceBottle) SetLabel(label *string) {
	jb.label = label
}

func (jb *JuiceBottle) GetMaterial() *string {
	return jb.material
}

func (jb *JuiceBottle) GetContent() *string {
	return jb.content
}

func (jb *JuiceBottle) GetCap() *string {
	return jb.cap
}

func (jb *JuiceBottle) GetLabel() *string {
	return jb.label
}
