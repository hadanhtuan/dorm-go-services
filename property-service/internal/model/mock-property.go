package model

import (
	"time"

	"github.com/hadanhtuan/go-sdk/db/orm"
	"gorm.io/gorm"
)

// use for crawl property from airbnb
type MockProperty struct {
	ID        string     `json:"id" gorm:"default:gen_random_uuid()"`
	CreatedAt time.Time  `json:"createdAt,omitempty"`
	UpdatedAt time.Time  `json:"updatedAt,omitempty"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" gorm:"index"`

	ListingUrl                              string  `gorm:"column:listing_url"`
	ScrapeId                                int64   `gorm:"column:scrape_id"`
	LastScraped                             string  `gorm:"column:last_scraped"`
	Source                                  string  `gorm:"column:source"`
	Name                                    string  `gorm:"column:name"`
	Description                             string  `gorm:"column:description"`
	Neighborhood_overview                   string  `gorm:"column:neighborhood_overview"`
	PictureUrl                              string  `gorm:"column:picture_url"`
	HostId                                  int64   `gorm:"column:host_id"`
	HostUrl                                 string  `gorm:"column:host_url"`
	HostName                                string  `gorm:"column:host_name"`
	HostSince                               string  `gorm:"column:host_since"`
	HostLocation                            string  `gorm:"column:host_location"`
	HostAbout                               string  `gorm:"column:host_about"`
	HostResponseTime                        string  `gorm:"column:host_response_time"`
	HostResponseRate                        string  `gorm:"column:host_response_rate"`
	HostAcceptanceRate                      string  `gorm:"column:host_acceptance_rate"`
	HostIsSuperHost                         string  `gorm:"column:host_is_superhost"`
	HostThumbnailUrl                        string  `gorm:"column:host_thumbnail_url"`
	HostPictureUrl                          string  `gorm:"column:host_picture_url"`
	HostNeighbourhood                       string  `gorm:"column:host_neighbourhood"`
	HostListingsCount                       int64   `gorm:"column:host_listings_count"`
	HostTotalListingsCount                  int64   `gorm:"column:host_total_listings_count"`
	HostVerifications                       string  `gorm:"column:host_verifications"`
	HostHasProfilePic                       string  `gorm:"column:host_has_profile_pic"`
	HostIdentityVerified                    string  `gorm:"column:host_identity_verified"`
	Neighbourhood                           string  `gorm:"column:neighbourhood"`
	NeighbourhoodCleansed                   string  `gorm:"column:neighbourhood_cleansed"`
	NeighbourhoodGroupCleansed              string  `gorm:"column:neighbourhood_group_cleansed"`
	Latitude                                float64 `gorm:"column:latitude"`
	Longitude                               float64 `gorm:"column:longitude"`
	PropertyType                            string  `gorm:"column:property_type"`
	RoomType                                string  `gorm:"column:room_type"`
	Accommodates                            int64   `gorm:"column:accommodates"`
	Bathrooms                               float64 `gorm:"column:bathrooms"`
	BathroomsText                           string  `gorm:"column:bathrooms_text"`
	Bedrooms                                int64   `gorm:"column:bedrooms"`
	Beds                                    int64   `gorm:"column:beds"`
	Amenities                               any
	Price                                   string  `gorm:"column:price"`
	MinimumNights                           int64   `gorm:"column:minimum_nights"`
	MaximumNights                           int64   `gorm:"column:maximum_nights"`
	MinimumMinimumNights                    int64   `gorm:"column:minimum_minimum_nights"`
	MaximumMinimumNights                    int64   `gorm:"column:maximum_minimum_nights"`
	MinimumMaximumNights                    int64   `gorm:"column:minimum_maximum_nights"`
	MaximumMaximumNights                    int64   `gorm:"column:maximum_maximum_nights"`
	MinimumNightsAvgNtm                     float64 `gorm:"column:minimum_nights_avg_ntm"`
	MaximumNightsAvgNtm                     float64 `gorm:"column:maximum_nights_avg_ntm"`
	CalendarUpdated                         string  `gorm:"column:calendar_updated"`
	HasAvailability                         string  `gorm:"column:has_availability"`
	Availability30                          int64   `gorm:"column:availability_30"`
	Availability60                          int64   `gorm:"column:availability_60"`
	Availability90                          int64   `gorm:"column:availability_90"`
	Availability365                         int64   `gorm:"column:availability_365"`
	Calendar_last_scraped                   string  `gorm:"column:calendar_last_scraped"`
	NumberOfReviews                         int64   `gorm:"column:number_of_reviews"`
	NumberOfReviewsLtm                      int64   `gorm:"column:number_of_reviews_ltm"`
	NumberOfReviewsL30d                     int64   `gorm:"column:number_of_reviews_l30d"`
	FirstReview                             string  `gorm:"column:first_review"`
	LastReview                              string  `gorm:"column:last_review"`
	ReviewScoresRating                      float64 `gorm:"column:review_scores_rating"`
	ReviewScoresAccuracy                    float64 `gorm:"column:review_scores_accuracy"`
	ReviewScoresCleanliness                 float64 `gorm:"column:review_scores_cleanliness"`
	ReviewScoresCheckin                     float64 `gorm:"column:review_scores_checkin"`
	ReviewScoresCommunication               float64 `gorm:"column:review_scores_communication"`
	ReviewScoresLocation                    float64 `gorm:"column:review_scores_location"`
	ReviewScoresValue                       float64 `gorm:"column:review_scores_value"`
	License                                 string  `gorm:"column:license"`
	InstantBookable                         string  `gorm:"column:instant_bookable"`
	CalculatedHostListingsCount             int64   `gorm:"column:calculated_host_listings_count"`
	CalculatedHostListingsCountEntireHomes  int64   `gorm:"column:calculated_host_listings_count_entire_homes"`
	CalculatedHostListingsCountPrivateRooms int64   `gorm:"column:calculated_host_listings_count_private_rooms"`
	CalculatedHostListingsCountSharedRooms  int64   `gorm:"column:calculated_host_listings_count_shared_rooms"`
	ReviewsPerMonth                         float64 `gorm:"column:reviews_per_month"`
}

func (MockProperty) TableName() string {
	return "mock_property"
}

var MockPropertyDB = &orm.Instance{
	TableName: "mock_property",
	Model:     &MockProperty{},
}

func InitTableMockProperty(db *gorm.DB) {
	db.AutoMigrate(&MockProperty{})
	MockPropertyDB.ApplyDatabase(db)
}
