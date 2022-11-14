package valuerenderer_test

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	_ "cosmossdk.io/api/cosmos/bank/v1beta1"
	_ "cosmossdk.io/api/cosmos/crypto/ed25519"
	_ "cosmossdk.io/api/cosmos/crypto/multisig"
	_ "cosmossdk.io/api/cosmos/crypto/secp256k1"
	textualv1 "cosmossdk.io/api/cosmos/msg/textual/v1"
	signingv1beta1 "cosmossdk.io/api/cosmos/tx/signing/v1beta1"
	txv1beta1 "cosmossdk.io/api/cosmos/tx/v1beta1"
	"cosmossdk.io/tx/textual/valuerenderer"
)

// txJsonTestTx represents the type that in the JSON test
// cases `proto` field. The inner contents are protojson
// encoded, so we represent them as []byte here, and decode
// them inside the test.
type txJsonTestTx struct {
	Body       json.RawMessage
	AuthInfo   json.RawMessage `json:"auth_info"`
	SignerData json.RawMessage `json:"signer_data"`
}

type txJsonTest struct {
	Proto   txJsonTestTx
	Error   bool
	Screens []valuerenderer.Screen
}

func TestTxJsonTestcases(t *testing.T) {
	raw, err := os.ReadFile("../internal/testdata/tx.json")
	require.NoError(t, err)

	var testcases []txJsonTest
	err = json.Unmarshal(raw, &testcases)
	require.NoError(t, err)

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {

			textualData := createTextualData(t, tc.Proto)

			tr := valuerenderer.NewTextual(EmptyCoinMetadataQuerier)
			rend := valuerenderer.NewTxValueRenderer(&tr)

			val := protoreflect.ValueOf(textualData.ProtoReflect())
			screens, err := rend.Format(context.Background(), val)
			if tc.Error {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tc.Screens, screens)

			// Round trip.
			// parsedVal, err := rend.Parse(context.Background(), screens)
			// require.NoError(t, err)
			// require.Equal(t, val.Interface(), parsedVal.Interface())
		})
	}
}

// createTextualData creates a Textual data give then JSON
// test case.
func createTextualData(t *testing.T, jsonTx txJsonTestTx) *textualv1.TextualData {
	txBody := &txv1beta1.TxBody{}
	txAuthInfo := &txv1beta1.AuthInfo{}
	signerData := &signingv1beta1.SignerData{}

	// We unmarshal from protojson to the protobuf types.
	err := protojson.Unmarshal(jsonTx.Body, txBody)
	require.NoError(t, err)
	err = protojson.Unmarshal(jsonTx.AuthInfo, txAuthInfo)
	require.NoError(t, err)
	err = protojson.Unmarshal(jsonTx.SignerData, signerData)
	require.NoError(t, err)

	// We marshal body and auth_info
	bodyBz, err := proto.Marshal(txBody)
	require.NoError(t, err)
	authInfoBz, err := proto.Marshal(txAuthInfo)
	require.NoError(t, err)

	return &textualv1.TextualData{
		BodyBytes:     bodyBz,
		AuthInfoBytes: authInfoBz,
		SignerData:    signerData,
	}
}
