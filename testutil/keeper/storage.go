package keeper

import (
	"testing"
	"time"

	tmdb "github.com/cometbft/cometbft-db"
	"github.com/cometbft/cometbft/libs/log"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	authzkeeper "github.com/cosmos/cosmos-sdk/x/authz/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	crosschainkeeper "github.com/cosmos/cosmos-sdk/x/crosschain/keeper"
	crosschaintypes "github.com/cosmos/cosmos-sdk/x/crosschain/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/cosmos/cosmos-sdk/x/group"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	oracletypes "github.com/cosmos/cosmos-sdk/x/oracle/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	"github.com/stretchr/testify/require"

	"github.com/bnb-chain/greenfield/app"
	paymentmodulekeeper "github.com/bnb-chain/greenfield/x/payment/keeper"
	paymentmoduletypes "github.com/bnb-chain/greenfield/x/payment/types"
	permissionmodulekeeper "github.com/bnb-chain/greenfield/x/permission/keeper"
	permissionmoduletypes "github.com/bnb-chain/greenfield/x/permission/types"
	spkeeper "github.com/bnb-chain/greenfield/x/sp/keeper"
	sptypes "github.com/bnb-chain/greenfield/x/sp/types"
	"github.com/bnb-chain/greenfield/x/storage/keeper"
	"github.com/bnb-chain/greenfield/x/storage/types"
)

var (
	storageMaccPerms = map[string][]string{
		authtypes.Minter:               {authtypes.Minter},
		authtypes.FeeCollectorName:     {authtypes.Minter, authtypes.Staking},
		stakingtypes.BondedPoolName:    {authtypes.Burner, authtypes.Staking},
		stakingtypes.NotBondedPoolName: {authtypes.Burner, authtypes.Staking},
		govtypes.ModuleName:            {authtypes.Burner},
		sptypes.ModuleName:             {authtypes.Staking},
		types.ModuleName:               {authtypes.Staking},
		paymentmoduletypes.ModuleName:  {},
	}
)

type StorageDepKeepers struct {
	PaymentKeeper *paymentmodulekeeper.Keeper
	SpKeeper      *spkeeper.Keeper
	BankKeeper    *bankkeeper.BaseKeeper
	AccountKeeper *authkeeper.AccountKeeper
}

func StorageKeeper(t testing.TB) (*keeper.Keeper, StorageDepKeepers, sdk.Context) {
	storeKeys := sdk.NewKVStoreKeys(authtypes.StoreKey, authz.ModuleName, banktypes.StoreKey, stakingtypes.StoreKey,
		minttypes.StoreKey, distrtypes.StoreKey, slashingtypes.StoreKey, govtypes.StoreKey,
		paramstypes.StoreKey, upgradetypes.StoreKey, feegrant.StoreKey, evidencetypes.StoreKey,
		group.StoreKey,
		crosschaintypes.StoreKey,
		sptypes.StoreKey,
		paymentmoduletypes.StoreKey,
		permissionmoduletypes.StoreKey,
		oracletypes.StoreKey, types.StoreKey)

	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)

	storeKey := storetypes.NewKVStoreKey(types.StoreKey)
	tStorekey := storetypes.NewTransientStoreKey(types.TStoreKey)

	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, storetypes.StoreTypeMemory, nil)
	stateStore.MountStoreWithDB(storeKeys[paramstypes.StoreKey], storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(storeKeys[authtypes.StoreKey], storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(storeKeys[banktypes.StoreKey], storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(storeKeys[paymentmoduletypes.StoreKey], storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(storeKeys[sptypes.StoreKey], storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(storeKeys[crosschaintypes.StoreKey], storetypes.StoreTypeIAVL, db)
	require.NoError(t, stateStore.LoadLatestVersion())

	cdcConfig := app.MakeEncodingConfig()
	cdc := cdcConfig.Marshaler

	accountKeeper := authkeeper.NewAccountKeeper(
		cdc,
		storeKeys[authtypes.StoreKey],
		authtypes.ProtoBaseAccount,
		storageMaccPerms,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)

	authzKeeper := authzkeeper.NewKeeper(
		storeKeys[authz.ModuleName],
		cdc,
		baseapp.NewMsgServiceRouter(),
		accountKeeper,
	)

	bankKeeper := bankkeeper.NewBaseKeeper(
		cdc,
		storeKeys[banktypes.StoreKey],
		accountKeeper,
		nil,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)

	spKeeper := spkeeper.NewKeeper(
		cdc,
		storeKeys[sptypes.ModuleName],
		accountKeeper,
		bankKeeper,
		authzKeeper,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)

	paymentKeeper := paymentmodulekeeper.NewKeeper(
		cdc,
		storeKeys[paymentmoduletypes.StoreKey],

		bankKeeper,
		accountKeeper,
		spKeeper,

		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)
	crossChainKeeper := crosschainkeeper.NewKeeper(
		cdc,
		storeKeys[crosschaintypes.StoreKey],
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)

	permissionKeeper := permissionmodulekeeper.NewKeeper(
		cdc,
		storeKeys[permissionmoduletypes.ModuleName],
		accountKeeper,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)

	k := keeper.NewKeeper(
		cdc,
		storeKey,
		tStorekey,
		accountKeeper,
		spKeeper,
		paymentKeeper,
		permissionKeeper,
		crossChainKeeper,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, nil, log.NewNopLogger())
	ctx = ctx.WithBlockTime(time.Now())

	// Initialize params
	err := k.SetParams(ctx, types.DefaultParams())
	require.NoError(t, err)
	err = accountKeeper.SetParams(ctx, authtypes.DefaultParams())
	require.NoError(t, err)
	err = spKeeper.SetParams(ctx, sptypes.DefaultParams())
	require.NoError(t, err)
	err = paymentKeeper.SetParams(ctx, paymentmoduletypes.DefaultParams())
	require.NoError(t, err)

	// Initialize module accounts
	paymentModulePool := accountKeeper.GetModuleAccount(ctx, paymentmoduletypes.ModuleName)
	accountKeeper.SetModuleAccount(ctx, paymentModulePool)

	amount := sdk.NewIntFromUint64(1e19)
	err = bankKeeper.MintCoins(ctx, authtypes.Minter, sdk.Coins{sdk.Coin{
		Denom:  "BNB",
		Amount: amount,
	}})
	if err != nil {
		panic("mint coins error")
	}

	return k, StorageDepKeepers{
		SpKeeper:      spKeeper,
		PaymentKeeper: paymentKeeper,
		BankKeeper:    &bankKeeper,
		AccountKeeper: &accountKeeper,
	}, ctx
}
