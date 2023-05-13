<!--
order: 7
-->

# Parameters

The Elysium module contains the following parameters:

| Key                    | Type   | Default Value                                                |
| ---------------------- | ------ | ------------------------------------------------------------ |
| `IbcElyDenom`          | string | `"ibc/6B5A664BF0AF4F71B2F0BAA33141E2F1321242FBD5D19762F541EC971ACB0865"` |
| `IbcTimeout`           | uint64 | `86400000000000`                                             |
| `ElysiumAdmin`          | string | `""`                                                         |
| `EnableAutoDeployment` | bool   | `false`                                                      |

- `IbcElyDenom` Specifies the IBC token that should be converted to gas token upon arrival automatically.

  When update this parameter at runtime, the tokens are not migrated magically, might need to handle the token migration explicitly, e.g. some token swap mechanism.

- `IbcTimeout` The timeout value to use when Elysium module transfer tokens to IBC channels.

  Can be updated at runtime.

- `ElysiumAdmin` The account that is authorized to manage token mapping through message, empty means no admin, should be a valid bech32 cosmos address if specified.

  Can be updated at runtime.

- `EnableAutoDeployment` Specifies if the auto-deployment feature is enabled. 

  When disabled and there's no external contract mapped for the token, new coming tokens are kept as native tokens, user can transfer them back using cosmos native messages.

  Can be updated at runtime, after disabled at runtime, the previous deposited tokens can still be withdrawn.
