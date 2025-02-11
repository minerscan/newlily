// Code generated by: `make actors-gen`. DO NOT EDIT.

package market

import (
	"bytes"
	"fmt"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lily/chain/actors"
	"github.com/ipfs/go-cid"
	cbg "github.com/whyrusleeping/cbor-gen"
	"golang.org/x/xerrors"

	"github.com/filecoin-project/lotus/chain/actors/adt"

	market9 "github.com/filecoin-project/go-state-types/builtin/v9/market"
	markettypes "github.com/filecoin-project/go-state-types/builtin/v9/market"
	adt9 "github.com/filecoin-project/go-state-types/builtin/v9/util/adt"
)

var _ State = (*state9)(nil)

func load9(store adt.Store, root cid.Cid) (State, error) {
	out := state9{store: store}
	err := store.Get(store.Context(), root, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func make9(store adt.Store) (State, error) {
	out := state9{store: store}

	s, err := market9.ConstructState(store)
	if err != nil {
		return nil, err
	}

	out.State = *s

	return &out, nil
}

type state9 struct {
	market9.State
	store adt.Store
}

func (s *state9) StatesChanged(otherState State) (bool, error) {
	otherState9, ok := otherState.(*state9)
	if !ok {
		// there's no way to compare different versions of the state, so let's
		// just say that means the state of balances has changed
		return true, nil
	}
	return !s.State.States.Equals(otherState9.State.States), nil
}

func (s *state9) States() (DealStates, error) {
	stateArray, err := adt9.AsArray(s.store, s.State.States, market9.StatesAmtBitwidth)
	if err != nil {
		return nil, err
	}
	return &dealStates9{stateArray}, nil
}

func (s *state9) ProposalsChanged(otherState State) (bool, error) {
	otherState9, ok := otherState.(*state9)
	if !ok {
		// there's no way to compare different versions of the state, so let's
		// just say that means the state of balances has changed
		return true, nil
	}
	return !s.State.Proposals.Equals(otherState9.State.Proposals), nil
}

func (s *state9) Proposals() (DealProposals, error) {
	proposalArray, err := adt9.AsArray(s.store, s.State.Proposals, market9.ProposalsAmtBitwidth)
	if err != nil {
		return nil, err
	}
	return &dealProposals9{proposalArray}, nil
}

type dealStates9 struct {
	adt.Array
}

func (s *dealStates9) Get(dealID abi.DealID) (*DealState, bool, error) {
	var deal9 market9.DealState
	found, err := s.Array.Get(uint64(dealID), &deal9)
	if err != nil {
		return nil, false, err
	}
	if !found {
		return nil, false, nil
	}
	deal := fromV9DealState(deal9)
	return &deal, true, nil
}

func (s *dealStates9) ForEach(cb func(dealID abi.DealID, ds DealState) error) error {
	var ds9 market9.DealState
	return s.Array.ForEach(&ds9, func(idx int64) error {
		return cb(abi.DealID(idx), fromV9DealState(ds9))
	})
}

func (s *dealStates9) decode(val *cbg.Deferred) (*DealState, error) {
	var ds9 market9.DealState
	if err := ds9.UnmarshalCBOR(bytes.NewReader(val.Raw)); err != nil {
		return nil, err
	}
	ds := fromV9DealState(ds9)
	return &ds, nil
}

func (s *dealStates9) array() adt.Array {
	return s.Array
}

func fromV9DealState(v9 market9.DealState) DealState {

	return (DealState)(v9)

}

type dealProposals9 struct {
	adt.Array
}

func (s *dealProposals9) Get(dealID abi.DealID) (*DealProposal, bool, error) {
	var proposal9 market9.DealProposal
	found, err := s.Array.Get(uint64(dealID), &proposal9)
	if err != nil {
		return nil, false, err
	}
	if !found {
		return nil, false, nil
	}

	proposal, err := fromV9DealProposal(proposal9)
	if err != nil {
		return nil, true, xerrors.Errorf("decoding proposal: %w", err)
	}

	return &proposal, true, nil
}

func (s *dealProposals9) ForEach(cb func(dealID abi.DealID, dp DealProposal) error) error {
	var dp9 market9.DealProposal
	return s.Array.ForEach(&dp9, func(idx int64) error {
		dp, err := fromV9DealProposal(dp9)
		if err != nil {
			return xerrors.Errorf("decoding proposal: %w", err)
		}

		return cb(abi.DealID(idx), dp)
	})
}

func (s *dealProposals9) decode(val *cbg.Deferred) (*DealProposal, error) {
	var dp9 market9.DealProposal
	if err := dp9.UnmarshalCBOR(bytes.NewReader(val.Raw)); err != nil {
		return nil, err
	}

	dp, err := fromV9DealProposal(dp9)
	if err != nil {
		return nil, err
	}

	return &dp, nil
}

func (s *dealProposals9) array() adt.Array {
	return s.Array
}

func fromV9DealProposal(v9 market9.DealProposal) (DealProposal, error) {

	label, err := fromV9Label(v9.Label)

	if err != nil {
		return DealProposal{}, xerrors.Errorf("error setting deal label: %w", err)
	}

	return DealProposal{
		PieceCID:     v9.PieceCID,
		PieceSize:    v9.PieceSize,
		VerifiedDeal: v9.VerifiedDeal,
		Client:       v9.Client,
		Provider:     v9.Provider,

		Label: label,

		StartEpoch:           v9.StartEpoch,
		EndEpoch:             v9.EndEpoch,
		StoragePricePerEpoch: v9.StoragePricePerEpoch,

		ProviderCollateral: v9.ProviderCollateral,
		ClientCollateral:   v9.ClientCollateral,
	}, nil
}

func (s *state9) DealProposalsAmtBitwidth() int {
	return market9.ProposalsAmtBitwidth
}

func (s *state9) DealStatesAmtBitwidth() int {
	return market9.StatesAmtBitwidth
}

func (s *state9) ActorKey() string {
	return actors.MarketKey
}

func (s *state9) ActorVersion() actors.Version {
	return actors.Version9
}

func (s *state9) Code() cid.Cid {
	code, ok := actors.GetActorCodeID(s.ActorVersion(), s.ActorKey())
	if !ok {
		panic(fmt.Errorf("didn't find actor %v code id for actor version %d", s.ActorKey(), s.ActorVersion()))
	}

	return code
}

func fromV9Label(v9 market9.DealLabel) (DealLabel, error) {
	if v9.IsString() {
		str, err := v9.ToString()
		if err != nil {
			return markettypes.EmptyDealLabel, xerrors.Errorf("failed to convert string label to string: %w", err)
		}
		return markettypes.NewLabelFromString(str)
	}

	bs, err := v9.ToBytes()
	if err != nil {
		return markettypes.EmptyDealLabel, xerrors.Errorf("failed to convert bytes label to bytes: %w", err)
	}
	return markettypes.NewLabelFromBytes(bs)
}
