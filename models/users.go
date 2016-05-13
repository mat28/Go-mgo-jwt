package models

import (
	"gopkg.in/mgo.v2/bson"
)


const (
	CollectionPeople = "peoples"
)


type PEOPLESortant struct {
	Id bson.ObjectId `bson:"_id,omitempty"`
	EmailAddress string
	Password string

}

type People struct {
	Id          bson.ObjectId `json:"_id,omitempty"bson:"_id,omitempty"`
	EmailAddress string `json:"emailAddress,omitempty"`
	Password     string `json:"password,omitempty"`
	FirstName    string `json:"firstName,omitempty"`
	Headline     string `json:"headline,omitempty"`
	Industry     string `json:"industry,omitempty"`
	LastName     string `json:"lastName,omitempty"`
	Location     struct {
			     Country struct {
					     Code string `json:"code,omitempty"`
				     } `json:"country,omitempty"`
			     Name string `json:"name,omitempty"`
	} `json:"location,omitempty"`
	NumConnections int `json:"numConnections,omitempty"`
	PictureUrls    struct {
			     _total int      `json:"_total,omitempty"`
			     Values []string `json:"values,omitempty"`
	} `json:"pictureUrls,omitempty"`
	Positions struct {
			     _total int `json:"_total,omitempty"`
			     Values []struct {
				     Company struct {
						     LinkedinId int    `json:"id,omitempty"`
						     Industry   string `json:"industry,omitempty"`
						     Name       string `json:"name,omitempty"`
						     Size       string `json:"size,omitempty"`
						     Type       string `json:"type,omitempty"`
					     } `json:"company,omitempty"`
				     ID        int  `json:"id,omitempty"`
				     IsCurrent bool `json:"isCurrent,omitempty"`
				     Location  struct {
						     Country struct {
								     Code string `json:"code,omitempty"`
								     Name string `json:"name,omitempty"`
							     } `json:"country,omitempty"`
						     Name string `json:"name,omitempty"`
					     } `json:"location,omitempty"`
				     StartDate struct {
						     Month int `json:"month,omitempty"`
						     Year  int `json:"year,omitempty"`
					     } `json:"startDate,omitempty"`
				     Summary string `json:"summary,omitempty"`
				     Title   string `json:"title,omitempty"`
			     } `json:"values,omitempty"`
	} `json:"positions,omitempty"`
	SiteStandardProfileRequest struct {
			     URL string `json:"url,omitempty"`
	} `json:"siteStandardProfileRequest,omitempty"`
	Newsletter bool `json:"newsletter,omitempty"`
	NbApplies int `json:"nbapplies,omitempty"`
	RoleName string `json:"rolename,omitempty"`
	RoleScope string `json:"rolescope,omitempty"`
}

type Ids struct {
	Candidat int `json:"candidat,omitempty"`
	Admin int `json:"admin,omitempty"`
	Prospect int `json:"prospect,omitempty"`
}

type ConfirmationUser struct {
	Date string `json:"date,omitempty"`
	IpAddress string `json:"ip,omitempty"`
}

type TokenActivation struct {
	Token string `json:"token,omitempty"`
}

var SettingsPeople = []string{"emailAddress","userType","creation.Date","newsletter","birthday"}

var CryptoPeople = []string{"Birthday","FirstName","LastName","PhoneNumber","IpAddress"}

