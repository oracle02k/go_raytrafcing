package util3d

type HitableList struct {
	list []Hitable
}

func NewHitableList() *HitableList {
	return &HitableList{list: make([]Hitable, 0, 32)}
}

func (h *HitableList) AddHitable(hitable Hitable) {
	h.list = append(h.list, hitable)
}

func (h *HitableList) Hit(r *Ray, t_min, t_max float64, rec *HitRecord) bool {
	tmpRec := HitRecord{}
	hitAnything := false
	closetSoFar := t_max
	for i := 0; i < len(h.list); i++ {
		if h.list[i].Hit(r, t_min, closetSoFar, rec) {
			hitAnything = true
			closetSoFar = tmpRec.t
			rec = &tmpRec
		}
	}
	return hitAnything
}
