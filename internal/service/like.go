package service

type ILike interface {
}

var localLike ILike

func RegisterLike(i ILike) {
	localLike = i
}

func Like() ILike {
	if localLike == nil {
		panic("implement not found for interface ILike, forgot register?")
	}

	return localLike
}
