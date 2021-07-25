package data

var Jobs = map[string][]string{
	"jobs": {"农业管理人员", "农业生产人员", "农业技师", "农机修护人员", "农业生产人员", "其他农业技师人员", "林业管理人员", "林木种植人员", "森林砍伐工人", "木材加工工人", "林业管理人员", "林木种植人员", "畜牧业管理人员", "畜牧业工人", "畜牧业管理人员", "渔业管理人员", "远洋渔业工人", "渔业管理人员", "内陆渔业工人", "水族馆工人", "渔业实验人员", "渔业工人", "煤炭开采和洗选业管理人员", "煤炭开采和洗选业地下工人", "煤炭开采和洗选业地上工人", "金属矿采选管理人员", "金属矿采选工人", "金属矿技术人员", "其他矿采选业管理人员", "其他矿采选业技术人员", "其他矿采选业工人", "矿山工程技术人员", "其他矿采选业工人", "油气开采管理人员", "油气开采技术人员", "油气开采工人", "油气运输工人", "食品、饮料制造管理人员", "食品、饮料制造技术人员", "食品、饮料制造工人", "食品加工工人", "烟草制造管理人员", "烟草制造机器操作人员", "烟草制造工人", "纺织、成衣制造管理人员", "纺织、成衣制造技术人员", "纺织、成衣制造工人", "皮革业管理人员", "皮革制造工人", "羽绒、羽毛制造工人", "竹、木、藤、草编制管理人员", "竹、木、藤、草编制工人", "家俱制造管理人员", "家俱制造技术人员", "家俱制造工人", "造纸业管理人员", "造纸业技术人员", "造纸工人", "印刷、文教、体育用品制造管理人员", "印刷、文教、体育用品制造技术人员", "印刷、文教、体育用品制造工人", "核工业燃料制造管理人员", "核工业燃料制造技术人员", "核工业燃料制造工人", "石油加工、炼焦管理人员", "石油加工技术人员", "石油加工工人", "炼焦技术人员", "炼焦工人", "化学原料及化学制品制造业管理人员", "化学原料及化学制品制造技术人员", "硫酸、盐酸、硝酸等有毒品制造工人", "化学原料及化学制品制造工人", "火药、爆竹制造处理人员(爆竹烟火)", "医药制造管理人员", "医药制造技术人员", "医药制造工人", "化学纤维制造管理人员", "化纤制造技术人员", "化纤制造工人", "橡胶制品管理人员", "橡胶制品技术人员", "橡胶制品工人", "塑料制品管理人员", "塑料制品技术人员", "塑料制品工人", "非金属矿物制品管理人员", "水泥、石灰制造技术人员", "水泥、石灰制造工人", "金属冶炼及压延加工管理人员", "金属冶炼技术人员", "金属冶炼工人", "金属制品管理人员", "金属制品技术人员", "金属制造工人", "其他金属制造技术工人", "其他金属制造工人", "设备、机械制造管理人员", "家电、电机技术人员", "家电、电机工人", "汽车、摩托车、自行车制造人员", "造修船技术人员", "造修船工人", "通信设备、计算机及其他电子设备制造管理人员", "电子技术人员", "电子工人", "电线电缆工人", "有关高压电工人", "仪器仪表、文化、办公用机械及其它制造管理人员", "仪器仪表、文化、办公用机械及其它制造技术人员", "仪器仪表制造人员", "文化、办公机械制造人员", "工艺品及其他制造业管理人员", "手工艺品制造技术人员", "手工艺品制造工人", "玻璃制造工人", "砖木、陶瓷制造工人", "音像、乐器、眼镜制造工人", "废弃资源和废旧材料重加工管理人员", "废旧材料加工技术人员", "废旧材料加工工人", "其他加工制造人员", "水电生产和供应业管理人员", "电力热力的生产和供应技术人员", "电力热力的生产和供应工人", "燃气生产和供应管理人员", "燃气生产和供应技术人员", "燃气公司工人", "水的生产和供应业管理人员", "水利工程设施技术人员", "水的生产和供应工人", "建筑业管理人员", "建筑业工人", "室内装璜人员", "室外装璜人员", "电梯升降机", "电梯操作员", "电梯维修工人", "建筑业设计人员", "铁路运输管理人员", "铁路道路修护人员", "铁路、道路铺设工人", "铁路车站工作人员", "铁路运输修护人员", "城市公共交通管理人员", "客运司机及随车工作人员", "客运公司工作人员", "道路运输管理人员", "人力三轮车夫", "机动三轮车摩托车驾驶员", "货车司机及随车工作人员", "货运公司工作人员", "其他货车司机", "水上运输管理人员", "远洋航运工作人员（非船上）", "远洋航运船上作业人员", "港口作业人员", "港口吊车人员", "港口缉私稽查人员", "救难船员", "内陆、沿海航运人员", "关务人员", "航空运输管理人员", "航空业工作人员", "机场工作人员", "航空机上工作人员", "地面维修人员", "管道运输管理人员", "油气输送工作人员", "装卸搬运和其他运输服务管理人员", "装卸搬运工人", "其他装卸搬运工人", "仓库管理人员", "仓库工人", "邮政管理人员", "邮政内勤人员", "邮政外勤人员", "电信和其他信息传输服务管理人员", "电信电力工程技术人员", "电力工程设施架设人员", "电力高压电工程设施人员", "电台天线维护人员", "计算机服务及软件管理人员", "计算机维护人员", "硬件开发及生产人员", "软件开发及生产人员", "批发零售管理人员", "室内批发零售人员", "室外批发零售人员", "流动摊贩", "燃气矿物油销售人员", "医药销售及推广人员", "旅馆餐饮业管理人员", "旅馆工作人员", "餐饮工作人员", "金融管理人员", "金融内勤人员", "金融外勤人员", "保险管理人员", "保险内勤人员", "保险外勤人员", "房地产管理人员", "房地产室内工作人员", "房地产开发工作人员", "房地产租赁销售人员", "社会服务及居民服务管理人员", "商务中心工作人员", "社区物业管理人员", "美容美发洗浴按摩人员", "中介服务人员", "殡葬服务人员", "家政服务人员", "清洁工", "其他清洁工", "采购人员", "车辆修理人员", "  ", "其他社会服务及居民服务工作人员", "科学研究与技术服务管理人员", "科学研究人员", "专业技术人员", "其他专业技术人员", "旅游管理人员", "旅游工作人员", "旅游外勤人员", "地质勘查管理人员", "地质探测测绘人员", "桥梁、隧道工程人员", "潜水、爆破工程人员", "环境和公共设施管理人员", "环境保护技术人员", "环境保护工作人员", "核工业工程环保人员", "安全质量检查人员", "教育机关(学校)管理人员", "教师", "学生", "文教相关工作人员", "医务行政管理人员", "医生", "护士", "医技人员", "医院后勤人员", "其他卫生机构人员", "社会保障社会福利管理人员", "社会保障社会福利工作人员", "新闻出版与广播、电视、电影及音像业管理人员", "武打演员", "特技演员", "广播、电影、电视从业人员", "新闻媒体工作人员", "报纸杂志从业人员", "摄影记者", "战地记者", "文化艺术广告业管理人员", "文化艺术工作人员", "文艺、书画创作人员", "考古及文物保护人员", "广告公司工作人员", "广告招牌架设人员", "体育管理人员", "职业教练员", "职业运动员", "职业运动相关人员", "运动场馆工作人员", "其他职业教练员1", "其他职业教练员2", "其他职业教练员3", "其他职业教练员4", "其他职业教练员5", "其他职业运动员1", "其他职业运动员2", "其他职业运动员3", "其他职业运动员4", "其他职业运动员5", "娱乐业管理人员", "歌舞厅、酒吧工作人员", "歌舞厅、酒吧保安", "驯兽师", "艺术演艺人员", "马戏团杂技表演人员", "高尔夫球场工作人员", "保龄球馆工作人员", "撞球场工作人员", "游泳池工作人员", "海水浴场工作人员", "海水浴场救生员", "其他娱乐场所工作人员", "党政机关工作人员", "社会团体工作人员", "寺庙及宗教管理人员", "僧尼,道士及传教人员", "国际组织工作人员", "文职军人", "军官", "特种兵(水中爆破兵、情治单位伞兵、化学兵、爆破任负有特殊任务者…等)", "现役陆地军人", "军事单位武器、弹药研究及管理人员", "现役海军、空军军人", "警务行政及内勤人员", "其他警务人员", "警务特勤人员", "防暴警察", "自由职业者内勤", "自由职业者外勤", "家庭主妇", "离退休人员", "无业、待业人员", "学龄前儿童"},
}
