package database_test

import (
	"github.com/forbole/bdjuno/v2/database"
	_ "github.com/lib/pq"
)

var expectedAppliedMigrations = []database.Migration{
	{ID: int64(1), Name: "000-initial_schema.sql", CreatedAt: int64(0)},
	{ID: int64(2), Name: "001-workers_storage.sql", CreatedAt: int64(0)},
	{ID: int64(3), Name: "002-inflation_calculation.sql", CreatedAt: int64(0)},
	{ID: int64(4), Name: "003-nft_module.sql", CreatedAt: int64(0)},
	{ID: int64(5), Name: "004-distinct_message_query_func.sql", CreatedAt: int64(0)},
	{ID: int64(6), Name: "005-group_module.sql", CreatedAt: int64(0)},
	{ID: int64(7), Name: "006-marketplace_module.sql", CreatedAt: int64(0)},
	{ID: int64(8), Name: "007-cw20token_module.sql", CreatedAt: int64(0)},
	{ID: int64(9), Name: "008-block_parsed_data.sql", CreatedAt: int64(0)},
	{ID: int64(10), Name: "009-cw20token_update.sql", CreatedAt: int64(0)},
	{ID: int64(11), Name: "010-nft-uniq-id.sql", CreatedAt: int64(0)},
	{ID: int64(12), Name: "011-nft-migrate-uniq-id-values.sql", CreatedAt: int64(0)},
	{ID: int64(13), Name: "012-marketplace-nft-id-column-unique.sql", CreatedAt: int64(0)},
}

func (suite *DbTestSuite) TestExecuteMigrations() {
	var rows []database.Migration
	suite.Require().NoError(suite.database.Sqlx.Select(&rows, `SELECT id, name FROM migrations`))
	suite.Require().Equal(expectedAppliedMigrations, rows)
}
