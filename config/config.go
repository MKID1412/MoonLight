package config

var RootDir = "../../"

// request请求相关配置
var ReqHeaders = map[string]string{
	"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:124.0) Gecko/20100101 Firefox/124.0",
}
var RedirectCount = 5 //最多重定向次数

// log文件存放配置
var ErrorLog = RootDir + "log/error.log"
var ReqErrLog = RootDir + "log/utilsRequestError.log"
var ImgDayErrLog = RootDir + "log/ImgDayError.log"
var CommonLog = RootDir + "log/log.log"
var LogSize = 3000 //当日志大小超过3000kb时会清除一次

// 备案查询相关配置
var BeianConfigFile = RootDir + "config/beianconfig.yaml"
var BeianDataHeader = []string{"域名", "ip", "备案单位", "性质", "法人", "权重", "行业", "注册资金", "备案号"}

// 空间测绘相关配置
var SpaceEngKeyFileName = RootDir + "config/spacekey.yaml"
var HunterDataHeader = []string{"ip", "标题", "url", "备案信息", "位置"}
var FofaDataHeader = []string{"ip", "标题", "url", "服务", "位置"}

// 漏洞测试相关配置
var VulScanDataHeader = []string{"url", "存在漏洞", "备案单位", "单位性质", "权重"}
var WebFingersJson = RootDir + "QtGui/nday/webfingers.json"
var PocDirName = RootDir + "QtGui/nday/pocs/"
var RequestTimeout = 20

// 休闲修养相关配置
var Wish = "愿芸芸众生,皆免其殇,安得善终"
var ImgDayFirstFileName = ":/xhc/img/13.jpg"
var ImgDayTitle = "无"

// var ImgDayUrl="https://api.vvhan.com/api/bing?rand=sj" //随机图片但是无标题
var ImgDayRandomUrl = "https://api.vvhan.com/api/bing?type=json&rand=sj"
var ImgDayFileName = RootDir + "QtGui/img/ImgDay.jpg"
var ImgDayNextFileName = RootDir + "QtGui/img/ImgDayNext.jpg"
var PcInfoUrl = "https://api.vvhan.com/api/visitor.info"

// 扩展功能相关
var ExpandConfigFileName = RootDir + "config/expand.yaml"
var ExpandExplain="选择的go文件会被复制到config目录中，因此也可以手动将go文件放到config目录，功能名就是左边列表中要显示的名字，函数名就是自己定义好的界面后能返回*widgets.QWidget类型的函数，右边就会显示对应界面，可以查看expand.go中的一个界面的例子ExampleGui函数，函数名首字母注意大写同时目前只能写为ExpandGui3~ExpandGui8之前，后续不够需要添加的话在expand/expand.go的GetFuncAndWidget函数中添加case，具体详情可以查看该函数，最后需要重新qtdeploy编译一次：qtdeploy build desktop main.go"

var AizhanExpress = map[string]string{
	"Alldata_express":       `(?s)<td class="thead">主办单位名称</td>(.*?)<h4>域名备案历史记录</h4>`,   //含有下面所有数据的匹配语法,即“主办单位名称</td>”到“公司地址”地址之间的所有数据
	"Master_express":        `(?s)<td class="thead">主办单位名称</td>\s*<td>(.*?)&nbsp;&nbsp;`, //备案单位正则匹配语法
	"LegalHuman_express":    `法定代表人</span><span>(.*?)</span></td>`,                       //法定代表人正则匹配语法
	"Industry_express":      `行业</span><span>(.*?)</span></td>`,                          //所属行业正则匹配语法
	"Nature_express":        `(?s)<td class="thead">主办单位性质</td>\s*<td>(.*?)</td>`,
	"BeianNumber_express":   `<td class="thead">网站备案/许可证号</td>\s*<td><span>(.*?)</span>`, //备案号正则匹配语法
	"RegisterMoney_express": `注册资本</span><span>(.*?)</span></td>`,
}
var ChinazExpress = map[string]string{
	"Alldata_express":       `(?s)工商信息</span><span class="sub-count">(.*?)<div id="datalog" class="company-content"`, //默认情况下, . 不匹配换行符，但是使用了 (?s) 后, . 将匹配一切字符
	"Master_express":        ``,                                                                                      //备案单位
	"LegalHuman_express":    `<td width="15%">法定代表人</td>\s*<td width="30%">(.*?)</td>`,                               //法定代表人
	"Industry_express":      `<td>所属行业</td>\s*<td>(.*?)</td>`,
	"Nature_express":        ``,
	"BeianNumber_express":   ``,
	"RegisterMoney_express": `<td width="15%">注册资本</td>\s*<td width="20%">(.*?)</td>`, //备案号                                                  //所属行业
}
var Chinaz_Master_express=`data-company="(.*?)"`
var Chinaz_Nature_express=`<div class="contactPhone">(.*?)</div>`
var ChinazSEOUrl="https://seo.chinaz.com/"
var ChinazSEO_Nature_express=`性质：\s*<i class="color-63">(.*?)</i>`
var ChinazSEO_Master_express=`<a href="//data\.chinaz\.com/company/t0-p0-c0-i0-d0-s-(.*?)"`
var WeightExpress = map[string]string{
	"PC_Br_express": `<img src="//statics\.aizhan\.com/images/br/(\w+)\.png"`,
	"M_Br_express":  `<img src="//statics\.aizhan\.com/images/mbr/(\w+)\.png"`,
	"Pr_express":    `<img src="//statics\.aizhan\.com/images/pr/(\w+)\.png"`,
}

//------------------------------------beianconfig.yaml
// aizhan:
//   url: https://icp.aizhan.com/
//   cookie: userId=0000000  #userId值可以随便一点
//   seo_url: https://www.aizhan.com/cha/
//   br_url: https://rank.aizhan.com/
//   pr_url: https://pr.aizhan.com/
//   weightapi_url: https://apistore.aizhan.com/baidurank/siteinfos/
//   weightapi_key:
// chinaz:
//   url: https://icp.chinaz.com/
//   cookie: "toolUserGrade=DA558BECA59696EB6D6F7073658259097496A34F9B3E8B35F3075E72A88B4A26CBAD406B542FEEF4A7B817C737962796746C7AC7384A4811136F863242F69BE99A6DC128CA2894975194E08B408B07AA2FE270DB81B779A720710205372A7E3A0522C075AD31A5B950B9552E0E0F299DA6400E54A5768622; bbsmax_user=7b90201e-7f31-71f3-f9c4-ccc6dfd2078B"   #主要是toolUserGrade和bbsmax_user值，bbsmax_user值可以随便一点
//   getip_url: https://ip.tool.chinaz.com/api/getResolvedIps?keyword=
