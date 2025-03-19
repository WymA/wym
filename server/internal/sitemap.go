package internal

import (
	"encoding/xml"
	"io"
	"time"
)

// Sitemap represents a sitemap with URLs
type Sitemap struct {
	URLs []URL
}

// URL represents a single URL entry in the sitemap
type URL struct {
	Loc        string
	LastMod    *time.Time
	ChangeFreq string
	Priority   float32
}

// ChangeFreq constants for sitemap
const (
	ChangeFreqAlways  = "always"
	ChangeFreqHourly  = "hourly"
	ChangeFreqDaily   = "daily"
	ChangeFreqWeekly  = "weekly"
	ChangeFreqMonthly = "monthly"
	ChangeFreqYearly  = "yearly"
	ChangeFreqNever   = "never"
)

// NewSitemap creates a new sitemap
func NewSitemap() *Sitemap {
	return &Sitemap{
		URLs: make([]URL, 0),
	}
}

// AddURL adds a URL to the sitemap
func (s *Sitemap) AddURL(loc string, lastMod *time.Time, changeFreq string, priority float32) {
	s.URLs = append(s.URLs, URL{
		Loc:        loc,
		LastMod:    lastMod,
		ChangeFreq: changeFreq,
		Priority:   priority,
	})
}

// xmlSitemap is used for XML marshaling
type xmlSitemap struct {
	XMLName xml.Name `xml:"urlset"`
	XMLNS   string   `xml:"xmlns,attr"`
	URLs    []xmlURL `xml:"url"`
}

// xmlURL is used for XML marshaling of URL entries
type xmlURL struct {
	Loc        string  `xml:"loc"`
	LastMod    string  `xml:"lastmod,omitempty"`
	ChangeFreq string  `xml:"changefreq,omitempty"`
	Priority   float32 `xml:"priority,omitempty"`
}

// Generate creates an XML sitemap
func (s *Sitemap) Generate(w io.Writer) error {
	sitemap := xmlSitemap{
		XMLNS: "http://www.sitemaps.org/schemas/sitemap/0.9",
		URLs:  make([]xmlURL, len(s.URLs)),
	}

	for i, url := range s.URLs {
		xmlURL := xmlURL{
			Loc:        url.Loc,
			ChangeFreq: url.ChangeFreq,
			Priority:   url.Priority,
		}

		if url.LastMod != nil {
			xmlURL.LastMod = url.LastMod.Format("2006-01-02T15:04:05-07:00")
		}

		sitemap.URLs[i] = xmlURL
	}

	// Add XML header
	w.Write([]byte(xml.Header))

	encoder := xml.NewEncoder(w)
	encoder.Indent("", "  ")
	return encoder.Encode(sitemap)
}

// GenerateToString returns the sitemap as a string
func (s *Sitemap) GenerateToString() (string, error) {
	data, err := xml.MarshalIndent(s.Generate, "", "  ")
	if err != nil {
		return "", err
	}
	return xml.Header + string(data), nil
}

func GenerateSitemapFromHistory(history []History) *Sitemap {
	sitemap := NewSitemap()
	for _, h := range history {
		// Convert string date to time.Time
		var lastMod *time.Time
		if h.Date != "" {
			parsedTime, err := time.Parse("20060102", h.Date)
			if err == nil {
				lastMod = &parsedTime
			}
		}
		sitemap.AddURL(BASE_HISTORY_URL+h.Filename, lastMod, ChangeFreqDaily, 0.5)
	}
	return sitemap
}
