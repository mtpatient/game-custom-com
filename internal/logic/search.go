package logic

import (
	"context"
	"game-custom-com/internal/dao"
	"game-custom-com/internal/service"
	"github.com/go-ego/gse"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type sSearch struct {
}

var seg gse.Segmenter

func init() {
	service.RegisterSearch(&sSearch{})

	err := seg.LoadDict()
	if err != nil {
		g.Log().Error(context.Background(), err)
		return
	}
}

func (s sSearch) Save(ctx context.Context, uid int, keyWord string) error {

	_, err := dao.SearchHistory.Ctx(ctx).Insert(g.Map{
		"user_id": uid,
		"keyword": keyWord,
	})
	if err != nil {
		g.Log().Error(ctx, err)
		return gerror.New("保存搜索历史失败")
	}

	return nil
}

func KeywordCut(str string) (res []string) {
	res = seg.Cut(str, true)
	return
}
