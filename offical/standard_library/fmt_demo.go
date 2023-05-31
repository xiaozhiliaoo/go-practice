package main

import (
	"errors"
	"fmt"
	"strings"
)

type Reservation struct {
	Method string `json:"method"`
	Args   struct {
		HotelCity   string `json:"hotel_city,omitempty"`
		CheckInDate string `json:"check_in_date,omitempty"`
		OrderId     string `json:"order_id,omitempty"`
	} `json:"args,omitempty"`
}

func ParseReservationsString(s string) ([]*Reservation, error) {
	reservations := []*Reservation{}
	reservationStrings := strings.Split(s[1:len(s)-1], ",")
	for _, reservationString := range reservationStrings {
		reservation, err := parseReservationString(reservationString)
		if err != nil {
			return nil, err
		}
		reservations = append(reservations, reservation)
	}
	return reservations, nil
}

func parseReservationString(s string) (*Reservation, error) {
	// 解析方法名和参数列表
	parts := strings.Split(s, "(")
	if len(parts) != 2 {
		return nil, errors.New("invalid reservation string")
	}
	method := strings.TrimSpace(parts[0])
	argsString := strings.TrimRight(strings.TrimSpace(parts[1]), ")")

	reservation := &Reservation{
		Method: method,
	}
	if len(argsString) > 0 {
		// 解析参数列表中的键值对
		args := make(map[string]string)
		kvStrings := strings.Split(argsString, ",")
		for _, kvString := range kvStrings {
			kvParts := strings.Split(kvString, "=")
			if len(kvParts) != 2 {
				return nil, errors.New("invalid reservation string")
			}
			k := strings.TrimSpace(kvParts[0])
			v := strings.TrimSpace(kvParts[1])
			if strings.HasPrefix(v, "\"") && strings.HasSuffix(v, "\"") {
				v = v[1 : len(v)-1]
			}
			args[k] = v
		}
		// 将键值对的值填充至对应的参数字段
		switch method {
		case "reserve_hotel":
			reservation.Args.HotelCity = args["hotel_city"]
			reservation.Args.CheckInDate = args["check_in_date"]
		case "modify_order":
			reservation.Args.OrderId = args["order_id"]
		default:
			return nil, fmt.Errorf("unknown reservation method: %s", method)
		}
	}
	return reservation, nil
}

func main() {
	input := "[reserve_hotel(hotel_city=\"北京\",check_in_date=\"下周三\"),modify_order(order_id=\"123456\")]"
	reservations, err := ParseReservationsString(input)
	if err != nil {
		panic(err)
	}
	fmt.Println(len(reservations)) // 输出：2
	fmt.Println(reservations[0].Method, reservations[0].Args.HotelCity, reservations[0].Args.CheckInDate)
	// 输出：reserve_hotel 北京 下周三
	fmt.Println(reservations[1].Method, reservations[1].Args.OrderId)
}
