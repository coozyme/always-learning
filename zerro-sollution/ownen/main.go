package main

import (
	"log"
)

type TimeDetails []TimeDetail
type TimeDetail struct {
	TimePickUpStart   string `json:"time_pick_up_start"`
	TimePickUpEnd     string `json:"time_pick_up_end"`
	TimeDeliveryStart string `json:"time_delivery_start"`
	TimeDeliveryEnd   string `json:"time_delivery_end"`
	Service           string `json:"service"`
}
type response struct {
	Message    string       `json:"message"`
	StatusCode int64        `json:"status_code"`
	Data       []TimeDetail `json:"data"`
}

type price struct {
	total int
}

func main() {
	newRes := response{
		Message:    "OK",
		StatusCode: 200,
		Data: []TimeDetail{
			{TimePickUpStart: "08.00", TimePickUpEnd: "09.00", TimeDeliveryStart: "08.00", TimeDeliveryEnd: "09.00", Service: "same_day"},
			{TimePickUpStart: "13.00", TimePickUpEnd: "14.00", TimeDeliveryStart: "13.00", TimeDeliveryEnd: "14.00", Service: "same_day"},
			{TimePickUpStart: "15.00", TimePickUpEnd: "16.00", TimeDeliveryStart: "15.00", TimeDeliveryEnd: "16.00", Service: "same_day"},
			{TimePickUpStart: "10.00", TimePickUpEnd: "11.00", TimeDeliveryStart: "10.00", TimeDeliveryEnd: "11.00", Service: "next_day"},
			// {TimePickUpStart: "12.00", TimePickUpEnd: "13.00", TimeDeliveryStart: "12.00", TimeDeliveryEnd: "13.00", Service: "reguler"},
			{TimePickUpStart: "14.00", TimePickUpEnd: "15.00", TimeDeliveryStart: "14.00", TimeDeliveryEnd: "15.00", Service: "next_day"},
		},
	}

	filtered := TimeDetails{}
	mapTimeDetial := make(map[string]TimeDetail)

	for _, v := range newRes.Data {
		// log.Println("LOG-F", idx, (newRes.Data[idx]))
		// if v.Service == "same_day" && idx >= 1 {
		// 	continue
		// }

		// if filtered[idx].Service == "same_day" {
		// 	continue
		// }

		mapTimeDetial[v.Service] = v

	}
	for _, v := range v {

	}
	filtered = append(filtered, mapTimeDetial["same_day"])
	filtered = append(filtered, mapTimeDetial["next_day"])
	filtered = append(filtered, mapTimeDetial["reguler"])

	// heap.Push(&filtered, mapTimeDetial)
	// log.Println("LOG-FILR", filtered)

	// newP := price{total: 1000}
	// newP.total -= 500
	// log.Println("LOG-PT", newP)
	log.Println("LOG-filtered", filtered)

}

func (pq TimeDetails) Len() int { return len(pq) }

func (pq *TimeDetails) Push(x any) {
	*pq = append(*pq, x.(TimeDetail))
}

func (pq TimeDetails) Less(i, j int) bool {
	if pq[i].Service == pq[j].Service {
		if pq[i].TimeDeliveryStart == pq[j].TimeDeliveryStart {
			return pq[i].TimeDeliveryEnd < pq[j].TimeDeliveryEnd
		}
		return pq[i].TimePickUpStart < pq[j].TimePickUpStart
	}
	return pq[i].Service < pq[j].Service
}

func (pq *TimeDetails) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]

	return item
}

func (pq TimeDetails) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
