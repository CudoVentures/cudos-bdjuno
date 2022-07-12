package database

func (db *Db) SaveMsgMintNFT(txHash string, tokenID uint64, denomID string) error {
	_, err := db.Sql.Exec(`INSERT INTO nft_mint (transaction_hash, token_id, denom_id) 
		VALUES($1, $2, $3) ON CONFLICT DO NOTHING`, txHash, tokenID, denomID)
	return err
}