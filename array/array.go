package util

// InArrInter 是否在slice内
func InArrInter(slice []interface{}, val interface{}) (flag bool) {
	for _, item := range slice {
		if item == val {
			flag = true
			break
		}
	}
	return
}

func InArrStr(slice []string, val string) (flag bool) {
	for _, item := range slice {
		if item == val {
			flag = true
			break
		}
	}
	return
}

func InArrInt32(slice []int32, val int32) (flag bool) {
	for _, item := range slice {
		if item == val {
			flag = true
			break
		}
	}
	return
}

// InterArrStr 交集
func InterArrStr(before []string, after []string) (inter []string) {
	bl := len(before)
	al := len(after)
	max := bl
	if max < al {
		max = al
	}
	inter = make([]string, 0, max)

	tmp := make(map[string]struct{}, bl)
	for _, item := range before {
		tmp[item] = struct{}{}
	}

	for _, item := range after {
		if _, ok := tmp[item]; ok {
			inter = append(inter, item)
		}
	}

	return
}

// DiffArrStr 差集
func DiffArrStr(before []string, after []string) (diff []string) {
	bl := len(before)
	al := len(after)
	max := bl
	if max < al {
		max = al
	}
	diff = make([]string, 0, bl)

	tmp := make(map[string]struct{}, bl)
	for _, item := range before {
		tmp[item] = struct{}{}
	}

	for _, item := range after {
		if _, ok := tmp[item]; !ok {
			diff = append(diff, item)
		}
	}

	return
}

// UniqueArrStr 唯一数组
func UniqueArrStr(arr []string) (uniqArr []string) {

	uniqArr = make([]string, 0, len(arr))
	uniqMap := make(map[string]struct{}, len(arr))

	for _, item := range arr {
		if _, ok := uniqMap[item]; !ok {
			uniqMap[item] = struct{}{}
			uniqArr = append(uniqArr, item)
		}
	}

	return
}
