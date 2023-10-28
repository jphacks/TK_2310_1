package algo

import "time"

type event struct {
	start_at time.Time
	end_at   time.Time
}

// 検索範囲が1日のみの場合
func imos() []event {
	// 必要データの作成(処理のに必要な情報尾)
	events := make([]event, 0)                                              //指定された日に開催されるイベント
	my_events := make([]event, 0)                                           //指定された日に既に参加確定しているイベント
	serach_date, _ := time.Parse(time.RFC3339, "2023-01-01T00:00:00+09:00") //検索日時の0時0分0秒
	imos_size := 24 * 60 / 10                                               //累積和の範囲 24時間 * 60分 を 10分の離散時間にする。検索範囲が1日の時はこれ

	// 結果の格納
	result := make([]event, imos_size, 0)

	// imos table
	table := make([]int, imos_size, 0)

	// imos operate
	for _, v := range my_events {
		start := int(v.start_at.Sub(serach_date).Minutes() / 10)
		end := int(v.end_at.Sub(serach_date).Minutes() / 10)
		table[start]++
		table[end]--
	}

	// シミュレーション
	for i, _ := range table {
		if 0 < i {
			table[i] += table[i-1]
		}
	}

	for _, v := range events {
		start := int(v.start_at.Sub(serach_date).Minutes() / 10)
		end := int(v.end_at.Sub(serach_date).Minutes() / 10)
		if table[start] == 1 && table[end] == 1 {
			result = append(result, v)
		}
	}

	return result
}
