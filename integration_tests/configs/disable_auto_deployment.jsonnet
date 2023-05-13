local config = import 'default.jsonnet';

config {
  'elysium_777-1'+: {
    'start-flags': '--trace --inv-check-period 5',
    'app-config'+: {
      'minimum-gas-prices':: super['minimum-gas-prices'],
      'json-rpc'+: {
        api:: super['api'],
      },
    },
    accounts: [{
      name: 'community',
      coins: '10000000000000000000000basetely',
      mnemonic: '${COMMUNITY_MNEMONIC}',
    }],
    genesis+: {
      app_state+: {
        elysium: {
          params: {
            elysium_admin: 'did:fury:e12luku6uxehhak02py4rcz65zu0swh7wjsrw0pp',
            enable_auto_deployment: false,
          },
        },
        transfer:: super['transfer'],
      },
      consensus_params+: {
        block+: {
           time_iota_ms: '2000',
        },
      },
    },
  },
}
