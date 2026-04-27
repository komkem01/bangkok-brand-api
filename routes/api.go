package routes

import (
	"fmt"
	"net/http"

	"bangkok-brand/app/modules"

	"github.com/gin-gonic/gin"
)

func WarpH(router *gin.RouterGroup, prefix string, handler http.Handler) {
	router.Any(fmt.Sprintf("%s/*w", prefix), gin.WrapH(http.StripPrefix(fmt.Sprintf("%s%s", router.BasePath(), prefix), handler)))
}

func api(r *gin.RouterGroup, mod *modules.Modules) {
	r.GET("/example/:id", mod.Example.Ctl.Get)
	r.GET("/example-http", mod.Example.Ctl.GetHttpReq)
	r.POST("/example", mod.Example.Ctl.Create)
}

func apiSystem(r *gin.RouterGroup, mod *modules.Modules) {
	systems := r.Group("/systems", authRequired(mod))
	{
		genders := systems.Group("/genders")
		{
			genders.GET("", mod.Gender.Ctl.List)
			genders.GET("/:id", mod.Gender.Ctl.Info)
			genders.PATCH("/:id", mod.Gender.Ctl.Update)
			genders.DELETE("/:id", mod.Gender.Ctl.Delete)
		}
		prefixes := systems.Group("/prefixes")
		{
			prefixes.GET("", mod.Prefix.Ctl.List)
			prefixes.GET("/:id", mod.Prefix.Ctl.Info)
			prefixes.PATCH("/:id", mod.Prefix.Ctl.Update)
			prefixes.DELETE("/:id", mod.Prefix.Ctl.Delete)
		}
		provinces := systems.Group("/provinces")
		{
			provinces.GET("", mod.Province.Ctl.List)
			provinces.GET("/:id", mod.Province.Ctl.Info)
			provinces.PATCH("/:id", mod.Province.Ctl.Update)
			provinces.DELETE("/:id", mod.Province.Ctl.Delete)
		}
		districts := systems.Group("/districts")
		{
			districts.GET("", mod.District.Ctl.List)
			districts.GET("/:id", mod.District.Ctl.Info)
			districts.PATCH("/:id", mod.District.Ctl.Update)
			districts.DELETE("/:id", mod.District.Ctl.Delete)
		}
		subdistricts := systems.Group("/subdistricts")
		{
			subdistricts.GET("", mod.Subdistrict.Ctl.List)
			subdistricts.GET("/:id", mod.Subdistrict.Ctl.Info)
			subdistricts.PATCH("/:id", mod.Subdistrict.Ctl.Update)
			subdistricts.DELETE("/:id", mod.Subdistrict.Ctl.Delete)
		}
		zipcodes := systems.Group("/zipcodes")
		{
			zipcodes.GET("", mod.Zipcode.Ctl.List)
			zipcodes.GET("/:id", mod.Zipcode.Ctl.Info)
			zipcodes.PATCH("/:id", mod.Zipcode.Ctl.Update)
			zipcodes.DELETE("/:id", mod.Zipcode.Ctl.Delete)
		}
		banks := systems.Group("/banks")
		{
			banks.POST("", mod.Bank.Ctl.Create)
			banks.GET("", mod.Bank.Ctl.List)
			banks.GET("/:id", mod.Bank.Ctl.Info)
			banks.PATCH("/:id", mod.Bank.Ctl.Update)
			banks.DELETE("/:id", mod.Bank.Ctl.Delete)
		}
		categories := systems.Group("/categories")
		{
			categories.POST("", mod.Category.Ctl.Create)
			categories.GET("", mod.Category.Ctl.List)
			categories.GET("/:id", mod.Category.Ctl.Info)
			categories.PATCH("/:id", mod.Category.Ctl.Update)
			categories.DELETE("/:id", mod.Category.Ctl.Delete)
		}
		brands := systems.Group("/brands")
		{
			brands.POST("", mod.Brand.Ctl.Create)
			brands.GET("", mod.Brand.Ctl.List)
			brands.GET("/:id", mod.Brand.Ctl.Info)
			brands.PATCH("/:id", mod.Brand.Ctl.Update)
			brands.DELETE("/:id", mod.Brand.Ctl.Delete)
		}
		products := systems.Group("/products")
		{
			products.POST("", mod.Product.Ctl.Create)
			products.GET("", mod.Product.Ctl.List)
			products.GET("/:id", mod.Product.Ctl.Info)
			products.PATCH("/:id", mod.Product.Ctl.Update)
			products.DELETE("/:id", mod.Product.Ctl.Delete)
		}
		productImages := systems.Group("/product-images")
		{
			productImages.POST("", mod.ProductImage.Ctl.Create)
			productImages.GET("", mod.ProductImage.Ctl.List)
			productImages.GET("/:id", mod.ProductImage.Ctl.Info)
			productImages.PATCH("/:id", mod.ProductImage.Ctl.Update)
			productImages.DELETE("/:id", mod.ProductImage.Ctl.Delete)
		}
		productAttributes := systems.Group("/product-attributes")
		{
			productAttributes.POST("", mod.ProductAttribute.Ctl.Create)
			productAttributes.GET("", mod.ProductAttribute.Ctl.List)
			productAttributes.GET("/:id", mod.ProductAttribute.Ctl.Info)
			productAttributes.PATCH("/:id", mod.ProductAttribute.Ctl.Update)
			productAttributes.DELETE("/:id", mod.ProductAttribute.Ctl.Delete)
		}
		productAttributeValues := systems.Group("/product-attribute-values")
		{
			productAttributeValues.POST("", mod.ProductAttributeValue.Ctl.Create)
			productAttributeValues.GET("", mod.ProductAttributeValue.Ctl.List)
			productAttributeValues.GET("/:id", mod.ProductAttributeValue.Ctl.Info)
			productAttributeValues.PATCH("/:id", mod.ProductAttributeValue.Ctl.Update)
			productAttributeValues.DELETE("/:id", mod.ProductAttributeValue.Ctl.Delete)
		}
		productStocks := systems.Group("/product-stocks")
		{
			productStocks.POST("", mod.ProductStock.Ctl.Create)
			productStocks.GET("", mod.ProductStock.Ctl.List)
			productStocks.GET("/:id", mod.ProductStock.Ctl.Info)
			productStocks.PATCH("/:id", mod.ProductStock.Ctl.Update)
			productStocks.DELETE("/:id", mod.ProductStock.Ctl.Delete)
		}
		carts := systems.Group("/carts")
		{
			carts.POST("", mod.Cart.Ctl.Create)
			carts.GET("", mod.Cart.Ctl.List)
			carts.GET("/:id", mod.Cart.Ctl.Info)
			carts.PATCH("/:id", mod.Cart.Ctl.Update)
			carts.DELETE("/:id", mod.Cart.Ctl.Delete)
		}
		cartItems := systems.Group("/cart-items")
		{
			cartItems.POST("", mod.CartItem.Ctl.Create)
			cartItems.GET("", mod.CartItem.Ctl.List)
			cartItems.GET("/:id", mod.CartItem.Ctl.Info)
			cartItems.PATCH("/:id", mod.CartItem.Ctl.Update)
			cartItems.DELETE("/:id", mod.CartItem.Ctl.Delete)
		}
		coupons := systems.Group("/coupons")
		{
			coupons.POST("", mod.Coupon.Ctl.Create)
			coupons.GET("", mod.Coupon.Ctl.List)
			coupons.GET("/:id", mod.Coupon.Ctl.Info)
			coupons.PATCH("/:id", mod.Coupon.Ctl.Update)
			coupons.DELETE("/:id", mod.Coupon.Ctl.Delete)
		}
		reviews := systems.Group("/reviews")
		{
			reviews.POST("", mod.Review.Ctl.Create)
			reviews.GET("", mod.Review.Ctl.List)
			reviews.GET("/:id", mod.Review.Ctl.Info)
			reviews.PATCH("/:id", mod.Review.Ctl.Update)
			reviews.DELETE("/:id", mod.Review.Ctl.Delete)
		}
		pointTransactions := systems.Group("/point-transactions")
		{
			pointTransactions.POST("", mod.Loyalty.Ctl.Create)
			pointTransactions.GET("", mod.Loyalty.Ctl.List)
			pointTransactions.GET("/:id", mod.Loyalty.Ctl.Info)
			pointTransactions.PATCH("/:id", mod.Loyalty.Ctl.Update)
			pointTransactions.DELETE("/:id", mod.Loyalty.Ctl.Delete)
		}
		chatRooms := systems.Group("/chat-rooms")
		{
			chatRooms.POST("", mod.Chat.Ctl.Create)
			chatRooms.GET("", mod.Chat.Ctl.List)
			chatRooms.GET("/:id", mod.Chat.Ctl.Info)
			chatRooms.PATCH("/:id", mod.Chat.Ctl.Update)
			chatRooms.DELETE("/:id", mod.Chat.Ctl.Delete)
		}
		kycVerifications := systems.Group("/kyc-verifications")
		{
			kycVerifications.POST("", mod.KYC.Ctl.Create)
			kycVerifications.GET("", mod.KYC.Ctl.List)
			kycVerifications.GET("/:id", mod.KYC.Ctl.Info)
			kycVerifications.PATCH("/:id", mod.KYC.Ctl.Update)
			kycVerifications.DELETE("/:id", mod.KYC.Ctl.Delete)
		}
		logisticsProviders := systems.Group("/logistics-providers")
		{
			logisticsProviders.POST("", mod.Logistics.Ctl.Create)
			logisticsProviders.GET("", mod.Logistics.Ctl.List)
			logisticsProviders.GET("/:id", mod.Logistics.Ctl.Info)
			logisticsProviders.PATCH("/:id", mod.Logistics.Ctl.Update)
			logisticsProviders.DELETE("/:id", mod.Logistics.Ctl.Delete)
		}
		shops := systems.Group("/shops")
		{
			shops.POST("", mod.Shop.Ctl.Create)
			shops.GET("", mod.Shop.Ctl.List)
			shops.GET("/:id", mod.Shop.Ctl.Info)
			shops.PATCH("/:id", mod.Shop.Ctl.Update)
			shops.DELETE("/:id", mod.Shop.Ctl.Delete)
		}
		variants := systems.Group("/variants")
		{
			variants.POST("", mod.Variant.Ctl.Create)
			variants.GET("", mod.Variant.Ctl.List)
			variants.GET("/:id", mod.Variant.Ctl.Info)
			variants.PATCH("/:id", mod.Variant.Ctl.Update)
			variants.DELETE("/:id", mod.Variant.Ctl.Delete)
		}
		shippingZones := systems.Group("/shipping-zones")
		{
			shippingZones.POST("", mod.Shipping.Ctl.Create)
			shippingZones.GET("", mod.Shipping.Ctl.List)
			shippingZones.GET("/:id", mod.Shipping.Ctl.Info)
			shippingZones.PATCH("/:id", mod.Shipping.Ctl.Update)
			shippingZones.DELETE("/:id", mod.Shipping.Ctl.Delete)
		}
		settlements := systems.Group("/settlements")
		{
			settlements.POST("", mod.Settlement.Ctl.Create)
			settlements.GET("", mod.Settlement.Ctl.List)
			settlements.GET("/:id", mod.Settlement.Ctl.Info)
			settlements.PATCH("/:id", mod.Settlement.Ctl.Update)
			settlements.DELETE("/:id", mod.Settlement.Ctl.Delete)
		}
		returnRequests := systems.Group("/return-requests")
		{
			returnRequests.POST("", mod.Returns.Ctl.Create)
			returnRequests.GET("", mod.Returns.Ctl.List)
			returnRequests.GET("/:id", mod.Returns.Ctl.Info)
			returnRequests.PATCH("/:id", mod.Returns.Ctl.Update)
			returnRequests.DELETE("/:id", mod.Returns.Ctl.Delete)
		}
		invoices := systems.Group("/invoices")
		{
			invoices.POST("", mod.Invoice.Ctl.Create)
			invoices.GET("", mod.Invoice.Ctl.List)
			invoices.GET("/:id", mod.Invoice.Ctl.Info)
			invoices.PATCH("/:id", mod.Invoice.Ctl.Update)
			invoices.DELETE("/:id", mod.Invoice.Ctl.Delete)
		}
		notifications := systems.Group("/notifications")
		{
			notifications.POST("", mod.Notification.Ctl.Create)
			notifications.GET("", mod.Notification.Ctl.List)
			notifications.GET("/:id", mod.Notification.Ctl.Info)
			notifications.PATCH("/:id", mod.Notification.Ctl.Update)
			notifications.DELETE("/:id", mod.Notification.Ctl.Delete)
		}
		flashSaleEvents := systems.Group("/flash-sale-events")
		{
			flashSaleEvents.POST("", mod.FlashSale.Ctl.Create)
			flashSaleEvents.GET("", mod.FlashSale.Ctl.List)
			flashSaleEvents.GET("/:id", mod.FlashSale.Ctl.Info)
			flashSaleEvents.PATCH("/:id", mod.FlashSale.Ctl.Update)
			flashSaleEvents.DELETE("/:id", mod.FlashSale.Ctl.Delete)
		}
		wishlists := systems.Group("/wishlists")
		{
			wishlists.POST("", mod.Wishlist.Ctl.Create)
			wishlists.GET("", mod.Wishlist.Ctl.List)
			wishlists.GET("/:id", mod.Wishlist.Ctl.Info)
			wishlists.PATCH("/:id", mod.Wishlist.Ctl.Update)
			wishlists.DELETE("/:id", mod.Wishlist.Ctl.Delete)
		}
		searchHistories := systems.Group("/search-histories")
		{
			searchHistories.POST("", mod.Search.Ctl.Create)
			searchHistories.GET("", mod.Search.Ctl.List)
			searchHistories.GET("/:id", mod.Search.Ctl.Info)
			searchHistories.PATCH("/:id", mod.Search.Ctl.Update)
			searchHistories.DELETE("/:id", mod.Search.Ctl.Delete)
		}
		productViews := systems.Group("/product-views")
		{
			productViews.POST("", mod.View.Ctl.Create)
			productViews.GET("", mod.View.Ctl.List)
			productViews.GET("/:id", mod.View.Ctl.Info)
			productViews.PATCH("/:id", mod.View.Ctl.Update)
			productViews.DELETE("/:id", mod.View.Ctl.Delete)
		}
		idempotencyKeys := systems.Group("/idempotency-keys")
		{
			idempotencyKeys.POST("", mod.Idempotency.Ctl.Create)
			idempotencyKeys.GET("", mod.Idempotency.Ctl.List)
			idempotencyKeys.GET("/:id", mod.Idempotency.Ctl.Info)
			idempotencyKeys.PATCH("/:id", mod.Idempotency.Ctl.Update)
			idempotencyKeys.DELETE("/:id", mod.Idempotency.Ctl.Delete)
		}
		webhookEvents := systems.Group("/webhook-events")
		{
			webhookEvents.POST("", mod.Webhook.Ctl.Create)
			webhookEvents.GET("", mod.Webhook.Ctl.List)
			webhookEvents.GET("/:id", mod.Webhook.Ctl.Info)
			webhookEvents.PATCH("/:id", mod.Webhook.Ctl.Update)
			webhookEvents.DELETE("/:id", mod.Webhook.Ctl.Delete)
		}
		auditLogs := systems.Group("/audit-logs")
		{
			auditLogs.POST("", mod.Audit.Ctl.Create)
			auditLogs.GET("", mod.Audit.Ctl.List)
			auditLogs.GET("/:id", mod.Audit.Ctl.Info)
			auditLogs.PATCH("/:id", mod.Audit.Ctl.Update)
			auditLogs.DELETE("/:id", mod.Audit.Ctl.Delete)
		}
	}
}

func apiStorage(r *gin.RouterGroup, mod *modules.Modules) {
	storages := r.Group("/storages", authRequired(mod))
	{
		storages.POST("/upload", mod.Storage.Ctl.Upload)
		storages.GET("", mod.Storage.Ctl.List)
		storages.GET("/:id", mod.Storage.Ctl.Info)
		storages.GET("/:id/presign", mod.Storage.Ctl.Presign)
		storages.DELETE("/:id", mod.Storage.Ctl.Delete)
	}
}

func apiMember(r *gin.RouterGroup, mod *modules.Modules) {
	members := r.Group("/members", authRequired(mod))
	{
		members.GET("/me", mod.Member.Ctl.InfoMe)
		members.GET("", mod.Member.Ctl.List)
		members.GET("/:id", mod.Member.Ctl.Info)
		members.PATCH("/:id", mod.Member.Ctl.Update)
		members.DELETE("/:id", mod.Member.Ctl.Delete)
	}
	memberContacts := members.Group("/member-contacts")
	{
		memberContacts.POST("", mod.Contact.Ctl.Create)
		memberContacts.GET("", mod.Contact.Ctl.List)
		memberContacts.GET("/:id", mod.Contact.Ctl.Info)
		memberContacts.PATCH("/:id", mod.Contact.Ctl.Update)
		memberContacts.DELETE("/:id", mod.Contact.Ctl.Delete)
	}
	memberAddresses := members.Group("/member-addresses")
	{
		memberAddresses.POST("", mod.Address.Ctl.Create)
		memberAddresses.GET("", mod.Address.Ctl.List)
		memberAddresses.GET("/:id", mod.Address.Ctl.Info)
		memberAddresses.PATCH("/:id", mod.Address.Ctl.Update)
		memberAddresses.DELETE("/:id", mod.Address.Ctl.Delete)
	}
	memberBankAccounts := members.Group("/member-bank-accounts")
	{
		memberBankAccounts.POST("", mod.BankAccount.Ctl.Create)
		memberBankAccounts.GET("", mod.BankAccount.Ctl.List)
		memberBankAccounts.GET("/:id", mod.BankAccount.Ctl.Info)
		memberBankAccounts.PATCH("/:id", mod.BankAccount.Ctl.Update)
		memberBankAccounts.DELETE("/:id", mod.BankAccount.Ctl.Delete)
	}
}

func apiPublic(r *gin.RouterGroup, mod *modules.Modules) {
	public := r.Group("/public")
	{
		auth := public.Group("/auth")
		{
			auth.POST("/google", mod.Auth.Ctl.GoogleLogin)
			auth.GET("/google/callback", mod.Auth.Ctl.GoogleCallback)
			auth.POST("/register", mod.Auth.Ctl.Register)
			auth.POST("/login", mod.Auth.Ctl.Login)
			auth.POST("/refresh", mod.Auth.Ctl.RefreshToken)
			auth.POST("/logout", mod.Auth.Ctl.Logout)
		}
		products := public.Group("/products")
		{
			products.GET("", mod.Product.Ctl.List)
			products.GET("/:id", mod.Product.Ctl.Info)
		}
		categories := public.Group("/categories")
		{
			categories.GET("", mod.Category.Ctl.List)
			categories.GET("/:id", mod.Category.Ctl.Info)
		}
		brands := public.Group("/brands")
		{
			brands.GET("", mod.Brand.Ctl.List)
			brands.GET("/:id", mod.Brand.Ctl.Info)
		}
		shops := public.Group("/shops")
		{
			shops.GET("", mod.Shop.Ctl.List)
			shops.GET("/:id", mod.Shop.Ctl.Info)
		}
	}
}
