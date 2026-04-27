package config

import (
	"bangkok-brand/app/modules/address"
	"bangkok-brand/app/modules/addresstype"
	"bangkok-brand/app/modules/adminactionlog"
	"bangkok-brand/app/modules/audit"
	"bangkok-brand/app/modules/auth"
	"bangkok-brand/app/modules/bank"
	"bangkok-brand/app/modules/bankaccount"
	"bangkok-brand/app/modules/brand"
	"bangkok-brand/app/modules/cart"
	"bangkok-brand/app/modules/cartitem"
	"bangkok-brand/app/modules/category"
	"bangkok-brand/app/modules/chat"
	"bangkok-brand/app/modules/chatmessage"
	"bangkok-brand/app/modules/chatparticipant"
	"bangkok-brand/app/modules/contact"
	"bangkok-brand/app/modules/contacttype"
	"bangkok-brand/app/modules/coupon"
	"bangkok-brand/app/modules/couponusage"
	"bangkok-brand/app/modules/disputecase"
	"bangkok-brand/app/modules/disputemessage"
	"bangkok-brand/app/modules/district"
	"bangkok-brand/app/modules/example"
	exampletwo "bangkok-brand/app/modules/example-two"
	"bangkok-brand/app/modules/flashsale"
	"bangkok-brand/app/modules/flashsaleitem"
	"bangkok-brand/app/modules/gender"
	"bangkok-brand/app/modules/idempotency"
	"bangkok-brand/app/modules/invoice"
	"bangkok-brand/app/modules/invoiceitem"
	"bangkok-brand/app/modules/kyc"
	"bangkok-brand/app/modules/kycdocument"
	"bangkok-brand/app/modules/kycstatushistory"
	"bangkok-brand/app/modules/logistics"
	"bangkok-brand/app/modules/loyalty"
	"bangkok-brand/app/modules/member"
	"bangkok-brand/app/modules/memberdevice"
	"bangkok-brand/app/modules/notification"
	"bangkok-brand/app/modules/order"
	"bangkok-brand/app/modules/orderitem"
	"bangkok-brand/app/modules/ordershipment"
	"bangkok-brand/app/modules/orderstatushistory"
	"bangkok-brand/app/modules/payment"
	"bangkok-brand/app/modules/pointsetting"
	"bangkok-brand/app/modules/prefix"
	"bangkok-brand/app/modules/product"
	"bangkok-brand/app/modules/productattribute"
	"bangkok-brand/app/modules/productattributevalue"
	"bangkok-brand/app/modules/productimage"
	"bangkok-brand/app/modules/productreviewimage"
	"bangkok-brand/app/modules/productstock"
	"bangkok-brand/app/modules/productvariantstock"
	"bangkok-brand/app/modules/productvariantvalue"
	"bangkok-brand/app/modules/province"
	"bangkok-brand/app/modules/refundtransaction"
	"bangkok-brand/app/modules/returnitem"
	"bangkok-brand/app/modules/returns"
	"bangkok-brand/app/modules/review"
	"bangkok-brand/app/modules/reward"
	"bangkok-brand/app/modules/rewardredemption"
	"bangkok-brand/app/modules/search"
	"bangkok-brand/app/modules/sentry"
	"bangkok-brand/app/modules/settlement"
	"bangkok-brand/app/modules/settlementitem"
	"bangkok-brand/app/modules/shipmenttrackinghistory"
	"bangkok-brand/app/modules/shipping"
	"bangkok-brand/app/modules/shippingmethod"
	"bangkok-brand/app/modules/shippingraterule"
	"bangkok-brand/app/modules/shippingzonearea"
	"bangkok-brand/app/modules/shop"
	"bangkok-brand/app/modules/shopmember"
	"bangkok-brand/app/modules/shopsetting"
	"bangkok-brand/app/modules/shopshippingmethod"
	"bangkok-brand/app/modules/shopwallettransaction"
	"bangkok-brand/app/modules/specs"
	"bangkok-brand/app/modules/storage"
	"bangkok-brand/app/modules/subdistrict"
	"bangkok-brand/app/modules/variant"
	"bangkok-brand/app/modules/view"
	"bangkok-brand/app/modules/webhook"
	"bangkok-brand/app/modules/wishlist"
	"bangkok-brand/app/modules/zipcode"
	"bangkok-brand/internal/kafka"
	"bangkok-brand/internal/log"
	"bangkok-brand/internal/otel/collector"
	"bangkok-brand/internal/s3"
)

// Config is a struct that contains all the configuration of the application.
type Config struct {
	Database Database

	AppName     string
	AppKey      string
	Environment string
	Specs       specs.Config
	Debug       bool

	Port           int
	HttpJsonNaming string

	SslCaPath      string
	SslPrivatePath string
	SslCertPath    string

	Otel   collector.Config
	Sentry sentry.Config

	Kafka kafka.Config
	Log   log.Option

	Example example.Config

	ExampleTwo              exampletwo.Config
	Auth                    auth.Config
	Contact                 contact.Config
	ContactType             contacttype.Config
	AddressType             addresstype.Config
	Address                 address.Config
	MemberDevice            memberdevice.Config
	AdminActionLog          adminactionlog.Config
	ChatMessage             chatmessage.Config
	ChatParticipant         chatparticipant.Config
	CouponUsage             couponusage.Config
	DisputeCase             disputecase.Config
	DisputeMessage          disputemessage.Config
	FlashSaleItem           flashsaleitem.Config
	InvoiceItem             invoiceitem.Config
	KYCDocument             kycdocument.Config
	KYCStatusHistory        kycstatushistory.Config
	OrderShipment           ordershipment.Config
	PointSetting            pointsetting.Config
	ProductReviewImage      productreviewimage.Config
	ProductVariantStock     productvariantstock.Config
	ProductVariantValue     productvariantvalue.Config
	Order                   order.Config
	OrderItem               orderitem.Config
	OrderStatusHistory      orderstatushistory.Config
	Payment                 payment.Config
	RefundTransaction       refundtransaction.Config
	ReturnItem              returnitem.Config
	RewardRedemption        rewardredemption.Config
	Reward                  reward.Config
	SettlementItem          settlementitem.Config
	ShipmentTrackingHistory shipmenttrackinghistory.Config
	ShopMember              shopmember.Config
	ShopSetting             shopsetting.Config
	ShippingZoneArea        shippingzonearea.Config
	ShippingMethod          shippingmethod.Config
	ShopShippingMethod      shopshippingmethod.Config
	ShippingRateRule        shippingraterule.Config
	ShopWalletTransaction   shopwallettransaction.Config

	Gender                gender.Config
	Prefix                prefix.Config
	Province              province.Config
	District              district.Config
	Subdistrict           subdistrict.Config
	Zipcode               zipcode.Config
	Storage               storage.Config
	Member                member.Config
	S3                    s3.Config
	Banks                 bank.Config
	BankAccount           bankaccount.Config
	Coupon                coupon.Config
	Review                review.Config
	Loyalty               loyalty.Config
	Chat                  chat.Config
	KYC                   kyc.Config
	Logistics             logistics.Config
	Shop                  shop.Config
	Variant               variant.Config
	Shipping              shipping.Config
	Settlement            settlement.Config
	Returns               returns.Config
	Invoice               invoice.Config
	Notification          notification.Config
	FlashSale             flashsale.Config
	Wishlist              wishlist.Config
	Search                search.Config
	View                  view.Config
	Idempotency           idempotency.Config
	Webhook               webhook.Config
	Audit                 audit.Config
	Brand                 brand.Config
	Category              category.Config
	Product               product.Config
	ProductAttribute      productattribute.Config
	ProductAttributeValue productattributevalue.Config
	ProductImage          productimage.Config
	ProductStock          productstock.Config
	Cart                  cart.Config
	CartItem              cartitem.Config
}

var App = Config{
	Specs: specs.Config{
		Version: "v1",
	},
	Database: database,
	Kafka:    kafkaConf,
	S3: s3.Config{
		Endpoint:        "localhost:9000",
		AccessKeyId:     "minioadmin",
		SecretAccessKey: "minioadmin",
		BucketName:      "default",
		UseSSL:          false,
	},

	AppName: "go_app",
	Port:    8080,
	AppKey:  "secret",
	Debug:   false,

	HttpJsonNaming: "snake_case",

	SslCaPath:      "bangkok-brand/cert/ca.pem",
	SslPrivatePath: "bangkok-brand/cert/server.pem",
	SslCertPath:    "bangkok-brand/cert/server-key.pem",

	Otel: collector.Config{
		CollectorEndpoint: "",
		LogMode:           "noop",
		TraceMode:         "noop",
		MetricMode:        "noop",
		TraceRatio:        0.01,
	},
}
