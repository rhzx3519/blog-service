/**
@author ZhengHao Lou
Date    2022/02/09
*/
package limiter

import (
    "github.com/gin-gonic/gin"
    "github.com/juju/ratelimit"
    "strings"
)

type MethodLimiter struct {
    *Limiter
}

func NewMethodLimiter() LimiterIface {
    return MethodLimiter{&Limiter{map[string]*ratelimit.Bucket{}}}
}

func (l MethodLimiter) Key(c *gin.Context) string {
    uri := c.Request.RequestURI
    index := strings.Index(uri, "?")
    if index == -1 {
        return uri
    }
    return uri[:index]
}

func (l MethodLimiter) GetBucket(key string) (*ratelimit.Bucket, bool) {
    bucket, ok := l.LimiterBuckets[key]
    return bucket, ok
}

func (l MethodLimiter) AddBuckets(rules ...LimiterBucketRule) LimiterIface {
    for _, rule := range rules {
        if _, ok := l.LimiterBuckets[rule.Key]; !ok {
            bucket := ratelimit.NewBucketWithQuantum(
                rule.FillInterval,
                rule.Capacity,
                rule.Quantum,
            )
            l.LimiterBuckets[rule.Key] = bucket
        }
    }
    return l
}
