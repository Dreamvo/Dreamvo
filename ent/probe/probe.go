// Code generated by entc, DO NOT EDIT.

package probe

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the probe type in the database.
	Label = "probe"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldFilename holds the string denoting the filename field in the database.
	FieldFilename = "filename"
	// FieldMimetype holds the string denoting the mimetype field in the database.
	FieldMimetype = "mimetype"
	// FieldFilesize holds the string denoting the filesize field in the database.
	FieldFilesize = "filesize"
	// FieldChecksumSha256 holds the string denoting the checksum_sha256 field in the database.
	FieldChecksumSha256 = "checksum_sha256"
	// FieldAspectRatio holds the string denoting the aspect_ratio field in the database.
	FieldAspectRatio = "aspect_ratio"
	// FieldWidth holds the string denoting the width field in the database.
	FieldWidth = "width"
	// FieldHeight holds the string denoting the height field in the database.
	FieldHeight = "height"
	// FieldDurationSeconds holds the string denoting the duration_seconds field in the database.
	FieldDurationSeconds = "duration_seconds"
	// FieldVideoBitrate holds the string denoting the video_bitrate field in the database.
	FieldVideoBitrate = "video_bitrate"
	// FieldAudioBitrate holds the string denoting the audio_bitrate field in the database.
	FieldAudioBitrate = "audio_bitrate"
	// FieldFramerate holds the string denoting the framerate field in the database.
	FieldFramerate = "framerate"
	// FieldFormat holds the string denoting the format field in the database.
	FieldFormat = "format"
	// FieldNbStreams holds the string denoting the nb_streams field in the database.
	FieldNbStreams = "nb_streams"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"

	// EdgeMedia holds the string denoting the media edge name in mutations.
	EdgeMedia = "media"

	// Table holds the table name of the probe in the database.
	Table = "meida_probe"
	// MediaTable is the table the holds the media relation/edge.
	MediaTable = "meida_probe"
	// MediaInverseTable is the table name for the Media entity.
	// It exists in this package in order to avoid circular dependency with the "media" package.
	MediaInverseTable = "media"
	// MediaColumn is the table column denoting the media relation/edge.
	MediaColumn = "media"
)

// Columns holds all SQL columns for probe fields.
var Columns = []string{
	FieldID,
	FieldFilename,
	FieldMimetype,
	FieldFilesize,
	FieldChecksumSha256,
	FieldAspectRatio,
	FieldWidth,
	FieldHeight,
	FieldDurationSeconds,
	FieldVideoBitrate,
	FieldAudioBitrate,
	FieldFramerate,
	FieldFormat,
	FieldNbStreams,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Probe type.
var ForeignKeys = []string{
	"media",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// FilenameValidator is a validator for the "filename" field. It is called by the builders before save.
	FilenameValidator func(string) error
	// MimetypeValidator is a validator for the "mimetype" field. It is called by the builders before save.
	MimetypeValidator func(string) error
	// DefaultFilesize holds the default value on creation for the filesize field.
	DefaultFilesize int
	// FilesizeValidator is a validator for the "filesize" field. It is called by the builders before save.
	FilesizeValidator func(int) error
	// ChecksumSha256Validator is a validator for the "checksum_sha256" field. It is called by the builders before save.
	ChecksumSha256Validator func(string) error
	// DefaultAspectRatio holds the default value on creation for the aspect_ratio field.
	DefaultAspectRatio string
	// AspectRatioValidator is a validator for the "aspect_ratio" field. It is called by the builders before save.
	AspectRatioValidator func(string) error
	// WidthValidator is a validator for the "width" field. It is called by the builders before save.
	WidthValidator func(int) error
	// HeightValidator is a validator for the "height" field. It is called by the builders before save.
	HeightValidator func(int) error
	// DefaultDurationSeconds holds the default value on creation for the duration_seconds field.
	DefaultDurationSeconds float64
	// DurationSecondsValidator is a validator for the "duration_seconds" field. It is called by the builders before save.
	DurationSecondsValidator func(float64) error
	// DefaultVideoBitrate holds the default value on creation for the video_bitrate field.
	DefaultVideoBitrate int
	// VideoBitrateValidator is a validator for the "video_bitrate" field. It is called by the builders before save.
	VideoBitrateValidator func(int) error
	// DefaultAudioBitrate holds the default value on creation for the audio_bitrate field.
	DefaultAudioBitrate int
	// AudioBitrateValidator is a validator for the "audio_bitrate" field. It is called by the builders before save.
	AudioBitrateValidator func(int) error
	// FramerateValidator is a validator for the "framerate" field. It is called by the builders before save.
	FramerateValidator func(int) error
	// FormatValidator is a validator for the "format" field. It is called by the builders before save.
	FormatValidator func(string) error
	// DefaultNbStreams holds the default value on creation for the nb_streams field.
	DefaultNbStreams int
	// NbStreamsValidator is a validator for the "nb_streams" field. It is called by the builders before save.
	NbStreamsValidator func(int) error
	// DefaultCreatedAt holds the default value on creation for the created_at field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the updated_at field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the id field.
	DefaultID func() uuid.UUID
)
