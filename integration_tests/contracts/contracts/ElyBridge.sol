pragma solidity ^0.6.6;

contract ElyBridge {

    event __ElysiumSendElyToIbc(address sender, string recipient, uint256 amount);

    // Pay the contract a certain ELY amount and trigger a ELY transfer
    // from the contract to recipient through IBC
    function send_ely_to_crypto_org(string memory recipient) public payable {
        emit __ElysiumSendElyToIbc(msg.sender, recipient, msg.value);
    }
}
