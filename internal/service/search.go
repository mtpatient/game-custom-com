package service

import "context"

type ISearch interface {
	Save(ctx context.Context, uid int, keyWord string) error
}

var localSearch ISearch

func RegisterSearch(i ISearch) {
	localSearch = i
}

func Search() ISearch {
	if localSearch == nil {
		panic("implement not found for Search, forgot register?")
	}
	return localSearch
}
