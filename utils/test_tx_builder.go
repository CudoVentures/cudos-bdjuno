package utils

import (
	"fmt"
	"strings"
	"time"

	abcitypes "github.com/tendermint/tendermint/abci/types"

	juno "github.com/forbole/juno/v2/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/group"
)

type testTxBuilder struct {
	events    []abcitypes.Event
	errors    []string
	timestamp time.Time
}

func NewTestTx(timestamp time.Time) *testTxBuilder {
	return &testTxBuilder{timestamp: timestamp}
}

func (builder *testTxBuilder) WithEventCreateGroup(groupID uint64, address string) *testTxBuilder {
	if address == "" {
		builder.errors = append(builder.errors, "error while building testTx: empty group address")

	}
	eventCreateGroup, err := sdk.TypedEventToEvent(&group.EventCreateGroup{GroupId: groupID})
	if err != nil {
		builder.errors = append(builder.errors, err.Error())
	}

	eventCreateGroupPolicy, err := sdk.TypedEventToEvent(&group.EventCreateGroupPolicy{Address: address})
	if err != nil {
		builder.errors = append(builder.errors, err.Error())
	}

	builder.events = append(builder.events, abcitypes.Event(eventCreateGroup), abcitypes.Event(eventCreateGroupPolicy))
	return builder
}

func (builder *testTxBuilder) WithEventSubmitProposal(proposalID uint64) *testTxBuilder {
	eventSubmitProposal, err := sdk.TypedEventToEvent(&group.EventSubmitProposal{ProposalId: proposalID})
	if err != nil {
		builder.errors = append(builder.errors, err.Error())
	}

	builder.events = append(builder.events, abcitypes.Event(eventSubmitProposal))
	return builder
}

func (builder *testTxBuilder) WithEventExec(result group.ProposalExecutorResult) *testTxBuilder {
	eventExec, err := sdk.TypedEventToEvent(&group.EventExec{Result: result, Logs: "1"})
	if err != nil {
		builder.errors = append(builder.errors, err.Error())
	}

	builder.events = append(builder.events, abcitypes.Event(eventExec))
	return builder
}

func (builder *testTxBuilder) WithEventVote() *testTxBuilder {
	eventVote, err := sdk.TypedEventToEvent(&group.EventVote{})
	if err != nil {
		builder.errors = append(builder.errors, err.Error())
	}

	builder.events = append(builder.events, abcitypes.Event(eventVote))
	return builder
}

func (builder *testTxBuilder) WithEventWithdrawProposal() *testTxBuilder {
	eventWithdraw, err := sdk.TypedEventToEvent(&group.EventWithdrawProposal{})
	if err != nil {
		builder.errors = append(builder.errors, err.Error())
	}

	builder.events = append(builder.events, abcitypes.Event(eventWithdraw))
	return builder
}

func (builder *testTxBuilder) Build() (*juno.Tx, error) {
	if len(builder.errors) > 0 {
		return &juno.Tx{}, fmt.Errorf(`error while building testTx: %s`, strings.Join(builder.errors, "\n"))
	}
	txLog := sdk.ABCIMessageLogs{{MsgIndex: 0, Events: sdk.StringifyEvents(builder.events)}}
	txResponse := sdk.TxResponse{
		TxHash:    "1",
		Logs:      txLog,
		Timestamp: builder.timestamp.Format(time.RFC3339),
		Height:    1,
	}

	return &juno.Tx{TxResponse: &txResponse}, nil
}