package utils

// GetLimitAndPageSize 获取起始页和页大小
func GetLimitAndPageSize(page, pageSize int) (int, int) {
	//将page从1开始
	if page <= 0 {
		page = 1
	}
	//pageSize默认最大200
	if pageSize > 200 {
		pageSize = 200
	}
	if pageSize <= 0 {
		pageSize = 20
	}
	size := (page - 1) * pageSize

	return size, pageSize
}

// GetPageSize 获取起始页和页大小
func GetPageSize(page, pageSize int) (int, int) {
	//将page从1开始
	if page <= 0 {
		page = 1
	}
	//pageSize默认最大200
	if pageSize > 200 {
		pageSize = 200
	}
	if pageSize <= 0 {
		pageSize = 20
	}
	size := (page - 1) * pageSize

	return size, pageSize
}
