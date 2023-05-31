package main

import (
	"fmt"
	"regexp"
)

const TransferManProtocolRegexp = `(<a href="[\s\S]*?">[\s\S]+?<\/a>)|(\[(不可用的转人工链接|[^\]]+)\])`

func main() {

	str := `1、这是普通文本啊啊全球领先的中文搜索引擎、致力于让网民更便捷地获取123。
2、这是超链接<a href="http://www.baidu.com">百度</a>qwerty
3、这是图片<img src="https://p.qpic.cn/qidian_pic/2852199668/20230327635d70da644ae207704a4da3e1925e89/0?type=image-material" class="industry-rich-image"/>
4、这是表情<img class="editor-emoji" src="https://cdn.xiaowei.qq.com/assets/images/Expression_5.png" data-value="%5B%E5%BE%97%E6%84%8F%5D" alt="得意" style="width: 24px; min-width: 24px; height: 24px; display: inline-block;" width="24" height="24"/><img class="editor-emoji" src="https://cdn.xiaowei.qq.com/assets/images/Expression_3.png" data-value="%5B%E8%89%B2%5D" alt="色" style="width: 24px; min-width: 24px; height: 24px; display: inline-block;" width="24" height="24"/> 
5、这是文件<a href="https://qidian-ea-material-1251316161.cos.ap-guangzhou.myqcloud.com/material/file/doc/v2/28521996682799182_e1dcae7c09e383e43d06a157391645f4.docx?ci-process=doc-preview" target="_blank">新建 Microsoft Word 文档.docx</a>
6、这是小程序链接
<a href="javascript:void(0)" data-miniprogram-appid="wx_dfjskfsihf" data-miniprogram-path="/pages/home/index">小程序链接</a>  7、这是转人工
<a href="qdim://webim/message/extsendc2b?content=转人工&amp;ext={&quot;action&quot;:&quot;8&quot;,&quot;reception_type&quot;:0,&quot;reception_id&quot;:3007447581}&amp;svrId=2000000">点我转接人工客服</a>
7、这是普通文本啊啊全球领先的中文搜索引擎、致力于让网民更便捷地获取信息，找到所求。百度超过千亿的中文网页数据库，可以瞬间找到相关的搜索结果。
[不可用的转人工链接]ddddddddd[不可用的转人工链接]阿里巴巴IDST和腾讯PK[不可用的转人工链接]123456简单`

	strings := SplitText(str, 500)
	for i, item := range strings {

		fmt.Printf("%d----%s\n", i, item)
	}

	compile := regexp.MustCompile(TransferManProtocolRegexp)

	fmt.Println(compile.MatchString(`<a href="dddd">ddd</a>`))
	fmt.Println(compile.MatchString(`[不可用的转人工链接]`))

}

func completeInterval(size int, source [][]int) [][]int {
	var target [][]int
	prevEnd := 0

	for _, pair := range source {
		start, end := pair[0], pair[1]
		target = append(target, []int{prevEnd, start})
		target = append(target, []int{start, end})
		prevEnd = end
	}
	if prevEnd <= size {
		target = append(target, []int{prevEnd, size})
	}
	return target
}

func SplitText(str string, maxByteSize int) []string {
	compile := regexp.MustCompile(TransferManProtocolRegexp)
	sourceInterval := compile.FindAllStringIndex(str, -1)
	interval := completeInterval(len(str), sourceInterval)

	var res []string
	for _, item := range interval {
		res = append(res, str[item[0]:item[1]])
	}

	return merge(res, maxByteSize)

}

func merge(str []string, maxByteSize int) []string {

	currentString := ""
	var targetString []string

	for _, str := range str {
		if len(currentString)+len(str) > maxByteSize {
			targetString = append(targetString, currentString)
			currentString = ""
		}
		currentString += str // 累积字符串
	}

	if currentString != "" {
		targetString = append(targetString, currentString)
	}

	return targetString
}
