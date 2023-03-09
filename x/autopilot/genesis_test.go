package autopilot_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/Stride-Labs/stride/v6/app/apptesting"
	"github.com/Stride-Labs/stride/v6/x/autopilot"
	"github.com/Stride-Labs/stride/v6/x/autopilot/types"
)

func TestGenesis(t *testing.T) {
	expectedGenesisState := types.GenesisState{
		Params: types.Params{Active: true},
	}

	s := apptesting.SetupSuitelessTestHelper()
	autopilot.InitGenesis(s.Ctx, s.App.AutopilotKeeper, expectedGenesisState)

	actualGenesisState := autopilot.ExportGenesis(s.Ctx, s.App.AutopilotKeeper)
	require.NotNil(t, actualGenesisState)
	require.Equal(t, expectedGenesisState.Params, actualGenesisState.Params)
}