<!--
order: 2
-->

# State

The `x/elysium` module keeps the following objects in state:

|                         | Key                                    | Value                      |
| ----------------------- | -------------------------------------- | -------------------------- |
| DenomToExternalContract | `[]byte{1} + []byte(denom)`            | `[]byte(contract_address)` |
| DenomToAutoContract     | `[]byte{2} + []byte(denom)`            | `[]byte(contract_address)` |
| ContractToDenom         | `[]byte{3} + []byte(contract_address)` | `[]byte(denom)`            |

- `DenomToExternalContract` stores a map from denom to external FRC20 contract.
- `DenomToAutoContract` stores a map from denom to auto-deployed FRC20 contract.
- `ContractToDenom` stores the reversed map for both external and auto-deployed contracts.
