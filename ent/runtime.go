// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/dreamvo/gilfoyle/ent/media"
	"github.com/dreamvo/gilfoyle/ent/mediafile"
	"github.com/dreamvo/gilfoyle/ent/probe"
	"github.com/dreamvo/gilfoyle/ent/schema"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	mediaFields := schema.Media{}.Fields()
	_ = mediaFields
	// mediaDescTitle is the schema descriptor for title field.
	mediaDescTitle := mediaFields[1].Descriptor()
	// media.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	media.TitleValidator = func() func(string) error {
		validators := mediaDescTitle.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
			validators[2].(func(string) error),
		}
		return func(title string) error {
			for _, fn := range fns {
				if err := fn(title); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// mediaDescOriginalFilename is the schema descriptor for original_filename field.
	mediaDescOriginalFilename := mediaFields[2].Descriptor()
	// media.DefaultOriginalFilename holds the default value on creation for the original_filename field.
	media.DefaultOriginalFilename = mediaDescOriginalFilename.Default.(string)
	// media.OriginalFilenameValidator is a validator for the "original_filename" field. It is called by the builders before save.
	media.OriginalFilenameValidator = mediaDescOriginalFilename.Validators[0].(func(string) error)
	// mediaDescMessage is the schema descriptor for message field.
	mediaDescMessage := mediaFields[4].Descriptor()
	// media.DefaultMessage holds the default value on creation for the message field.
	media.DefaultMessage = mediaDescMessage.Default.(string)
	// media.MessageValidator is a validator for the "message" field. It is called by the builders before save.
	media.MessageValidator = mediaDescMessage.Validators[0].(func(string) error)
	// mediaDescCreatedAt is the schema descriptor for created_at field.
	mediaDescCreatedAt := mediaFields[5].Descriptor()
	// media.DefaultCreatedAt holds the default value on creation for the created_at field.
	media.DefaultCreatedAt = mediaDescCreatedAt.Default.(func() time.Time)
	// mediaDescUpdatedAt is the schema descriptor for updated_at field.
	mediaDescUpdatedAt := mediaFields[6].Descriptor()
	// media.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	media.DefaultUpdatedAt = mediaDescUpdatedAt.Default.(func() time.Time)
	// media.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	media.UpdateDefaultUpdatedAt = mediaDescUpdatedAt.UpdateDefault.(func() time.Time)
	// mediaDescID is the schema descriptor for id field.
	mediaDescID := mediaFields[0].Descriptor()
	// media.DefaultID holds the default value on creation for the id field.
	media.DefaultID = mediaDescID.Default.(func() uuid.UUID)
	mediafileFields := schema.MediaFile{}.Fields()
	_ = mediafileFields
	// mediafileDescRenditionName is the schema descriptor for rendition_name field.
	mediafileDescRenditionName := mediafileFields[1].Descriptor()
	// mediafile.RenditionNameValidator is a validator for the "rendition_name" field. It is called by the builders before save.
	mediafile.RenditionNameValidator = func() func(string) error {
		validators := mediafileDescRenditionName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(rendition_name string) error {
			for _, fn := range fns {
				if err := fn(rendition_name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// mediafileDescFormat is the schema descriptor for format field.
	mediafileDescFormat := mediafileFields[2].Descriptor()
	// mediafile.FormatValidator is a validator for the "format" field. It is called by the builders before save.
	mediafile.FormatValidator = mediafileDescFormat.Validators[0].(func(string) error)
	// mediafileDescTargetBandwidth is the schema descriptor for target_bandwidth field.
	mediafileDescTargetBandwidth := mediafileFields[3].Descriptor()
	// mediafile.DefaultTargetBandwidth holds the default value on creation for the target_bandwidth field.
	mediafile.DefaultTargetBandwidth = mediafileDescTargetBandwidth.Default.(uint64)
	// mediafileDescVideoBitrate is the schema descriptor for video_bitrate field.
	mediafileDescVideoBitrate := mediafileFields[4].Descriptor()
	// mediafile.VideoBitrateValidator is a validator for the "video_bitrate" field. It is called by the builders before save.
	mediafile.VideoBitrateValidator = mediafileDescVideoBitrate.Validators[0].(func(int64) error)
	// mediafileDescResolutionWidth is the schema descriptor for resolution_width field.
	mediafileDescResolutionWidth := mediafileFields[5].Descriptor()
	// mediafile.ResolutionWidthValidator is a validator for the "resolution_width" field. It is called by the builders before save.
	mediafile.ResolutionWidthValidator = func() func(uint16) error {
		validators := mediafileDescResolutionWidth.Validators
		fns := [...]func(uint16) error{
			validators[0].(func(uint16) error),
			validators[1].(func(uint16) error),
		}
		return func(resolution_width uint16) error {
			for _, fn := range fns {
				if err := fn(resolution_width); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// mediafileDescResolutionHeight is the schema descriptor for resolution_height field.
	mediafileDescResolutionHeight := mediafileFields[6].Descriptor()
	// mediafile.ResolutionHeightValidator is a validator for the "resolution_height" field. It is called by the builders before save.
	mediafile.ResolutionHeightValidator = func() func(uint16) error {
		validators := mediafileDescResolutionHeight.Validators
		fns := [...]func(uint16) error{
			validators[0].(func(uint16) error),
			validators[1].(func(uint16) error),
		}
		return func(resolution_height uint16) error {
			for _, fn := range fns {
				if err := fn(resolution_height); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// mediafileDescFramerate is the schema descriptor for framerate field.
	mediafileDescFramerate := mediafileFields[7].Descriptor()
	// mediafile.FramerateValidator is a validator for the "framerate" field. It is called by the builders before save.
	mediafile.FramerateValidator = func() func(uint8) error {
		validators := mediafileDescFramerate.Validators
		fns := [...]func(uint8) error{
			validators[0].(func(uint8) error),
			validators[1].(func(uint8) error),
		}
		return func(framerate uint8) error {
			for _, fn := range fns {
				if err := fn(framerate); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// mediafileDescDurationSeconds is the schema descriptor for duration_seconds field.
	mediafileDescDurationSeconds := mediafileFields[8].Descriptor()
	// mediafile.DurationSecondsValidator is a validator for the "duration_seconds" field. It is called by the builders before save.
	mediafile.DurationSecondsValidator = mediafileDescDurationSeconds.Validators[0].(func(float64) error)
	// mediafileDescMessage is the schema descriptor for message field.
	mediafileDescMessage := mediafileFields[11].Descriptor()
	// mediafile.DefaultMessage holds the default value on creation for the message field.
	mediafile.DefaultMessage = mediafileDescMessage.Default.(string)
	// mediafile.MessageValidator is a validator for the "message" field. It is called by the builders before save.
	mediafile.MessageValidator = mediafileDescMessage.Validators[0].(func(string) error)
	// mediafileDescEntryFile is the schema descriptor for entry_file field.
	mediafileDescEntryFile := mediafileFields[12].Descriptor()
	// mediafile.DefaultEntryFile holds the default value on creation for the entry_file field.
	mediafile.DefaultEntryFile = mediafileDescEntryFile.Default.(string)
	// mediafile.EntryFileValidator is a validator for the "entry_file" field. It is called by the builders before save.
	mediafile.EntryFileValidator = mediafileDescEntryFile.Validators[0].(func(string) error)
	// mediafileDescMimetype is the schema descriptor for mimetype field.
	mediafileDescMimetype := mediafileFields[13].Descriptor()
	// mediafile.DefaultMimetype holds the default value on creation for the mimetype field.
	mediafile.DefaultMimetype = mediafileDescMimetype.Default.(string)
	// mediafile.MimetypeValidator is a validator for the "mimetype" field. It is called by the builders before save.
	mediafile.MimetypeValidator = mediafileDescMimetype.Validators[0].(func(string) error)
	// mediafileDescCreatedAt is the schema descriptor for created_at field.
	mediafileDescCreatedAt := mediafileFields[14].Descriptor()
	// mediafile.DefaultCreatedAt holds the default value on creation for the created_at field.
	mediafile.DefaultCreatedAt = mediafileDescCreatedAt.Default.(func() time.Time)
	// mediafileDescUpdatedAt is the schema descriptor for updated_at field.
	mediafileDescUpdatedAt := mediafileFields[15].Descriptor()
	// mediafile.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	mediafile.DefaultUpdatedAt = mediafileDescUpdatedAt.Default.(func() time.Time)
	// mediafile.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	mediafile.UpdateDefaultUpdatedAt = mediafileDescUpdatedAt.UpdateDefault.(func() time.Time)
	// mediafileDescID is the schema descriptor for id field.
	mediafileDescID := mediafileFields[0].Descriptor()
	// mediafile.DefaultID holds the default value on creation for the id field.
	mediafile.DefaultID = mediafileDescID.Default.(func() uuid.UUID)
	probeFields := schema.Probe{}.Fields()
	_ = probeFields
	// probeDescFilename is the schema descriptor for filename field.
	probeDescFilename := probeFields[1].Descriptor()
	// probe.FilenameValidator is a validator for the "filename" field. It is called by the builders before save.
	probe.FilenameValidator = func() func(string) error {
		validators := probeDescFilename.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(filename string) error {
			for _, fn := range fns {
				if err := fn(filename); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// probeDescMimetype is the schema descriptor for mimetype field.
	probeDescMimetype := probeFields[2].Descriptor()
	// probe.MimetypeValidator is a validator for the "mimetype" field. It is called by the builders before save.
	probe.MimetypeValidator = func() func(string) error {
		validators := probeDescMimetype.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(mimetype string) error {
			for _, fn := range fns {
				if err := fn(mimetype); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// probeDescFilesize is the schema descriptor for filesize field.
	probeDescFilesize := probeFields[3].Descriptor()
	// probe.DefaultFilesize holds the default value on creation for the filesize field.
	probe.DefaultFilesize = probeDescFilesize.Default.(int)
	// probe.FilesizeValidator is a validator for the "filesize" field. It is called by the builders before save.
	probe.FilesizeValidator = probeDescFilesize.Validators[0].(func(int) error)
	// probeDescChecksumSha256 is the schema descriptor for checksum_sha256 field.
	probeDescChecksumSha256 := probeFields[4].Descriptor()
	// probe.ChecksumSha256Validator is a validator for the "checksum_sha256" field. It is called by the builders before save.
	probe.ChecksumSha256Validator = func() func(string) error {
		validators := probeDescChecksumSha256.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(checksum_sha256 string) error {
			for _, fn := range fns {
				if err := fn(checksum_sha256); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// probeDescAspectRatio is the schema descriptor for aspect_ratio field.
	probeDescAspectRatio := probeFields[5].Descriptor()
	// probe.DefaultAspectRatio holds the default value on creation for the aspect_ratio field.
	probe.DefaultAspectRatio = probeDescAspectRatio.Default.(string)
	// probe.AspectRatioValidator is a validator for the "aspect_ratio" field. It is called by the builders before save.
	probe.AspectRatioValidator = func() func(string) error {
		validators := probeDescAspectRatio.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(aspect_ratio string) error {
			for _, fn := range fns {
				if err := fn(aspect_ratio); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// probeDescWidth is the schema descriptor for width field.
	probeDescWidth := probeFields[6].Descriptor()
	// probe.WidthValidator is a validator for the "width" field. It is called by the builders before save.
	probe.WidthValidator = probeDescWidth.Validators[0].(func(int) error)
	// probeDescHeight is the schema descriptor for height field.
	probeDescHeight := probeFields[7].Descriptor()
	// probe.HeightValidator is a validator for the "height" field. It is called by the builders before save.
	probe.HeightValidator = probeDescHeight.Validators[0].(func(int) error)
	// probeDescDurationSeconds is the schema descriptor for duration_seconds field.
	probeDescDurationSeconds := probeFields[8].Descriptor()
	// probe.DefaultDurationSeconds holds the default value on creation for the duration_seconds field.
	probe.DefaultDurationSeconds = probeDescDurationSeconds.Default.(float64)
	// probe.DurationSecondsValidator is a validator for the "duration_seconds" field. It is called by the builders before save.
	probe.DurationSecondsValidator = probeDescDurationSeconds.Validators[0].(func(float64) error)
	// probeDescVideoBitrate is the schema descriptor for video_bitrate field.
	probeDescVideoBitrate := probeFields[9].Descriptor()
	// probe.DefaultVideoBitrate holds the default value on creation for the video_bitrate field.
	probe.DefaultVideoBitrate = probeDescVideoBitrate.Default.(int)
	// probe.VideoBitrateValidator is a validator for the "video_bitrate" field. It is called by the builders before save.
	probe.VideoBitrateValidator = probeDescVideoBitrate.Validators[0].(func(int) error)
	// probeDescAudioBitrate is the schema descriptor for audio_bitrate field.
	probeDescAudioBitrate := probeFields[10].Descriptor()
	// probe.DefaultAudioBitrate holds the default value on creation for the audio_bitrate field.
	probe.DefaultAudioBitrate = probeDescAudioBitrate.Default.(int)
	// probe.AudioBitrateValidator is a validator for the "audio_bitrate" field. It is called by the builders before save.
	probe.AudioBitrateValidator = probeDescAudioBitrate.Validators[0].(func(int) error)
	// probeDescFramerate is the schema descriptor for framerate field.
	probeDescFramerate := probeFields[11].Descriptor()
	// probe.FramerateValidator is a validator for the "framerate" field. It is called by the builders before save.
	probe.FramerateValidator = probeDescFramerate.Validators[0].(func(int) error)
	// probeDescFormat is the schema descriptor for format field.
	probeDescFormat := probeFields[12].Descriptor()
	// probe.FormatValidator is a validator for the "format" field. It is called by the builders before save.
	probe.FormatValidator = probeDescFormat.Validators[0].(func(string) error)
	// probeDescNbStreams is the schema descriptor for nb_streams field.
	probeDescNbStreams := probeFields[13].Descriptor()
	// probe.DefaultNbStreams holds the default value on creation for the nb_streams field.
	probe.DefaultNbStreams = probeDescNbStreams.Default.(int)
	// probe.NbStreamsValidator is a validator for the "nb_streams" field. It is called by the builders before save.
	probe.NbStreamsValidator = probeDescNbStreams.Validators[0].(func(int) error)
	// probeDescCreatedAt is the schema descriptor for created_at field.
	probeDescCreatedAt := probeFields[14].Descriptor()
	// probe.DefaultCreatedAt holds the default value on creation for the created_at field.
	probe.DefaultCreatedAt = probeDescCreatedAt.Default.(func() time.Time)
	// probeDescUpdatedAt is the schema descriptor for updated_at field.
	probeDescUpdatedAt := probeFields[15].Descriptor()
	// probe.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	probe.DefaultUpdatedAt = probeDescUpdatedAt.Default.(func() time.Time)
	// probe.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	probe.UpdateDefaultUpdatedAt = probeDescUpdatedAt.UpdateDefault.(func() time.Time)
	// probeDescID is the schema descriptor for id field.
	probeDescID := probeFields[0].Descriptor()
	// probe.DefaultID holds the default value on creation for the id field.
	probe.DefaultID = probeDescID.Default.(func() uuid.UUID)
}
