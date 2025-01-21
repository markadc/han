package kss

import "log"

// Nget 实现了map多层取值，KEY不存在则返回默认值
func Nget(data A, args []string, failed any) any {
	temp := data
	for i, a := range args {
		// 检查 key 是否存在
		if val, exists := temp[a]; exists {
			// 如果这是最后一个 key，返回对应的值
			if i == len(args)-1 {
				return val
			}
			// 如果值是 map 类型，继续下一步，否则返回 failed
			if nested, ok := val.(map[string]any); ok {
				temp = nested
			} else {
				log.Printf("KEY %q VALUE %v not is map, %v", a, val, args)
				return failed
			}
		} else {
			log.Printf("KEY %q miss, %v", a, args)
			return failed
		}
	}
	return data
}
