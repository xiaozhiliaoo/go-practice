package main

import (
	"fmt"
	"reflect"
	"strings"
)

// MsgRecordInfo 消息记录信息
type MsgRecordInfo struct {
	ID           string `json:"id"`
	SessionID    string `json:"session_id"`     // 会话ID
	SessionIDURI string `json:"session_id_uri"` // 会话ID链接
	RecordID     string `json:"record_id"`      // 记录ID
	IsHappy      bool   `json:"is_happy"`       // 记录ID
}

func filterSliceByFields(slice interface{}, fieldNames string) ([][]string, error) {
	sliceValue := reflect.ValueOf(slice)
	if sliceValue.Kind() != reflect.Slice {
		return nil, fmt.Errorf("expected a slice, got %v", sliceValue.Kind())
	}

	fields := strings.Split(fieldNames, ",")
	result := make([][]string, 0, sliceValue.Len())
	for i := 0; i < sliceValue.Len(); i++ {
		item := sliceValue.Index(i)
		itemData := make([]string, 0, len(fields))
		for _, fieldName := range fields {
			field := item.FieldByName(fieldName)
			if !field.IsValid() {
				return nil, fmt.Errorf("field %s not found in %v", fieldName, item.Type())
			}
			itemData = append(itemData, fmt.Sprintf("%v", field.Interface()))
		}
		result = append(result, itemData)
	}

	return result, nil
}

func main() {
	data := []MsgRecordInfo{
		{ID: "1", SessionID: "s1", SessionIDURI: "s1_uri", RecordID: "r1", IsHappy: true},
		{ID: "2", SessionID: "s2", SessionIDURI: "s2_uri", RecordID: "r2", IsHappy: false},
		{ID: "3", SessionID: "s3", SessionIDURI: "s3_uri", RecordID: "r3", IsHappy: true},
	}

	filteredData, err := filterSliceByFields(data, "ID,IsHappy")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(filteredData)
}
