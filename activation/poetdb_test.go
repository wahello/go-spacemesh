package activation

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/spacemeshos/go-spacemesh/codec"
	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/hash"
	"github.com/spacemeshos/go-spacemesh/log/logtest"
	"github.com/spacemeshos/go-spacemesh/sql"
)

var memberHash = []byte{0x17, 0x51, 0xac, 0x12, 0xe7, 0xe, 0x15, 0xb4, 0xf7, 0x6c, 0x16, 0x77, 0x5c, 0xd3, 0x29, 0xae, 0x55, 0x97, 0x3b, 0x61, 0x25, 0x21, 0xda, 0xb2, 0xde, 0x82, 0x8a, 0x5c, 0xdb, 0x6c, 0x8a, 0xb3}

func readPoetProofFromDisk(t *testing.T) *types.PoetProofMessage {
	file, err := os.Open(filepath.Join("test_resources", "poet.proof"))
	require.NoError(t, err)

	var poetProof types.PoetProof
	_, err = codec.DecodeFrom(file, &poetProof)
	require.NoError(t, err)
	require.EqualValues(t, [][]byte{memberHash}, poetProof.Members)
	poetID := []byte("poet_id_123456")
	roundID := "1337"
	return &types.PoetProofMessage{
		PoetProof:     poetProof,
		PoetServiceID: poetID,
		RoundID:       roundID,
		Signature:     nil,
	}
}

func TestPoetDbHappyFlow(t *testing.T) {
	msg := readPoetProofFromDisk(t)
	poetDb := NewPoetDb(sql.InMemory(), logtest.New(t))

	require.NoError(t, poetDb.Validate(msg.PoetProof, msg.PoetServiceID, msg.RoundID, nil))
	ref, err := msg.Ref()
	require.NoError(t, err)

	proofBytes, err := codec.Encode(&msg.PoetProof)
	require.NoError(t, err)
	expectedRef := hash.Sum(proofBytes)
	require.Equal(t, types.PoetProofRef(types.CalcHash32(expectedRef[:]).Bytes()), ref)

	require.NoError(t, poetDb.StoreProof(context.TODO(), ref, msg))
	got, err := poetDb.GetProofRef(msg.PoetServiceID, msg.RoundID)
	require.NoError(t, err)
	assert.Equal(t, ref, got)

	membership, err := poetDb.GetMembershipMap(ref)
	require.NoError(t, err)
	assert.True(t, membership[types.BytesToHash(memberHash)])
	assert.False(t, membership[types.BytesToHash([]byte("5"))])
}

func TestPoetDbPoetProofNoMembers(t *testing.T) {
	r := require.New(t)

	poetDb := NewPoetDb(sql.InMemory(), logtest.New(t))

	file, err := os.Open(filepath.Join("test_resources", "poet.proof"))
	r.NoError(err)

	var poetProof types.PoetProof
	_, err = codec.DecodeFrom(file, &poetProof)
	r.NoError(err)
	r.EqualValues([][]byte{memberHash}, poetProof.Members)
	poetID := []byte("poet_id_123456")
	roundID := "1337"
	poetProof.Root = []byte("some other root")

	poetProof.Members = nil

	err = poetDb.Validate(poetProof, poetID, roundID, nil)
	r.NoError(err)
	r.False(types.IsProcessingError(err))
}

func TestPoetDbInvalidPoetProof(t *testing.T) {
	msg := readPoetProofFromDisk(t)
	poetDb := NewPoetDb(sql.InMemory(), logtest.New(t))
	msg.PoetProof.Root = []byte("some other root")

	err := poetDb.Validate(msg.PoetProof, msg.PoetServiceID, msg.RoundID, nil)
	require.EqualError(t, err, fmt.Sprintf("failed to validate poet proof for poetID %x round 1337: validate PoET: merkle proof not valid",
		msg.PoetServiceID[:5]))
	assert.False(t, types.IsProcessingError(err))
}

func TestPoetDbNonExistingKeys(t *testing.T) {
	msg := readPoetProofFromDisk(t)
	poetDb := NewPoetDb(sql.InMemory(), logtest.New(t))

	_, err := poetDb.GetProofRef(msg.PoetServiceID, "0")
	require.EqualError(t, err, fmt.Sprintf("could not fetch poet proof for poet ID %x in round %v: get value: database: not found", msg.PoetServiceID[:5], "0"))

	ref := []byte("abcde")
	_, err = poetDb.GetMembershipMap(ref)
	require.EqualError(t, err, fmt.Sprintf("could not fetch poet proof for ref %x: get proof from store: get value: database: not found", ref[:5]))
}
