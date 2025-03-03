// Code generated by github.com/spacemeshos/go-scale/scalegen. DO NOT EDIT.

// nolint
package types

import (
	"github.com/spacemeshos/go-scale"
)

func (t *NIPostBuilderState) EncodeScale(enc *scale.Encoder) (total int, err error) {
	{
		n, err := scale.EncodeByteArray(enc, t.Challenge[:])
		if err != nil {
			return total, err
		}
		total += n
	}
	{
		n, err := scale.EncodeOption(enc, t.NIPost)
		if err != nil {
			return total, err
		}
		total += n
	}
	{
		n, err := scale.EncodeStructSlice(enc, t.PoetRequests)
		if err != nil {
			return total, err
		}
		total += n
	}
	{
		n, err := scale.EncodeByteSlice(enc, t.PoetProofRef)
		if err != nil {
			return total, err
		}
		total += n
	}
	return total, nil
}

func (t *NIPostBuilderState) DecodeScale(dec *scale.Decoder) (total int, err error) {
	{
		n, err := scale.DecodeByteArray(dec, t.Challenge[:])
		if err != nil {
			return total, err
		}
		total += n
	}
	{
		field, n, err := scale.DecodeOption[NIPost](dec)
		if err != nil {
			return total, err
		}
		total += n
		t.NIPost = field
	}
	{
		field, n, err := scale.DecodeStructSlice[PoetRequest](dec)
		if err != nil {
			return total, err
		}
		total += n
		t.PoetRequests = field
	}
	{
		field, n, err := scale.DecodeByteSlice(dec)
		if err != nil {
			return total, err
		}
		total += n
		t.PoetProofRef = field
	}
	return total, nil
}

func (t *PoetRequest) EncodeScale(enc *scale.Encoder) (total int, err error) {
	{
		n, err := scale.EncodeOption(enc, t.PoetRound)
		if err != nil {
			return total, err
		}
		total += n
	}
	{
		n, err := scale.EncodeByteSlice(enc, t.PoetServiceID)
		if err != nil {
			return total, err
		}
		total += n
	}
	return total, nil
}

func (t *PoetRequest) DecodeScale(dec *scale.Decoder) (total int, err error) {
	{
		field, n, err := scale.DecodeOption[PoetRound](dec)
		if err != nil {
			return total, err
		}
		total += n
		t.PoetRound = field
	}
	{
		field, n, err := scale.DecodeByteSlice(dec)
		if err != nil {
			return total, err
		}
		total += n
		t.PoetServiceID = field
	}
	return total, nil
}
