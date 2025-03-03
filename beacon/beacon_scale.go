// Code generated by github.com/spacemeshos/go-scale/scalegen. DO NOT EDIT.

// nolint
package beacon

import (
	"github.com/spacemeshos/go-scale"
)

func (t *BuildProposalMessage) EncodeScale(enc *scale.Encoder) (total int, err error) {
	{
		n, err := scale.EncodeString(enc, string(t.Prefix))
		if err != nil {
			return total, err
		}
		total += n
	}
	{
		n, err := scale.EncodeCompact32(enc, uint32(t.Epoch))
		if err != nil {
			return total, err
		}
		total += n
	}
	return total, nil
}

func (t *BuildProposalMessage) DecodeScale(dec *scale.Decoder) (total int, err error) {
	{
		field, n, err := scale.DecodeString(dec)
		if err != nil {
			return total, err
		}
		total += n
		t.Prefix = string(field)
	}
	{
		field, n, err := scale.DecodeCompact32(dec)
		if err != nil {
			return total, err
		}
		total += n
		t.Epoch = uint32(field)
	}
	return total, nil
}
