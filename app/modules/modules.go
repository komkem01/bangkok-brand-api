package modules

import (
	"bangkok-brand/app/modules/address"
	"bangkok-brand/app/modules/audit"
	"bangkok-brand/app/modules/bank"
	"bangkok-brand/app/modules/bankaccount"
	"bangkok-brand/app/modules/brand"
	"bangkok-brand/app/modules/cart"
	"bangkok-brand/app/modules/cartitem"
	"bangkok-brand/app/modules/category"
	"bangkok-brand/app/modules/chat"
	"bangkok-brand/app/modules/coupon"
	"bangkok-brand/app/modules/flashsale"
	"bangkok-brand/app/modules/idempotency"
	"bangkok-brand/app/modules/invoice"
	"bangkok-brand/app/modules/kyc"
	"bangkok-brand/app/modules/logistics"
	"bangkok-brand/app/modules/loyalty"
	"bangkok-brand/app/modules/notification"
	"bangkok-brand/app/modules/returns"
	"bangkok-brand/app/modules/review"
	"bangkok-brand/app/modules/search"
	"bangkok-brand/app/modules/settlement"
	"bangkok-brand/app/modules/shipping"
	"bangkok-brand/app/modules/shop"
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
	"bangkok-brand/app/modules/prefix"
	"bangkok-brand/app/modules/product"
	"bangkok-brand/app/modules/productattribute"
	"bangkok-brand/app/modules/productattributevalue"
	"bangkok-brand/app/modules/productimage"
	"bangkok-brand/app/modules/productstock"
	"bangkok-brand/app/modules/province"
	"bangkok-brand/app/modules/sentry"
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
	Example               *example.Module
	Example2              *exampletwo.Module
	Gender                *gender.Module
	Prefix                *prefix.Module
	Province              *province.Module
	District              *district.Module
	Subdistrict           *subdistrict.Module
	Zipcode               *zipcode.Module
	Bank                  *bank.Module
	Storage               *storage.Module
	Member                *member.Module
	Auth                  *auth.Module
	Contact               *contact.Module
	Address               *address.Module
	BankAccount           *bankaccount.Module
	Brand                 *brand.Module
	Category              *category.Module
	Product               *product.Module
	ProductAttribute      *productattribute.Module
	ProductAttributeValue *productattributevalue.Module
	ProductImage          *productimage.Module
	ProductStock          *productstock.Module
	Cart                  *cart.Module
	CartItem              *cartitem.Module
	Coupon                *coupon.Module
	Review                *review.Module
	Loyalty               *loyalty.Module
	Chat                  *chat.Module
	KYC                   *kyc.Module
	Logistics             *logistics.Module
	Shop                  *shop.Module
	Variant               *variant.Module
	Shipping              *shipping.Module
	Settlement            *settlement.Module
	Returns               *returns.Module
	Invoice               *invoice.Module
	Notification          *notification.Module
	FlashSale             *flashsale.Module
	Wishlist              *wishlist.Module
	Search                *search.Module
	View                  *view.Module
	Idempotency           *idempotency.Module
	Webhook               *webhook.Module
	Audit                 *audit.Module
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
	authMod := auth.New(config.Conf[auth.Config](confMod.Svc), entitiesMod.Svc)
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
	addressMod := address.New(config.Conf[address.Config](confMod.Svc), entitiesMod.Svc)
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
	couponMod := coupon.New(config.Conf[coupon.Config](confMod.Svc), entitiesMod.Svc)
	reviewMod := review.New(config.Conf[review.Config](confMod.Svc), entitiesMod.Svc)
	loyaltyMod := loyalty.New(config.Conf[loyalty.Config](confMod.Svc), entitiesMod.Svc)
	chatMod := chat.New(config.Conf[chat.Config](confMod.Svc), entitiesMod.Svc)
	kycMod := kyc.New(config.Conf[kyc.Config](confMod.Svc), entitiesMod.Svc)
	logisticsMod := logistics.New(config.Conf[logistics.Config](confMod.Svc), entitiesMod.Svc)
	shopMod := shop.New(config.Conf[shop.Config](confMod.Svc), entitiesMod.Svc)
	variantMod := variant.New(config.Conf[variant.Config](confMod.Svc), entitiesMod.Svc)
	shippingMod := shipping.New(config.Conf[shipping.Config](confMod.Svc), entitiesMod.Svc)
	settlementMod := settlement.New(config.Conf[settlement.Config](confMod.Svc), entitiesMod.Svc)
	returnsMod := returns.New(config.Conf[returns.Config](confMod.Svc), entitiesMod.Svc)
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
		Conf:                  confMod,
		Specs:                 specsMod,
		Log:                   logMod,
		OTEL:                  otel,
		Sentry:                sentryMod,
		DB:                    db,
		ENT:                   entitiesMod,
		Example:               exampleMod,
		Example2:              exampleMod2,
		Auth:                  authMod,
		Gender:                genderMod,
		Prefix:                prefixMod,
		Province:              provinceMod,
		District:              districtMod,
		Subdistrict:           subdistrictMod,
		Zipcode:               zipcodeMod,
		Bank:                  bankMod,
		Storage:               storageMod,
		Member:                memberMod,
		Contact:               contactMod,
		Address:               addressMod,
		BankAccount:           bankAccountMod,
		Brand:                 brandMod,
		Category:              categoryMod,
		Product:               productMod,
		ProductAttribute:      productAttributeMod,
		ProductAttributeValue: productAttributeValueMod,
		ProductImage:          productImageMod,
		ProductStock:          productStockMod,
		Cart:                  cartMod,
		CartItem:              cartItemMod,
		Coupon:                couponMod,
		Review:                reviewMod,
		Loyalty:               loyaltyMod,
		Chat:                  chatMod,
		KYC:                   kycMod,
		Logistics:             logisticsMod,
		Shop:                  shopMod,
		Variant:               variantMod,
		Shipping:              shippingMod,
		Settlement:            settlementMod,
		Returns:               returnsMod,
		Invoice:               invoiceMod,
		Notification:          notificationMod,
		FlashSale:             flashSaleMod,
		Wishlist:              wishlistMod,
		Search:                searchMod,
		View:                  viewMod,
		Idempotency:           idempotencyMod,
		Webhook:               webhookMod,
		Audit:                 auditMod,
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
