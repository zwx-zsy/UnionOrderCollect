package lib

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/yancyzhou/unionsdk/JdunionSdk"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type JdOrderResponse struct {
	Id                 bson.ObjectId `bson:"_id,omitempty"`
	FinishTime         int64         `bson:"finishTime"` //订单完成时间(时间戳，毫秒)
	OrderEmt           int           `bson:"orderEmt"`   //下单设备(1:PC,2:无线)
	OrderId            int64         `bson:"orderId"`    //订单ID
	OrderTime          time.Duration `bson:"orderTime"`  //下单时间(时间戳，毫秒)
	ParentId           int64         `bson:"parentId"`   //父单的订单ID，仅当发生订单拆分时返回， 0：未拆分，有值则表示此订单为子订单
	PayMonth           int           `bson:"payMonth"`   //订单维度预估结算时间（格式：yyyyMMdd），0：未结算，订单的预估结算时间仅供参考。账号未通过资质审核或订单发生售后，会影响订单实际结算时间。
	Plus               int           `bson:"plus"`       //下单用户是否为PLUS会员 0：否，1：是
	PopId              int           `bson:"popId"`      //商家ID
	UserCommissionRate float64       `json:"userCommissionRate"`
	SkuList            []struct {
		//订单包含的商品信息列表
		ActualCosPrice      float64 `bson:"actualCosPrice"`      //实际计算佣金的金额。订单完成后，会将误扣除的运费券金额更正。如订单完成后发生退款，此金额会更新。
		ActualFee           float64 `bson:"actualFee"`           //推客获得的实际佣金（实际计佣金额*佣金比例*最终比例）。如订单完成后发生退款，此金额会更新。
		CommissionRate      float64 `bson:"commissionRate"`      //佣金比例
		EstimateCosPrice    float64 `bson:"estimateCosPrice"`    //预估计佣金额，即用户下单的金额(已扣除优惠券、白条、支付优惠、进口税，未扣除红包和京豆)，有时会误扣除运费券金额，完成结算时会在实际计佣金额中更正。如订单完成前发生退款，此金额不会更新
		EstimateFee         float64 `bson:"estimateFee"`         //推客的预估佣金（预估计佣金额*佣金比例*最终比例），如订单完成前发生退款，此金额不会更新
		FinalRate           float64 `bson:"finalRate"`           //最终比例（分成比例+补贴比例）
		Cid1                int64   `bson:"cid1"`                //一级类目ID
		FrozenSkuNum        int64   `bson:"frozenSkuNum"`        //商品售后中数量
		Pid                 string  `bson:"pid"`                 //联盟子站长身份标识，格式：子站长ID_子站长网站ID_子站长推广位ID
		PositionId          int64   `bson:"positionId"`          //推广位ID,0代表无推广位
		Cid2                int64   `bson:"cid2"`                //二级类目ID
		SiteId              int64   `bson:"siteId"`              //网站ID，0：无网站
		SkuId               int64   `bson:"skuId"`               //商品ID
		SkuNum              int64   `bson:"skuNum"`              //商品数量
		SkuReturnNum        int64   `bson:"skuReturnNum"`        //商品已退货数量
		Cid3                int64   `bson:"cid3"`                //三级类目ID
		UnionAlias          string  `bson:"unionAlias"`          //PID所属母账号平台名称（原第三方服务商来源）
		UnionTag            string  `bson:"unionTag"`            //联盟标签数据（整型的二进制字符串(32位)，目前只返回8位：00000001。数据从右向左进行，每一位为1表示符合联盟的标签特征，第1位：京喜红包，第2位：组合推广订单，第3位：拼购订单，第5位：有效首次购订单（00011XXX表示有效首购，最终奖励活动结算金额会结合订单状态判断，以联盟后台对应活动效果数据报表https://union.jd.com/active为准）。例如：00000001:京喜红包订单，00000010:组合推广订单，00000100:拼购订单，00011000:有效首购，00000111：京喜红包+组合推广+拼购等）
		UnionTrafficGroup   int     `bson:"unionTrafficGroup"`   //渠道组 1：1号店，其他：京东
		ValidCode           int     `bson:"validCode"`           //sku维度的有效码（-1：未知,2.无效-拆单,3.无效-取消,4.无效-京东帮帮主订单,5.无效-账号异常,6.无效-赠品类目不返佣,7.无效-校园订单,8.无效-企业订单,9.无效-团购订单,10.无效-开增值税专用发票订单,11.无效-乡村推广员下单,12.无效-自己推广自己下单,13.无效-违规订单,14.无效-来源与备案网址不符,15.待付款,16.已付款,17.已完成,18.已结算（5.9号不再支持结算状态回写展示））注：自2018/7/13起，自己推广自己下单已经允许返佣，故12无效码仅针对历史数据有效
		SubUnionId          string  `bson:"subUnionId"`          //子联盟ID(需要联系运营开放白名单才能拿到数据)
		TraceType           int     `bson:"traceType"`           //2：同店；3：跨店
		PayMonth            int     `bson:"payMonth"`            //订单行维度预估结算时间（格式：yyyyMMdd） ，0：未结算。订单的预估结算时间仅供参考。账号未通过资质审核或订单发生售后，会影响订单实际结算时间。
		PopId               int64   `bson:"popId"`               //商家ID，订单行维度
		Ext1                string  `bson:"ext1"`                //推客生成推广链接时传入的扩展字段（需要联系运营开放白名单才能拿到数据）。&lt;订单行维度&gt;
		Price               float64 `bson:"price"`               //商品单价
		SkuName             string  `bson:"skuName"`             //商品名称
		SubSideRate         float64 `bson:"subSideRate"`         //分成比例
		SubsidyRate         float64 `bson:"subsidyRate"`         //补贴比例
		GiftCouponKey       string  `json:"giftCouponKey"`       //礼金批次ID：使用礼金的订单会有该值
		GiftCouponOcsAmount float64 `json:"giftCouponOcsAmount"` //礼金分摊金额：使用礼金的订单会有该值
		UnionRole           int     `json:"unionRole"`           //站长角色：1 推客 2 团长 3内容服务商
	}
	UnionId   int64  `bson:"unionId"`   //推客的联盟ID
	Ext1      string `bson:"ext1"`      //推客生成推广链接时传入的扩展字段，订单维度（需要联系运营开放白名单才能拿到数据）
	ValidCode int    `bson:"validCode"` //订单维度的有效码（-1：未知,2.无效-拆单,3.无效-取消,4.无效-京东帮帮主订单,5.无效-账号异常,6.无效-赠品类目不返佣,7.无效-校园订单,8.无效-企业订单,9.无效-团购订单,10.无效-开增值税专用发票订单,11.无效-乡村推广员下单,12.无效-自己推广自己下单,13.无效-违规订单,14.无效-来源与备案网址不符,15.待付款,16.已付款,17.已完成,18.已结算（5.9号不再支持结算状态回写展示））注：自2018/7/13起，自己推广自己下单已经允许返佣，故12无效码仅针对历史数据有效
	//HasMore   bool   `bson:"hasMore"`   //是否还有更多,true：还有数据；false:已查询完毕，没有数据
}

func JdOrderTask() {
	var Newclient CronSession
	Newclient.NewClient()
	CronOrder(Newclient.Csession, 1)
}

func GetOrder(s *mgo.Session, PageNo int) (Pagecurrent int, ispaegs bool) {
	ParamJsons := JdunionSdk.OrderParam{OrderReq: JdunionSdk.OrderReq{}}
	ParamJsons.OrderReq.PageSize = 500
	ParamJsons.OrderReq.PageNo = PageNo
	// m, _ := time.ParseDuration("-1m")
	ParamJsons.OrderReq.Time = "202012270211" //time.Now().Add(m).Format("200601021504")
	ParamJsons.OrderReq.Type = 3
	orders := J.GetOrders(ParamJsons)
	var db *mgo.Database
	if len(orders.Data) > 0 {
		localSession := s.Copy()
		defer localSession.Close()
		db = localSession.DB(ServerConf.DBConf.DatabaseName)
	}
	for _, value := range orders.Data {
		var OneDoc JdOrderResponse
		OneDocBytes, _ := json.Marshal(value)
		json.Unmarshal(OneDocBytes, &OneDoc)
		for k, v := range OneDoc.SkuList {
			if v.SubUnionId != "" {
				splits := strings.Split(v.SubUnionId, "_")
				OneDoc.SkuList[k].SubUnionId = splits[0]
				var UserCommissionRate float64
				if len(splits) < 2 {
					UserCommissionRate = 0.00
				} else {
					UserCommissionRate, _ = strconv.ParseFloat(splits[1], 64)
				}

				OneDoc.UserCommissionRate = UserCommissionRate
				//计算最终返利（*当前会员等级的返利比例）
				OneDoc.SkuList[k].ActualFee = FloatRound(OneDoc.SkuList[k].ActualFee*UserCommissionRate, 2)
			}
		}
		_, err := db.C("JdUnion_Order").Upsert(bson.M{"orderId": OneDoc.OrderId}, OneDoc)
		if err != nil {
			fmt.Println(err)
		}
	}

	return PageNo, orders.HasMore
}

func CronOrder(s *mgo.Session, pages int) {
	Pagecurrent, ispaegs := GetOrder(s, pages)
	if ispaegs {
		CronOrder(s, Pagecurrent)
	}
}
