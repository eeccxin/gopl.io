//测试bufio.Scanner,参考：https://blog.csdn.net/weixin_29662775/article/details/112099878
package main

import (
	"bufio"
	"errors"
	"fmt"
	"strings"
)

// // 1. 需要补充更多的数据
// func main() {
// 	input := "abcdefghijkl"
// 	scanner := bufio.NewScanner(strings.NewReader(input))
// 	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
// 		fmt.Printf("%t\t%d\t%s\n", atEOF, len(data), data)
// 		return 0, nil, nil
// 	}
// 	scanner.Split(split)
// 	buf := make([]byte, 2)
// 	scanner.Buffer(buf, bufio.MaxScanTokenSize)
// 	for scanner.Scan() {
// 		fmt.Printf("%s\n", scanner.Text())
// 	}
// }

//2. 已找到字符标记（token）
//下面的这个函数将只会寻找连续的 foo 串
// func main() {
// 	input := "foofoofoo"
// 	scanner := bufio.NewScanner(strings.NewReader(input))
// 	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
// 		if bytes.Equal(data[:3], []byte{'f', 'o', 'o'}) {
// 			return 3, []byte{'F'}, nil
// 		}
// 		if atEOF {
// 			return 0, nil, io.EOF
// 		}
// 		return 0, nil, nil
// 	}
// 	scanner.Split(split)
// 	for scanner.Scan() {
// 		fmt.Printf("%s\n", scanner.Text())
// 	}
// }

//3. 报错
//如果 split 函数返回了错误那么扫描器就会停止工作
func main() {
	input := "abcdefghijkl"
	scanner := bufio.NewScanner(strings.NewReader(input))
	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		return 0, nil, errors.New("bad luck")
	}
	scanner.Split(split)
	for scanner.Scan() {
		fmt.Printf("%sn", scanner.Text())
	}
	if scanner.Err() != nil {
		fmt.Printf("error: %sn", scanner.Err())
	}
}
