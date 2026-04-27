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
		contactTypes := systems.Group("/contact-types")
		{
			contactTypes.POST("", mod.ContactType.Ctl.Create)
			contactTypes.GET("", mod.ContactType.Ctl.List)
			contactTypes.GET("/:id", mod.ContactType.Ctl.Info)
			contactTypes.PATCH("/:id", mod.ContactType.Ctl.Update)
			contactTypes.DELETE("/:id", mod.ContactType.Ctl.Delete)
		}
		addressTypes := systems.Group("/address-types")
		{
			addressTypes.POST("", mod.AddressType.Ctl.Create)
			addressTypes.GET("", mod.AddressType.Ctl.List)
			addressTypes.GET("/:id", mod.AddressType.Ctl.Info)
			addressTypes.PATCH("/:id", mod.AddressType.Ctl.Update)
			addressTypes.DELETE("/:id", mod.AddressType.Ctl.Delete)
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
		orders := systems.Group("/orders")
		{
			orders.POST("", mod.Order.Ctl.Create)
			orders.GET("", mod.Order.Ctl.List)
			orders.GET("/:id", mod.Order.Ctl.Info)
			orders.PATCH("/:id", mod.Order.Ctl.Update)
			orders.DELETE("/:id", mod.Order.Ctl.Delete)
		}
		orderItems := systems.Group("/order-items")
		{
			orderItems.POST("", mod.OrderItem.Ctl.Create)
			orderItems.GET("", mod.OrderItem.Ctl.List)
			orderItems.GET("/:id", mod.OrderItem.Ctl.Info)
			orderItems.PATCH("/:id", mod.OrderItem.Ctl.Update)
			orderItems.DELETE("/:id", mod.OrderItem.Ctl.Delete)
		}
		orderStatusHistories := systems.Group("/order-status-histories")
		{
			orderStatusHistories.POST("", mod.OrderStatusHistory.Ctl.Create)
			orderStatusHistories.GET("", mod.OrderStatusHistory.Ctl.List)
			orderStatusHistories.GET("/:id", mod.OrderStatusHistory.Ctl.Info)
			orderStatusHistories.PATCH("/:id", mod.OrderStatusHistory.Ctl.Update)
			orderStatusHistories.DELETE("/:id", mod.OrderStatusHistory.Ctl.Delete)
		}
		payments := systems.Group("/payments")
		{
			payments.POST("", mod.Payment.Ctl.Create)
			payments.GET("", mod.Payment.Ctl.List)
			payments.GET("/:id", mod.Payment.Ctl.Info)
			payments.PATCH("/:id", mod.Payment.Ctl.Update)
			payments.DELETE("/:id", mod.Payment.Ctl.Delete)
		}
		coupons := systems.Group("/coupons")
		{
			coupons.POST("", mod.Coupon.Ctl.Create)
			coupons.GET("", mod.Coupon.Ctl.List)
			coupons.GET("/:id", mod.Coupon.Ctl.Info)
			coupons.PATCH("/:id", mod.Coupon.Ctl.Update)
			coupons.DELETE("/:id", mod.Coupon.Ctl.Delete)
		}
		couponUsages := systems.Group("/coupon-usages")
		{
			couponUsages.POST("", mod.CouponUsage.Ctl.Create)
			couponUsages.GET("", mod.CouponUsage.Ctl.List)
			couponUsages.GET("/:id", mod.CouponUsage.Ctl.Info)
			couponUsages.PATCH("/:id", mod.CouponUsage.Ctl.Update)
			couponUsages.DELETE("/:id", mod.CouponUsage.Ctl.Delete)
		}
		reviews := systems.Group("/reviews")
		{
			reviews.POST("", mod.Review.Ctl.Create)
			reviews.GET("", mod.Review.Ctl.List)
			reviews.GET("/:id", mod.Review.Ctl.Info)
			reviews.PATCH("/:id", mod.Review.Ctl.Update)
			reviews.DELETE("/:id", mod.Review.Ctl.Delete)
		}
		productReviewImages := systems.Group("/product-review-images")
		{
			productReviewImages.POST("", mod.ProductReviewImage.Ctl.Create)
			productReviewImages.GET("", mod.ProductReviewImage.Ctl.List)
			productReviewImages.GET("/:id", mod.ProductReviewImage.Ctl.Info)
			productReviewImages.PATCH("/:id", mod.ProductReviewImage.Ctl.Update)
			productReviewImages.DELETE("/:id", mod.ProductReviewImage.Ctl.Delete)
		}
		pointTransactions := systems.Group("/point-transactions")
		{
			pointTransactions.POST("", mod.Loyalty.Ctl.Create)
			pointTransactions.GET("", mod.Loyalty.Ctl.List)
			pointTransactions.GET("/:id", mod.Loyalty.Ctl.Info)
			pointTransactions.PATCH("/:id", mod.Loyalty.Ctl.Update)
			pointTransactions.DELETE("/:id", mod.Loyalty.Ctl.Delete)
		}
		pointSettings := systems.Group("/point-settings")
		{
			pointSettings.POST("", mod.PointSetting.Ctl.Create)
			pointSettings.GET("", mod.PointSetting.Ctl.List)
			pointSettings.GET("/:id", mod.PointSetting.Ctl.Info)
			pointSettings.PATCH("/:id", mod.PointSetting.Ctl.Update)
			pointSettings.DELETE("/:id", mod.PointSetting.Ctl.Delete)
		}
		chatRooms := systems.Group("/chat-rooms")
		{
			chatRooms.POST("", mod.Chat.Ctl.Create)
			chatRooms.GET("", mod.Chat.Ctl.List)
			chatRooms.GET("/:id", mod.Chat.Ctl.Info)
			chatRooms.PATCH("/:id", mod.Chat.Ctl.Update)
			chatRooms.DELETE("/:id", mod.Chat.Ctl.Delete)
		}
		chatParticipants := systems.Group("/chat-participants")
		{
			chatParticipants.POST("", mod.ChatParticipant.Ctl.Create)
			chatParticipants.GET("", mod.ChatParticipant.Ctl.List)
			chatParticipants.GET("/:id", mod.ChatParticipant.Ctl.Info)
			chatParticipants.PATCH("/:id", mod.ChatParticipant.Ctl.Update)
			chatParticipants.DELETE("/:id", mod.ChatParticipant.Ctl.Delete)
		}
		chatMessages := systems.Group("/chat-messages")
		{
			chatMessages.POST("", mod.ChatMessage.Ctl.Create)
			chatMessages.GET("", mod.ChatMessage.Ctl.List)
			chatMessages.GET("/:id", mod.ChatMessage.Ctl.Info)
			chatMessages.PATCH("/:id", mod.ChatMessage.Ctl.Update)
			chatMessages.DELETE("/:id", mod.ChatMessage.Ctl.Delete)
		}
		disputeMessages := systems.Group("/dispute-messages")
		{
			disputeMessages.POST("", mod.DisputeMessage.Ctl.Create)
			disputeMessages.GET("", mod.DisputeMessage.Ctl.List)
			disputeMessages.GET("/:id", mod.DisputeMessage.Ctl.Info)
			disputeMessages.PATCH("/:id", mod.DisputeMessage.Ctl.Update)
			disputeMessages.DELETE("/:id", mod.DisputeMessage.Ctl.Delete)
		}
		kycVerifications := systems.Group("/kyc-verifications")
		{
			kycVerifications.POST("", mod.KYC.Ctl.Create)
			kycVerifications.GET("", mod.KYC.Ctl.List)
			kycVerifications.GET("/:id", mod.KYC.Ctl.Info)
			kycVerifications.PATCH("/:id", mod.KYC.Ctl.Update)
			kycVerifications.DELETE("/:id", mod.KYC.Ctl.Delete)
		}
		kycDocuments := systems.Group("/kyc-documents")
		{
			kycDocuments.POST("", mod.KYCDocument.Ctl.Create)
			kycDocuments.GET("", mod.KYCDocument.Ctl.List)
			kycDocuments.GET("/:id", mod.KYCDocument.Ctl.Info)
			kycDocuments.PATCH("/:id", mod.KYCDocument.Ctl.Update)
			kycDocuments.DELETE("/:id", mod.KYCDocument.Ctl.Delete)
		}
		kycStatusHistories := systems.Group("/kyc-status-histories")
		{
			kycStatusHistories.POST("", mod.KYCStatusHistory.Ctl.Create)
			kycStatusHistories.GET("", mod.KYCStatusHistory.Ctl.List)
			kycStatusHistories.GET("/:id", mod.KYCStatusHistory.Ctl.Info)
			kycStatusHistories.PATCH("/:id", mod.KYCStatusHistory.Ctl.Update)
			kycStatusHistories.DELETE("/:id", mod.KYCStatusHistory.Ctl.Delete)
		}
		logisticsProviders := systems.Group("/logistics-providers")
		{
			logisticsProviders.POST("", mod.Logistics.Ctl.Create)
			logisticsProviders.GET("", mod.Logistics.Ctl.List)
			logisticsProviders.GET("/:id", mod.Logistics.Ctl.Info)
			logisticsProviders.PATCH("/:id", mod.Logistics.Ctl.Update)
			logisticsProviders.DELETE("/:id", mod.Logistics.Ctl.Delete)
		}
		shippingMethods := systems.Group("/shipping-methods")
		{
			shippingMethods.POST("", mod.ShippingMethod.Ctl.Create)
			shippingMethods.GET("", mod.ShippingMethod.Ctl.List)
			shippingMethods.GET("/:id", mod.ShippingMethod.Ctl.Info)
			shippingMethods.PATCH("/:id", mod.ShippingMethod.Ctl.Update)
			shippingMethods.DELETE("/:id", mod.ShippingMethod.Ctl.Delete)
		}
		shops := systems.Group("/shops")
		{
			shops.POST("", mod.Shop.Ctl.Create)
			shops.GET("", mod.Shop.Ctl.List)
			shops.GET("/:id", mod.Shop.Ctl.Info)
			shops.PATCH("/:id", mod.Shop.Ctl.Update)
			shops.DELETE("/:id", mod.Shop.Ctl.Delete)
		}
		shopMembers := systems.Group("/shop-members")
		{
			shopMembers.POST("", mod.ShopMember.Ctl.Create)
			shopMembers.GET("", mod.ShopMember.Ctl.List)
			shopMembers.GET("/:id", mod.ShopMember.Ctl.Info)
			shopMembers.PATCH("/:id", mod.ShopMember.Ctl.Update)
			shopMembers.DELETE("/:id", mod.ShopMember.Ctl.Delete)
		}
		shopSettings := systems.Group("/shop-settings")
		{
			shopSettings.POST("", mod.ShopSetting.Ctl.Create)
			shopSettings.GET("", mod.ShopSetting.Ctl.List)
			shopSettings.GET("/:id", mod.ShopSetting.Ctl.Info)
			shopSettings.PATCH("/:id", mod.ShopSetting.Ctl.Update)
			shopSettings.DELETE("/:id", mod.ShopSetting.Ctl.Delete)
		}
		adminActionLogs := systems.Group("/admin-action-logs")
		{
			adminActionLogs.POST("", mod.AdminActionLog.Ctl.Create)
			adminActionLogs.GET("", mod.AdminActionLog.Ctl.List)
			adminActionLogs.GET("/:id", mod.AdminActionLog.Ctl.Info)
			adminActionLogs.PATCH("/:id", mod.AdminActionLog.Ctl.Update)
			adminActionLogs.DELETE("/:id", mod.AdminActionLog.Ctl.Delete)
		}
		variants := systems.Group("/variants")
		{
			variants.POST("", mod.Variant.Ctl.Create)
			variants.GET("", mod.Variant.Ctl.List)
			variants.GET("/:id", mod.Variant.Ctl.Info)
			variants.PATCH("/:id", mod.Variant.Ctl.Update)
			variants.DELETE("/:id", mod.Variant.Ctl.Delete)
		}
		productVariantValues := systems.Group("/product-variant-values")
		{
			productVariantValues.POST("", mod.ProductVariantValue.Ctl.Create)
			productVariantValues.GET("", mod.ProductVariantValue.Ctl.List)
			productVariantValues.GET("/:id", mod.ProductVariantValue.Ctl.Info)
			productVariantValues.PATCH("/:id", mod.ProductVariantValue.Ctl.Update)
			productVariantValues.DELETE("/:id", mod.ProductVariantValue.Ctl.Delete)
		}
		productVariantStocks := systems.Group("/product-variant-stocks")
		{
			productVariantStocks.POST("", mod.ProductVariantStock.Ctl.Create)
			productVariantStocks.GET("", mod.ProductVariantStock.Ctl.List)
			productVariantStocks.GET("/:id", mod.ProductVariantStock.Ctl.Info)
			productVariantStocks.PATCH("/:id", mod.ProductVariantStock.Ctl.Update)
			productVariantStocks.DELETE("/:id", mod.ProductVariantStock.Ctl.Delete)
		}
		disputeCases := systems.Group("/dispute-cases")
		{
			disputeCases.POST("", mod.DisputeCase.Ctl.Create)
			disputeCases.GET("", mod.DisputeCase.Ctl.List)
			disputeCases.GET("/:id", mod.DisputeCase.Ctl.Info)
			disputeCases.PATCH("/:id", mod.DisputeCase.Ctl.Update)
			disputeCases.DELETE("/:id", mod.DisputeCase.Ctl.Delete)
		}
		shippingZones := systems.Group("/shipping-zones")
		{
			shippingZones.POST("", mod.Shipping.Ctl.Create)
			shippingZones.GET("", mod.Shipping.Ctl.List)
			shippingZones.GET("/:id", mod.Shipping.Ctl.Info)
			shippingZones.PATCH("/:id", mod.Shipping.Ctl.Update)
			shippingZones.DELETE("/:id", mod.Shipping.Ctl.Delete)
		}
		shippingZoneAreas := systems.Group("/shipping-zone-areas")
		{
			shippingZoneAreas.POST("", mod.ShippingZoneArea.Ctl.Create)
			shippingZoneAreas.GET("", mod.ShippingZoneArea.Ctl.List)
			shippingZoneAreas.GET("/:id", mod.ShippingZoneArea.Ctl.Info)
			shippingZoneAreas.PATCH("/:id", mod.ShippingZoneArea.Ctl.Update)
			shippingZoneAreas.DELETE("/:id", mod.ShippingZoneArea.Ctl.Delete)
		}
		shopShippingMethods := systems.Group("/shop-shipping-methods")
		{
			shopShippingMethods.POST("", mod.ShopShippingMethod.Ctl.Create)
			shopShippingMethods.GET("", mod.ShopShippingMethod.Ctl.List)
			shopShippingMethods.GET("/:id", mod.ShopShippingMethod.Ctl.Info)
			shopShippingMethods.PATCH("/:id", mod.ShopShippingMethod.Ctl.Update)
			shopShippingMethods.DELETE("/:id", mod.ShopShippingMethod.Ctl.Delete)
		}
		shippingRateRules := systems.Group("/shipping-rate-rules")
		{
			shippingRateRules.POST("", mod.ShippingRateRule.Ctl.Create)
			shippingRateRules.GET("", mod.ShippingRateRule.Ctl.List)
			shippingRateRules.GET("/:id", mod.ShippingRateRule.Ctl.Info)
			shippingRateRules.PATCH("/:id", mod.ShippingRateRule.Ctl.Update)
			shippingRateRules.DELETE("/:id", mod.ShippingRateRule.Ctl.Delete)
		}
		orderShipments := systems.Group("/order-shipments")
		{
			orderShipments.POST("", mod.OrderShipment.Ctl.Create)
			orderShipments.GET("", mod.OrderShipment.Ctl.List)
			orderShipments.GET("/:id", mod.OrderShipment.Ctl.Info)
			orderShipments.PATCH("/:id", mod.OrderShipment.Ctl.Update)
			orderShipments.DELETE("/:id", mod.OrderShipment.Ctl.Delete)
		}
		shipmentTrackingHistories := systems.Group("/shipment-tracking-histories")
		{
			shipmentTrackingHistories.POST("", mod.ShipmentTrackingHistory.Ctl.Create)
			shipmentTrackingHistories.GET("", mod.ShipmentTrackingHistory.Ctl.List)
			shipmentTrackingHistories.GET("/:id", mod.ShipmentTrackingHistory.Ctl.Info)
			shipmentTrackingHistories.PATCH("/:id", mod.ShipmentTrackingHistory.Ctl.Update)
			shipmentTrackingHistories.DELETE("/:id", mod.ShipmentTrackingHistory.Ctl.Delete)
		}
		settlements := systems.Group("/settlements")
		{
			settlements.POST("", mod.Settlement.Ctl.Create)
			settlements.GET("", mod.Settlement.Ctl.List)
			settlements.GET("/:id", mod.Settlement.Ctl.Info)
			settlements.PATCH("/:id", mod.Settlement.Ctl.Update)
			settlements.DELETE("/:id", mod.Settlement.Ctl.Delete)
		}
		shopWalletTransactions := systems.Group("/shop-wallet-transactions")
		{
			shopWalletTransactions.POST("", mod.ShopWalletTransaction.Ctl.Create)
			shopWalletTransactions.GET("", mod.ShopWalletTransaction.Ctl.List)
			shopWalletTransactions.GET("/:id", mod.ShopWalletTransaction.Ctl.Info)
			shopWalletTransactions.PATCH("/:id", mod.ShopWalletTransaction.Ctl.Update)
			shopWalletTransactions.DELETE("/:id", mod.ShopWalletTransaction.Ctl.Delete)
		}
		settlementItems := systems.Group("/settlement-items")
		{
			settlementItems.POST("", mod.SettlementItem.Ctl.Create)
			settlementItems.GET("", mod.SettlementItem.Ctl.List)
			settlementItems.GET("/:id", mod.SettlementItem.Ctl.Info)
			settlementItems.PATCH("/:id", mod.SettlementItem.Ctl.Update)
			settlementItems.DELETE("/:id", mod.SettlementItem.Ctl.Delete)
		}
		returnRequests := systems.Group("/return-requests")
		{
			returnRequests.POST("", mod.Returns.Ctl.Create)
			returnRequests.GET("", mod.Returns.Ctl.List)
			returnRequests.GET("/:id", mod.Returns.Ctl.Info)
			returnRequests.PATCH("/:id", mod.Returns.Ctl.Update)
			returnRequests.DELETE("/:id", mod.Returns.Ctl.Delete)
		}
		returnItems := systems.Group("/return-items")
		{
			returnItems.POST("", mod.ReturnItem.Ctl.Create)
			returnItems.GET("", mod.ReturnItem.Ctl.List)
			returnItems.GET("/:id", mod.ReturnItem.Ctl.Info)
			returnItems.PATCH("/:id", mod.ReturnItem.Ctl.Update)
			returnItems.DELETE("/:id", mod.ReturnItem.Ctl.Delete)
		}
		refundTransactions := systems.Group("/refund-transactions")
		{
			refundTransactions.POST("", mod.RefundTransaction.Ctl.Create)
			refundTransactions.GET("", mod.RefundTransaction.Ctl.List)
			refundTransactions.GET("/:id", mod.RefundTransaction.Ctl.Info)
			refundTransactions.PATCH("/:id", mod.RefundTransaction.Ctl.Update)
			refundTransactions.DELETE("/:id", mod.RefundTransaction.Ctl.Delete)
		}
		rewards := systems.Group("/rewards")
		{
			rewards.POST("", mod.Reward.Ctl.Create)
			rewards.GET("", mod.Reward.Ctl.List)
			rewards.GET("/:id", mod.Reward.Ctl.Info)
			rewards.PATCH("/:id", mod.Reward.Ctl.Update)
			rewards.DELETE("/:id", mod.Reward.Ctl.Delete)
		}
		rewardRedemptions := systems.Group("/reward-redemptions")
		{
			rewardRedemptions.POST("", mod.RewardRedemption.Ctl.Create)
			rewardRedemptions.GET("", mod.RewardRedemption.Ctl.List)
			rewardRedemptions.GET("/:id", mod.RewardRedemption.Ctl.Info)
			rewardRedemptions.PATCH("/:id", mod.RewardRedemption.Ctl.Update)
			rewardRedemptions.DELETE("/:id", mod.RewardRedemption.Ctl.Delete)
		}
		invoices := systems.Group("/invoices")
		{
			invoices.POST("", mod.Invoice.Ctl.Create)
			invoices.GET("", mod.Invoice.Ctl.List)
			invoices.GET("/:id", mod.Invoice.Ctl.Info)
			invoices.PATCH("/:id", mod.Invoice.Ctl.Update)
			invoices.DELETE("/:id", mod.Invoice.Ctl.Delete)
		}
		invoiceItems := systems.Group("/invoice-items")
		{
			invoiceItems.POST("", mod.InvoiceItem.Ctl.Create)
			invoiceItems.GET("", mod.InvoiceItem.Ctl.List)
			invoiceItems.GET("/:id", mod.InvoiceItem.Ctl.Info)
			invoiceItems.PATCH("/:id", mod.InvoiceItem.Ctl.Update)
			invoiceItems.DELETE("/:id", mod.InvoiceItem.Ctl.Delete)
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
		flashSaleItems := systems.Group("/flash-sale-items")
		{
			flashSaleItems.POST("", mod.FlashSaleItem.Ctl.Create)
			flashSaleItems.GET("", mod.FlashSaleItem.Ctl.List)
			flashSaleItems.GET("/:id", mod.FlashSaleItem.Ctl.Info)
			flashSaleItems.PATCH("/:id", mod.FlashSaleItem.Ctl.Update)
			flashSaleItems.DELETE("/:id", mod.FlashSaleItem.Ctl.Delete)
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
	memberDevices := members.Group("/member-devices")
	{
		memberDevices.POST("", mod.MemberDevice.Ctl.Create)
		memberDevices.GET("", mod.MemberDevice.Ctl.List)
		memberDevices.GET("/:id", mod.MemberDevice.Ctl.Info)
		memberDevices.PATCH("/:id", mod.MemberDevice.Ctl.Update)
		memberDevices.DELETE("/:id", mod.MemberDevice.Ctl.Delete)
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
		systems := public.Group("/systems")
		{
			genders := systems.Group("/genders")
			{
				genders.GET("", mod.Gender.Ctl.List)
				genders.GET("/:id", mod.Gender.Ctl.Info)
			}
			prefixes := systems.Group("/prefixes")
			{
				prefixes.GET("", mod.Prefix.Ctl.List)
				prefixes.GET("/:id", mod.Prefix.Ctl.Info)
			}
			provinces := systems.Group("/provinces")
			{
				provinces.GET("", mod.Province.Ctl.List)
				provinces.GET("/:id", mod.Province.Ctl.Info)
			}
			districts := systems.Group("/districts")
			{
				districts.GET("", mod.District.Ctl.List)
				districts.GET("/:id", mod.District.Ctl.Info)
			}
			subdistricts := systems.Group("/subdistricts")
			{
				subdistricts.GET("", mod.Subdistrict.Ctl.List)
				subdistricts.GET("/:id", mod.Subdistrict.Ctl.Info)
			}
			zipcodes := systems.Group("/zipcodes")
			{
				zipcodes.GET("", mod.Zipcode.Ctl.List)
				zipcodes.GET("/:id", mod.Zipcode.Ctl.Info)
			}
		}
		auth := public.Group("/auth")
		{
			auth.POST("/google", mod.Auth.Ctl.GoogleLogin)
			auth.GET("/google/callback", mod.Auth.Ctl.GoogleCallback)
			auth.POST("/register", mod.Auth.Ctl.Register)
			auth.POST("/login", mod.Auth.Ctl.Login)
			auth.POST("/verify", mod.Auth.Ctl.Verify)
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
