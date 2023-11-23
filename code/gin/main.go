package main

import (
	"fmt"
	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"html/template"
	"net/http"
	_ "net/http/pprof"
	"strings"
	"time"
)

func main() {

	res := map[string]any{"data": "1111", "record_id": "qqqq", "stopChat": true}
	s, _ := jsoniter.MarshalToString(res)
	substr := "\"stopChat\":true"
	contains := strings.Contains(s, substr)
	fmt.Println(contains)

	r := gin.New()

	html := template.Must(template.ParseFiles("index.tmpl"))
	r.SetHTMLTemplate(html)

	stream := &Event{
		NewClients:    make(chan chan string),
		ClosedClients: make(chan chan string),
		TotalClients:  make(map[chan string]bool),
	}

	go stream.listen()

	r.POST("/chat", StreamV2Middleware, stream.serve(), ChatStream)

	r.POST("/stream", Stream)

	r.GET("/stop", StopStream)

	r.GET("/test", Test)

	r.GET("/file", File)

	r.POST("/hello",
		PreProcessing,
		Hello,
		PostProcessing)

	r.POST("/hello-timeout2",
		PreProcessing,
		timeout.New(
			timeout.WithTimeout(5*time.Second),
			timeout.WithHandler(Hello),
			timeout.WithResponse(PostProcessingTimeOut),
		),
		PostProcessing)

	r.POST("/hello-timeout", timeout.New(
		timeout.WithTimeout(100*time.Microsecond),
		timeout.WithHandler(emptySuccessResponse),
		timeout.WithHandler(nextFunc),
	))

	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"ID":           "141852",
			"UserID":       "50",
			"SessionID":    "e065194d-888b-fce9-fba9-aec91823714e",
			"Score":        "0",
			"Reason":       "",
			"Question":     "李玟最近怎么了",
			"Answer":       "输入参数  HTTP 请求头公共参数请参见签名验证章节的 [公共参数说明](https://cloud.tencent.com/document/product/1095/42413)。  | 参数名称 | 必选 | 参数类型 | 参数描述 | | --- | --- | --- | --- | | operator_userid | 是 | String | 操作者用户 ID，必须为企业下具有操作资源权限的注册用户。 | | type | 是 | Integer | 发起会议类型：0：企业下全部用户可发起会议1：企业下部分用户可发起会议 | | users | 否 | String 数组 | 企业成员 userid 数组，必须为企业下的注册用户，仅当 type 为1时此字段有效。 |  ## 输出参数  成功返回空消息体，失败返回 [错误码](https://cloud.tencent.com/document/product/1095/43704) 和错误信息。",
			"StaffName":    "lili",
			"Prompt":       "请用下文信息推理来回答下面问题，如果问题与后文无关，请回答不相关 文本:[示例1] 设置企业成员发起会议的权限 => ## 设置企业成员发起会议的权限  ## 接口描述  **描述**：当账户类型为企业版时，可设置企业成员发起会议的权限。企业成员发起会议的权限包含2种：企业下全部用户可发起会议和企业下部分用户可发起会议，当权限为部分用户可发起会议时，可设置具有权限的用户列表，新增的用户会在已有的用户列表上进行追加。不支持 OAuth2.0 鉴权访问。  **调用方式**：POST  **接口请求域名**：  ```lang-Plaintext https://api.meeting.qq.com/v1/corp-resource/book-meeting/authorized-users ```   [示例2] 设置企业成员发起会议的权限 => 输入参数  HTTP 请求头公共参数请参见签名验证章节的 [公共参数说明](https://cloud.tencent.com/document/product/1095/42413)。  | 参数名称 | 必选 | 参数类型 | 参数描述 | | --- | --- | --- | --- | | operator\\_userid | 是 | String | 操作者用户 ID，必须为企业下具有操作资源权限的注册用户。 | | type | 是 | Integer | 发起会议类型：0：企业下全部用户可发起会议1：企业下部分用户可发起会议 | | users | 否 | String 数组 | 企业成员 userid 数组，必须为企业下的注册用户，仅当 type 为1时此字段有效。 |  ## 输出参数  成功返回空消息体，失败返回 [错误码](https://cloud.tencent.com/document/product/1095/43704) 和错误信息。   [示例3] 设置企业成员发起会议的权限 => 示例  #### 输入示例  ```lang-plaintext POSThttps://api.meeting.qq.com/v1/corp-resource/book-meeting/authorized-users{  \"operator_userid\": \"meeting4529091\",  \"type\": 0,  \"users\" : [\"meeting4529091\",\"meeting4529092\"]} ```  #### 输出示例  ```lang-plaintext {} ```  [示例4] 设置企业成员发起大型会议权限 => ## 接口描述  **描述**：当账户类型为商业版或企业版时，可为企业下具有发起会议权限的成员设置发起大型会议的权限。若企业成员未配置过大型会议权限，则为成员设置可发起大型会议权限；若企业成员配置过大型会议权限，则调整用户大型会议权限，覆盖已有的权限信息。不支持 OAuth2.0 鉴权访问。  **调用方式**：POST  **接口请求域名**：  ```lang-Plaintext https://api.meeting.qq.com/v1/corp-resource/book-large-meeting/authorized-users ```   [示例5] 腾讯会议如何设置企业成员发起会议的权限 => 您好，腾讯会议支持调用 [API接口](https://cloud.tencent.com/document/product/1095/56788) 设置企业成员发起会议的权限，详情如下： 1. 当账户类型为企业版时，可设置企业成员发起会议的权限。 2. 企业成员发起会议的权限包含2种：企业下全部用户可发起会议和企业下部分用户可发起会议，当权限为部分用户可发起会议时，可设置具有权限的用户列表，新增的用户会在已有的用户列表上进行追加。 3. 不支持OAuth2.0鉴权访问。  更多详情请参考：[腾讯会议设置企业成员发起会议的权限](https://cloud.tencent.com/document/product/1095/56788) 问题:调用【设置企业成员发起会议的权限】接口授权",
			"IndexContent": replace("{\"contents\":[\"## 设置企业成员发起会议的权限\\n\\n## 接口描述\\n\\n**描述**：当账户类型为企业版时，可设置企业成员发起会议的权限。企业成员发起会议的权限包含2种：企业下全部用户可发起会议和企业下部分用户可发起会议，当权限为部分用户可发起会议时，可设置具有权限的用户列表，新增的用户会在已有的用户列表上进行追加。不支持 OAuth2.0 鉴权访问。\\n\\n**调用方式**：POST\\n\\n**接口请求域名**：\\n\\n```lang-Plaintext\\nhttps://api.meeting.qq.com/v1/corp-resource/book-meeting/authorized-users\\n```\\n\",\"输入参数\\n\\nHTTP 请求头公共参数请参见签名验证章节的 [公共参数说明](https://cloud.tencent.com/document/product/1095/42413#.E5.85.AC.E5.85.B1.E5.8F.82.E6.95.B0)。\\n\\n| 参数名称 | 必选 | 参数类型 | 参数描述 |\\n| --- | --- | --- | --- |\\n| operator\\\\_userid | 是 | String | 操作者用户 ID，必须为企业下具有操作资源权限的注册用户。 |\\n| type | 是 | Integer | 发起会议类型：0：企业下全部用户可发起会议1：企业下部分用户可发起会议 |\\n| users | 否 | String 数组 | 企业成员 userid 数组，必须为企业下的注册用户，仅当 type 为1时此字段有效。 |\\n\\n## 输出参数\\n\\n成功返回空消息体，失败返回 [错误码](https://cloud.tencent.com/document/product/1095/43704) 和错误信息。\\n\",\"示例\\n\\n#### 输入示例\\n\\n```lang-plaintext\\nPOSThttps://api.meeting.qq.com/v1/corp-resource/book-meeting/authorized-users{  \\\"operator_userid\\\": \\\"meeting4529091\\\",  \\\"type\\\": 0,  \\\"users\\\" : [\\\"meeting4529091\\\",\\\"meeting4529092\\\"]}\\n```\\n\\n#### 输出示例\\n\\n```lang-plaintext\\n{}\\n```\",\"## 接口描述\\n\\n**描述**：当账户类型为商业版或企业版时，可为企业下具有发起会议权限的成员设置发起大型会议的权限。若企业成员未配置过大型会议权限，则为成员设置可发起大型会议权限；若企业成员配置过大型会议权限，则调整用户大型会议权限，覆盖已有的权限信息。不支持 OAuth2.0 鉴权访问。\\n\\n**调用方式**：POST\\n\\n**接口请求域名**：\\n\\n```lang-Plaintext\\nhttps://api.meeting.qq.com/v1/corp-resource/book-large-meeting/authorized-users\\n```\\n\",\"您好，腾讯会议支持调用 [API接口](https://cloud.tencent.com/document/product/1095/56788) 设置企业成员发起会议的权限，详情如下：\\n1. 当账户类型为企业版时，可设置企业成员发起会议的权限。\\n2. 企业成员发起会议的权限包含2种：企业下全部用户可发起会议和企业下部分用户可发起会议，当权限为部分用户可发起会议时，可设置具有权限的用户列表，新增的用户会在已有的用户列表上进行追加。\\n3. 不支持OAuth2.0鉴权访问。\\n\\n更多详情请参考：[腾讯会议设置企业成员发起会议的权限](https://cloud.tencent.com/document/product/1095/56788)\"],\"sources\":[\"andon\",\"andon\",\"andon\",\"andon\",\"andon\"],\"ids\":[\"ff599a227dc06c3f27eea54c29c011dc\",\"ff599a227dc06c3f27eea54c29c011dc\",\"ff599a227dc06c3f27eea54c29c011dc\",\"a7cf29ce463e2208eb3db759e924194f\",\"47561\"],\"links\":[\"https://meeting.tencent.com/support/topic/745\",\"https://meeting.tencent.com/support/topic/745\",\"https://meeting.tencent.com/support/topic/745\",\"https://meeting.tencent.com/support/topic/746\",\"\"],\"texts\":[\"设置企业成员发起会议的权限\",\"设置企业成员发起会议的权限\",\"设置企业成员发起会议的权限\",\"设置企业成员发起大型会议权限\",\"腾讯会议如何设置企业成员发起会议的权限\"],\"scores\":[5.4360104,5.0825963,4.816036,4.360688,4.257387],\"source_types\":[\"andon\",\"andon\",\"andon\",\"andon\",\"andon\"],\"doc_types\":[\"qa\",\"doc\",\"doc\",\"qa\",\"qa\"],\"page_contents\":[\"## 设置企业成员发起会议的权限\\n\\n## 接口描述\\n\\n**描述**：当账户类型为企业版时，可设置企业成员发起会议的权限。企业成员发起会议的权限包含2种：企业下全部用户可发起会议和企业下部分用户可发起会议，当权限为部分用户可发起会议时，可设置具有权限的用户列表，新增的用户会在已有的用户列表上进行追加。不支持 OAuth2.0 鉴权访问。\\n\\n**调用方式**：POST\\n\\n**接口请求域名**：\\n\\n```lang-Plaintext\\nhttps://api.meeting.qq.com/v1/corp-resource/book-meeting/authorized-users\\n```\\n\",\"输入参数\\n\\nHTTP 请求头公共参数请参见签名验证章节的 [公共参数说明](https://cloud.tencent.com/document/product/1095/42413#.E5.85.AC.E5.85.B1.E5.8F.82.E6.95.B0)。\\n\\n| 参数名称 | 必选 | 参数类型 | 参数描述 |\\n| --- | --- | --- | --- |\\n| operator\\\\_userid | 是 | String | 操作者用户 ID，必须为企业下具有操作资源权限的注册用户。 |\\n| type | 是 | Integer | 发起会议类型：0：企业下全部用户可发起会议1：企业下部分用户可发起会议 |\\n| users | 否 | String 数组 | 企业成员 userid 数组，必须为企业下的注册用户，仅当 type 为1时此字段有效。 |\\n\\n## 输出参数\\n\\n成功返回空消息体，失败返回 [错误码](https://cloud.tencent.com/document/product/1095/43704) 和错误信息。\\n\",\"示例\\n\\n#### 输入示例\\n\\n```lang-plaintext\\nPOSThttps://api.meeting.qq.com/v1/corp-resource/book-meeting/authorized-users{  \\\"operator_userid\\\": \\\"meeting4529091\\\",  \\\"type\\\": 0,  \\\"users\\\" : [\\\"meeting4529091\\\",\\\"meeting4529092\\\"]}\\n```\\n\\n#### 输出示例\\n\\n```lang-plaintext\\n{}\\n```\",\"## 接口描述\\n\\n**描述**：当账户类型为商业版或企业版时，可为企业下具有发起会议权限的成员设置发起大型会议的权限。若企业成员未配置过大型会议权限，则为成员设置可发起大型会议权限；若企业成员配置过大型会议权限，则调整用户大型会议权限，覆盖已有的权限信息。不支持 OAuth2.0 鉴权访问。\\n\\n**调用方式**：POST\\n\\n**接口请求域名**：\\n\\n```lang-Plaintext\\nhttps://api.meeting.qq.com/v1/corp-resource/book-large-meeting/authorized-users\\n```\\n\",\"您好，腾讯会议支持调用 [API接口](https://cloud.tencent.com/document/product/1095/56788) 设置企业成员发起会议的权限，详情如下：\\n1. 当账户类型为企业版时，可设置企业成员发起会议的权限。\\n2. 企业成员发起会议的权限包含2种：企业下全部用户可发起会议和企业下部分用户可发起会议，当权限为部分用户可发起会议时，可设置具有权限的用户列表，新增的用户会在已有的用户列表上进行追加。\\n3. 不支持OAuth2.0鉴权访问。\\n\\n更多详情请参考：[腾讯会议设置企业成员发起会议的权限](https://cloud.tencent.com/document/product/1095/56788)\"],\"org_datas\":[\"## 设置企业成员发起会议的权限\\n\\n## 接口描述\\n\\n**描述**：当账户类型为企业版时，可设置企业成员发起会议的权限。企业成员发起会议的权限包含2种：企业下全部用户可发起会议和企业下部分用户可发起会议，当权限为部分用户可发起会议时，可设置具有权限的用户列表，新增的用户会在已有的用户列表上进行追加。不支持 OAuth2.0 鉴权访问。\\n\\n**调用方式**：POST\\n\\n**接口请求域名**：\\n\\n```lang-Plaintext\\nhttps://api.meeting.qq.com/v1/corp-resource/book-meeting/authorized-users\\n```\\n\",\"输入参数\\n\\nHTTP 请求头公共参数请参见签名验证章节的 [公共参数说明](https://cloud.tencent.com/document/product/1095/42413)。\\n\\n| 参数名称 | 必选 | 参数类型 | 参数描述 |\\n| --- | --- | --- | --- |\\n| operator\\\\_userid | 是 | String | 操作者用户 ID，必须为企业下具有操作资源权限的注册用户。 |\\n| type | 是 | Integer | 发起会议类型：0：企业下全部用户可发起会议1：企业下部分用户可发起会议 |\\n| users | 否 | String 数组 | 企业成员 userid 数组，必须为企业下的注册用户，仅当 type 为1时此字段有效。 |\\n\\n## 输出参数\\n\\n成功返回空消息体，失败返回 [错误码](https://cloud.tencent.com/document/product/1095/43704) 和错误信息。\\n\",\"示例\\n\\n#### 输入示例\\n\\n```lang-plaintext\\nPOSThttps://api.meeting.qq.com/v1/corp-resource/book-meeting/authorized-users{  \\\"operator_userid\\\": \\\"meeting4529091\\\",  \\\"type\\\": 0,  \\\"users\\\" : [\\\"meeting4529091\\\",\\\"meeting4529092\\\"]}\\n```\\n\\n#### 输出示例\\n\\n```lang-plaintext\\n{}\\n```\",\"## 接口描述\\n\\n**描述**：当账户类型为商业版或企业版时，可为企业下具有发起会议权限的成员设置发起大型会议的权限。若企业成员未配置过大型会议权限，则为成员设置可发起大型会议权限；若企业成员配置过大型会议权限，则调整用户大型会议权限，覆盖已有的权限信息。不支持 OAuth2.0 鉴权访问。\\n\\n**调用方式**：POST\\n\\n**接口请求域名**：\\n\\n```lang-Plaintext\\nhttps://api.meeting.qq.com/v1/corp-resource/book-large-meeting/authorized-users\\n```\\n\",\"您好，腾讯会议支持调用 [API接口](https://cloud.tencent.com/document/product/1095/56788) 设置企业成员发起会议的权限，详情如下：\\n1. 当账户类型为企业版时，可设置企业成员发起会议的权限。\\n2. 企业成员发起会议的权限包含2种：企业下全部用户可发起会议和企业下部分用户可发起会议，当权限为部分用户可发起会议时，可设置具有权限的用户列表，新增的用户会在已有的用户列表上进行追加。\\n3. 不支持OAuth2.0鉴权访问。\\n\\n更多详情请参考：[腾讯会议设置企业成员发起会议的权限](https://cloud.tencent.com/document/product/1095/56788)\"],\"org_titles\":[\"设置企业成员发起会议的权限\",\"设置企业成员发起会议的权限\",\"设置企业成员发起会议的权限\",\"设置企业成员发起大型会议权限\",\"腾讯会议如何设置企业成员发起会议的权限\"],\"products\":[\"\",\"\",\"\",\"\",\"\"]}"),
			"ModelID":      65,
			"TraceID":      "4rrrfff",
			"CreateTime":   "2023-07-06T16:00:17+08:00",
			"UpdateTime":   "2023-07-06T16:00:18+08:00",
			"AllData":      `{"code":0,"message":"","data":{"ID":141852,"UserID":50,"SessionID":"e065194d-888b-fce9-fba9-aec91823714e","Score":0,"Reason":"","Question":"李玟最近怎么了","Answer":"李玟最近一直在努力应对抑郁症，她曾多次投入工作，尽力保持乐观。但不幸的是，她的努力并没有让她度过抑郁症这一难关。最终，她决定以自己的方式离开这个世界，令人痛心。","StaffName":"timjlin","Prompt":"基于以下已知信息，简洁和专业的来回答用户的问题。如果无法从中得到答案，忽略文段内容并用中文回答用户问题。\n已知内容:\n华语乐坛国际级歌手李玟罹患抑郁症轻生，终年48岁- BBC News 中文 在1990及2000年代享誉亚洲流行乐坛的华人歌手李玟（Coco Lee）日前离世，终年48岁。 李玟胞姐李思林在周三（7月5日）晚间向媒体公布妹妹的死讯，当中指李 \n李玟因忧郁症轻生离世\"首位进军美国乐坛华人歌手\" - RFI 歌手李玟7月5日因忧郁症缠身轻生去世。李玟生前创华语歌手多项纪录，被指是首位进军美国乐坛、在全球发行英语专辑的华语女歌手，也是首位登上奥斯卡 \n李玟因抑郁症轻生，不久前她才说：从来不怕任何困难 - 新浪财经 今年2月，李玟做手术治疗腿伤，她特意发文给自己打气：我从来不怕任何困难，一切问题总是一个一个来解决。在文中，她还对大家说：人总有脆弱的时候，这是 \n李玟轻生原因曝光！曾表示非常想要一个小孩，最后语音令人心碎…… - 新浪新闻 据报道，讣闻透露李玟因饱受抑郁症困扰轻生，据知李玟由于几年前受婚变问题缠绕，令原本性格开朗的她患上抑郁症。不过李玟一直努力医病，不时投入工作去 \n李玟人生的最后一年：未放弃最爱的音乐，为回归舞台接受手术 - 南方网 人生的最后一年，李玟仍从未放弃最爱的音乐事业，甚至为了回归舞台接受 ... 令人难过的是，就在7月2日，歌迷会微博还发布了一段来自李玟的最新语音。\n\n问题:\n李玟最近怎么了","IndexContent":"{\"contents\":null,\"sources\":[\"google\",\"google\",\"google\",\"google\",\"google\"],\"ids\":[\"0\",\"0\",\"0\",\"0\",\"0\"],\"links\":[\"https://www.bbc.com/zhongwen/simp/chinese-news-66117881\",\"https://www.rfi.fr/cn/%E4%B8%AD%E5%9B%BD/20230705-%E6%9D%8E%E7%8E%9F%E5%9B%A0%E5%BF%A7%E9%83%81%E7%97%87%E8%BD%BB%E7%94%9F%E7%A6%BB%E4%B8%96-%E9%A6%96%E4%BD%8D%E8%BF%9B%E5%86%9B%E7%BE%8E%E5%9B%BD%E4%B9%90%E5%9D%9B%E5%8D%8E%E4%BA%BA%E6%AD%8C%E6%89%8B\",\"https://finance.sina.com.cn/tech/roll/2023-07-05/doc-imyzrzqf2834817.shtml\",\"https://news.sina.com.cn/c/2023-07-06/doc-imyzsfwh9414356.shtml\",\"https://news.southcn.com/node_17a07e5926/d18d552cf2.shtml\"],\"texts\":[\"华语乐坛国际级歌手李玟罹患抑郁症轻生，终年48岁- BBC News 中文\",\"李玟因忧郁症轻生离世\\\"首位进军美国乐坛华人歌手\\\" - RFI\",\"李玟因抑郁症轻生，不久前她才说：从来不怕任何困难 - 新浪财经\",\"李玟轻生原因曝光！曾表示非常想要一个小孩，最后语音令人心碎…… - 新浪新闻\",\"李玟人生的最后一年：未放弃最爱的音乐，为回归舞台接受手术 - 南方网\"],\"scores\":[0,0,0,0,0],\"source_types\":[\"\",\"\",\"\",\"\",\"\"],\"doc_types\":[\"\",\"\",\"\",\"\",\"\"],\"page_contents\":[\"在1990及2000年代享誉亚洲流行乐坛的华人歌手李玟（Coco Lee）日前离世，终年48岁。 李玟胞姐李思林在周三（7月5日）晚间向媒体公布妹妹的死讯，当中指李 \",\"歌手李玟7月5日因忧郁症缠身轻生去世。李玟生前创华语歌手多项纪录，被指是首位进军美国乐坛、在全球发行英语专辑的华语女歌手，也是首位登上奥斯卡 \",\"今年2月，李玟做手术治疗腿伤，她特意发文给自己打气：我从来不怕任何困难，一切问题总是一个一个来解决。在文中，她还对大家说：人总有脆弱的时候，这是 \",\"据报道，讣闻透露李玟因饱受抑郁症困扰轻生，据知李玟由于几年前受婚变问题缠绕，令原本性格开朗的她患上抑郁症。不过李玟一直努力医病，不时投入工作去 \",\"人生的最后一年，李玟仍从未放弃最爱的音乐事业，甚至为了回归舞台接受 ... 令人难过的是，就在7月2日，歌迷会微博还发布了一段来自李玟的最新语音。\"],\"org_datas\":[\"在1990及2000年代享誉亚洲流行乐坛的华人歌手李玟（Coco Lee）日前离世，终年48岁。 李玟胞姐李思林在周三（7月5日）晚间向媒体公布妹妹的死讯，当中指李 \",\"歌手李玟7月5日因忧郁症缠身轻生去世。李玟生前创华语歌手多项纪录，被指是首位进军美国乐坛、在全球发行英语专辑的华语女歌手，也是首位登上奥斯卡 \",\"今年2月，李玟做手术治疗腿伤，她特意发文给自己打气：我从来不怕任何困难，一切问题总是一个一个来解决。在文中，她还对大家说：人总有脆弱的时候，这是 \",\"据报道，讣闻透露李玟因饱受抑郁症困扰轻生，据知李玟由于几年前受婚变问题缠绕，令原本性格开朗的她患上抑郁症。不过李玟一直努力医病，不时投入工作去 \",\"人生的最后一年，李玟仍从未放弃最爱的音乐事业，甚至为了回归舞台接受 ... 令人难过的是，就在7月2日，歌迷会微博还发布了一段来自李玟的最新语音。\"],\"org_titles\":[\"华语乐坛国际级歌手李玟罹患抑郁症轻生，终年48岁- BBC News 中文\",\"李玟因忧郁症轻生离世\\\"首位进军美国乐坛华人歌手\\\" - RFI\",\"李玟因抑郁症轻生，不久前她才说：从来不怕任何困难 - 新浪财经\",\"李玟轻生原因曝光！曾表示非常想要一个小孩，最后语音令人心碎…… - 新浪新闻\",\"李玟人生的最后一年：未放弃最爱的音乐，为回归舞台接受手术 - 南方网\"],\"products\":[\"\",\"\",\"\",\"\",\"\"]}","ModelID":65,"TraceID":"851117b0-5783-4ab4-ac61-47f299d0c9f7","CreateTime":"2023-07-06T16:00:17+08:00","UpdateTime":"2023-07-06T16:00:18+08:00"}}`,
		})
	})

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}

func replace(str string) string {
	return strings.ReplaceAll(str, "\\n", "\n")
}

func emptySuccessResponse(c *gin.Context) {
	time.Sleep(30000 * time.Microsecond)
	c.String(http.StatusOK, "")
}

func nextFunc(c *gin.Context) {
	c.String(http.StatusOK, "quick timeout")
}