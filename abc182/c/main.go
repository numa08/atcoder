package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	cnt := make([]int, 3) // cnt[0], cnt[1], cnt[2] が作られることを想定
	for n != 0 {
		cnt[n % 10 % 3]++ // 各桁を３で割ったあまりの数をカウント
		n /= 10
	}
	// 各桁を３で割った数のあまりの合計が計算できる（すごい）
	// 各桁を合計して3で割ったあまりと、各桁を3で割ったあまりの合計を3で割ったあまりは同じになる（不思議）
	// すでに計算した各桁を3で割ったあまりを足し合わせると再計算しなくてよいので嬉しい
	cur := (cnt[1] + 2 * cnt[2]) % 3
	fmt.Printf("cur %d\n", cur)
	k := cnt[0] + cnt[1] + cnt[2]
	var res int
	if cur == 0 {
		res = 0
	} else if (cur == 1) {
		if cnt[1] > 0 {
			if k == 1 {
				res = -1
			} else {
				res = 1
			}
		} else {
			if k == 2 {
				res = -1
			} else {
				res = 2
			}
		}
	} else {
		if cnt[2] > 0 {
			if k == 1 {
				res = -1
			} else {
				res = 1
			}
		} else {
			if k == 2 {
				res = -1
			} else {
				res = 2
			}
		}
	}
	fmt.Printf("%d", res)
}
