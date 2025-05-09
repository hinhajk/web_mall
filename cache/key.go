package cache

import (
	"fmt"
	"strconv"
)

const (
	RankKey = "rank"
)

// ProductViewKey 构建商品标识
func ProductViewKey(id uint) string {
	return fmt.Sprintf("view product:%s", strconv.Itoa(int(id)))
}
