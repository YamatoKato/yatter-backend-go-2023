package object

import (
	"net/url"
	"strconv"
)

const QUERY_ONLY_MEDIA = "only_media"
const QUERY_MAX_ID = "max_id"
const QUERY_SINCE_ID = "since_id"
const QUERY_LIMIT = "limit"
const ONLY_MEDIA_DEFAULT bool = false
const MAX_ID_DEFAULT int64 = 0
const SINCE_ID_DEFAULT int64 = 0
const LIMIT_DEFAULT int64 = 40
const LIMIT_MAX int64 = 80

type Timeline struct {
	Statuses []*Status `json:"statuses,omitempty"`
}

type TimelineOptions struct {
	OnlyMedia bool
	MaxID     int64
	SinceID   int64
	Limit     int64
}

func ParseTimelineOptions(queries url.Values) (*TimelineOptions, error) {
	onlyMediaStr := queries.Get(QUERY_ONLY_MEDIA)
	maxIdStr := queries.Get(QUERY_MAX_ID)
	sinceIdStr := queries.Get(QUERY_SINCE_ID)
	limitStr := queries.Get(QUERY_LIMIT)

	onlyMedia, err := parseBool(onlyMediaStr, ONLY_MEDIA_DEFAULT)
	if err != nil {
		return nil, err
	}

	maxId, err := parseInt64(maxIdStr, MAX_ID_DEFAULT)
	if err != nil {
		return nil, err
	}

	sinceId, err := parseInt64(sinceIdStr, SINCE_ID_DEFAULT)
	if err != nil {
		return nil, err
	}

	limit, err := parseInt64(limitStr, LIMIT_DEFAULT)
	if limit > LIMIT_MAX {
		limit = LIMIT_MAX
	}
	if err != nil {
		return nil, err
	}

	return &TimelineOptions{
		OnlyMedia: onlyMedia,
		MaxID:     maxId,
		SinceID:   sinceId,
		Limit:     limit,
	}, nil
}

func parseBool(s string, initDefault bool) (bool, error) {
	if s == "" {
		return initDefault, nil
	}
	b, err := strconv.ParseBool(s)
	if err != nil {
		return false, err
	}
	return b, nil
}

func parseInt64(s string, initDefault int64) (int64, error) {
	if s == "" {
		return initDefault, nil
	}
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return i, nil
}
