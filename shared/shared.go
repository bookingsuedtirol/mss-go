package shared

import (
	"encoding/xml"
	"time"
)

type Date struct{ time.Time }

const dateLayout = "2006-01-02"

func (d Date) MarshalXML(element *xml.Encoder, start xml.StartElement) error {
	return element.EncodeElement(d.Format(dateLayout), start)
}

func (d *Date) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var str string
	err := decoder.DecodeElement(&str, &start)
	if err != nil {
		return err
	}

	value, err := ParseDateTime(dateLayout, str)
	if err != nil {
		return err
	}

	if value != nil {
		*d = Date{*value}
	}

	return nil
}

func ParseDateTime(layout string, value string) (*time.Time, error) {
	if emptyDateTime(value) {
		return nil, nil
	}

	parsed, err := time.Parse(layout, value)
	if err != nil {
		return nil, err
	}

	return &parsed, nil
}

func ParseLocalDateTime(layout string, value string) (*time.Time, error) {
	if emptyDateTime(value) {
		return nil, nil
	}

	loc, err := loadLocation()
	if err != nil {
		return nil, err
	}

	parsed, err := time.ParseInLocation(layout, value, loc)
	if err != nil {
		return nil, err
	}

	return &parsed, nil
}

// location caches the loaded time zone.
// Calling time.LoadLocation before each time.ParseInLocation
// would decrease performance.
var location *time.Location

// loadLocation returns the location Europe/Rome
// which is the local time zone of MSS.
func loadLocation() (*time.Location, error) {
	if location != nil {
		return location, nil
	}

	loc, err := time.LoadLocation("Europe/Rome")
	if err != nil {
		return nil, err
	}

	location = loc
	return loc, nil
}

// emptyDateTime checks if the MSS XML output
// is an empty or zero date/time.
func emptyDateTime(t string) bool {
	return t == "" || t == "0000-00-00" || t == "0000-00-00 00:00:00"
}

type LTSHotelStatus int

const (
	LTSHotelStatusUndefined LTSHotelStatus = iota
	LTSHotelActive
	LTSHotelInactive
)

var LTSHotelStatuses = []LTSHotelStatus{
	LTSHotelStatusUndefined,
	LTSHotelActive,
	LTSHotelInactive,
}

type MemberOfTourismAssociation int

const (
	MemberOfTourismAssociationUndefined MemberOfTourismAssociation = iota
	MemberOfTourismAssociationFalse
	MemberOfTourismAssociationTrue
)

var MemberOfTourismAssociations = []MemberOfTourismAssociation{
	MemberOfTourismAssociationUndefined,
	MemberOfTourismAssociationFalse,
	MemberOfTourismAssociationTrue,
}

type LTSHotelRepresentation int

const (
	LTSHotelRepresentationUndefined LTSHotelRepresentation = iota
	LTSHotelRepresentationDoNotDisplay
	LTSHotelRepresentationMinimal
	LTSHotelRepresentationComplete
)

var LTSHotelRepresentations = []LTSHotelRepresentation{
	LTSHotelRepresentationUndefined,
	LTSHotelRepresentationDoNotDisplay,
	LTSHotelRepresentationMinimal,
	LTSHotelRepresentationComplete,
}

type LTSData struct {
	A0Ene LTSHotelStatus             `xml:"A0Ene"`
	A0MTV MemberOfTourismAssociation `xml:"A0MTV"`
	A0Rep LTSHotelRepresentation     `xml:"A0Rep"`
}

type RoomType int

const (
	RoomTypeRoom RoomType = 1 << iota
	RoomTypeApartment
	RoomTypeCampingPitch
)

var RoomTypes = []RoomType{
	RoomTypeRoom,
	RoomTypeApartment,
	RoomTypeCampingPitch,
}

type HotelType int

const (
	HotelTypeHotel HotelType = 1 << iota
	HotelTypeSkiSchool
	HotelTypeResidence
	HotelTypeBBAndAppartmentsPriv HotelType = 2 << iota
	HotelTypeFarmVacation
	HotelTypeMountainInn
	HotelTypeCampingSite
	HotelTypeHolidayHome
	HotelTypeYouthHostel
	HotelTypeGuesthouse
	HotelTypeRefuge
	HotelTypeGarni
	HotelTypeInn
)

var HotelTypes = []HotelType{
	HotelTypeHotel,
	HotelTypeSkiSchool,
	HotelTypeResidence,
	HotelTypeBBAndAppartmentsPriv,
	HotelTypeFarmVacation,
	HotelTypeMountainInn,
	HotelTypeCampingSite,
	HotelTypeHolidayHome,
	HotelTypeYouthHostel,
	HotelTypeGuesthouse,
	HotelTypeRefuge,
	HotelTypeGarni,
	HotelTypeInn,
}

type HotelStars float64

const (
	HotelStarsUndefined HotelStars = 0
	HotelStars1         HotelStars = 1
	HotelStars2         HotelStars = 2
	HotelStars3         HotelStars = 3
	HotelStars3S        HotelStars = 3.5
	HotelStars4         HotelStars = 4
	HotelStars4S        HotelStars = 4.5
	HotelStars5         HotelStars = 5
)

var HotelStarsList = []HotelStars{
	HotelStarsUndefined,
	HotelStars1,
	HotelStars2,
	HotelStars3,
	HotelStars3S,
	HotelStars4,
	HotelStars4S,
	HotelStars5,
}

type HotelFeature int

const (
	HotelFeatureGarage HotelFeature = 1 << iota
	HotelFeatureElevatorLift
	HotelFeatureRestaurant
	HotelFeatureGym
	HotelFeatureWellness
	HotelFeatureSpaCuisineHealthFoods
	HotelFeatureContinentalBreakfastLuncheon
	HotelFeatureBreakfastBuffet
	HotelFeatureOutdoorPool
	HotelFeatureIndoorPool
	HotelFeatureBar
	HotelFeatureBarrierFree
	HotelFeatureWlan
	HotelFeatureShuttleService
	HotelFeatureChildcare
	HotelFeatureSmallPetsAllowed
	HotelFeatureBeautyFarm
	HotelFeatureCentralLocation HotelFeature = 2 << iota
	HotelFeatureCoveredParking
	HotelFeatureOpenParking
	HotelFeatureMassages
	HotelFeatureSauna
	HotelFeatureSteamBath
	HotelFeaturePublicBar
	HotelFeatureDogsAllowed
)

var HotelFeatures = []HotelFeature{
	HotelFeatureGarage,
	HotelFeatureElevatorLift,
	HotelFeatureRestaurant,
	HotelFeatureGym,
	HotelFeatureWellness,
	HotelFeatureSpaCuisineHealthFoods,
	HotelFeatureContinentalBreakfastLuncheon,
	HotelFeatureBreakfastBuffet,
	HotelFeatureOutdoorPool,
	HotelFeatureIndoorPool,
	HotelFeatureBar,
	HotelFeatureBarrierFree,
	HotelFeatureWlan,
	HotelFeatureShuttleService,
	HotelFeatureChildcare,
	HotelFeatureSmallPetsAllowed,
	HotelFeatureBeautyFarm,
	HotelFeatureCentralLocation,
	HotelFeatureCoveredParking,
	HotelFeatureOpenParking,
	HotelFeatureMassages,
	HotelFeatureSauna,
	HotelFeatureSteamBath,
	HotelFeaturePublicBar,
	HotelFeatureDogsAllowed,
}

type Theme int

const (
	ThemeFamily Theme = 1 << iota
	ThemeWellness
	ThemeHiking
	ThemeMotorcycle
	ThemeBike
	ThemeGolf
	ThemeRiding
	ThemeRomantic
	ThemeSki
	ThemeMeeting
	ThemeCrossCountrySkiing
	ThemeCulture
	ThemeSnowshoeing
)

var Themes = []Theme{
	ThemeFamily,
	ThemeWellness,
	ThemeHiking,
	ThemeMotorcycle,
	ThemeBike,
	ThemeGolf,
	ThemeRiding,
	ThemeRomantic,
	ThemeSki,
	ThemeMeeting,
	ThemeCrossCountrySkiing,
	ThemeCulture,
	ThemeSnowshoeing,
}

type Board int

const (
	BoardAny Board = iota
	BoardWithoutBoard
	BoardWithBreakfast
	BoardHalfBoard
	BoardFullBoard
	BoardAllInclusive
)

var Boards = []Board{
	BoardAny,
	BoardWithoutBoard,
	BoardWithBreakfast,
	BoardHalfBoard,
	BoardFullBoard,
	BoardAllInclusive,
}

type OfferType int

const (
	OfferTypePriceListStandardPrice       OfferType = 10
	OfferTypeBasedOnAgeOfPeople           OfferType = 20
	OfferTypeBasedOnNumberOfPeople        OfferType = 21
	OfferTypeBasedOnStaying               OfferType = 22
	OfferTypeBasedOnDateOfBooking         OfferType = 23
	OfferTypeBasedOnWeekday               OfferType = 24
	OfferTypeNoReference                  OfferType = 25
	OfferTypeSpecialBasedOnAgeOfPeople    OfferType = 50
	OfferTypeSpecialBasedOnNumberOfPeople OfferType = 51
	OfferTypeSpecialBasedOnStaying        OfferType = 52
	OfferTypeSpecialBasedOnDateOfBooking  OfferType = 53
	OfferTypeSpecialBasedOnWeekday        OfferType = 54
	OfferTypeSpecialNoReference           OfferType = 55
)

var OfferTypes = []OfferType{
	OfferTypePriceListStandardPrice,
	OfferTypeBasedOnAgeOfPeople,
	OfferTypeBasedOnNumberOfPeople,
	OfferTypeBasedOnStaying,
	OfferTypeBasedOnDateOfBooking,
	OfferTypeBasedOnWeekday,
	OfferTypeNoReference,
	OfferTypeSpecialBasedOnAgeOfPeople,
	OfferTypeSpecialBasedOnNumberOfPeople,
	OfferTypeSpecialBasedOnStaying,
	OfferTypeSpecialBasedOnDateOfBooking,
	OfferTypeSpecialBasedOnWeekday,
	OfferTypeSpecialNoReference,
}

type SpecialType int

const (
	SpecialTypePackages SpecialType = 1 << iota
	SpecialTypeSpecials
	SpecialTypeMasterpackages
)

var SpecialTypes = []SpecialType{
	SpecialTypePackages,
	SpecialTypeSpecials,
	SpecialTypeMasterpackages,
}

type SpecialPremium int

const (
	SpecialPremiumVitalpina SpecialPremium = 1 << iota
	SpecialPremiumFamilyHotelsPremium
	SpecialPremiumVinumHotelsPremium
	SpecialPremiumSüdtirolBalancePremium
	SpecialPremiumVitalpinaDurchatmen
	SpecialPremiumVitalpinaWohlfühlen
	SpecialPremiumVitaplinaErnährung
	SpecialPremiumVitaplinaAktiv
	SpecialPremiumVitalpinaPremium
	SpecialPremiumBikehotelsMountainbike
	SpecialPremiumBikehotelsBikeTouringAndEbike
	SpecialPremiumBikehotelsRoadbike
	SpecialPremiumBikehotelsPremium
	SpecialPremiumArchitectureDays
	SpecialPremiumVinumHotels
	SpecialPremiumFamilienHotels
	SpecialPremiumFamilienHotelsNaturdetektivSommer
	SpecialPremiumFamilienHotelsNaturdetektivWinter
)

var SpecialPremiums = []SpecialPremium{
	SpecialPremiumVitalpina,
	SpecialPremiumFamilyHotelsPremium,
	SpecialPremiumVinumHotelsPremium,
	SpecialPremiumSüdtirolBalancePremium,
	SpecialPremiumVitalpinaDurchatmen,
	SpecialPremiumVitalpinaWohlfühlen,
	SpecialPremiumVitaplinaErnährung,
	SpecialPremiumVitaplinaAktiv,
	SpecialPremiumVitalpinaPremium,
	SpecialPremiumBikehotelsMountainbike,
	SpecialPremiumBikehotelsBikeTouringAndEbike,
	SpecialPremiumBikehotelsRoadbike,
	SpecialPremiumBikehotelsPremium,
	SpecialPremiumArchitectureDays,
	SpecialPremiumVinumHotels,
	SpecialPremiumFamilienHotels,
	SpecialPremiumFamilienHotelsNaturdetektivSommer,
	SpecialPremiumFamilienHotelsNaturdetektivWinter,
}

type ThemeID int

const (
	ThemeIDHiking ThemeID = iota + 1
	ThemeIDCyclingMountainbike
	ThemeIDFamily
	ThemeIDWellnessHealth
	ThemeIDFoodAndDrink
	ThemeIDGolf
	ThemeIDCulture
	ThemeIDMotorsport
	ThemeIDCarFreeHolidays
	ThemeIDSkiSnowboard
	ThemeIDSummerActivities
	ThemeIDEvents
	ThemeIDChristmasMarkets
	ThemeIDActiveWinter
	ThemeIDVitalpina
	ThemeIDVitalpinaBreathe
	ThemeIDBikeHotelsEBike
	ThemeIDBikeHotelsFreeride
	ThemeIDBikeHotelsMountainbike ThemeID = iota + 2
	ThemeIDBikeHotelsBikeTours
	ThemeIDBikeHotelsRacingBike
	ThemeIDFamilyHotels
	ThemeIDFamilyHotelsNatureDetective
	ThemeIDFamilyHotel ThemeID = iota + 3
	ThemeIDNatureDetectivSummer
	ThemeIDNatureDetectivWinter
	ThemeIDEcologicHoliday ThemeID = iota + 53
	ThemeIDHorseBackRiding
	ThemeIDLuxuryHoliday
	ThemeIDPetsFriendlyHoliday
	ThemeIDRoadBike
	ThemeIDRomanticHoliday
	ThemeIDWine
	ThemeIDBicycleTouring
	ThemeIDEBike
)

var ThemeIDs = []ThemeID{
	ThemeIDHiking,
	ThemeIDCyclingMountainbike,
	ThemeIDFamily,
	ThemeIDWellnessHealth,
	ThemeIDFoodAndDrink,
	ThemeIDGolf,
	ThemeIDCulture,
	ThemeIDMotorsport,
	ThemeIDCarFreeHolidays,
	ThemeIDSkiSnowboard,
	ThemeIDSummerActivities,
	ThemeIDEvents,
	ThemeIDChristmasMarkets,
	ThemeIDActiveWinter,
	ThemeIDVitalpina,
	ThemeIDVitalpinaBreathe,
	ThemeIDBikeHotelsEBike,
	ThemeIDBikeHotelsFreeride,
	ThemeIDBikeHotelsMountainbike,
	ThemeIDBikeHotelsBikeTours,
	ThemeIDBikeHotelsRacingBike,
	ThemeIDFamilyHotels,
	ThemeIDFamilyHotelsNatureDetective,
	ThemeIDFamilyHotel,
	ThemeIDNatureDetectivSummer,
	ThemeIDNatureDetectivWinter,
	ThemeIDEcologicHoliday,
	ThemeIDHorseBackRiding,
	ThemeIDLuxuryHoliday,
	ThemeIDPetsFriendlyHoliday,
	ThemeIDRoadBike,
	ThemeIDRomanticHoliday,
	ThemeIDWine,
	ThemeIDBicycleTouring,
	ThemeIDEBike,
}

type LocationType string

const (
	// Area = an Italian province or region (e.g. South Tyrol)
	LocationTypeArea LocationType = "ara"
	// Region is an area
	LocationTypeRegion LocationType = "reg"
	// Community = municipality
	LocationTypeCommunity LocationType = "com"
	// Location is a subdivision of a municipality
	LocationTypeLocation LocationType = "cit"
	// VirtualLocation = arbitrary group of one or more locations
	LocationTypeVirtualLocation LocationType = "vir"
)

var LocationTypes = []LocationType{
	LocationTypeArea,
	LocationTypeRegion,
	LocationTypeCommunity,
	LocationTypeLocation,
	LocationTypeVirtualLocation,
}

type PaymentMethod int

const (
	PaymentMethodDepositByCreditCard PaymentMethod = 1 << iota
	PaymentMethodCreditCardAsSecurity
	PaymentMethodDepositByBankTransfer
	PaymentMethodPaymentByCreditCard
	PaymentMethodPaymentByBankTransfer
	PaymentMethodPaymentAtTheHotel
)

var PaymentMethods = []PaymentMethod{
	PaymentMethodDepositByCreditCard,
	PaymentMethodCreditCardAsSecurity,
	PaymentMethodDepositByBankTransfer,
	PaymentMethodPaymentByCreditCard,
	PaymentMethodPaymentByBankTransfer,
	PaymentMethodPaymentAtTheHotel,
}

type Gender string

const (
	GenderUnknown Gender = ""
	GenderMale    Gender = "m"
	GenderFemale  Gender = "f"
)

var Genders = []Gender{
	GenderUnknown,
	GenderMale,
	GenderFemale,
}

type CouponType string

const (
	CouponTypeVoucher   CouponType = "voucher"
	CouponTypePromotion CouponType = "promotion"
)

var CouponTypes = []CouponType{
	CouponTypeVoucher,
	CouponTypePromotion,
}
