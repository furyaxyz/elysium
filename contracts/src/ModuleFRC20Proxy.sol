pragma solidity ^0.6.8;

import "ds-math/math.sol";
import "./ModuleFRC20.sol";

contract ModuleFRC20Proxy is DSMath {
    // sha256('elysium-evm')[:20]
    address constant module_address = 0x89A7EF2F08B1c018D5Cc88836249b84Dd5392905;
    ModuleFRC20 frc20Contract;
    bool isSource;

    event __ElysiumSendToIbc(address indexed sender, uint256 indexed channel_id, string recipient, uint256 amount, bytes extraData);
    event __ElysiumSendToEvmChain(address indexed sender, address indexed recipient, uint256 indexed chain_id, uint256 amount, uint256 bridge_fee, bytes extraData);
    event __ElysiumCancelSendToEvmChain(address indexed sender, uint256 id);

    /**
        Instantiate a ModuleFRC20Proxy contract. Need to set manually the frc20 contract authority to be the proxy
        like the following call:
        frc20Contract.setAuthority(DSAuthority(address(new ModuleFRC20ProxyAuthority(address(this)))));
    **/
    constructor(address frc20Contract_, bool isSource_) public {
        frc20Contract = ModuleFRC20(frc20Contract_);
        isSource = isSource_;
    }

    /**
        views
    **/
    function frc20() public view returns (address) {
        return address(frc20Contract);
    }

    function is_source() public view returns (bool) {
        return isSource;
    }


    /**
        Internal functions to be called by elysium module.
    **/
    function mint_by_elysium_module(address addr, uint amount) public {
        require(msg.sender == module_address);
        frc20Contract.mint(addr, amount);
    }

    function burn_by_elysium_module(address addr, uint amount) public {
        require(msg.sender == module_address);
        frc20_burn(addr, amount);
    }

    function transfer_by_elysium_module(address addr, uint amount) public {
        require(msg.sender == module_address);
        frc20Contract.move(addr, module_address, amount);
    }

    function transfer_from_elysium_module(address addr, uint amount) public {
        require(msg.sender == module_address);
        frc20Contract.move(address(this), addr, amount);
    }


    /**
        Evm hooks functions
    **/

    // send to another chain through gravity bridge, require approval for the burn.
    function send_to_evm_chain(address recipient, uint amount, uint chain_id, uint bridge_fee, bytes calldata extraData) external {
        // transfer back the token to the proxy account
        if (isSource) {
            frc20Contract.move(msg.sender, address(this), add(amount, bridge_fee));
        } else {
            frc20_burn(msg.sender, add(amount, bridge_fee));
        }
        emit __ElysiumSendToEvmChain(msg.sender, recipient, chain_id, amount, bridge_fee, extraData);
    }

    // send an "amount" of the contract token to recipient through IBC
    function send_to_ibc(string memory recipient, uint amount, uint channel_id, bytes memory extraData) public {
        if (isSource) {
            frc20Contract.move(msg.sender, address(this), amount);
        } else {
            frc20_burn(msg.sender, amount);
        }
        emit __ElysiumSendToIbc(msg.sender, channel_id, recipient, amount, extraData);
    }

    // cancel a send to chain transaction considering if it hasnt been batched yet.
    function cancel_send_to_evm_chain(uint256 id) external {
        emit __ElysiumCancelSendToEvmChain(msg.sender, id);
    }

    /**
        Internal functions
    **/

    // burn the token on behalf of the user. requires approval
    function frc20_burn(address addr, uint amount) internal {
        frc20Contract.move(addr, address(this), amount);
        frc20Contract.burn(amount);
    }
}