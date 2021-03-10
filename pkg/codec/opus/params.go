package opus

import (
	"time"

	"github.com/pion/mediadevices/pkg/codec"
	"github.com/pion/mediadevices/pkg/io/audio"
	"github.com/pion/mediadevices/pkg/prop"
	"github.com/pion/mediadevices/pkg/wave/mixer"
)

// Params stores opus specific encoding parameters.
type Params struct {
	codec.BaseParams
	// ChannelMixer is a mixer to be used if number of given and expected channels differ.
	ChannelMixer mixer.ChannelMixer
}

// NewParams returns default opus codec specific parameters.
func NewParams() (Params, error) {
	return Params{
		BaseParams: codec.BaseParams{
			Latency: 20 * time.Millisecond,
		},
	}, nil
}

// RTPCodec represents the codec metadata
func (p *Params) RTPCodec() *codec.RTPCodec {
	c := codec.NewRTPOpusCodec(48000)
	c.Latency = p.Latency
	return c
}

// BuildAudioEncoder builds opus encoder with given params
func (p *Params) BuildAudioEncoder(r audio.Reader, property prop.Media) (codec.ReadCloser, error) {
	return newEncoder(r, property, *p)
}
