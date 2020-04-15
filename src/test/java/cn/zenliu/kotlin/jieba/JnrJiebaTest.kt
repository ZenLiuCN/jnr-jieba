package cn.zenliu.kotlin.jieba

import cn.zenliu.kotlin.jieba.JnrJieba.TokenizerMode.CUT_HMM
import org.junit.jupiter.api.Test

import org.junit.jupiter.api.Assertions.*
import kotlin.system.measureTimeMillis

internal class JnrJiebaTest {
	val art = listOf(
			"　原标题：在这里工作45年后，李先念女儿退休",
			"　　来源：北京青年报",
			"　　撰文 | 董鑫",
			"　　中国人民对外友好协会微信公众号4月13日消息，4月9日，中国人民对外友好协会全国理事会常务理事会召开会议。",
			"　　会议通过决议，林松添同志（中国前驻南非大使）接替李小林同志担任中国人民对外友好协会全国理事会常务理事、会长。",
			"　　上述消息意味着，在对外友协工作45年之后，李小林正式退休。",
			"　　李小林出生于1953年10月，湖北红安人，她是中国前国家主席李先念的女儿。在致友好人士、友好组织的一封信中，李小林写道：",
			"　　“我将愉快地结束在友协45年的工作生涯，开启我人生的新一段体验享受轻松惬意的退休生活。45年前，我的父亲帮助我选择了友协这份工作。他当时对我说，和各国的朋友在一起，一定会感到无比快乐。”",
			"　　全国友协唯一女会长",
			"　　李先念主管国家经济工作26年，但他不允许自己的儿女经商，李家的四个子女没有一人下海。",
			"　　在接受媒体采访时，李小林回忆说，改革开放后，李先念有一次在饭桌上对孩子们严厉地说：“你们谁要经商，打断你们的腿。”",
			"　　作为李先念和夫人林佳楣最小的女儿，李小林本想成为一名护士，但李先念认为她的性格比较外向，就帮她选择了民间友好的工作。",
			"　1975年，李小林从武汉大学英语系毕业，到对外友协当了一名翻译。2011年10月，李小林接任陈昊苏，成为对外友协的第九任会长，同时也是对外友协成立以来唯一一位女会长。",
			"　　中国人民对外友好协会是从事民间活动的社会团体，成立于1954年。",
			"　　在中美两国建交之前，对外友协是中国唯一能够接待美国代表团的民间组织，承担了中美两国人民交往的大部分工作。",
			"　民间外交女杰",
			"　　工作至今，除了赴美攻读硕士、出任两年中国驻美国大使馆一等秘书外，李小林一直在对外友协工作，并常年负责对美洲和大洋洲地区的民间外交事务，有海外华文报纸称她为“中国民间外交女杰”。",
			"　　在美国遭受“9·11”恐怖袭击之后，李小林率领的对外友协代表团先于其他中国政要访问美国，出席美中友协在华盛顿召开的第13届美中关系研讨会和第18届年会，并频繁地拜访各界政要，带去了中国人民对美国人民的同情与问候。",
			"　　“非典”时期，世界各地一度引起恐慌，许多国家和地区拒绝中国人到访。李小林第一个率团访美，出席全球健康卫生理事会年会，并争取到了会议主办者的支持，在会议期间专设了中国非典情况报告会。",
			"　“5·12”汶川大地震时，李小林正在大洋洲访问，得知四川汶川发生特大地震，立即致电当时全国友协的会长陈昊苏，商讨成立抗震救灾工作领导小组，积极配合中央抗震救灾工作的开展。",
			"　　回顾45年的工作，在4月9日发出的致友好人士、友好组织的一封信中，李小林表示，“父亲的话一点儿也没有错！”她结识了千千万万的朋友，让她的人生充满了精彩，也深深体会到，为世界和平人类进步而贡献自己的青春与才华是多么神圣的事业。",
			"　　湖北省长致信感谢",
			"　　2020年新冠肺炎疫情期间，也能发现李小林忙碌的身影。",
			"　　政知圈（微信ID：wepolitics）在全国友协的官网上，看到一封湖北省省长王晓东于今年3月29日写给李小林的感谢信。",
			"　　信中提到，抗疫期间，李小林和全国友协全体干部职工通过各种渠道协调捐赠了大量医疗防护物资和生活物资援助湖北，在非常时期发挥了友协的特殊作用，架起了一座联通海外、凝心聚力的友谊桥梁，极大地增强了湖北人民战胜疫情的信心，为湖北打赢疫情防控阻击战发挥了积极作用。",
			"　3月30日，“同舟共济，携手抗疫医用口罩和防护服捐赠仪式”在京举办。",
			"　　全国对外友协联合江西省抚州市人民政府、孚真堂等国内机构，通过西班牙驻华使馆向西班牙阿尔卡拉市和其他机构捐赠4万只一次性医用口罩、1.5万件医用防护服等价值140万元人民币的防疫物资，用于支援当地抗击新冠肺炎疫情。",
			"　　李小林出席物资捐赠仪式并致辞。",
			"　4月9日，在致友好人士、友好组织的一封信中，李小林表示，回望过去的45年，世界已经发生了翻天覆地的变化。我们既欣喜地看到全球化和科技进步给人类带来的巨大利益，也不得不忧虑地面对气候变化、病毒肆虐等非传统安全问题导致的严峻挑战。",
			"　　让她尤其感到不安的是，在应对全球挑战的关键时刻，孤立主义、民粹主义的倾向又开始抬头，以邻为壑、落井下石的言论和行为甚至屡屡见于一些大国的舆论引领者。实现不同国家人民之间的互谅互信、互利合作依然还有漫长的道路要走，还需要大家继续长期艰苦的努力。",
			"　继任者是谁？",
			"　　李小林的继任者林松添出生于1960年5月，福建省人。2017年至今，林松添担任中国驻南非大使（副部级）。",
			"　　林松添曾先后代表中国出使利比里亚和马拉维，访问了30个非洲国家，并从2014年6月起担任中国外交部非洲司司长和中非合作论坛中方后续行动委员会秘书长。",
			"　政知圈（微信ID：wepolitics）注意到，一直到今年3月底，林松添还在南非就中国抗击新冠肺炎疫情相关情况接受采访，出席中国政府、中资企业、华侨华人向南非捐赠物资的交接仪式等。",
			"　　2月初，正当中国面临新冠肺炎疫情肆虐时，在非洲进行首访的美国国务卿蓬佩奥却对中国进行了攻击。在亚的斯亚贝巴发表演讲时，蓬佩奥提醒非洲国家“应当警惕（中国）作出的空洞的承诺”，他还告诉非洲国家“与美国的经济伙伴关系才是非洲走向真正解放的正确道路”。",
			"　　2月18日，林松添在驻南非大使馆举行大型记者会，20多家中南和国际媒体参加，当地主流电视台还对记者会进行了全程直播。"
	)
	@Test
	fun tokenizer() {
		JnrJieba.initial("dict.big.gz")
		JnrJieba.tokenizer(art[18],mode = CUT_HMM)
				.apply {
					println(this)
				}
	}
	@Test
	fun tokenizerBenchmark() {
		JnrJieba.initial()
		JnrJieba.tokenizer(art[18])
				.forEach {
					println(it)
				}
		measureTimeMillis {
			(0..1000).forEach {_ ->
				JnrJieba.tokenizer(art[(art.indices).random()]).apply(::println)
			}
		}.apply {
			println("jnr jieba ${this/1000.0} ms/op")
		}

	}
}