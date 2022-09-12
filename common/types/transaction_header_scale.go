// Code generated by github.com/spacemeshos/go-scale/scalegen. DO NOT EDIT.

// nolint
package types

import (
	"github.com/spacemeshos/go-scale"
)

func (t *TxHeader) EncodeScale(enc *scale.Encoder) (total int, err error) {
	{
		n, err := scale.EncodeByteArray(enc, t.Principal[:])
		if err != nil {
			return total, err
		}
		total += n
	}
	{
		n, err := scale.EncodeByteArray(enc, t.TemplateAddress[:])
		if err != nil {
			return total, err
		}
		total += n
	}
	{
		n, err := scale.EncodeCompact8(enc, uint8(t.Method))
		if err != nil {
			return total, err
		}
		total += n
	}
	{
		n, err := t.Nonce.EncodeScale(enc)
		if err != nil {
			return total, err
		}
		total += n
	}
	{
		n, err := t.LayerLimits.EncodeScale(enc)
		if err != nil {
			return total, err
		}
		total += n
	}
	{
		n, err := scale.EncodeCompact64(enc, uint64(t.MaxGas))
		if err != nil {
			return total, err
		}
		total += n
	}
	{
		n, err := scale.EncodeCompact64(enc, uint64(t.GasPrice))
		if err != nil {
			return total, err
		}
		total += n
	}
	{
		n, err := scale.EncodeCompact64(enc, uint64(t.MaxSpend))
		if err != nil {
			return total, err
		}
		total += n
	}
	return total, nil
}

func (t *TxHeader) DecodeScale(dec *scale.Decoder) (total int, err error) {
	{
		n, err := scale.DecodeByteArray(dec, t.Principal[:])
		if err != nil {
			return total, err
		}
		total += n
	}
	{
		n, err := scale.DecodeByteArray(dec, t.TemplateAddress[:])
		if err != nil {
			return total, err
		}
		total += n
	}
	{
		field, n, err := scale.DecodeCompact8(dec)
		if err != nil {
			return total, err
		}
		total += n
		t.Method = uint8(field)
	}
	{
		n, err := t.Nonce.DecodeScale(dec)
		if err != nil {
			return total, err
		}
		total += n
	}
	{
		n, err := t.LayerLimits.DecodeScale(dec)
		if err != nil {
			return total, err
		}
		total += n
	}
	{
		field, n, err := scale.DecodeCompact64(dec)
		if err != nil {
			return total, err
		}
		total += n
		t.MaxGas = uint64(field)
	}
	{
		field, n, err := scale.DecodeCompact64(dec)
		if err != nil {
			return total, err
		}
		total += n
		t.GasPrice = uint64(field)
	}
	{
		field, n, err := scale.DecodeCompact64(dec)
		if err != nil {
			return total, err
		}
		total += n
		t.MaxSpend = uint64(field)
	}
	return total, nil
}

func (t *LayerLimits) EncodeScale(enc *scale.Encoder) (total int, err error) {
	{
		n, err := scale.EncodeCompact32(enc, uint32(t.Min))
		if err != nil {
			return total, err
		}
		total += n
	}
	{
		n, err := scale.EncodeCompact32(enc, uint32(t.Max))
		if err != nil {
			return total, err
		}
		total += n
	}
	return total, nil
}

func (t *LayerLimits) DecodeScale(dec *scale.Decoder) (total int, err error) {
	{
		field, n, err := scale.DecodeCompact32(dec)
		if err != nil {
			return total, err
		}
		total += n
		t.Min = uint32(field)
	}
	{
		field, n, err := scale.DecodeCompact32(dec)
		if err != nil {
			return total, err
		}
		total += n
		t.Max = uint32(field)
	}
	return total, nil
}

func (t *Nonce) EncodeScale(enc *scale.Encoder) (total int, err error) {
	{
		n, err := scale.EncodeCompact64(enc, uint64(t.Counter))
		if err != nil {
			return total, err
		}
		total += n
	}
	{
		n, err := scale.EncodeCompact8(enc, uint8(t.Bitfield))
		if err != nil {
			return total, err
		}
		total += n
	}
	return total, nil
}

func (t *Nonce) DecodeScale(dec *scale.Decoder) (total int, err error) {
	{
		field, n, err := scale.DecodeCompact64(dec)
		if err != nil {
			return total, err
		}
		total += n
		t.Counter = uint64(field)
	}
	{
		field, n, err := scale.DecodeCompact8(dec)
		if err != nil {
			return total, err
		}
		total += n
		t.Bitfield = uint8(field)
	}
	return total, nil
}
