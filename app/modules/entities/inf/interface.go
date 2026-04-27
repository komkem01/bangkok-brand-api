package entitiesinf

import (
	"context"
	"time"

	"bangkok-brand/app/modules/entities/ent"

	"github.com/google/uuid"
)

// ObjectEntity defines the interface for object entity operations such as create, retrieve, update, and soft delete.
type ExampleEntity interface {
	CreateExample(ctx context.Context, userID uuid.UUID) (*ent.Example, error)
	GetExampleByID(ctx context.Context, id uuid.UUID) (*ent.Example, error)
	UpdateExampleByID(ctx context.Context, id uuid.UUID, status ent.ExampleStatus) (*ent.Example, error)
	SoftDeleteExampleByID(ctx context.Context, id uuid.UUID) error
	ListExamplesByStatus(ctx context.Context, status ent.ExampleStatus) ([]*ent.Example, error)
}
type ExampleTwoEntity interface {
	CreateExampleTwo(ctx context.Context, userID uuid.UUID) (*ent.Example, error)
}

// GenderEntity defines CRUD operations for the genders table.
type GenderEntity interface {
	ListGenders(ctx context.Context) ([]*ent.Gender, error)
	GetGenderByID(ctx context.Context, id uuid.UUID) (*ent.Gender, error)
	UpdateGenderByID(ctx context.Context, id uuid.UUID, nameTh, nameEn string, isActive bool) (*ent.Gender, error)
	DeleteGenderByID(ctx context.Context, id uuid.UUID) error
}

// PrefixEntity defines CRUD operations for the prefixes table.
type PrefixEntity interface {
	ListPrefixes(ctx context.Context) ([]*ent.Prefix, error)
	GetPrefixByID(ctx context.Context, id uuid.UUID) (*ent.Prefix, error)
	UpdatePrefixByID(ctx context.Context, id uuid.UUID, genderID *uuid.UUID, nameTh, nameEn string, isActive bool) (*ent.Prefix, error)
	DeletePrefixByID(ctx context.Context, id uuid.UUID) error
}

// ProvinceEntity defines CRUD operations for the provinces table.
type ProvinceEntity interface {
	ListProvinces(ctx context.Context) ([]*ent.Province, error)
	GetProvinceByID(ctx context.Context, id uuid.UUID) (*ent.Province, error)
	UpdateProvinceByID(ctx context.Context, id uuid.UUID, name string, isActive bool) (*ent.Province, error)
	DeleteProvinceByID(ctx context.Context, id uuid.UUID) error
}

// DistrictEntity defines CRUD operations for the districts table.
type DistrictEntity interface {
	ListDistricts(ctx context.Context) ([]*ent.District, error)
	GetDistrictByID(ctx context.Context, id uuid.UUID) (*ent.District, error)
	UpdateDistrictByID(ctx context.Context, id uuid.UUID, provinceID *uuid.UUID, name string, isActive bool) (*ent.District, error)
	DeleteDistrictByID(ctx context.Context, id uuid.UUID) error
}

// SubdistrictEntity defines CRUD operations for the sub_districts table.
type SubdistrictEntity interface {
	ListSubdistricts(ctx context.Context) ([]*ent.Subdistrict, error)
	GetSubdistrictByID(ctx context.Context, id uuid.UUID) (*ent.Subdistrict, error)
	UpdateSubdistrictByID(ctx context.Context, id uuid.UUID, districtID *uuid.UUID, name string, isActive bool) (*ent.Subdistrict, error)
	DeleteSubdistrictByID(ctx context.Context, id uuid.UUID) error
}

// ZipcodeEntity defines CRUD operations for the zipcodes table.
type ZipcodeEntity interface {
	ListZipcodes(ctx context.Context) ([]*ent.Zipcode, error)
	GetZipcodeByID(ctx context.Context, id uuid.UUID) (*ent.Zipcode, error)
	UpdateZipcodeByID(ctx context.Context, id uuid.UUID, subDistrictID *uuid.UUID, name string, isActive bool) (*ent.Zipcode, error)
	DeleteZipcodeByID(ctx context.Context, id uuid.UUID) error
}

// BankEntity defines CRUD operations for the banks table.
type BankEntity interface {
	ListBanks(ctx context.Context) ([]*ent.Bank, error)
	GetBankByID(ctx context.Context, id uuid.UUID) (*ent.Bank, error)
	CreateBank(ctx context.Context, b *ent.Bank) (*ent.Bank, error)
	UpdateBankByID(ctx context.Context, id uuid.UUID, nameTh, nameEn, code string, isActive bool) (*ent.Bank, error)
	DeleteBankByID(ctx context.Context, id uuid.UUID) error
}

// StorageEntity defines CRUD operations for the storages table.
type StorageEntity interface {
	ListStorages(ctx context.Context) ([]*ent.Storage, error)
	GetStorageByID(ctx context.Context, id uuid.UUID) (*ent.Storage, error)
	CreateStorage(ctx context.Context, s *ent.Storage) (*ent.Storage, error)
	DeleteStorageByID(ctx context.Context, id uuid.UUID) error
}

// MemberEntity defines CRUD operations for the members table.
type MemberEntity interface {
	ListMembers(ctx context.Context) ([]*ent.Member, error)
	GetMemberByID(ctx context.Context, id uuid.UUID) (*ent.Member, error)
	GetMemberByEmail(ctx context.Context, email string) (*ent.Member, error)
	CreateMember(ctx context.Context, member *ent.Member) (*ent.Member, error)
	UpdateMemberLastLoginByID(ctx context.Context, id uuid.UUID, lastedLogin *time.Time) error
	UpdateMemberByID(ctx context.Context, id uuid.UUID, member *ent.Member) (*ent.Member, error)
	DeleteMemberByID(ctx context.Context, id uuid.UUID) error
}

// ContactEntity defines CRUD operations for the member_contacts table.
type ContactEntity interface {
	ListContacts(ctx context.Context) ([]*ent.MemberContact, error)
	ListContactsByMemberID(ctx context.Context, memberID uuid.UUID) ([]*ent.MemberContact, error)
	GetContactByID(ctx context.Context, id uuid.UUID) (*ent.MemberContact, error)
	CreateContact(ctx context.Context, c *ent.MemberContact) (*ent.MemberContact, error)
	UpdateContactByID(ctx context.Context, id uuid.UUID, c *ent.MemberContact) (*ent.MemberContact, error)
	DeleteContactByID(ctx context.Context, id uuid.UUID) error
}

// ContactTypeEntity defines CRUD operations for the contact_types table.
type ContactTypeEntity interface {
	ListContactTypes(ctx context.Context) ([]*ent.ContactType, error)
	GetContactTypeByID(ctx context.Context, id uuid.UUID) (*ent.ContactType, error)
	GetContactTypeByNameEn(ctx context.Context, nameEn string) (*ent.ContactType, error)
	CreateContactType(ctx context.Context, c *ent.ContactType) (*ent.ContactType, error)
	UpdateContactTypeByID(ctx context.Context, id uuid.UUID, nameTh, nameEn string, isActive bool) (*ent.ContactType, error)
	DeleteContactTypeByID(ctx context.Context, id uuid.UUID) error
}

// AddressTypeEntity defines CRUD operations for the address_types table.
type AddressTypeEntity interface {
	ListAddressTypes(ctx context.Context) ([]*ent.AddressType, error)
	GetAddressTypeByID(ctx context.Context, id uuid.UUID) (*ent.AddressType, error)
	GetAddressTypeByNameEn(ctx context.Context, nameEn string) (*ent.AddressType, error)
	CreateAddressType(ctx context.Context, c *ent.AddressType) (*ent.AddressType, error)
	UpdateAddressTypeByID(ctx context.Context, id uuid.UUID, nameTh, nameEn string, isActive bool) (*ent.AddressType, error)
	DeleteAddressTypeByID(ctx context.Context, id uuid.UUID) error
}

// AddressEntity defines CRUD operations for the member_addresses table.
type AddressEntity interface {
	ListAddresses(ctx context.Context) ([]*ent.MemberAddress, error)
	ListAddressesByMemberID(ctx context.Context, memberID uuid.UUID) ([]*ent.MemberAddress, error)
	GetAddressByID(ctx context.Context, id uuid.UUID) (*ent.MemberAddress, error)
	CreateAddress(ctx context.Context, a *ent.MemberAddress) (*ent.MemberAddress, error)
	UpdateAddressByID(ctx context.Context, id uuid.UUID, a *ent.MemberAddress) (*ent.MemberAddress, error)
	DeleteAddressByID(ctx context.Context, id uuid.UUID) error
}

// MemberDeviceEntity defines CRUD operations for the member_devices table.
type MemberDeviceEntity interface {
	ListMemberDevices(ctx context.Context) ([]*ent.MemberDevice, error)
	ListMemberDevicesByMemberID(ctx context.Context, memberID uuid.UUID) ([]*ent.MemberDevice, error)
	GetMemberDeviceByID(ctx context.Context, id uuid.UUID) (*ent.MemberDevice, error)
	CreateMemberDevice(ctx context.Context, item *ent.MemberDevice) (*ent.MemberDevice, error)
	UpdateMemberDeviceByID(ctx context.Context, id uuid.UUID, item *ent.MemberDevice) (*ent.MemberDevice, error)
	DeleteMemberDeviceByID(ctx context.Context, id uuid.UUID) error
}

// ProductEntity defines CRUD operations for the products table.
type ProductEntity interface {
	ListProducts(ctx context.Context) ([]*ent.Product, error)
	GetProductByID(ctx context.Context, id uuid.UUID) (*ent.Product, error)
	CreateProduct(ctx context.Context, p *ent.Product) (*ent.Product, error)
	UpdateProductByID(ctx context.Context, id uuid.UUID, p *ent.Product) (*ent.Product, error)
	DeleteProductByID(ctx context.Context, id uuid.UUID) error
}

// ProductImageEntity defines CRUD operations for the product_images table.
type ProductImageEntity interface {
	ListProductImages(ctx context.Context) ([]*ent.ProductImage, error)
	GetProductImageByID(ctx context.Context, id uuid.UUID) (*ent.ProductImage, error)
	CreateProductImage(ctx context.Context, p *ent.ProductImage) (*ent.ProductImage, error)
	UpdateProductImageByID(ctx context.Context, id uuid.UUID, p *ent.ProductImage) (*ent.ProductImage, error)
	DeleteProductImageByID(ctx context.Context, id uuid.UUID) error
}

// ProductAttributeEntity defines CRUD operations for the product_attributes table.
type ProductAttributeEntity interface {
	ListProductAttributes(ctx context.Context) ([]*ent.ProductAttribute, error)
	GetProductAttributeByID(ctx context.Context, id uuid.UUID) (*ent.ProductAttribute, error)
	CreateProductAttribute(ctx context.Context, p *ent.ProductAttribute) (*ent.ProductAttribute, error)
	UpdateProductAttributeByID(ctx context.Context, id uuid.UUID, p *ent.ProductAttribute) (*ent.ProductAttribute, error)
	DeleteProductAttributeByID(ctx context.Context, id uuid.UUID) error
}

// ProductAttributeValueEntity defines CRUD operations for the product_attribute_values table.
type ProductAttributeValueEntity interface {
	ListProductAttributeValues(ctx context.Context) ([]*ent.ProductAttributeValue, error)
	GetProductAttributeValueByID(ctx context.Context, id uuid.UUID) (*ent.ProductAttributeValue, error)
	CreateProductAttributeValue(ctx context.Context, p *ent.ProductAttributeValue) (*ent.ProductAttributeValue, error)
	UpdateProductAttributeValueByID(ctx context.Context, id uuid.UUID, p *ent.ProductAttributeValue) (*ent.ProductAttributeValue, error)
	DeleteProductAttributeValueByID(ctx context.Context, id uuid.UUID) error
}

// ProductStockEntity defines CRUD operations for the product_stocks table.
type ProductStockEntity interface {
	ListProductStocks(ctx context.Context) ([]*ent.ProductStock, error)
	GetProductStockByID(ctx context.Context, id uuid.UUID) (*ent.ProductStock, error)
	CreateProductStock(ctx context.Context, p *ent.ProductStock) (*ent.ProductStock, error)
	UpdateProductStockByID(ctx context.Context, id uuid.UUID, p *ent.ProductStock) (*ent.ProductStock, error)
	DeleteProductStockByID(ctx context.Context, id uuid.UUID) error
}

// CartEntity defines CRUD operations for the carts table.
type CartEntity interface {
	ListCarts(ctx context.Context) ([]*ent.Cart, error)
	GetCartByID(ctx context.Context, id uuid.UUID) (*ent.Cart, error)
	CreateCart(ctx context.Context, c *ent.Cart) (*ent.Cart, error)
	UpdateCartByID(ctx context.Context, id uuid.UUID, c *ent.Cart) (*ent.Cart, error)
	DeleteCartByID(ctx context.Context, id uuid.UUID) error
}

// CartItemEntity defines CRUD operations for the cart_items table.
type CartItemEntity interface {
	ListCartItems(ctx context.Context) ([]*ent.CartItem, error)
	GetCartItemByID(ctx context.Context, id uuid.UUID) (*ent.CartItem, error)
	CreateCartItem(ctx context.Context, c *ent.CartItem) (*ent.CartItem, error)
	UpdateCartItemByID(ctx context.Context, id uuid.UUID, c *ent.CartItem) (*ent.CartItem, error)
	DeleteCartItemByID(ctx context.Context, id uuid.UUID) error
}

// OrderEntity defines CRUD operations for the orders table.
type OrderEntity interface {
	ListOrders(ctx context.Context) ([]*ent.Order, error)
	GetOrderByID(ctx context.Context, id uuid.UUID) (*ent.Order, error)
	CreateOrder(ctx context.Context, o *ent.Order) (*ent.Order, error)
	UpdateOrderByID(ctx context.Context, id uuid.UUID, o *ent.Order) (*ent.Order, error)
	DeleteOrderByID(ctx context.Context, id uuid.UUID) error
}

// OrderItemEntity defines CRUD operations for the order_items table.
type OrderItemEntity interface {
	ListOrderItems(ctx context.Context) ([]*ent.OrderItem, error)
	GetOrderItemByID(ctx context.Context, id uuid.UUID) (*ent.OrderItem, error)
	CreateOrderItem(ctx context.Context, o *ent.OrderItem) (*ent.OrderItem, error)
	UpdateOrderItemByID(ctx context.Context, id uuid.UUID, o *ent.OrderItem) (*ent.OrderItem, error)
	DeleteOrderItemByID(ctx context.Context, id uuid.UUID) error
}

// OrderStatusHistoryEntity defines CRUD operations for the order_status_histories table.
type OrderStatusHistoryEntity interface {
	ListOrderStatusHistories(ctx context.Context) ([]*ent.OrderStatusHistory, error)
	GetOrderStatusHistoryByID(ctx context.Context, id uuid.UUID) (*ent.OrderStatusHistory, error)
	CreateOrderStatusHistory(ctx context.Context, o *ent.OrderStatusHistory) (*ent.OrderStatusHistory, error)
	UpdateOrderStatusHistoryByID(ctx context.Context, id uuid.UUID, o *ent.OrderStatusHistory) (*ent.OrderStatusHistory, error)
	DeleteOrderStatusHistoryByID(ctx context.Context, id uuid.UUID) error
}

// PaymentEntity defines CRUD operations for the payments table.
type PaymentEntity interface {
	ListPayments(ctx context.Context) ([]*ent.Payment, error)
	GetPaymentByID(ctx context.Context, id uuid.UUID) (*ent.Payment, error)
	CreatePayment(ctx context.Context, p *ent.Payment) (*ent.Payment, error)
	UpdatePaymentByID(ctx context.Context, id uuid.UUID, p *ent.Payment) (*ent.Payment, error)
	DeletePaymentByID(ctx context.Context, id uuid.UUID) error
}

// BrandEntity defines CRUD operations for the brands table.
type BrandEntity interface {
	ListBrands(ctx context.Context) ([]*ent.Brand, error)
	GetBrandByID(ctx context.Context, id uuid.UUID) (*ent.Brand, error)
	CreateBrand(ctx context.Context, b *ent.Brand) (*ent.Brand, error)
	UpdateBrandByID(ctx context.Context, id uuid.UUID, b *ent.Brand) (*ent.Brand, error)
	DeleteBrandByID(ctx context.Context, id uuid.UUID) error
}

// CategoryEntity defines CRUD operations for the categories table.
type CategoryEntity interface {
	ListCategories(ctx context.Context) ([]*ent.Category, error)
	GetCategoryByID(ctx context.Context, id uuid.UUID) (*ent.Category, error)
	CreateCategory(ctx context.Context, c *ent.Category) (*ent.Category, error)
	UpdateCategoryByID(ctx context.Context, id uuid.UUID, c *ent.Category) (*ent.Category, error)
	DeleteCategoryByID(ctx context.Context, id uuid.UUID) error
}

// MemberBankAccountEntity defines CRUD operations for the member_bank_accounts table.
type MemberBankAccountEntity interface {
	ListMemberBankAccounts(ctx context.Context) ([]*ent.MemberBankAccount, error)
	GetMemberBankAccountByID(ctx context.Context, id uuid.UUID) (*ent.MemberBankAccount, error)
	CreateMemberBankAccount(ctx context.Context, a *ent.MemberBankAccount) (*ent.MemberBankAccount, error)
	UpdateMemberBankAccountByID(ctx context.Context, id uuid.UUID, a *ent.MemberBankAccount) (*ent.MemberBankAccount, error)
	DeleteMemberBankAccountByID(ctx context.Context, id uuid.UUID) error
}

type NotificationEntity interface {
	ListNotifications(ctx context.Context) ([]*ent.Notification, error)
	GetNotificationByID(ctx context.Context, id uuid.UUID) (*ent.Notification, error)
	CreateNotification(ctx context.Context, item *ent.Notification) (*ent.Notification, error)
	UpdateNotificationByID(ctx context.Context, id uuid.UUID, item *ent.Notification) (*ent.Notification, error)
	DeleteNotificationByID(ctx context.Context, id uuid.UUID) error
}

type FlashSaleEventEntity interface {
	ListFlashSaleEvents(ctx context.Context) ([]*ent.FlashSaleEvent, error)
	GetFlashSaleEventByID(ctx context.Context, id uuid.UUID) (*ent.FlashSaleEvent, error)
	CreateFlashSaleEvent(ctx context.Context, item *ent.FlashSaleEvent) (*ent.FlashSaleEvent, error)
	UpdateFlashSaleEventByID(ctx context.Context, id uuid.UUID, item *ent.FlashSaleEvent) (*ent.FlashSaleEvent, error)
	DeleteFlashSaleEventByID(ctx context.Context, id uuid.UUID) error
}

type WishlistEntity interface {
	ListWishlists(ctx context.Context) ([]*ent.Wishlist, error)
	GetWishlistByID(ctx context.Context, id uuid.UUID) (*ent.Wishlist, error)
	CreateWishlist(ctx context.Context, item *ent.Wishlist) (*ent.Wishlist, error)
	UpdateWishlistByID(ctx context.Context, id uuid.UUID, item *ent.Wishlist) (*ent.Wishlist, error)
	DeleteWishlistByID(ctx context.Context, id uuid.UUID) error
}

type SearchHistoryEntity interface {
	ListSearchHistories(ctx context.Context) ([]*ent.SearchHistory, error)
	GetSearchHistoryByID(ctx context.Context, id uuid.UUID) (*ent.SearchHistory, error)
	CreateSearchHistory(ctx context.Context, item *ent.SearchHistory) (*ent.SearchHistory, error)
	UpdateSearchHistoryByID(ctx context.Context, id uuid.UUID, item *ent.SearchHistory) (*ent.SearchHistory, error)
	DeleteSearchHistoryByID(ctx context.Context, id uuid.UUID) error
}

type ProductViewEntity interface {
	ListProductViews(ctx context.Context) ([]*ent.ProductView, error)
	GetProductViewByID(ctx context.Context, id uuid.UUID) (*ent.ProductView, error)
	CreateProductView(ctx context.Context, item *ent.ProductView) (*ent.ProductView, error)
	UpdateProductViewByID(ctx context.Context, id uuid.UUID, item *ent.ProductView) (*ent.ProductView, error)
	DeleteProductViewByID(ctx context.Context, id uuid.UUID) error
}

type IdempotencyKeyEntity interface {
	ListIdempotencyKeys(ctx context.Context) ([]*ent.IdempotencyKey, error)
	GetIdempotencyKeyByID(ctx context.Context, id uuid.UUID) (*ent.IdempotencyKey, error)
	CreateIdempotencyKey(ctx context.Context, item *ent.IdempotencyKey) (*ent.IdempotencyKey, error)
	UpdateIdempotencyKeyByID(ctx context.Context, id uuid.UUID, item *ent.IdempotencyKey) (*ent.IdempotencyKey, error)
	DeleteIdempotencyKeyByID(ctx context.Context, id uuid.UUID) error
}

type WebhookEventEntity interface {
	ListWebhookEvents(ctx context.Context) ([]*ent.WebhookEvent, error)
	GetWebhookEventByID(ctx context.Context, id uuid.UUID) (*ent.WebhookEvent, error)
	CreateWebhookEvent(ctx context.Context, item *ent.WebhookEvent) (*ent.WebhookEvent, error)
	UpdateWebhookEventByID(ctx context.Context, id uuid.UUID, item *ent.WebhookEvent) (*ent.WebhookEvent, error)
	DeleteWebhookEventByID(ctx context.Context, id uuid.UUID) error
}

type AuditLogEntity interface {
	ListAuditLogs(ctx context.Context) ([]*ent.AuditLog, error)
	GetAuditLogByID(ctx context.Context, id uuid.UUID) (*ent.AuditLog, error)
	CreateAuditLog(ctx context.Context, item *ent.AuditLog) (*ent.AuditLog, error)
	UpdateAuditLogByID(ctx context.Context, id uuid.UUID, item *ent.AuditLog) (*ent.AuditLog, error)
	DeleteAuditLogByID(ctx context.Context, id uuid.UUID) error
}

type CouponEntity interface {
	ListCoupons(ctx context.Context) ([]*ent.Coupon, error)
	GetCouponByID(ctx context.Context, id uuid.UUID) (*ent.Coupon, error)
	CreateCoupon(ctx context.Context, item *ent.Coupon) (*ent.Coupon, error)
	UpdateCouponByID(ctx context.Context, id uuid.UUID, item *ent.Coupon) (*ent.Coupon, error)
	DeleteCouponByID(ctx context.Context, id uuid.UUID) error
}

type ReviewEntity interface {
	ListReviews(ctx context.Context) ([]*ent.ProductReview, error)
	GetReviewByID(ctx context.Context, id uuid.UUID) (*ent.ProductReview, error)
	CreateReview(ctx context.Context, item *ent.ProductReview) (*ent.ProductReview, error)
	UpdateReviewByID(ctx context.Context, id uuid.UUID, item *ent.ProductReview) (*ent.ProductReview, error)
	DeleteReviewByID(ctx context.Context, id uuid.UUID) error
}

type PointTransactionEntity interface {
	ListPointTransactions(ctx context.Context) ([]*ent.PointTransaction, error)
	GetPointTransactionByID(ctx context.Context, id uuid.UUID) (*ent.PointTransaction, error)
	CreatePointTransaction(ctx context.Context, item *ent.PointTransaction) (*ent.PointTransaction, error)
	UpdatePointTransactionByID(ctx context.Context, id uuid.UUID, item *ent.PointTransaction) (*ent.PointTransaction, error)
	DeletePointTransactionByID(ctx context.Context, id uuid.UUID) error
}

type ChatRoomEntity interface {
	ListChatRooms(ctx context.Context) ([]*ent.ChatRoom, error)
	GetChatRoomByID(ctx context.Context, id uuid.UUID) (*ent.ChatRoom, error)
	CreateChatRoom(ctx context.Context, item *ent.ChatRoom) (*ent.ChatRoom, error)
	UpdateChatRoomByID(ctx context.Context, id uuid.UUID, item *ent.ChatRoom) (*ent.ChatRoom, error)
	DeleteChatRoomByID(ctx context.Context, id uuid.UUID) error
}

type KYCVerificationEntity interface {
	ListKYCVerifications(ctx context.Context) ([]*ent.KYCVerification, error)
	GetKYCVerificationByID(ctx context.Context, id uuid.UUID) (*ent.KYCVerification, error)
	CreateKYCVerification(ctx context.Context, item *ent.KYCVerification) (*ent.KYCVerification, error)
	UpdateKYCVerificationByID(ctx context.Context, id uuid.UUID, item *ent.KYCVerification) (*ent.KYCVerification, error)
	DeleteKYCVerificationByID(ctx context.Context, id uuid.UUID) error
}

type LogisticsProviderEntity interface {
	ListLogisticsProviders(ctx context.Context) ([]*ent.LogisticsProvider, error)
	GetLogisticsProviderByID(ctx context.Context, id uuid.UUID) (*ent.LogisticsProvider, error)
	CreateLogisticsProvider(ctx context.Context, item *ent.LogisticsProvider) (*ent.LogisticsProvider, error)
	UpdateLogisticsProviderByID(ctx context.Context, id uuid.UUID, item *ent.LogisticsProvider) (*ent.LogisticsProvider, error)
	DeleteLogisticsProviderByID(ctx context.Context, id uuid.UUID) error
}

type ShopEntity interface {
	ListShops(ctx context.Context) ([]*ent.Shop, error)
	GetShopByID(ctx context.Context, id uuid.UUID) (*ent.Shop, error)
	CreateShop(ctx context.Context, item *ent.Shop) (*ent.Shop, error)
	UpdateShopByID(ctx context.Context, id uuid.UUID, item *ent.Shop) (*ent.Shop, error)
	DeleteShopByID(ctx context.Context, id uuid.UUID) error
}

type VariantEntity interface {
	ListVariants(ctx context.Context) ([]*ent.ProductVariant, error)
	GetVariantByID(ctx context.Context, id uuid.UUID) (*ent.ProductVariant, error)
	CreateVariant(ctx context.Context, item *ent.ProductVariant) (*ent.ProductVariant, error)
	UpdateVariantByID(ctx context.Context, id uuid.UUID, item *ent.ProductVariant) (*ent.ProductVariant, error)
	DeleteVariantByID(ctx context.Context, id uuid.UUID) error
}

type ShippingZoneEntity interface {
	ListShippingZones(ctx context.Context) ([]*ent.ShippingZone, error)
	GetShippingZoneByID(ctx context.Context, id uuid.UUID) (*ent.ShippingZone, error)
	CreateShippingZone(ctx context.Context, item *ent.ShippingZone) (*ent.ShippingZone, error)
	UpdateShippingZoneByID(ctx context.Context, id uuid.UUID, item *ent.ShippingZone) (*ent.ShippingZone, error)
	DeleteShippingZoneByID(ctx context.Context, id uuid.UUID) error
}

type SettlementEntity interface {
	ListSettlements(ctx context.Context) ([]*ent.SettlementBatch, error)
	GetSettlementByID(ctx context.Context, id uuid.UUID) (*ent.SettlementBatch, error)
	CreateSettlement(ctx context.Context, item *ent.SettlementBatch) (*ent.SettlementBatch, error)
	UpdateSettlementByID(ctx context.Context, id uuid.UUID, item *ent.SettlementBatch) (*ent.SettlementBatch, error)
	DeleteSettlementByID(ctx context.Context, id uuid.UUID) error
}

type ReturnRequestEntity interface {
	ListReturnRequests(ctx context.Context) ([]*ent.ReturnRequest, error)
	GetReturnRequestByID(ctx context.Context, id uuid.UUID) (*ent.ReturnRequest, error)
	CreateReturnRequest(ctx context.Context, item *ent.ReturnRequest) (*ent.ReturnRequest, error)
	UpdateReturnRequestByID(ctx context.Context, id uuid.UUID, item *ent.ReturnRequest) (*ent.ReturnRequest, error)
	DeleteReturnRequestByID(ctx context.Context, id uuid.UUID) error
}

type InvoiceEntity interface {
	ListInvoices(ctx context.Context) ([]*ent.Invoice, error)
	GetInvoiceByID(ctx context.Context, id uuid.UUID) (*ent.Invoice, error)
	CreateInvoice(ctx context.Context, item *ent.Invoice) (*ent.Invoice, error)
	UpdateInvoiceByID(ctx context.Context, id uuid.UUID, item *ent.Invoice) (*ent.Invoice, error)
	DeleteInvoiceByID(ctx context.Context, id uuid.UUID) error
}

type AdminActionLogEntity interface {
	ListAdminActionLogs(ctx context.Context) ([]*ent.AdminActionLog, error)
	GetAdminActionLogByID(ctx context.Context, id uuid.UUID) (*ent.AdminActionLog, error)
	CreateAdminActionLog(ctx context.Context, item *ent.AdminActionLog) (*ent.AdminActionLog, error)
	UpdateAdminActionLogByID(ctx context.Context, id uuid.UUID, item *ent.AdminActionLog) (*ent.AdminActionLog, error)
	DeleteAdminActionLogByID(ctx context.Context, id uuid.UUID) error
}

type ChatMessageEntity interface {
	ListChatMessages(ctx context.Context) ([]*ent.ChatMessage, error)
	GetChatMessageByID(ctx context.Context, id uuid.UUID) (*ent.ChatMessage, error)
	CreateChatMessage(ctx context.Context, item *ent.ChatMessage) (*ent.ChatMessage, error)
	UpdateChatMessageByID(ctx context.Context, id uuid.UUID, item *ent.ChatMessage) (*ent.ChatMessage, error)
	DeleteChatMessageByID(ctx context.Context, id uuid.UUID) error
}

type ChatParticipantEntity interface {
	ListChatParticipants(ctx context.Context) ([]*ent.ChatParticipant, error)
	GetChatParticipantByID(ctx context.Context, id uuid.UUID) (*ent.ChatParticipant, error)
	CreateChatParticipant(ctx context.Context, item *ent.ChatParticipant) (*ent.ChatParticipant, error)
	UpdateChatParticipantByID(ctx context.Context, id uuid.UUID, item *ent.ChatParticipant) (*ent.ChatParticipant, error)
	DeleteChatParticipantByID(ctx context.Context, id uuid.UUID) error
}

type CouponUsageEntity interface {
	ListCouponUsages(ctx context.Context) ([]*ent.CouponUsage, error)
	GetCouponUsageByID(ctx context.Context, id uuid.UUID) (*ent.CouponUsage, error)
	CreateCouponUsage(ctx context.Context, item *ent.CouponUsage) (*ent.CouponUsage, error)
	UpdateCouponUsageByID(ctx context.Context, id uuid.UUID, item *ent.CouponUsage) (*ent.CouponUsage, error)
	DeleteCouponUsageByID(ctx context.Context, id uuid.UUID) error
}

type DisputeCaseEntity interface {
	ListDisputeCases(ctx context.Context) ([]*ent.DisputeCase, error)
	GetDisputeCaseByID(ctx context.Context, id uuid.UUID) (*ent.DisputeCase, error)
	CreateDisputeCase(ctx context.Context, item *ent.DisputeCase) (*ent.DisputeCase, error)
	UpdateDisputeCaseByID(ctx context.Context, id uuid.UUID, item *ent.DisputeCase) (*ent.DisputeCase, error)
	DeleteDisputeCaseByID(ctx context.Context, id uuid.UUID) error
}

type DisputeMessageEntity interface {
	ListDisputeMessages(ctx context.Context) ([]*ent.DisputeMessage, error)
	GetDisputeMessageByID(ctx context.Context, id uuid.UUID) (*ent.DisputeMessage, error)
	CreateDisputeMessage(ctx context.Context, item *ent.DisputeMessage) (*ent.DisputeMessage, error)
	UpdateDisputeMessageByID(ctx context.Context, id uuid.UUID, item *ent.DisputeMessage) (*ent.DisputeMessage, error)
	DeleteDisputeMessageByID(ctx context.Context, id uuid.UUID) error
}

type FlashSaleItemEntity interface {
	ListFlashSaleItems(ctx context.Context) ([]*ent.FlashSaleItem, error)
	GetFlashSaleItemByID(ctx context.Context, id uuid.UUID) (*ent.FlashSaleItem, error)
	CreateFlashSaleItem(ctx context.Context, item *ent.FlashSaleItem) (*ent.FlashSaleItem, error)
	UpdateFlashSaleItemByID(ctx context.Context, id uuid.UUID, item *ent.FlashSaleItem) (*ent.FlashSaleItem, error)
	DeleteFlashSaleItemByID(ctx context.Context, id uuid.UUID) error
}

type InvoiceItemEntity interface {
	ListInvoiceItems(ctx context.Context) ([]*ent.InvoiceItem, error)
	GetInvoiceItemByID(ctx context.Context, id uuid.UUID) (*ent.InvoiceItem, error)
	CreateInvoiceItem(ctx context.Context, item *ent.InvoiceItem) (*ent.InvoiceItem, error)
	UpdateInvoiceItemByID(ctx context.Context, id uuid.UUID, item *ent.InvoiceItem) (*ent.InvoiceItem, error)
	DeleteInvoiceItemByID(ctx context.Context, id uuid.UUID) error
}

type KYCDocumentEntity interface {
	ListKYCDocuments(ctx context.Context) ([]*ent.KYCDocument, error)
	GetKYCDocumentByID(ctx context.Context, id uuid.UUID) (*ent.KYCDocument, error)
	CreateKYCDocument(ctx context.Context, item *ent.KYCDocument) (*ent.KYCDocument, error)
	UpdateKYCDocumentByID(ctx context.Context, id uuid.UUID, item *ent.KYCDocument) (*ent.KYCDocument, error)
	DeleteKYCDocumentByID(ctx context.Context, id uuid.UUID) error
}

type KYCStatusHistoryEntity interface {
	ListKYCStatusHistories(ctx context.Context) ([]*ent.KYCStatusHistory, error)
	GetKYCStatusHistoryByID(ctx context.Context, id uuid.UUID) (*ent.KYCStatusHistory, error)
	CreateKYCStatusHistory(ctx context.Context, item *ent.KYCStatusHistory) (*ent.KYCStatusHistory, error)
	UpdateKYCStatusHistoryByID(ctx context.Context, id uuid.UUID, item *ent.KYCStatusHistory) (*ent.KYCStatusHistory, error)
	DeleteKYCStatusHistoryByID(ctx context.Context, id uuid.UUID) error
}

type OrderShipmentEntity interface {
	ListOrderShipments(ctx context.Context) ([]*ent.OrderShipment, error)
	GetOrderShipmentByID(ctx context.Context, id uuid.UUID) (*ent.OrderShipment, error)
	CreateOrderShipment(ctx context.Context, item *ent.OrderShipment) (*ent.OrderShipment, error)
	UpdateOrderShipmentByID(ctx context.Context, id uuid.UUID, item *ent.OrderShipment) (*ent.OrderShipment, error)
	DeleteOrderShipmentByID(ctx context.Context, id uuid.UUID) error
}

type PointSettingEntity interface {
	ListPointSettings(ctx context.Context) ([]*ent.PointSetting, error)
	GetPointSettingByID(ctx context.Context, id uuid.UUID) (*ent.PointSetting, error)
	CreatePointSetting(ctx context.Context, item *ent.PointSetting) (*ent.PointSetting, error)
	UpdatePointSettingByID(ctx context.Context, id uuid.UUID, item *ent.PointSetting) (*ent.PointSetting, error)
	DeletePointSettingByID(ctx context.Context, id uuid.UUID) error
}

type ProductReviewImageEntity interface {
	ListProductReviewImages(ctx context.Context) ([]*ent.ProductReviewImage, error)
	GetProductReviewImageByID(ctx context.Context, id uuid.UUID) (*ent.ProductReviewImage, error)
	CreateProductReviewImage(ctx context.Context, item *ent.ProductReviewImage) (*ent.ProductReviewImage, error)
	UpdateProductReviewImageByID(ctx context.Context, id uuid.UUID, item *ent.ProductReviewImage) (*ent.ProductReviewImage, error)
	DeleteProductReviewImageByID(ctx context.Context, id uuid.UUID) error
}

type ProductVariantStockEntity interface {
	ListProductVariantStocks(ctx context.Context) ([]*ent.ProductVariantStock, error)
	GetProductVariantStockByID(ctx context.Context, id uuid.UUID) (*ent.ProductVariantStock, error)
	CreateProductVariantStock(ctx context.Context, item *ent.ProductVariantStock) (*ent.ProductVariantStock, error)
	UpdateProductVariantStockByID(ctx context.Context, id uuid.UUID, item *ent.ProductVariantStock) (*ent.ProductVariantStock, error)
	DeleteProductVariantStockByID(ctx context.Context, id uuid.UUID) error
}

type ProductVariantValueEntity interface {
	ListProductVariantValues(ctx context.Context) ([]*ent.ProductVariantValue, error)
	GetProductVariantValueByID(ctx context.Context, id uuid.UUID) (*ent.ProductVariantValue, error)
	CreateProductVariantValue(ctx context.Context, item *ent.ProductVariantValue) (*ent.ProductVariantValue, error)
	UpdateProductVariantValueByID(ctx context.Context, id uuid.UUID, item *ent.ProductVariantValue) (*ent.ProductVariantValue, error)
	DeleteProductVariantValueByID(ctx context.Context, id uuid.UUID) error
}

type RefundTransactionEntity interface {
	ListRefundTransactions(ctx context.Context) ([]*ent.RefundTransaction, error)
	GetRefundTransactionByID(ctx context.Context, id uuid.UUID) (*ent.RefundTransaction, error)
	CreateRefundTransaction(ctx context.Context, item *ent.RefundTransaction) (*ent.RefundTransaction, error)
	UpdateRefundTransactionByID(ctx context.Context, id uuid.UUID, item *ent.RefundTransaction) (*ent.RefundTransaction, error)
	DeleteRefundTransactionByID(ctx context.Context, id uuid.UUID) error
}

type ReturnItemEntity interface {
	ListReturnItems(ctx context.Context) ([]*ent.ReturnItem, error)
	GetReturnItemByID(ctx context.Context, id uuid.UUID) (*ent.ReturnItem, error)
	CreateReturnItem(ctx context.Context, item *ent.ReturnItem) (*ent.ReturnItem, error)
	UpdateReturnItemByID(ctx context.Context, id uuid.UUID, item *ent.ReturnItem) (*ent.ReturnItem, error)
	DeleteReturnItemByID(ctx context.Context, id uuid.UUID) error
}

type RewardRedemptionEntity interface {
	ListRewardRedemptions(ctx context.Context) ([]*ent.RewardRedemption, error)
	GetRewardRedemptionByID(ctx context.Context, id uuid.UUID) (*ent.RewardRedemption, error)
	CreateRewardRedemption(ctx context.Context, item *ent.RewardRedemption) (*ent.RewardRedemption, error)
	UpdateRewardRedemptionByID(ctx context.Context, id uuid.UUID, item *ent.RewardRedemption) (*ent.RewardRedemption, error)
	DeleteRewardRedemptionByID(ctx context.Context, id uuid.UUID) error
}

type RewardEntity interface {
	ListRewards(ctx context.Context) ([]*ent.Reward, error)
	GetRewardByID(ctx context.Context, id uuid.UUID) (*ent.Reward, error)
	CreateReward(ctx context.Context, item *ent.Reward) (*ent.Reward, error)
	UpdateRewardByID(ctx context.Context, id uuid.UUID, item *ent.Reward) (*ent.Reward, error)
	DeleteRewardByID(ctx context.Context, id uuid.UUID) error
}

type SettlementItemEntity interface {
	ListSettlementItems(ctx context.Context) ([]*ent.SettlementItem, error)
	GetSettlementItemByID(ctx context.Context, id uuid.UUID) (*ent.SettlementItem, error)
	CreateSettlementItem(ctx context.Context, item *ent.SettlementItem) (*ent.SettlementItem, error)
	UpdateSettlementItemByID(ctx context.Context, id uuid.UUID, item *ent.SettlementItem) (*ent.SettlementItem, error)
	DeleteSettlementItemByID(ctx context.Context, id uuid.UUID) error
}

type ShipmentTrackingHistoryEntity interface {
	ListShipmentTrackingHistories(ctx context.Context) ([]*ent.ShipmentTrackingHistory, error)
	GetShipmentTrackingHistoryByID(ctx context.Context, id uuid.UUID) (*ent.ShipmentTrackingHistory, error)
	CreateShipmentTrackingHistory(ctx context.Context, item *ent.ShipmentTrackingHistory) (*ent.ShipmentTrackingHistory, error)
	UpdateShipmentTrackingHistoryByID(ctx context.Context, id uuid.UUID, item *ent.ShipmentTrackingHistory) (*ent.ShipmentTrackingHistory, error)
	DeleteShipmentTrackingHistoryByID(ctx context.Context, id uuid.UUID) error
}

type ShopMemberEntity interface {
	ListShopMembers(ctx context.Context) ([]*ent.ShopMember, error)
	GetShopMemberByID(ctx context.Context, id uuid.UUID) (*ent.ShopMember, error)
	CreateShopMember(ctx context.Context, item *ent.ShopMember) (*ent.ShopMember, error)
	UpdateShopMemberByID(ctx context.Context, id uuid.UUID, item *ent.ShopMember) (*ent.ShopMember, error)
	DeleteShopMemberByID(ctx context.Context, id uuid.UUID) error
}

type ShopSettingEntity interface {
	ListShopSettings(ctx context.Context) ([]*ent.ShopSetting, error)
	GetShopSettingByID(ctx context.Context, id uuid.UUID) (*ent.ShopSetting, error)
	CreateShopSetting(ctx context.Context, item *ent.ShopSetting) (*ent.ShopSetting, error)
	UpdateShopSettingByID(ctx context.Context, id uuid.UUID, item *ent.ShopSetting) (*ent.ShopSetting, error)
	DeleteShopSettingByID(ctx context.Context, id uuid.UUID) error
}

type ShippingZoneAreaEntity interface {
	ListShippingZoneAreas(ctx context.Context) ([]*ent.ShippingZoneArea, error)
	GetShippingZoneAreaByID(ctx context.Context, id uuid.UUID) (*ent.ShippingZoneArea, error)
	CreateShippingZoneArea(ctx context.Context, item *ent.ShippingZoneArea) (*ent.ShippingZoneArea, error)
	UpdateShippingZoneAreaByID(ctx context.Context, id uuid.UUID, item *ent.ShippingZoneArea) (*ent.ShippingZoneArea, error)
	DeleteShippingZoneAreaByID(ctx context.Context, id uuid.UUID) error
}

type ShopShippingMethodEntity interface {
	ListShopShippingMethods(ctx context.Context) ([]*ent.ShopShippingMethod, error)
	GetShopShippingMethodByID(ctx context.Context, id uuid.UUID) (*ent.ShopShippingMethod, error)
	CreateShopShippingMethod(ctx context.Context, item *ent.ShopShippingMethod) (*ent.ShopShippingMethod, error)
	UpdateShopShippingMethodByID(ctx context.Context, id uuid.UUID, item *ent.ShopShippingMethod) (*ent.ShopShippingMethod, error)
	DeleteShopShippingMethodByID(ctx context.Context, id uuid.UUID) error
}

type ShippingRateRuleEntity interface {
	ListShippingRateRules(ctx context.Context) ([]*ent.ShippingRateRule, error)
	GetShippingRateRuleByID(ctx context.Context, id uuid.UUID) (*ent.ShippingRateRule, error)
	CreateShippingRateRule(ctx context.Context, item *ent.ShippingRateRule) (*ent.ShippingRateRule, error)
	UpdateShippingRateRuleByID(ctx context.Context, id uuid.UUID, item *ent.ShippingRateRule) (*ent.ShippingRateRule, error)
	DeleteShippingRateRuleByID(ctx context.Context, id uuid.UUID) error
}

type ShippingMethodEntity interface {
	ListShippingMethods(ctx context.Context) ([]*ent.ShippingMethod, error)
	GetShippingMethodByID(ctx context.Context, id uuid.UUID) (*ent.ShippingMethod, error)
	CreateShippingMethod(ctx context.Context, item *ent.ShippingMethod) (*ent.ShippingMethod, error)
	UpdateShippingMethodByID(ctx context.Context, id uuid.UUID, item *ent.ShippingMethod) (*ent.ShippingMethod, error)
	DeleteShippingMethodByID(ctx context.Context, id uuid.UUID) error
}

type ShopWalletTransactionEntity interface {
	ListShopWalletTransactions(ctx context.Context) ([]*ent.ShopWalletTransaction, error)
	GetShopWalletTransactionByID(ctx context.Context, id uuid.UUID) (*ent.ShopWalletTransaction, error)
	CreateShopWalletTransaction(ctx context.Context, item *ent.ShopWalletTransaction) (*ent.ShopWalletTransaction, error)
	UpdateShopWalletTransactionByID(ctx context.Context, id uuid.UUID, item *ent.ShopWalletTransaction) (*ent.ShopWalletTransaction, error)
	DeleteShopWalletTransactionByID(ctx context.Context, id uuid.UUID) error
}
