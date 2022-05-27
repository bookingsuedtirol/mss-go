package shared

import (
	"encoding/xml"
	"time"
)

type Date time.Time

func (input Date) MarshalXML(element *xml.Encoder, start xml.StartElement) error {
	timeValue := time.Time(input)
	timeString := timeValue.Format("2006-01-02")
	return element.EncodeElement(timeString, start)
}

func (input *Date) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var str string
	err := decoder.DecodeElement(&str, &start)
	if err != nil {
		return err
	}

	value, err := ParseDateTime("2006-01-02", str)
	if err != nil {
		return err
	}

	if value != nil {
		*input = Date(*value)
	}

	return nil
}

func ParseDateTime(layout string, value string) (*time.Time, error) {
	if value == "" || value == "0000-00-00" {
		return nil, nil
	}

	parsed, err := time.Parse(layout, value)
	if err != nil {
		return nil, err
	}

	return &parsed, nil
}

func (date Date) String() string {
	return time.Time(date).String()
}

type LTSData struct {
	A0Ene int `xml:"A0Ene"`
	A0MTV int `xml:"A0MTV"`
	A0Rep int `xml:"A0Rep"`
}

type RoomType int

const (
	RoomTypeRoom RoomType = 1 << iota
	RoomTypeApartment
	RoomTypeCampingPitch
)

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

type Board int

const (
	BoardWithoutBoard Board = iota + 1
	BoardWithBreakfast
	BoardHalfBoard
	BoardFullBoard
	BoardAllInclusive
)

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

type SpecialType int

const (
	SpecialTypePackages SpecialType = 1 << iota
	SpecialTypeSpecials
	SpecialTypeMasterpackages
)

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
)

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

type PaymentMethod int

const (
	PaymentMethodDepositByCreditCard PaymentMethod = 1 << iota
	PaymentMethodCreditCardAsSecurity
	PaymentMethodDepositByBankTransfer
	PaymentMethodPaymentByCreditCard
	PaymentMethodPaymentByBankTransfer
	PaymentMethodPaymentAtTheHotel
)

type SEOType string

const (
	SEOTypeGeneral  SEOType = "gen"
	SEOTypeLocation SEOType = "loc"
	SEOTypeTheme    SEOType = "the"
)

type Gender string

const (
	GenderUnknown Gender = ""
	GenderMale    Gender = "m"
	GenderFemale  Gender = "f"
)
