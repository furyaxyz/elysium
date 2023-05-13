package client

import (
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"

	"github.com/furyaxyz/elysium/v2/x/elysium/client/cli"
)

// ProposalHandler is the token mapping change proposal handler.
var ProposalHandler = govclient.NewProposalHandler(cli.NewSubmitTokenMappingChangeProposalTxCmd)
