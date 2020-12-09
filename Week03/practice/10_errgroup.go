package main

import (
	"context"
	"errors"
	"fmt"

	"golang.org/x/sync/errgroup"
)

func main() {
	g, ctx := errgroup.WithContext(context.Background())

	var a, b, c []int

	// 调用广告服务
	g.Go(func() error {
		a = []int{0}
		return nil
	})

	// 调用AI服务
	g.Go(func() error {
		b = []int{1}
		return errors.New("ai error")
	})

	// 调用运营平台
	g.Go(func() error {
		c = []int{2}
		return nil
	})

	// 等所有g.Go 都执行完，返回第一个报错
	err := g.Wait()
	fmt.Println(err)

	// a + b + c merge起来
	// 防止data race 所以没有用同一个 [],而是最后merge合并

	fmt.Println(ctx.Err())

	// 处理场景：
	// 1. 有一个报错，全部取消
	// 2. 报错后降级
}
