package nft

import (
	"fmt"
	"strconv"

	marketplaceTypes "github.com/CudoVentures/cudos-node/x/marketplace/types"
	nftTypes "github.com/CudoVentures/cudos-node/x/nft/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/forbole/bdjuno/v2/modules/utils"
	juno "github.com/forbole/juno/v2/types"
)

// HandleMsg implements MessageModule
func (m *Module) HandleMsg(index int, msg sdk.Msg, tx *juno.Tx) error {
	if tx.Code != 0 {
		return nil
	}

	switch cosmosMsg := msg.(type) {
	case *nftTypes.MsgIssueDenom:
		return m.handleMsgIssueDenom(tx, cosmosMsg)
	case *nftTypes.MsgTransferDenom:
		return m.handleMsgTransferDenom(cosmosMsg)
	case *nftTypes.MsgMintNFT:
		return m.handleMsgMintNFT(index, tx, cosmosMsg)
	case *nftTypes.MsgEditNFT:
		return m.handleMsgEditNFT(cosmosMsg)
	case *nftTypes.MsgTransferNft:
		return m.handleMsgTransferNFT(cosmosMsg)
	case *nftTypes.MsgBurnNFT:
		return m.handleMsgBurnNFT(cosmosMsg)
	case *marketplaceTypes.MsgCreateCollection:
		return m.handleMsgCreateCollection(tx, cosmosMsg)
	default:
		return nil
	}
}

func (m *Module) handleMsgIssueDenom(tx *juno.Tx, msg *nftTypes.MsgIssueDenom) error {
	dataJSON, dataText := getData(msg.Data)

	return m.db.SaveDenom(tx.TxHash, msg.Id, msg.Name, msg.Schema, msg.Symbol, msg.Sender, msg.ContractAddressSigner,
		msg.Traits, msg.Minter, msg.Description, dataText, utils.SanitizeUTF8(dataJSON))
}

func (m *Module) handleMsgTransferDenom(msg *nftTypes.MsgTransferDenom) error {
	return m.db.UpdateDenom(msg.Id, msg.Recipient)
}

func (m *Module) handleMsgMintNFT(index int, tx *juno.Tx, msg *nftTypes.MsgMintNFT) error {
	tokenIDStr := utils.GetValueFromLogs(uint32(index), tx.Logs, nftTypes.EventTypeMintNFT, nftTypes.AttributeKeyTokenID)
	if tokenIDStr == "" {
		return fmt.Errorf("token id not found in tx %s", tx.TxHash)
	}

	tokenID, err := strconv.ParseUint(tokenIDStr, 10, 64)
	if err != nil {
		return err
	}

	dataJSON, dataText := getData(msg.Data)

	return m.db.SaveNFT(tx.TxHash, tokenID, msg.DenomId, msg.Name, msg.URI, utils.SanitizeUTF8(dataJSON), dataText, msg.Recipient, msg.Sender, msg.ContractAddressSigner)
}

func (m *Module) handleMsgEditNFT(msg *nftTypes.MsgEditNFT) error {
	dataJSON, dataText := getData(msg.Data)

	return m.db.UpdateNFT(msg.Id, msg.DenomId, msg.Name, msg.URI, utils.SanitizeUTF8(dataJSON), dataText)
}

func (m *Module) handleMsgTransferNFT(msg *nftTypes.MsgTransferNft) error {
	return m.db.UpdateNFTOwner(msg.TokenId, msg.DenomId, msg.To)
}

func (m *Module) handleMsgBurnNFT(msg *nftTypes.MsgBurnNFT) error {
	return m.db.BurnNFT(msg.Id, msg.DenomId)
}

func (m *Module) handleMsgCreateCollection(tx *juno.Tx, msg *marketplaceTypes.MsgCreateCollection) error {
	dataJSON, dataText := getData(msg.Data)

	return m.db.SaveDenom(tx.TxHash, msg.Id, msg.Name, msg.Schema, msg.Symbol, msg.Creator, "",
		msg.Traits, msg.Minter, msg.Description, dataText, utils.SanitizeUTF8(dataJSON))
}

func getData(data string) (string, string) {
	dataText := data
	dataJSON := "{}"

	if data != "" && utils.IsJSON(data) {
		dataJSON = data
		dataText = ""
	}

	return dataJSON, dataText
}