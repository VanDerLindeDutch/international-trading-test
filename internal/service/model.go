package service

type HashedValue struct {
	SHA256value string `json:"SHA256Value"`
	MD5value    string `json:"MD5Value"`
}

type GetValueDto struct {
	Input string `json:"input"`
}
