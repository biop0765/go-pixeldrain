package pd

import (
	"fmt"
	"time"
)

type ResponseDefault struct {
	StatusCode int    `json:"status_code"`
	Success    bool   `json:"success"`
	Value      string `json:"value,omitempty"`
	Message    string `json:"message,omitempty"`
}

type ResponseUpload struct {
	ID string `json:"id,omitempty"`
	ResponseDefault
}

// GetFileURL return the full URl to the uploaded file
func (rsp *ResponseUpload) GetFileURL() string {
	return fmt.Sprintf("%su/%s", BaseURL, rsp.ID)
}

type ResponseDownload struct {
	FilePath string `json:"file_path"`
	FileName string `json:"file_name"`
	FileSize int64  `json:"file_size"`
	ResponseDefault
}

type ResponseFileInfo struct {
	ID                string    `json:"id"`
	Name              string    `json:"name"`
	Size              int       `json:"size"`
	Views             int       `json:"views"`
	BandwidthUsed     int       `json:"bandwidth_used"`
	BandwidthUsedPaid int       `json:"bandwidth_used_paid"`
	Downloads         int       `json:"downloads"`
	DateUpload        time.Time `json:"date_upload"`
	DateLastView      time.Time `json:"date_last_view"`
	MimeType          string    `json:"mime_type"`
	ThumbnailHref     string    `json:"thumbnail_href"`
	HashSha256        string    `json:"hash_sha256"`
	CanEdit           bool      `json:"can_edit"`
	ResponseDefault
}

type ResponseThumbnail struct {
	FilePath string `json:"file_path"`
	FileName string `json:"file_name"`
	FileSize int64  `json:"file_size"`
	ResponseDefault
}

type ResponseDelete struct {
	ResponseDefault
}

type ResponseCreateList struct {
	ID string `json:"id"`
	ResponseDefault
}

type FileGetList struct {
	DetailHref    string    `json:"detail_href"`
	Description   string    `json:"description"`
	Success       bool      `json:"success"`
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	Size          int       `json:"size"`
	DateCreated   time.Time `json:"date_created"`
	DateLastView  time.Time `json:"date_last_view"`
	MimeType      string    `json:"mime_type"`
	Views         int       `json:"views"`
	BandwidthUsed int       `json:"bandwidth_used"`
	ThumbnailHref string    `json:"thumbnail_href"`
}

type ResponseGetList struct {
	ID          string        `json:"id"`
	Title       string        `json:"title"`
	DateCreated time.Time     `json:"date_created"`
	Files       []FileGetList `json:"files"`
	ResponseDefault
}

type FileGetUser struct {
	ID                  string    `json:"id"`
	Name                string    `json:"name"`
	Size                int       `json:"size"`
	Views               int       `json:"views"`
	BandwidthUsed       int       `json:"bandwidth_used"`
	BandwidthUsedPaid   int       `json:"bandwidth_used_paid"`
	Downloads           int       `json:"downloads"`
	DateUpload          time.Time `json:"date_upload"`
	DateLastView        time.Time `json:"date_last_view"`
	MimeType            string    `json:"mime_type"`
	ThumbnailHref       string    `json:"thumbnail_href"`
	HashSha256          string    `json:"hash_sha256"`
	Availability        string    `json:"availability"`
	AvailabilityMessage string    `json:"availability_message"`
	AbuseType           string    `json:"abuse_type"`
	AbuseReporterName   string    `json:"abuse_reporter_name"`
	CanEdit             bool      `json:"can_edit"`
	ShowAds             bool      `json:"show_ads"`
	AllowVideoPlayer    bool      `json:"allow_video_player"`
	DownloadSpeedLimit  int       `json:"download_speed_limit"`
}

type ResponseGetUserFiles struct {
	Files []FileGetUser `json:"files"`
	ResponseDefault
}

type ListsGetUser struct {
	ID          string      `json:"id"`
	Title       string      `json:"title"`
	DateCreated time.Time   `json:"date_created"`
	FileCount   int         `json:"file_count"`
	Files       interface{} `json:"files"`
	CanEdit     bool        `json:"can_edit"`
}

type ResponseGetUserLists struct {
	Lists []ListsGetUser `json:"lists"`
	ResponseDefault
}
