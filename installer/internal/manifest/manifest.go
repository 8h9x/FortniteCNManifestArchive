package manifest

import (
	"encoding/json"
	"io"
)

type wrapper struct {
	JSONEncodeFiles    string    `json:"json_encode_files"` // []ManifestFile
	JSONEncodeFilesMds DepotHash `json:"json_encode_files_mds"`
}

type File struct {
	Name          string        `json:"name"`
	Length        int           `json:"length"`
	CreateTimeUtc string        `json:"create_time_utc"`
	Md5           string        `json:"md5"`
	Segments      []FileSegment `json:"segments"`
}

type FileSegment struct {
	Index int    `json:"index"`
	Md5   string `json:"md5"`
}

type DepotHash struct {
	Length int    `json:"length"`
	KeyID  string `json:"key_id"`
	MD5    string `json:"md5"`
}

type Depot struct {
	Files     []File
	DepotHash DepotHash
}

func (w *Depot) UnmarshalJSON(data []byte) error {
	var raw wrapper
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	w.DepotHash = raw.JSONEncodeFilesMds
	return json.Unmarshal([]byte(raw.JSONEncodeFiles), &w.Files)
}

func Decode(r io.Reader, out *Depot) error {
	if err := json.NewDecoder(r).Decode(&out); err != nil {
		return err
	}
	return nil
}
