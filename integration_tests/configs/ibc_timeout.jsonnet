local ibc = import 'ibc.jsonnet';

ibc {
  'elysium_777-1'+: {
    genesis+: {
      app_state+: {
        elysium+: {
          params+: {
            ibc_timeout: 0,
          },
        },
      },
    },
  },
}
