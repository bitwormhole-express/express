package service

// UUIDGenerator  是一个生成UUID的服务
// 【inject:"#the-uuid-generator"】
type UUIDGenerator interface {
	GenUUID(aTypeName, aID string) string
}
