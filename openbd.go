package openbd

import (
	"encoding/json"
	"net/http"
	"time"
)

const endpoint = "https://api.openbd.jp/v1/"

type OpenBD struct {
	Onix struct {
		RecordReference   string `json:"RecordReference"`
		ProductIdentifier struct {
			IDValue       string `json:"IDValue"`
			ProductIDType string `json:"ProductIDType"`
		} `json:"ProductIdentifier"`
		DescriptiveDetail struct {
			ProductComposition string `json:"ProductComposition"`
			ProductForm        string `json:"ProductForm"`
			ProductFormDetail  string `json:"ProductFormDetail"`
			Measure            []struct {
				MeasureType     string `json:"MeasureType"`
				MeasureUnitCode string `json:"MeasureUnitCode"`
				Measurement     string `json:"Measurement"`
			} `json:"Measure"`
			Collection struct {
				CollectionType string `json:"CollectionType"`
				TitleDetail    struct {
					TitleType    string `json:"TitleType"`
					TitleElement []struct {
						TitleElementLevel string `json:"TitleElementLevel"`
						TitleText         struct {
							Content      string `json:"content"`
							Collationkey string `json:"collationkey"`
						} `json:"TitleText"`
					} `json:"TitleElement"`
				} `json:"TitleDetail"`
			} `json:"Collection"`
			TitleDetail struct {
				TitleType    string `json:"TitleType"`
				TitleElement struct {
					TitleElementLevel string `json:"TitleElementLevel"`
					TitleText         struct {
						Content      string `json:"content"`
						Collationkey string `json:"collationkey"`
					} `json:"TitleText"`
				} `json:"TitleElement"`
			} `json:"TitleDetail"`
			Contributor []struct {
				ContributorRole []string `json:"ContributorRole"`
				PersonName      struct {
					Content      string `json:"content"`
					Collationkey string `json:"collationkey"`
				} `json:"PersonName"`
				BiographicalNote string `json:"BiographicalNote"`
				SequenceNumber   string `json:"SequenceNumber"`
			} `json:"Contributor"`
			Language []struct {
				LanguageCode string `json:"LanguageCode"`
				LanguageRole string `json:"LanguageRole"`
				CountryCode  string `json:"CountryCode"`
			} `json:"Language"`
			Audience []struct {
				AudienceCodeType  string `json:"AudienceCodeType"`
				AudienceCodeValue string `json:"AudienceCodeValue"`
			} `json:"Audience"`
			Extent []struct {
				ExtentValue string `json:"ExtentValue"`
				ExtentUnit  string `json:"ExtentUnit"`
				ExtentType  string `json:"ExtentType"`
			} `json:"Extent"`
			Subject []struct {
				SubjectSchemeIdentifier string `json:"SubjectSchemeIdentifier"`
				SubjectCode             string `json:"SubjectCode"`
			} `json:"Subject"`
		} `json:"DescriptiveDetail"`
		CollateralDetail struct {
			TextContent []struct {
				Text            string `json:"Text"`
				TextType        string `json:"TextType"`
				ContentAudience string `json:"ContentAudience"`
			} `json:"TextContent"`
			SupportingResource []struct {
				ResourceContentType string `json:"ResourceContentType"`
				ResourceMode        string `json:"ResourceMode"`
				ContentAudience     string `json:"ContentAudience"`
				ResourceVersion     []struct {
					ResourceLink           string `json:"ResourceLink"`
					ResourceForm           string `json:"ResourceForm"`
					ResourceVersionFeature []struct {
						ResourceVersionFeatureType string `json:"ResourceVersionFeatureType"`
						FeatureValue               string `json:"FeatureValue"`
					} `json:"ResourceVersionFeature"`
				} `json:"ResourceVersion"`
			} `json:"SupportingResource"`
		} `json:"CollateralDetail"`
		PublishingDetail struct {
			Imprint struct {
				ImprintName       string `json:"ImprintName"`
				ImprintIdentifier []struct {
					IDValue       string `json:"IDValue"`
					ImprintIDType string `json:"ImprintIDType"`
				} `json:"ImprintIdentifier"`
			} `json:"Imprint"`
			Publisher struct {
				PublisherIdentifier []struct {
					PublisherIDType string `json:"PublisherIDType"`
					IDValue         string `json:"IDValue"`
				} `json:"PublisherIdentifier"`
				PublisherName  string `json:"PublisherName"`
				PublishingRole string `json:"PublishingRole"`
			} `json:"Publisher"`
			PublishingDate []struct {
				Date               string `json:"Date"`
				PublishingDateRole string `json:"PublishingDateRole"`
			} `json:"PublishingDate"`
		} `json:"PublishingDetail"`
		ProductSupply struct {
			SupplyDetail struct {
				ReturnsConditions struct {
					ReturnsCode     string `json:"ReturnsCode"`
					ReturnsCodeType string `json:"ReturnsCodeType"`
				} `json:"ReturnsConditions"`
				ProductAvailability string `json:"ProductAvailability"`
				Price               []struct {
					PriceAmount  string `json:"PriceAmount"`
					CurrencyCode string `json:"CurrencyCode"`
					PriceType    string `json:"PriceType"`
				} `json:"Price"`
			} `json:"SupplyDetail"`
		} `json:"ProductSupply"`
		NotificationType string `json:"NotificationType"`
	} `json:"onix"`
	Hanmoto struct {
		Reviews []struct {
			Appearance string `json:"appearance"`
			Reviewer   string `json:"reviewer"`
			SourceID   int    `json:"source_id"`
			KubunID    int    `json:"kubun_id"`
			Source     string `json:"source"`
			Choyukan   string `json:"choyukan"`
			Han        string `json:"han"`
			Link       string `json:"link"`
			PostUser   string `json:"post_user"`
		} `json:"reviews"`
		Dateshuppan  string `json:"dateshuppan"`
		Datemodified string `json:"datemodified"`
		Datecreated  string `json:"datecreated"`
	} `json:"hanmoto"`
	Summary struct {
		Title     string `json:"title"`
		Author    string `json:"author"`
		Publisher string `json:"publisher"`
		Series    string `json:"series"`
		Volume    string `json:"volume"`
		Pubdate   string `json:"pubdate"`
		Isbn      string `json:"isbn"`
		Cover     string `json:"cover"`
	} `json:"summary"`
}

func GetOpenBD(isbn string, timeout time.Duration) ([]OpenBD, error) {
	client := &http.Client{
		Timeout: timeout,
	}
	url := endpoint + "get?isbn=" + isbn
	r, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	rtn := make([]OpenBD, 0)
	json.NewDecoder(r.Body).Decode(&rtn)
	return rtn, nil
}
