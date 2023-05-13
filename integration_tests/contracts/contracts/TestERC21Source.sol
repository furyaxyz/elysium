pragma solidity 0.8.10;
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract TestERC21Source is ERC20 {
	// sha256('elysium-evm')[:20]
	address constant module_address = 0x89A7EF2F08B1c018D5Cc88836249b84Dd5392905;
	string denom;
	bool isSource;

	event __ElysiumSendToIbc(address sender, string recipient, uint256 amount);
	event __ElysiumSendToIbc(address indexed sender, uint256 indexed channel_id, string recipient, uint256 amount, bytes extraData);
	event __ElysiumSendToEvmChain(address indexed sender, address indexed recipient, uint256 indexed chain_id, uint256 amount, uint256 bridge_fee, bytes extraData);
	event __ElysiumCancelSendToEvmChain(address indexed sender, uint256 id);

	constructor() ERC20("Doggo", "DOG") public {
		denom = "Doggo";
		isSource = true;
		_mint(msg.sender, 100000000000000000000000000);
	}

	/**
        views
    **/
	function native_denom() public view returns (string memory) {
		return denom;
	}

	function is_source() public view returns (bool) {
		return isSource;
	}


	/**
        Internal functions to be called by elysium module
    **/
	function mint_by_elysium_module(address addr, uint amount) public {
		require(msg.sender == module_address);
		_mint(addr, amount);
	}

	function burn_by_elysium_module(address addr, uint amount) public {
		require(msg.sender == module_address);
		unsafe_burn(addr, amount);
	}

	function transfer_by_elysium_module(address addr, uint amount) public {
		require(msg.sender == module_address);
		unsafe_transfer(addr, module_address, amount);
	}

	function transfer_from_elysium_module(address addr, uint amount) public {
		require(msg.sender == module_address);
		transfer(addr, amount);
	}

	/**
        Evm hooks functions
    **/

	// send an "amount" of the contract token to recipient through IBC
	function send_to_ibc(string memory recipient, uint amount) public {
		if (isSource) {
			transfer(module_address, amount);
		} else {
			unsafe_burn(msg.sender, amount);
		}
		emit __ElysiumSendToIbc(msg.sender, recipient, amount);
	}

	function send_to_ibc_v2(string memory recipient, uint amount, uint channel_id, bytes memory extraData) public {
		if (isSource) {
			transfer(module_address, amount);
		} else {
			unsafe_burn(msg.sender, amount);
		}
		emit __ElysiumSendToIbc(msg.sender, channel_id, recipient, amount, extraData);
	}

	// send to another chain through gravity bridge
	function send_to_evm_chain(address recipient, uint amount, uint chain_id, uint bridge_fee, bytes calldata extraData) external {
		if (isSource) {
			transfer(module_address, amount + bridge_fee);
		} else {
			unsafe_burn(msg.sender, amount + bridge_fee);
		}
		emit __ElysiumSendToEvmChain(msg.sender, recipient, chain_id, amount, bridge_fee, extraData);
	}

	// cancel a send to chain transaction considering if it hasnt been batched yet.
	function cancel_send_to_evm_chain(uint256 id) external {
		emit __ElysiumCancelSendToEvmChain(msg.sender, id);
	}

	/**
        Internal functions
    **/

	// unsafe_burn burn tokens without user's approval and authentication, used internally
	function unsafe_burn(address addr, uint amount) internal {
		_burn(addr, amount);
	}

	// unsafe_transfer transfer tokens without user's approval and authentication, used internally
	function unsafe_transfer(address src, address dst, uint amount) internal {
		_transfer(src, dst, amount);
	}
}