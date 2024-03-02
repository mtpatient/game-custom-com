package utility

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
	"time"
)

func GetCOSSignature(ctx context.Context, names []string, t time.Duration) ([]string, error) {
	cosURL := g.Cfg().MustGet(ctx, "TencentCOS.URL").String()
	ak := g.Cfg().MustGet(ctx, "TencentCOS.SecretID").String()
	sk := g.Cfg().MustGet(ctx, "TencentCOS.SecretKey").String()
	u, _ := url.Parse(cosURL)
	b := &cos.BaseURL{BucketURL: u}

	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  ak,
			SecretKey: sk,
		},
	})

	signatures := make([]string, len(names))

	for i, n := range names {
		signature, err := client.Object.GetPresignedURL(ctx, http.MethodPut, n, ak, sk, t, nil)
		if err != nil {
			return nil, gerror.New("COS Error")
		}
		signatures[i] = signature.String()
	}

	return signatures, nil
}
