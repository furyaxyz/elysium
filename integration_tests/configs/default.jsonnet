{
  dotenv: '../../scripts/.env',
  'elysium_777-1': {
    cmd: 'elysiumd',
    'start-flags': '--trace',
    config: {
      mempool: {
        version: 'v1',
      },
    },
    'app-config': {
      'app-db-backend': 'pebbledb',
      'minimum-gas-prices': '0basetely',
      'index-events': ['ethereum_tx.ethereumTxHash'],
      'iavl-lazy-loading': true,
      'json-rpc': {
        address: '0.0.0.0:{EVMRPC_PORT}',
        'ws-address': '0.0.0.0:{EVMRPC_PORT_WS}',
        api: 'eth,net,web3,debug,elysium',
        'feehistory-cap': 100,
        'block-range-cap': 10000,
        'logs-cap': 10000,
      },
    },
    validators: [{
      coins: '1000000000000000000stake,10000000000000000000000basetely',
      staked: '1000000000000000000stake',
      mnemonic: '${VALIDATOR1_MNEMONIC}',
      'app-config': {
        store: {
          memiavl: true,
          'memiavl-zero-copy': true,
          streamers: ['versiondb'],
        },
      },
    }, {
      coins: '1000000000000000000stake,10000000000000000000000basetely',
      staked: '1000000000000000000stake',
      mnemonic: '${VALIDATOR2_MNEMONIC}',
    }],
    accounts: [{
      name: 'community',
      coins: '10000000000000000000000basetely',
      mnemonic: '${COMMUNITY_MNEMONIC}',
    }, {
      name: 'signer1',
      coins: '20000000000000000000000basetely',
      mnemonic: '${SIGNER1_MNEMONIC}',
    }, {
      name: 'signer2',
      coins: '30000000000000000000000basetely',
      mnemonic: '${SIGNER2_MNEMONIC}',
    }],
    genesis: {
      consensus_params: {
        block: {
          max_bytes: '1048576',
          max_gas: '81500000',
        },
      },
      app_state: {
        evm: {
          params: {
            evm_denom: 'basetely',
          },
        },
        elysium: {
          params: {
            elysium_admin: '${ELYSIUM_ADMIN}',
            enable_auto_deployment: true,
            ibc_ely_denom: '${IBC_ELY_DENOM}',
          },
        },
        gov: {
          voting_params: {
            voting_period: '10s',
          },
          deposit_params: {
            max_deposit_period: '10s',
            min_deposit: [
              {
                denom: 'basetely',
                amount: '1',
              },
            ],
          },
        },
        transfer: {
          params: {
            receive_enabled: true,
            send_enabled: true,
          },
        },
        feemarket: {
          params: {
            no_base_fee: false,
            base_fee: '100000000000',
          },
        },
      },
    },
  },
}
