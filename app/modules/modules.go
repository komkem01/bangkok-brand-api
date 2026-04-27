package modules

import (
	"bangkok-brand/app/modules/address"
	"bangkok-brand/app/modules/addresstype"
	"bangkok-brand/app/modules/adminactionlog"
	"bangkok-brand/app/modules/audit"
	"bangkok-brand/app/modules/bank"
	"bangkok-brand/app/modules/bankaccount"
	"bangkok-brand/app/modules/brand"
	"bangkok-brand/app/modules/cart"
	"bangkok-brand/app/modules/cartitem"
	"bangkok-brand/app/modules/category"
	"bangkok-brand/app/modules/chat"
	"bangkok-brand/app/modules/chatmessage"
	"bangkok-brand/app/modules/chatparticipant"
	"bangkok-brand/app/modules/contacttype"
	"bangkok-brand/app/modules/coupon"
	"bangkok-brand/app/modules/couponusage"
	"bangkok-brand/app/modules/disputecase"
	"bangkok-brand/app/modules/disputemessage"
	"bangkok-brand/app/modules/flashsale"
	"bangkok-brand/app/modules/flashsaleitem"
	"bangkok-brand/app/modules/idempotency"
	"bangkok-brand/app/modules/invoice"
	"bangkok-brand/app/modules/invoiceitem"
	"bangkok-brand/app/modules/kyc"
	"bangkok-brand/app/modules/kycdocument"
	"bangkok-brand/app/modules/kycstatushistory"
	"bangkok-brand/app/modules/logistics"
	"bangkok-brand/app/modules/loyalty"
	"bangkok-brand/app/modules/notification"
	"bangkok-brand/app/modules/order"
	"bangkok-brand/app/modules/orderitem"
	"bangkok-brand/app/modules/ordershipment"
	"bangkok-brand/app/modules/orderstatushistory"
	"bangkok-brand/app/modules/payment"
	"bangkok-brand/app/modules/pointsetting"
	"bangkok-brand/app/modules/productreviewimage"
	"bangkok-brand/app/modules/productvariantstock"
	"bangkok-brand/app/modules/productvariantvalue"
	"bangkok-brand/app/modules/refundtransaction"
	"bangkok-brand/app/modules/returnitem"
	"bangkok-brand/app/modules/returns"
	"bangkok-brand/app/modules/review"
	"bangkok-brand/app/modules/reward"
	"bangkok-brand/app/modules/rewardredemption"
	"bangkok-brand/app/modules/search"
	"bangkok-brand/app/modules/settlement"
	"bangkok-brand/app/modules/settlementitem"
	"bangkok-brand/app/modules/shipmenttrackinghistory"
	"bangkok-brand/app/modules/shipping"
	"bangkok-brand/app/modules/shippingmethod"
	"bangkok-brand/app/modules/shippingraterule"
	"bangkok-brand/app/modules/shop"
	"bangkok-brand/app/modules/shopmember"
	"bangkok-brand/app/modules/shopsetting"
	"bangkok-brand/app/modules/shopshippingmethod"
	"bangkok-brand/app/modules/shopwallettransaction"
	"bangkok-brand/app/modules/variant"
	"bangkok-brand/app/modules/view"
	"bangkok-brand/app/modules/webhook"
	"bangkok-brand/app/modules/wishlist"
	"log/slog"
	"sync"

	"bangkok-brand/app/modules/auth"
	"bangkok-brand/app/modules/contact"
	"bangkok-brand/app/modules/district"
	"bangkok-brand/app/modules/entities"
	"bangkok-brand/app/modules/example"
	"bangkok-brand/app/modules/gender"
	"bangkok-brand/app/modules/member"
	"bangkok-brand/app/modules/memberdevice"
	"bangkok-brand/app/modules/prefix"
	"bangkok-brand/app/modules/product"
	"bangkok-brand/app/modules/productattribute"
	"bangkok-brand/app/modules/productattributevalue"
	"bangkok-brand/app/modules/productimage"
	"bangkok-brand/app/modules/productstock"
	"bangkok-brand/app/modules/province"
	"bangkok-brand/app/modules/sentry"
	"bangkok-brand/app/modules/shippingzonearea"
	"bangkok-brand/app/modules/specs"
	"bangkok-brand/app/modules/storage"
	"bangkok-brand/app/modules/subdistrict"
	"bangkok-brand/app/modules/zipcode"
	"bangkok-brand/internal/config"
	"bangkok-brand/internal/database"
	"bangkok-brand/internal/log"
	"bangkok-brand/internal/otel/collector"
	"bangkok-brand/internal/s3"

	exampletwo "bangkok-brand/app/modules/example-two"

	appConf "bangkok-brand/config"
	// "bangkok-brand/app/modules/kafka"
)

type Modules struct {
	Conf   *config.Module[appConf.Config]
	Specs  *specs.Module
	Log    *log.Module
	OTEL   *collector.Module
	Sentry *sentry.Module
	DB     *database.DatabaseModule
	ENT    *entities.Module
	// Kafka *kafka.Module
	Example                 *example.Module
	Example2                *exampletwo.Module
	Gender                  *gender.Module
	Prefix                  *prefix.Module
	Province                *province.Module
	District                *district.Module
	Subdistrict             *subdistrict.Module
	Zipcode                 *zipcode.Module
	Bank                    *bank.Module
	Storage                 *storage.Module
	Member                  *member.Module
	Auth                    *auth.Module
	Contact                 *contact.Module
	ContactType             *contacttype.Module
	AddressType             *addresstype.Module
	Address                 *address.Module
	MemberDevice            *memberdevice.Module
	BankAccount             *bankaccount.Module
	Brand                   *brand.Module
	Category                *category.Module
	Product                 *product.Module
	ProductAttribute        *productattribute.Module
	ProductAttributeValue   *productattributevalue.Module
	ProductImage            *productimage.Module
	ProductStock            *productstock.Module
	Cart                    *cart.Module
	CartItem                *cartitem.Module
	Order                   *order.Module
	OrderItem               *orderitem.Module
	OrderStatusHistory      *orderstatushistory.Module
	Payment                 *payment.Module
	Coupon                  *coupon.Module
	CouponUsage             *couponusage.Module
	Review                  *review.Module
	Loyalty                 *loyalty.Module
	Chat                    *chat.Module
	ChatMessage             *chatmessage.Module
	ChatParticipant         *chatparticipant.Module
	DisputeMessage          *disputemessage.Module
	FlashSaleItem           *flashsaleitem.Module
	InvoiceItem             *invoiceitem.Module
	KYC                     *kyc.Module
	KYCDocument             *kycdocument.Module
	KYCStatusHistory        *kycstatushistory.Module
	OrderShipment           *ordershipment.Module
	PointSetting            *pointsetting.Module
	ProductReviewImage      *productreviewimage.Module
	ProductVariantStock     *productvariantstock.Module
	ProductVariantValue     *productvariantvalue.Module
	RefundTransaction       *refundtransaction.Module
	ReturnItem              *returnitem.Module
	RewardRedemption        *rewardredemption.Module
	Reward                  *reward.Module
	SettlementItem          *settlementitem.Module
	ShipmentTrackingHistory *shipmenttrackinghistory.Module
	ShopMember              *shopmember.Module
	ShopSetting             *shopsetting.Module
	ShippingZoneArea        *shippingzonearea.Module
	ShippingMethod          *shippingmethod.Module
	ShopShippingMethod      *shopshippingmethod.Module
	ShippingRateRule        *shippingraterule.Module
	ShopWalletTransaction   *shopwallettransaction.Module
	Logistics               *logistics.Module
	Shop                    *shop.Module
	AdminActionLog          *adminactionlog.Module
	Variant                 *variant.Module
	Shipping                *shipping.Module
	Settlement              *settlement.Module
	Returns                 *returns.Module
	DisputeCase             *disputecase.Module
	Invoice                 *invoice.Module
	Notification            *notification.Module
	FlashSale               *flashsale.Module
	Wishlist                *wishlist.Module
	Search                  *search.Module
	View                    *view.Module
	Idempotency             *idempotency.Module
	Webhook                 *webhook.Module
	Audit                   *audit.Module
}

func modulesInit() {
	confMod := config.New(&appConf.App)
	specsMod := specs.New(config.Conf[specs.Config](confMod.Svc))
	conf := confMod.Svc.Config()

	logMod := log.New(config.Conf[log.Option](confMod.Svc))
	otel := collector.New(config.Conf[collector.Config](confMod.Svc))
	log := log.With(slog.String("module", "modules"))

	sentryMod := sentry.New(config.Conf[sentry.Config](confMod.Svc))

	db := database.New(conf.Database.Sql)
	entitiesMod := entities.New(db.Svc.DB())
	exampleMod := example.New(config.Conf[example.Config](confMod.Svc), entitiesMod.Svc)
	exampleMod2 := exampletwo.New(config.Conf[exampletwo.Config](confMod.Svc), entitiesMod.Svc)
	authMod := auth.New(config.Conf[auth.Config](confMod.Svc), entitiesMod.Svc, entitiesMod.Svc, entitiesMod.Svc)
	genderMod := gender.New(config.Conf[gender.Config](confMod.Svc), entitiesMod.Svc)
	prefixMod := prefix.New(config.Conf[prefix.Config](confMod.Svc), entitiesMod.Svc, entitiesMod.Svc)
	provinceMod := province.New(config.Conf[province.Config](confMod.Svc), entitiesMod.Svc)
	districtMod := district.New(config.Conf[district.Config](confMod.Svc), entitiesMod.Svc)
	subdistrictMod := subdistrict.New(config.Conf[subdistrict.Config](confMod.Svc), entitiesMod.Svc)
	zipcodeMod := zipcode.New(config.Conf[zipcode.Config](confMod.Svc), entitiesMod.Svc)
	bankMod := bank.New(config.Conf[bank.Config](confMod.Svc), entitiesMod.Svc)
	s3Mod := s3.New(config.Conf[s3.Config](confMod.Svc).Val)
	storageMod := storage.New(config.Conf[storage.Config](confMod.Svc), entitiesMod.Svc, s3Mod.Svc)
	memberMod := member.New(config.Conf[member.Config](confMod.Svc), entitiesMod.Svc)
	contactMod := contact.New(config.Conf[contact.Config](confMod.Svc), entitiesMod.Svc)
	contactTypeMod := contacttype.New(config.Conf[contacttype.Config](confMod.Svc), entitiesMod.Svc)
	addressTypeMod := addresstype.New(config.Conf[addresstype.Config](confMod.Svc), entitiesMod.Svc)
	addressMod := address.New(config.Conf[address.Config](confMod.Svc), entitiesMod.Svc)
	memberDeviceMod := memberdevice.New(config.Conf[memberdevice.Config](confMod.Svc), entitiesMod.Svc)
	bankAccountMod := bankaccount.New(config.Conf[bankaccount.Config](confMod.Svc), entitiesMod.Svc)
	brandMod := brand.New(config.Conf[brand.Config](confMod.Svc), entitiesMod.Svc)
	categoryMod := category.New(config.Conf[category.Config](confMod.Svc), entitiesMod.Svc)
	productMod := product.New(config.Conf[product.Config](confMod.Svc), entitiesMod.Svc)
	productAttributeMod := productattribute.New(config.Conf[productattribute.Config](confMod.Svc), entitiesMod.Svc)
	productAttributeValueMod := productattributevalue.New(config.Conf[productattributevalue.Config](confMod.Svc), entitiesMod.Svc)
	productImageMod := productimage.New(config.Conf[productimage.Config](confMod.Svc), entitiesMod.Svc)
	productStockMod := productstock.New(config.Conf[productstock.Config](confMod.Svc), entitiesMod.Svc)
	cartMod := cart.New(config.Conf[cart.Config](confMod.Svc), entitiesMod.Svc)
	cartItemMod := cartitem.New(config.Conf[cartitem.Config](confMod.Svc), entitiesMod.Svc)
	orderMod := order.New(config.Conf[order.Config](confMod.Svc), entitiesMod.Svc)
	orderItemMod := orderitem.New(config.Conf[orderitem.Config](confMod.Svc), entitiesMod.Svc)
	orderStatusHistoryMod := orderstatushistory.New(config.Conf[orderstatushistory.Config](confMod.Svc), entitiesMod.Svc)
	paymentMod := payment.New(config.Conf[payment.Config](confMod.Svc), entitiesMod.Svc)
	couponMod := coupon.New(config.Conf[coupon.Config](confMod.Svc), entitiesMod.Svc)
	couponUsageMod := couponusage.New(config.Conf[couponusage.Config](confMod.Svc), entitiesMod.Svc)
	reviewMod := review.New(config.Conf[review.Config](confMod.Svc), entitiesMod.Svc)
	loyaltyMod := loyalty.New(config.Conf[loyalty.Config](confMod.Svc), entitiesMod.Svc)
	chatMod := chat.New(config.Conf[chat.Config](confMod.Svc), entitiesMod.Svc)
	chatMessageMod := chatmessage.New(config.Conf[chatmessage.Config](confMod.Svc), entitiesMod.Svc)
	chatParticipantMod := chatparticipant.New(config.Conf[chatparticipant.Config](confMod.Svc), entitiesMod.Svc)
	disputeMessageMod := disputemessage.New(config.Conf[disputemessage.Config](confMod.Svc), entitiesMod.Svc)
	flashSaleItemMod := flashsaleitem.New(config.Conf[flashsaleitem.Config](confMod.Svc), entitiesMod.Svc)
	invoiceItemMod := invoiceitem.New(config.Conf[invoiceitem.Config](confMod.Svc), entitiesMod.Svc)
	kycMod := kyc.New(config.Conf[kyc.Config](confMod.Svc), entitiesMod.Svc)
	kycDocumentMod := kycdocument.New(config.Conf[kycdocument.Config](confMod.Svc), entitiesMod.Svc)
	kycStatusHistoryMod := kycstatushistory.New(config.Conf[kycstatushistory.Config](confMod.Svc), entitiesMod.Svc)
	orderShipmentMod := ordershipment.New(config.Conf[ordershipment.Config](confMod.Svc), entitiesMod.Svc)
	pointSettingMod := pointsetting.New(config.Conf[pointsetting.Config](confMod.Svc), entitiesMod.Svc)
	productReviewImageMod := productreviewimage.New(config.Conf[productreviewimage.Config](confMod.Svc), entitiesMod.Svc)
	productVariantStockMod := productvariantstock.New(config.Conf[productvariantstock.Config](confMod.Svc), entitiesMod.Svc)
	productVariantValueMod := productvariantvalue.New(config.Conf[productvariantvalue.Config](confMod.Svc), entitiesMod.Svc)
	refundTransactionMod := refundtransaction.New(config.Conf[refundtransaction.Config](confMod.Svc), entitiesMod.Svc)
	returnItemMod := returnitem.New(config.Conf[returnitem.Config](confMod.Svc), entitiesMod.Svc)
	rewardRedemptionMod := rewardredemption.New(config.Conf[rewardredemption.Config](confMod.Svc), entitiesMod.Svc)
	rewardMod := reward.New(config.Conf[reward.Config](confMod.Svc), entitiesMod.Svc)
	settlementItemMod := settlementitem.New(config.Conf[settlementitem.Config](confMod.Svc), entitiesMod.Svc)
	logisticsMod := logistics.New(config.Conf[logistics.Config](confMod.Svc), entitiesMod.Svc)
	shopMod := shop.New(config.Conf[shop.Config](confMod.Svc), entitiesMod.Svc)
	adminActionLogMod := adminactionlog.New(config.Conf[adminactionlog.Config](confMod.Svc), entitiesMod.Svc)
	variantMod := variant.New(config.Conf[variant.Config](confMod.Svc), entitiesMod.Svc)
	shippingMod := shipping.New(config.Conf[shipping.Config](confMod.Svc), entitiesMod.Svc)
	shippingMethodMod := shippingmethod.New(config.Conf[shippingmethod.Config](confMod.Svc), entitiesMod.Svc)
	shipmentTrackingHistoryMod := shipmenttrackinghistory.New(config.Conf[shipmenttrackinghistory.Config](confMod.Svc), entitiesMod.Svc)
	shopMemberMod := shopmember.New(config.Conf[shopmember.Config](confMod.Svc), entitiesMod.Svc)
	shopSettingMod := shopsetting.New(config.Conf[shopsetting.Config](confMod.Svc), entitiesMod.Svc)
	shippingZoneAreaMod := shippingzonearea.New(config.Conf[shippingzonearea.Config](confMod.Svc), entitiesMod.Svc)
	shopShippingMethodMod := shopshippingmethod.New(config.Conf[shopshippingmethod.Config](confMod.Svc), entitiesMod.Svc)
	shippingRateRuleMod := shippingraterule.New(config.Conf[shippingraterule.Config](confMod.Svc), entitiesMod.Svc)
	shopWalletTransactionMod := shopwallettransaction.New(config.Conf[shopwallettransaction.Config](confMod.Svc), entitiesMod.Svc)
	settlementMod := settlement.New(config.Conf[settlement.Config](confMod.Svc), entitiesMod.Svc)
	returnsMod := returns.New(config.Conf[returns.Config](confMod.Svc), entitiesMod.Svc)
	disputeCaseMod := disputecase.New(config.Conf[disputecase.Config](confMod.Svc), entitiesMod.Svc)
	invoiceMod := invoice.New(config.Conf[invoice.Config](confMod.Svc), entitiesMod.Svc)
	notificationMod := notification.New(config.Conf[notification.Config](confMod.Svc), entitiesMod.Svc)
	flashSaleMod := flashsale.New(config.Conf[flashsale.Config](confMod.Svc), entitiesMod.Svc)
	wishlistMod := wishlist.New(config.Conf[wishlist.Config](confMod.Svc), entitiesMod.Svc)
	searchMod := search.New(config.Conf[search.Config](confMod.Svc), entitiesMod.Svc)
	viewMod := view.New(config.Conf[view.Config](confMod.Svc), entitiesMod.Svc)
	idempotencyMod := idempotency.New(config.Conf[idempotency.Config](confMod.Svc), entitiesMod.Svc)
	webhookMod := webhook.New(config.Conf[webhook.Config](confMod.Svc), entitiesMod.Svc)
	auditMod := audit.New(config.Conf[audit.Config](confMod.Svc), entitiesMod.Svc)
	// kafka := kafka.New(&conf.Kafka)
	mod = &Modules{
		Conf:                    confMod,
		Specs:                   specsMod,
		Log:                     logMod,
		OTEL:                    otel,
		Sentry:                  sentryMod,
		DB:                      db,
		ENT:                     entitiesMod,
		Example:                 exampleMod,
		Example2:                exampleMod2,
		Auth:                    authMod,
		Gender:                  genderMod,
		Prefix:                  prefixMod,
		Province:                provinceMod,
		District:                districtMod,
		Subdistrict:             subdistrictMod,
		Zipcode:                 zipcodeMod,
		Bank:                    bankMod,
		Storage:                 storageMod,
		Member:                  memberMod,
		Contact:                 contactMod,
		ContactType:             contactTypeMod,
		AddressType:             addressTypeMod,
		Address:                 addressMod,
		MemberDevice:            memberDeviceMod,
		BankAccount:             bankAccountMod,
		Brand:                   brandMod,
		Category:                categoryMod,
		Product:                 productMod,
		ProductAttribute:        productAttributeMod,
		ProductAttributeValue:   productAttributeValueMod,
		ProductImage:            productImageMod,
		ProductStock:            productStockMod,
		Cart:                    cartMod,
		CartItem:                cartItemMod,
		Order:                   orderMod,
		OrderItem:               orderItemMod,
		OrderStatusHistory:      orderStatusHistoryMod,
		Payment:                 paymentMod,
		Coupon:                  couponMod,
		CouponUsage:             couponUsageMod,
		Review:                  reviewMod,
		Loyalty:                 loyaltyMod,
		Chat:                    chatMod,
		ChatMessage:             chatMessageMod,
		ChatParticipant:         chatParticipantMod,
		DisputeMessage:          disputeMessageMod,
		FlashSaleItem:           flashSaleItemMod,
		InvoiceItem:             invoiceItemMod,
		KYC:                     kycMod,
		KYCDocument:             kycDocumentMod,
		KYCStatusHistory:        kycStatusHistoryMod,
		OrderShipment:           orderShipmentMod,
		PointSetting:            pointSettingMod,
		ProductReviewImage:      productReviewImageMod,
		ProductVariantStock:     productVariantStockMod,
		ProductVariantValue:     productVariantValueMod,
		RefundTransaction:       refundTransactionMod,
		ReturnItem:              returnItemMod,
		RewardRedemption:        rewardRedemptionMod,
		Reward:                  rewardMod,
		SettlementItem:          settlementItemMod,
		Logistics:               logisticsMod,
		Shop:                    shopMod,
		AdminActionLog:          adminActionLogMod,
		Variant:                 variantMod,
		Shipping:                shippingMod,
		ShippingMethod:          shippingMethodMod,
		ShipmentTrackingHistory: shipmentTrackingHistoryMod,
		ShopMember:              shopMemberMod,
		ShopSetting:             shopSettingMod,
		ShippingZoneArea:        shippingZoneAreaMod,
		ShopShippingMethod:      shopShippingMethodMod,
		ShippingRateRule:        shippingRateRuleMod,
		ShopWalletTransaction:   shopWalletTransactionMod,
		Settlement:              settlementMod,
		Returns:                 returnsMod,
		DisputeCase:             disputeCaseMod,
		Invoice:                 invoiceMod,
		Notification:            notificationMod,
		FlashSale:               flashSaleMod,
		Wishlist:                wishlistMod,
		Search:                  searchMod,
		View:                    viewMod,
		Idempotency:             idempotencyMod,
		Webhook:                 webhookMod,
		Audit:                   auditMod,
		// Kafka: kafka,
	}

	log.Infof("all modules initialized")
}

var (
	once sync.Once
	mod  *Modules
)

func Get() *Modules {
	once.Do(modulesInit)

	return mod
}
