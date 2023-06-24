package main

import "fmt"

// ProjectFileDBInfo 文件解析信息
type ProjectFileDBInfo struct {
	ID  int64  `db:"id"`  // 文件ID
	Url string `db:"url"` // url

}

func addUrlPrefix(fileDBList []ProjectFileDBInfo) {
	for i, file := range fileDBList {
		fileDBList[i].Url = file.Url + "ddd222"
	}
}

func main() {
	infos := []ProjectFileDBInfo{{ID: 1, Url: "www.baidu.com"}, {ID: 2, Url: "www.alibaba.com"}}
	addUrlPrefix(infos)
	fmt.Printf("%+v", infos)

}
