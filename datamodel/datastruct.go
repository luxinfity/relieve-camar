package datamodel

import (
	"github.com/globalsign/mgo/bson"
	"time"
)

type Response struct {
	Status int `json:"status"`
	Data interface{} `json:"data"`
}

type EarthquakeDataSnapshoot struct {
	Title    string      `json:"title"`
	URL      string      `json:"url"`
	Location struct {
		Type        string    `json:"type"`
		Coordinates []float64 `json:"coordinates"`
	} `json:"location"`
	Mag      float64     `json:"mag"`
	Depth            float64    `json:"depth"`
	Place    string      `json:"place"`
	Time     int64       `json:"time"`
	Tsunami  int         `json:"tsunami"`
}
type CountryData struct {
	Languages   string `json:"languages"`
	Distance    string `json:"distance"`
	CountryCode string `json:"countryCode"`
	CountryName string `json:"countryName"`
} 
// GeoJSONEarthquakeData is the main struct to wrap data from GeoJSON endpoint.
type GeoJSON struct {
	BsonID     bson.ObjectId `bson:"_id" json:"bson_id"`
	Type       string        `json:"type"`
	Properties struct {
		Mag      float64     `json:"mag"`
		Place    string      `json:"place"`
		Time     int64       `json:"time"`
		Updated  int64       `json:"updated"`
		Tz       int         `json:"tz"`
		URL      string      `json:"url"`
		Felt     int         `json:"felt"`
		Cdi      float64     `json:"cdi"`
		Mmi      float64     `json:"mmi"`
		Alert    string      `json:"alert"`
		Status   string      `json:"status"`
		Tsunami  int         `json:"tsunami"`
		Sig      int         `json:"sig"`
		Net      string      `json:"net"`
		Code     string      `json:"code"`
		Ids      string      `json:"ids"`
		Sources  string      `json:"sources"`
		Types    string      `json:"types"`
		Nst      interface{} `json:"nst"`
		Dmin     float64     `json:"dmin"`
		Rms      float64     `json:"rms"`
		Gap      int         `json:"gap"`
		MagType  string      `json:"magType"`
		Type     string      `json:"type"`
		Title    string      `json:"title"`
		Products struct {
			Dyfi []struct {
				Indexid    string `json:"indexid"`
				IndexTime  int64  `json:"indexTime"`
				ID         string `json:"id"`
				Type       string `json:"type"`
				Code       string `json:"code"`
				Source     string `json:"source"`
				UpdateTime int64  `json:"updateTime"`
				Status     string `json:"status"`
				Properties struct {
					Depth            string    `json:"depth"`
					Eventsource      string    `json:"eventsource"`
					Eventsourcecode  string    `json:"eventsourcecode"`
					Eventtime        time.Time `json:"eventtime"`
					Latitude         string    `json:"latitude"`
					Longitude        string    `json:"longitude"`
					Magnitude        string    `json:"magnitude"`
					Maxmmi           string    `json:"maxmmi"`
					NumResponses     string    `json:"num-responses"`
					NumResp          string    `json:"numResp"`
					PdlClientVersion string    `json:"pdl-client-version"`
				} `json:"properties"`
				PreferredWeight int `json:"preferredWeight"`
				Contents        struct {
					CdiGeoTxt struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"cdi_geo.txt"`
					CdiGeoXML struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"cdi_geo.xml"`
					CdiGeo1KmTxt struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"cdi_geo_1km.txt"`
					CdiZipTxt struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"cdi_zip.txt"`
					CdiZipXML struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"cdi_zip.xml"`
					ContentsXML struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"contents.xml"`
					DyfiKmz struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"dyfi.kmz"`
					DyfiGeoKmz struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"dyfi_geo.kmz"`
					DyfiGeo10KmGeojson struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"dyfi_geo_10km.geojson"`
					DyfiGeo1KmGeojson struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"dyfi_geo_1km.geojson"`
					DyfiPlotAttenJSON struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"dyfi_plot_atten.json"`
					DyfiPlotNumrespJSON struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"dyfi_plot_numresp.json"`
					DyfiZipGeojson struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"dyfi_zip.geojson"`
					DyfiZipKmz struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"dyfi_zip.kmz"`
					EventDataXML struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"event_data.xml"`
					Us2000Ha1KCiimJpg struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"us2000ha1k_ciim.jpg"`
					Us2000Ha1KCiimPdf struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"us2000ha1k_ciim.pdf"`
					Us2000Ha1KCiimPs struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"us2000ha1k_ciim.ps"`
					Us2000Ha1KCiimGeoJpg struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"us2000ha1k_ciim_geo.jpg"`
					Us2000Ha1KCiimGeoPdf struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"us2000ha1k_ciim_geo.pdf"`
					Us2000Ha1KCiimGeoPs struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"us2000ha1k_ciim_geo.ps"`
					Us2000Ha1KCiimGeoImapHTML struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"us2000ha1k_ciim_geo_imap.html"`
					Us2000Ha1KCiimImapHTML struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"us2000ha1k_ciim_imap.html"`
					Us2000Ha1KPlotAttenJpg struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"us2000ha1k_plot_atten.jpg"`
					Us2000Ha1KPlotAttenPs struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"us2000ha1k_plot_atten.ps"`
					Us2000Ha1KPlotAttenTxt struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"us2000ha1k_plot_atten.txt"`
					Us2000Ha1KPlotNumrespJpg struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"us2000ha1k_plot_numresp.jpg"`
					Us2000Ha1KPlotNumrespPs struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"us2000ha1k_plot_numresp.ps"`
					Us2000Ha1KPlotNumrespTxt struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"us2000ha1k_plot_numresp.txt"`
				} `json:"contents"`
			} `json:"dyfi"`
			Geoserve []struct {
				Indexid    string `json:"indexid"`
				IndexTime  int64  `json:"indexTime"`
				ID         string `json:"id"`
				Type       string `json:"type"`
				Code       string `json:"code"`
				Source     string `json:"source"`
				UpdateTime int64  `json:"updateTime"`
				Status     string `json:"status"`
				Properties struct {
					Eventsource      string `json:"eventsource"`
					Eventsourcecode  string `json:"eventsourcecode"`
					Location         string `json:"location"`
					PdlClientVersion string `json:"pdl-client-version"`
					TsunamiFlag      string `json:"tsunamiFlag"`
					UtcOffset        string `json:"utcOffset"`
				} `json:"properties"`
				PreferredWeight int `json:"preferredWeight"`
				Contents        struct {
					GeoserveJSON struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"geoserve.json"`
				} `json:"contents"`
			} `json:"geoserve"`
			GroundFailure []struct {
				Indexid    string `json:"indexid"`
				IndexTime  int64  `json:"indexTime"`
				ID         string `json:"id"`
				Type       string `json:"type"`
				Code       string `json:"code"`
				Source     string `json:"source"`
				UpdateTime int64  `json:"updateTime"`
				Status     string `json:"status"`
				Properties struct {
					Depth                                string    `json:"depth"`
					Eventsource                          string    `json:"eventsource"`
					Eventsourcecode                      string    `json:"eventsourcecode"`
					Eventtime                            time.Time `json:"eventtime"`
					LandslideAlert                       string    `json:"landslide-alert"`
					LandslideHazardAlertColor            string    `json:"landslide-hazard-alert-color"`
					LandslideHazardAlertParameter        string    `json:"landslide-hazard-alert-parameter"`
					LandslideHazardAlertValue            string    `json:"landslide-hazard-alert-value"`
					LandslideMaximumLatitude             string    `json:"landslide-maximum-latitude"`
					LandslideMaximumLongitude            string    `json:"landslide-maximum-longitude"`
					LandslideMinimumLatitude             string    `json:"landslide-minimum-latitude"`
					LandslideMinimumLongitude            string    `json:"landslide-minimum-longitude"`
					LandslideOverlay                     string    `json:"landslide-overlay"`
					LandslidePopulationAlertColor        string    `json:"landslide-population-alert-color"`
					LandslidePopulationAlertParameter    string    `json:"landslide-population-alert-parameter"`
					LandslidePopulationAlertValue        string    `json:"landslide-population-alert-value"`
					Latitude                             string    `json:"latitude"`
					LiquefactionAlert                    string    `json:"liquefaction-alert"`
					LiquefactionHazardAlertColor         string    `json:"liquefaction-hazard-alert-color"`
					LiquefactionHazardAlertParameter     string    `json:"liquefaction-hazard-alert-parameter"`
					LiquefactionHazardAlertValue         string    `json:"liquefaction-hazard-alert-value"`
					LiquefactionMaximumLatitude          string    `json:"liquefaction-maximum-latitude"`
					LiquefactionMaximumLongitude         string    `json:"liquefaction-maximum-longitude"`
					LiquefactionMinimumLatitude          string    `json:"liquefaction-minimum-latitude"`
					LiquefactionMinimumLongitude         string    `json:"liquefaction-minimum-longitude"`
					LiquefactionOverlay                  string    `json:"liquefaction-overlay"`
					LiquefactionPopulationAlertColor     string    `json:"liquefaction-population-alert-color"`
					LiquefactionPopulationAlertParameter string    `json:"liquefaction-population-alert-parameter"`
					LiquefactionPopulationAlertValue     string    `json:"liquefaction-population-alert-value"`
					Longitude                            string    `json:"longitude"`
					Magnitude                            string    `json:"magnitude"`
					MaximumLatitude                      string    `json:"maximum-latitude"`
					MaximumLongitude                     string    `json:"maximum-longitude"`
					MinimumLatitude                      string    `json:"minimum-latitude"`
					MinimumLongitude                     string    `json:"minimum-longitude"`
					PdlClientVersion                     string    `json:"pdl-client-version"`
					RuptureWarning                       string    `json:"rupture-warning"`
					ShakemapVersion                      string    `json:"shakemap-version"`
				} `json:"properties"`
				PreferredWeight int `json:"preferredWeight"`
				Contents        struct {
					ContentsXML struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"contents.xml"`
					Godt2008Hdf5 struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"godt_2008.hdf5"`
					Godt2008Png struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"godt_2008.png"`
					Godt2008ModelFlt struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"godt_2008_model.flt"`
					Godt2008ModelHdr struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"godt_2008_model.hdr"`
					Godt2008ModelTif struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"godt_2008_model.tif"`
					InfoJSON struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"info.json"`
					Jessee2017Hdf5 struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"jessee_2017.hdf5"`
					Jessee2017Png struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"jessee_2017.png"`
					Jessee2017ModelFlt struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"jessee_2017_model.flt"`
					Jessee2017ModelHdr struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"jessee_2017_model.hdr"`
					Jessee2017ModelTif struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"jessee_2017_model.tif"`
					Nowicki2014GlobalHdf5 struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"nowicki_2014_global.hdf5"`
					Nowicki2014GlobalPng struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"nowicki_2014_global.png"`
					Nowicki2014GlobalModelFlt struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"nowicki_2014_global_model.flt"`
					Nowicki2014GlobalModelHdr struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"nowicki_2014_global_model.hdr"`
					Nowicki2014GlobalModelTif struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"nowicki_2014_global_model.tif"`
					Zhu2015Hdf5 struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"zhu_2015.hdf5"`
					Zhu2015ModelFlt struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"zhu_2015_model.flt"`
					Zhu2015ModelHdr struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"zhu_2015_model.hdr"`
					Zhu2015ModelTif struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"zhu_2015_model.tif"`
					Zhu2017GeneralHdf5 struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"zhu_2017_general.hdf5"`
					Zhu2017GeneralPng struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"zhu_2017_general.png"`
					Zhu2017GeneralModelFlt struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"zhu_2017_general_model.flt"`
					Zhu2017GeneralModelHdr struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"zhu_2017_general_model.hdr"`
					Zhu2017GeneralModelTif struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"zhu_2017_general_model.tif"`
				} `json:"contents"`
			} `json:"ground-failure"`
			Losspager []struct {
				Indexid    string `json:"indexid"`
				IndexTime  int64  `json:"indexTime"`
				ID         string `json:"id"`
				Type       string `json:"type"`
				Code       string `json:"code"`
				Source     string `json:"source"`
				UpdateTime int64  `json:"updateTime"`
				Status     string `json:"status"`
				Properties struct {
					Alertlevel       string    `json:"alertlevel"`
					Depth            string    `json:"depth"`
					Eventsource      string    `json:"eventsource"`
					Eventsourcecode  string    `json:"eventsourcecode"`
					Eventtime        time.Time `json:"eventtime"`
					Latitude         string    `json:"latitude"`
					Longitude        string    `json:"longitude"`
					Magnitude        string    `json:"magnitude"`
					Maxmmi           string    `json:"maxmmi"`
					PdlClientVersion string    `json:"pdl-client-version"`
				} `json:"properties"`
				PreferredWeight int `json:"preferredWeight"`
				Contents        struct {
					AlerteconPdf struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"alertecon.pdf"`
					AlerteconPng struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"alertecon.png"`
					AlerteconSmallPng struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"alertecon_small.png"`
					AlerteconSmallerPng struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"alertecon_smaller.png"`
					AlertfatalPdf struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"alertfatal.pdf"`
					AlertfatalPng struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"alertfatal.png"`
					AlertfatalSmallPng struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"alertfatal_small.png"`
					AlertfatalSmallerPng struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"alertfatal_smaller.png"`
					ContentsXML struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"contents.xml"`
					EventLog struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"event.log"`
					ExposurePdf struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"exposure.pdf"`
					ExposurePng struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"exposure.png"`
					GridXML struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"grid.xml"`
					JSONAlertsJSON struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"json/alerts.json"`
					JSONCitiesJSON struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"json/cities.json"`
					JSONCommentsJSON struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"json/comments.json"`
					JSONEventJSON struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"json/event.json"`
					JSONExposuresJSON struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"json/exposures.json"`
					JSONHistoricalEarthquakesJSON struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"json/historical_earthquakes.json"`
					JSONLossesJSON struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"json/losses.json"`
					OnepagerAux struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"onepager.aux"`
					OnepagerLog struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"onepager.log"`
					OnepagerPdf struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"onepager.pdf"`
					OnepagerTex struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"onepager.tex"`
					PagerXML struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"pager.xml"`
				} `json:"contents"`
			} `json:"losspager"`
			MomentTensor []struct {
				Indexid    string `json:"indexid"`
				IndexTime  int64  `json:"indexTime"`
				ID         string `json:"id"`
				Type       string `json:"type"`
				Code       string `json:"code"`
				Source     string `json:"source"`
				UpdateTime int64  `json:"updateTime"`
				Status     string `json:"status"`
				Properties struct {
					BeachballSource         string    `json:"beachball-source"`
					Depth                   string    `json:"depth"`
					DerivedDepth            string    `json:"derived-depth"`
					DerivedEventtime        time.Time `json:"derived-eventtime"`
					DerivedLatitude         string    `json:"derived-latitude"`
					DerivedLongitude        string    `json:"derived-longitude"`
					DerivedMagnitude        string    `json:"derived-magnitude"`
					DerivedMagnitudeType    string    `json:"derived-magnitude-type"`
					EvaluationStatus        string    `json:"evaluation-status"`
					EventParametersPublicID string    `json:"eventParametersPublicID"`
					Eventsource             string    `json:"eventsource"`
					Eventsourcecode         string    `json:"eventsourcecode"`
					Eventtime               time.Time `json:"eventtime"`
					Latitude                string    `json:"latitude"`
					Longitude               string    `json:"longitude"`
					NAxisAzimuth            string    `json:"n-axis-azimuth"`
					NAxisLength             string    `json:"n-axis-length"`
					NAxisPlunge             string    `json:"n-axis-plunge"`
					NodalPlane1Dip          string    `json:"nodal-plane-1-dip"`
					NodalPlane1Rake         string    `json:"nodal-plane-1-rake"`
					NodalPlane1Strike       string    `json:"nodal-plane-1-strike"`
					NodalPlane2Dip          string    `json:"nodal-plane-2-dip"`
					NodalPlane2Rake         string    `json:"nodal-plane-2-rake"`
					NodalPlane2Strike       string    `json:"nodal-plane-2-strike"`
					PAxisAzimuth            string    `json:"p-axis-azimuth"`
					PAxisLength             string    `json:"p-axis-length"`
					PAxisPlunge             string    `json:"p-axis-plunge"`
					PdlClientVersion        string    `json:"pdl-client-version"`
					PercentDoubleCouple     string    `json:"percent-double-couple"`
					QuakemlPublicid         string    `json:"quakeml-publicid"`
					ReviewStatus            string    `json:"review-status"`
					ScalarMoment            string    `json:"scalar-moment"`
					SourcetimeDecaytime     string    `json:"sourcetime-decaytime"`
					SourcetimeDuration      string    `json:"sourcetime-duration"`
					SourcetimeRisetime      string    `json:"sourcetime-risetime"`
					SourcetimeType          string    `json:"sourcetime-type"`
					TAxisAzimuth            string    `json:"t-axis-azimuth"`
					TAxisLength             string    `json:"t-axis-length"`
					TAxisPlunge             string    `json:"t-axis-plunge"`
					TensorMpp               string    `json:"tensor-mpp"`
					TensorMrp               string    `json:"tensor-mrp"`
					TensorMrr               string    `json:"tensor-mrr"`
					TensorMrt               string    `json:"tensor-mrt"`
					TensorMtp               string    `json:"tensor-mtp"`
					TensorMtt               string    `json:"tensor-mtt"`
				} `json:"properties"`
				PreferredWeight int `json:"preferredWeight"`
				Contents        struct {
					ContentsXML struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"contents.xml"`
					QuakemlXML struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"quakeml.xml"`
				} `json:"contents"`
			} `json:"moment-tensor"`
			Origin []struct {
				Indexid    string `json:"indexid"`
				IndexTime  int64  `json:"indexTime"`
				ID         string `json:"id"`
				Type       string `json:"type"`
				Code       string `json:"code"`
				Source     string `json:"source"`
				UpdateTime int64  `json:"updateTime"`
				Status     string `json:"status"`
				Properties struct {
					AzimuthalGap             string    `json:"azimuthal-gap"`
					Depth                    string    `json:"depth"`
					DepthType                string    `json:"depth-type"`
					ErrorEllipseAzimuth      string    `json:"error-ellipse-azimuth"`
					ErrorEllipseIntermediate string    `json:"error-ellipse-intermediate"`
					ErrorEllipseMajor        string    `json:"error-ellipse-major"`
					ErrorEllipseMinor        string    `json:"error-ellipse-minor"`
					ErrorEllipsePlunge       string    `json:"error-ellipse-plunge"`
					ErrorEllipseRotation     string    `json:"error-ellipse-rotation"`
					EvaluationStatus         string    `json:"evaluation-status"`
					EventType                string    `json:"event-type"`
					EventParametersPublicID  string    `json:"eventParametersPublicID"`
					Eventsource              string    `json:"eventsource"`
					Eventsourcecode          string    `json:"eventsourcecode"`
					Eventtime                time.Time `json:"eventtime"`
					EventtimeError           string    `json:"eventtime-error"`
					HorizontalError          string    `json:"horizontal-error"`
					Latitude                 string    `json:"latitude"`
					LatitudeError            string    `json:"latitude-error"`
					Longitude                string    `json:"longitude"`
					LongitudeError           string    `json:"longitude-error"`
					Magnitude                string    `json:"magnitude"`
					MagnitudeError           string    `json:"magnitude-error"`
					MagnitudeNumStationsUsed string    `json:"magnitude-num-stations-used"`
					MagnitudeSource          string    `json:"magnitude-source"`
					MagnitudeType            string    `json:"magnitude-type"`
					MinimumDistance          string    `json:"minimum-distance"`
					NumPhasesUsed            string    `json:"num-phases-used"`
					OriginSource             string    `json:"origin-source"`
					PdlClientVersion         string    `json:"pdl-client-version"`
					QuakemlMagnitudePublicid string    `json:"quakeml-magnitude-publicid"`
					QuakemlOriginPublicid    string    `json:"quakeml-origin-publicid"`
					QuakemlPublicid          string    `json:"quakeml-publicid"`
					ReviewStatus             string    `json:"review-status"`
					StandardError            string    `json:"standard-error"`
					VerticalError            string    `json:"vertical-error"`
				} `json:"properties"`
				PreferredWeight int `json:"preferredWeight"`
				Contents        struct {
					ContentsXML struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"contents.xml"`
					QuakemlXML struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"quakeml.xml"`
				} `json:"contents"`
			} `json:"origin"`
			PhaseData []struct {
				Indexid    string `json:"indexid"`
				IndexTime  int64  `json:"indexTime"`
				ID         string `json:"id"`
				Type       string `json:"type"`
				Code       string `json:"code"`
				Source     string `json:"source"`
				UpdateTime int64  `json:"updateTime"`
				Status     string `json:"status"`
				Properties struct {
					AzimuthalGap             string    `json:"azimuthal-gap"`
					Depth                    string    `json:"depth"`
					DepthType                string    `json:"depth-type"`
					ErrorEllipseAzimuth      string    `json:"error-ellipse-azimuth"`
					ErrorEllipseIntermediate string    `json:"error-ellipse-intermediate"`
					ErrorEllipseMajor        string    `json:"error-ellipse-major"`
					ErrorEllipseMinor        string    `json:"error-ellipse-minor"`
					ErrorEllipsePlunge       string    `json:"error-ellipse-plunge"`
					ErrorEllipseRotation     string    `json:"error-ellipse-rotation"`
					EvaluationStatus         string    `json:"evaluation-status"`
					EventType                string    `json:"event-type"`
					EventParametersPublicID  string    `json:"eventParametersPublicID"`
					Eventsource              string    `json:"eventsource"`
					Eventsourcecode          string    `json:"eventsourcecode"`
					Eventtime                time.Time `json:"eventtime"`
					EventtimeError           string    `json:"eventtime-error"`
					HorizontalError          string    `json:"horizontal-error"`
					Latitude                 string    `json:"latitude"`
					LatitudeError            string    `json:"latitude-error"`
					Longitude                string    `json:"longitude"`
					LongitudeError           string    `json:"longitude-error"`
					Magnitude                string    `json:"magnitude"`
					MagnitudeError           string    `json:"magnitude-error"`
					MagnitudeNumStationsUsed string    `json:"magnitude-num-stations-used"`
					MagnitudeSource          string    `json:"magnitude-source"`
					MagnitudeType            string    `json:"magnitude-type"`
					MinimumDistance          string    `json:"minimum-distance"`
					NumPhasesUsed            string    `json:"num-phases-used"`
					OriginSource             string    `json:"origin-source"`
					PdlClientVersion         string    `json:"pdl-client-version"`
					QuakemlMagnitudePublicid string    `json:"quakeml-magnitude-publicid"`
					QuakemlOriginPublicid    string    `json:"quakeml-origin-publicid"`
					QuakemlPublicid          string    `json:"quakeml-publicid"`
					ReviewStatus             string    `json:"review-status"`
					StandardError            string    `json:"standard-error"`
					VerticalError            string    `json:"vertical-error"`
				} `json:"properties"`
				PreferredWeight int `json:"preferredWeight"`
				Contents        struct {
					ContentsXML struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"contents.xml"`
					QuakemlXML struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"quakeml.xml"`
				} `json:"contents"`
			} `json:"phase-data"`
			Shakemap []struct {
				Indexid    string `json:"indexid"`
				IndexTime  int64  `json:"indexTime"`
				ID         string `json:"id"`
				Type       string `json:"type"`
				Code       string `json:"code"`
				Source     string `json:"source"`
				UpdateTime int64  `json:"updateTime"`
				Status     string `json:"status"`
				Properties struct {
					Depth            string    `json:"depth"`
					EventDescription string    `json:"event-description"`
					EventType        string    `json:"event-type"`
					Eventsource      string    `json:"eventsource"`
					Eventsourcecode  string    `json:"eventsourcecode"`
					Eventtime        time.Time `json:"eventtime"`
					Latitude         string    `json:"latitude"`
					Longitude        string    `json:"longitude"`
					Magnitude        string    `json:"magnitude"`
					MapStatus        string    `json:"map-status"`
					MaximumLatitude  string    `json:"maximum-latitude"`
					MaximumLongitude string    `json:"maximum-longitude"`
					Maxmmi           string    `json:"maxmmi"`
					MaxmmiGrid       string    `json:"maxmmi-grid"`
					Maxpga           string    `json:"maxpga"`
					MaxpgaGrid       string    `json:"maxpga-grid"`
					Maxpgv           string    `json:"maxpgv"`
					MaxpgvGrid       string    `json:"maxpgv-grid"`
					Maxpsa03         string    `json:"maxpsa03"`
					Maxpsa03Grid     string    `json:"maxpsa03-grid"`
					Maxpsa10         string    `json:"maxpsa10"`
					Maxpsa10Grid     string    `json:"maxpsa10-grid"`
					Maxpsa30         string    `json:"maxpsa30"`
					Maxpsa30Grid     string    `json:"maxpsa30-grid"`
					MinimumLatitude  string    `json:"minimum-latitude"`
					MinimumLongitude string    `json:"minimum-longitude"`
					OverlayHeight    string    `json:"overlayHeight"`
					OverlayWidth     string    `json:"overlayWidth"`
					PdlClientVersion string    `json:"pdl-client-version"`
					ProcessTimestamp time.Time `json:"process-timestamp"`
					Version          string    `json:"version"`
				} `json:"properties"`
				PreferredWeight int `json:"preferredWeight"`
				Contents        struct {
					AboutFormatsHTML struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"about_formats.html"`
					ContentsXML struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"contents.xml"`
					Download2000Ha1KKml struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"download/2000ha1k.kml"`
					DownloadContMiJSON struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"download/cont_mi.json"`
					DownloadContMiKmz struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"download/cont_mi.kmz"`
					DownloadContPgaJSON struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"download/cont_pga.json"`
					DownloadContPgaKmz struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"download/cont_pga.kmz"`
					DownloadContPgvJSON struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"download/cont_pgv.json"`
					DownloadContPgvKmz struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"download/cont_pgv.kmz"`
					DownloadContPsa03JSON struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"download/cont_psa03.json"`
					DownloadContPsa03Kmz struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"download/cont_psa03.kmz"`
					DownloadContPsa10JSON struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"download/cont_psa10.json"`
					DownloadContPsa10Kmz struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"download/cont_psa10.kmz"`
					DownloadContPsa30JSON struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"download/cont_psa30.json"`
					DownloadContPsa30Kmz struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"download/cont_psa30.kmz"`
					DownloadEpicenterKmz struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"download/epicenter.kmz"`
					DownloadEventTxt struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"download/event.txt"`
					DownloadGridXML struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"download/grid.xml"`
					DownloadGridXMLZip struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"download/grid.xml.zip"`
					DownloadGridXyzZip struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"download/grid.xyz.zip"`
					DownloadHazusZip struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"download/hazus.zip"`
					DownloadIiOverlayPng struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"download/ii_overlay.png"`
					DownloadIiThumbnailJpg struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"download/ii_thumbnail.jpg"`
					DownloadInfoJSON struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"download/info.json"`
					DownloadIntensityJpg struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"download/intensity.jpg"`
					DownloadIntensityPsZip struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"download/intensity.ps.zip"`
					DownloadMetadataTxt struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"download/metadata.txt"`
					DownloadMiRegrPng struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"download/mi_regr.png"`
					DownloadOverlayKmz struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"download/overlay.kmz"`
					DownloadPgaJpg struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"download/pga.jpg"`
					DownloadPgaPsZip struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"download/pga.ps.zip"`
					DownloadPgaRegrPng struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"download/pga_regr.png"`
					DownloadPgvJpg struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"download/pgv.jpg"`
					DownloadPgvPsZip struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"download/pgv.ps.zip"`
					DownloadPgvRegrPng struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"download/pgv_regr.png"`
					DownloadPolygonsMiKmz struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"download/polygons_mi.kmz"`
					DownloadPsa03Jpg struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"download/psa03.jpg"`
					DownloadPsa03PsZip struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"download/psa03.ps.zip"`
					DownloadPsa03RegrPng struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"download/psa03_regr.png"`
					DownloadPsa10Jpg struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"download/psa10.jpg"`
					DownloadPsa10PsZip struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"download/psa10.ps.zip"`
					DownloadPsa10RegrPng struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"download/psa10_regr.png"`
					DownloadPsa30Jpg struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"download/psa30.jpg"`
					DownloadPsa30PsZip struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"download/psa30.ps.zip"`
					DownloadPsa30RegrPng struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"download/psa30_regr.png"`
					DownloadRasterZip struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"download/raster.zip"`
					DownloadRockGridXMLZip struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"download/rock_grid.xml.zip"`
					DownloadSdJpg struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"download/sd.jpg"`
					DownloadShapeZip struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"download/shape.zip"`
					DownloadStationlistJSON struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"download/stationlist.json"`
					DownloadStationlistTxt struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"download/stationlist.txt"`
					DownloadStationlistXML struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"download/stationlist.xml"`
					DownloadStationsKmz struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"download/stations.kmz"`
					DownloadTvguideTxt struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"download/tvguide.txt"`
					DownloadTvmapJpg struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"download/tvmap.jpg"`
					DownloadTvmapPsZip struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"download/tvmap.ps.zip"`
					DownloadTvmapBareJpg struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"download/tvmap_bare.jpg"`
					DownloadTvmapBarePsZip struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"download/tvmap_bare.ps.zip"`
					DownloadUncertaintyXMLZip struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"download/uncertainty.xml.zip"`
					DownloadUratPgaJpg struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"download/urat_pga.jpg"`
					DownloadUratPgaPsZip struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"download/urat_pga.ps.zip"`
					DownloadUs2000Ha1KKml struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"download/us2000ha1k.kml"`
					DownloadUs2000Ha1KKmz struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"download/us2000ha1k.kmz"`
					IntensityHTML struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"intensity.html"`
					PgaHTML struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"pga.html"`
					PgvHTML struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"pgv.html"`
					ProductsHTML struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"products.html"`
					Psa03HTML struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"psa03.html"`
					Psa10HTML struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"psa10.html"`
					Psa30HTML struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"psa30.html"`
					StationlistHTML struct {
						ContentType  string `json:"contentType"`
						LastModified int64  `json:"lastModified"`
						Length       int    `json:"length"`
						URL          string `json:"url"`
					} `json:"stationlist.html"`
				} `json:"contents"`
			} `json:"shakemap"`
		} `json:"products"`
	} `json:"properties"`
	Geometry struct {
		Type        string    `json:"type"`
		Coordinates []float64 `json:"coordinates"`
	} `json:"geometry"`
	ID  string `json:"id"`
	URL string `json:"url"`
}
