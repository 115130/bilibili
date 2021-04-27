package net

import (
	"github.com/zhshch2002/goreq"
	"time"
)

func GetHtml(url string) string {
	client := goreq.NewClient(
		goreq.WithRandomUA(),
		goreq.WithRateLimiter(false, &goreq.RateLimiterOpinion{
			LimiterMatcher: goreq.LimiterMatcher{
				Glob: "*",
			},
			Rate: 2,
		}),
		goreq.WithDelayLimiter(false, &goreq.DelayLimiterOpinion{
			LimiterMatcher: goreq.LimiterMatcher{
				Glob: "*",
			},
			Delay: 1 * time.Second,
			// RandomDelay: 5 * time.Second,
		}))
	get := goreq.Get(url)
	get.SetClient(client)
	html, err := get.AddParam("P", "p").Do().Txt()
	if err != nil {
		panic(err.Error() + "网址错误")
	}
	return html
}
