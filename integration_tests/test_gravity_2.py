import pytest
from eth_account.account import Account
from hexbytes import HexBytes

from .gravity_utils import prepare_gravity, setup_cosmos_erc20_contract
from .network import setup_elysium, setup_geth
from .utils import (
    ADDRS,
    CONTRACTS,
    KEYS,
    deploy_contract,
    deploy_erc20,
    eth_to_bech32,
    send_to_cosmos,
    send_transaction,
    setup_token_mapping,
    wait_for_fn,
    wait_for_new_blocks,
)

pytestmark = pytest.mark.gravity

Account.enable_unaudited_hdwallet_features()


@pytest.fixture(scope="module")
def custom_geth(tmp_path_factory):
    yield from setup_geth(tmp_path_factory.mktemp("geth"), 8555)


@pytest.fixture(scope="module", params=[True, False])
def custom_elysium(request, tmp_path_factory):
    yield from setup_elysium(tmp_path_factory.mktemp("elysium"), 26600, request.param)


@pytest.fixture(scope="module")
def gravity(custom_elysium, custom_geth):
    yield from prepare_gravity(custom_elysium, custom_geth)


def test_gravity_proxy_contract(gravity):
    if not gravity.elysium.enable_auto_deployment:
        geth = gravity.geth

        # deploy test erc20 contract
        erc20 = deploy_contract(
            geth,
            CONTRACTS["TestERC20A"],
        )
        balance = erc20.caller.balanceOf(ADDRS["validator"])
        assert balance == 100000000000000000000000000

        denom = f"gravity{erc20.address}"

        # deploy frc20 contract
        w3 = gravity.elysium.w3
        frc20 = deploy_contract(w3, CONTRACTS["TestFRC20"])

        print("frc20 contract deployed at address: ", frc20.address)

        # setup the contract mapping
        elysium_cli = gravity.elysium.cosmos_cli()

        print("check the contract mapping not exists yet")
        with pytest.raises(AssertionError):
            elysium_cli.query_contract_by_denom(denom)

        rsp = elysium_cli.update_token_mapping(
            denom, frc20.address, "TEST", 18, from_="validator"
        )
        assert rsp["code"] == 0, rsp["raw_log"]
        wait_for_new_blocks(elysium_cli, 1)

        print("check the contract mapping exists now")
        rsp = elysium_cli.query_denom_by_contract(frc20.address)
        assert rsp["denom"] == denom

        # Send some tokens
        print("send to elysium frc20")
        amount = 1000
        recipient = HexBytes(ADDRS["community"])
        txreceipt = send_to_cosmos(
            gravity.contract, erc20, geth, recipient, amount, KEYS["validator"]
        )
        assert txreceipt.status == 1, "should success"
        assert erc20.caller.balanceOf(ADDRS["validator"]) == balance - amount

        def check_gravity_tokens():
            "check the balance of gravity native token"
            return frc20.caller.balanceOf(ADDRS["community"]) == amount

        wait_for_fn("check_gravity_tokens", check_gravity_tokens)

        # deploy frc20 proxy contract
        proxyfrc20 = deploy_contract(
            w3,
            CONTRACTS["TestFRC20Proxy"],
            (frc20.address, False),
        )

        print("proxyfrc20 contract deployed at address: ", proxyfrc20.address)
        assert not proxyfrc20.caller.is_source()
        assert proxyfrc20.caller.frc20() == frc20.address

        # change token mapping
        rsp = elysium_cli.update_token_mapping(
            denom, proxyfrc20.address, "DOG", 18, from_="validator"
        )
        assert rsp["code"] == 0, rsp["raw_log"]
        wait_for_new_blocks(elysium_cli, 1)

        print("check the contract mapping exists now")
        rsp = elysium_cli.query_denom_by_contract(proxyfrc20.address)
        assert rsp["denom"] == denom

        # Fund the proxy contract cosmos account with original supply
        # by sending tokens to dead address
        # (because mint to zero address is forbidden in ERC20 contract)
        print("restore original supply frc20 by sending token to dead address")
        amount = 1000
        balance = erc20.caller.balanceOf(ADDRS["validator"])
        dead_address = "0x000000000000000000000000000000000000dEaD"
        cosmos_dead_address = HexBytes(dead_address)
        txreceipt = send_to_cosmos(
            gravity.contract,
            erc20,
            geth,
            cosmos_dead_address,
            amount,
            KEYS["validator"],
        )
        assert txreceipt.status == 1, "should success"
        assert erc20.caller.balanceOf(ADDRS["validator"]) == balance - amount

        def check_dead_gravity_tokens():
            "check the balance of gravity token"
            return frc20.caller.balanceOf(dead_address) == amount

        wait_for_fn("check_dead_gravity_tokens", check_dead_gravity_tokens)

        # Try to send back token to ethereum
        amount = 500
        ethereum_receiver = ADDRS["validator2"]
        # community_balance_before_send = frc20.caller.balanceOf(community)
        balance_before_send_to_ethereum = erc20.caller.balanceOf(ethereum_receiver)

        print("send to ethereum")
        # First we need to approve the proxy contract to move asset
        tx = frc20.functions.approve(proxyfrc20.address, amount).build_transaction(
            {"from": ADDRS["community"]}
        )
        txreceipt = send_transaction(w3, tx, key=KEYS["community"])
        assert txreceipt.status == 1, "should success"
        assert frc20.caller.allowance(ADDRS["community"], proxyfrc20.address) == amount

        # Then trigger the send to evm chain
        sender = ADDRS["community"]
        community_balance_before_send = frc20.caller.balanceOf(sender)
        print(
            "sender address : ",
        )
        tx2 = proxyfrc20.functions.send_to_evm_chain(
            ethereum_receiver, amount, 1, 0, b""
        ).build_transaction({"from": ADDRS["community"]})
        txreceipt2 = send_transaction(w3, tx2, key=KEYS["community"])
        print("receipt : ", txreceipt2)
        assert txreceipt2.status == 1, "should success"
        # Check deduction
        assert frc20.caller.balanceOf(sender) == community_balance_before_send - amount

        balance_after_send_to_ethereum = balance_before_send_to_ethereum

        def check_ethereum_balance_change():
            nonlocal balance_after_send_to_ethereum
            balance_after_send_to_ethereum = erc20.caller.balanceOf(ethereum_receiver)
            print("balance dead address", frc20.caller.balanceOf(dead_address))
            return balance_before_send_to_ethereum != balance_after_send_to_ethereum

        wait_for_fn(
            "ethereum balance change", check_ethereum_balance_change, timeout=60
        )
        assert (
            balance_after_send_to_ethereum == balance_before_send_to_ethereum + amount
        )


def test_gravity_proxy_contract_source_token(gravity):
    if not gravity.elysium.enable_auto_deployment:
        # deploy contracts
        w3 = gravity.elysium.w3
        symbol = "TEST"
        contract, denom = setup_token_mapping(gravity.elysium, "TestFRC20", symbol)
        cosmos_erc20_contract = setup_cosmos_erc20_contract(
            gravity,
            denom,
            symbol,
        )
        # setup the contract mapping
        elysium_cli = gravity.elysium.cosmos_cli()

        # deploy frc20 proxy contract
        proxyfrc20 = deploy_contract(
            w3,
            CONTRACTS["TestFRC20Proxy"],
            (contract.address, True),
        )

        print("proxyfrc20 contract deployed at address: ", proxyfrc20.address)
        assert proxyfrc20.caller.is_source()
        assert proxyfrc20.caller.frc20() == contract.address

        # change token mapping
        rsp = elysium_cli.update_token_mapping(
            denom, proxyfrc20.address, symbol, 6, from_="validator"
        )
        assert rsp["code"] == 0, rsp["raw_log"]
        wait_for_new_blocks(elysium_cli, 1)

        print("check the contract mapping exists now")
        rsp = elysium_cli.query_denom_by_contract(proxyfrc20.address)
        assert rsp["denom"] == denom

        # Try to send token to ethereum
        amount = 500
        ethereum_receiver = ADDRS["validator"]
        sender = ADDRS["validator"]
        # community_balance_before_send = frc20.caller.balanceOf(community)
        balance_before_send_to_ethereum = cosmos_erc20_contract.caller.balanceOf(
            ethereum_receiver
        )

        print("send to ethereum")
        # First we need to approve the proxy contract to move asset
        tx = contract.functions.approve(proxyfrc20.address, amount).build_transaction(
            {"from": sender}
        )
        txreceipt = send_transaction(w3, tx, key=KEYS["validator"])
        assert txreceipt.status == 1, "should success"
        assert (
            contract.caller.allowance(ADDRS["validator"], proxyfrc20.address) == amount
        )

        # Then trigger the send to evm chain
        community_balance_before_send = contract.caller.balanceOf(sender)
        tx2 = proxyfrc20.functions.send_to_evm_chain(
            ethereum_receiver, amount, 1, 0, b""
        ).build_transaction({"from": sender})
        txreceipt2 = send_transaction(w3, tx2, key=KEYS["validator"])
        print("receipt : ", txreceipt2)
        assert txreceipt2.status == 1, "should success"
        # Check deduction
        assert (
            contract.caller.balanceOf(sender) == community_balance_before_send - amount
        )

        balance_after_send_to_ethereum = balance_before_send_to_ethereum

        def check_ethereum_balance_change():
            nonlocal balance_after_send_to_ethereum
            balance_after_send_to_ethereum = cosmos_erc20_contract.caller.balanceOf(
                ethereum_receiver
            )
            return balance_before_send_to_ethereum != balance_after_send_to_ethereum

        wait_for_fn(
            "ethereum balance change", check_ethereum_balance_change, timeout=60
        )
        assert (
            balance_after_send_to_ethereum == balance_before_send_to_ethereum + amount
        )

        # Send back token to elysium
        elysium_receiver = ADDRS["community"]
        balance_before_send_to_cosmos = contract.caller.balanceOf(elysium_receiver)
        amount = 15
        txreceipt = send_to_cosmos(
            gravity.contract,
            cosmos_erc20_contract,
            gravity.geth,
            HexBytes(elysium_receiver),
            amount,
            KEYS["validator"],
        )
        assert txreceipt.status == 1, "should success"

        balance_after_send_to_cosmos = balance_before_send_to_cosmos

        def check_elysium_balance_change():
            nonlocal balance_after_send_to_cosmos
            balance_after_send_to_cosmos = contract.caller.balanceOf(elysium_receiver)
            return balance_before_send_to_cosmos != balance_after_send_to_cosmos

        wait_for_fn("check elysium balance change", check_elysium_balance_change)
        assert balance_after_send_to_cosmos == balance_before_send_to_cosmos + amount


def test_gravity_detect_malicious_supply(gravity):
    if not gravity.elysium.enable_auto_deployment:
        geth = gravity.geth
        cli = gravity.elysium.cosmos_cli()

        # deploy fake contract to trigger the malicious supply
        # any transfer made with this contract will send an amount of token
        # equal to max uint256
        erc20 = deploy_contract(
            geth,
            CONTRACTS["TestMaliciousSupply"],
        )
        denom = f"gravity{erc20.address}"
        print(denom)

        # check that the bridge is activated
        activate = cli.query_gravity_params()["params"]["bridge_active"]
        assert activate is True

        max_int = 2**256 - 1
        print("send max_int to community address using gravity bridge")
        recipient = HexBytes(ADDRS["community"])
        txreceipt = send_to_cosmos(
            gravity.contract, erc20, geth, recipient, max_int, KEYS["validator"]
        )
        assert txreceipt.status == 1, "should success"

        # check that amount has been received
        def check_gravity_native_tokens():
            "check the balance of gravity native token"
            return cli.balance(eth_to_bech32(recipient), denom=denom) == max_int

        wait_for_fn("balance", check_gravity_native_tokens)

        # check that the bridge is still activated
        activate = cli.query_gravity_params()["params"]["bridge_active"]
        assert activate is True

        # need a random transferFrom to increment the counter in the contract
        # (see logic) to be able to redo a max uint256 transfer
        print("do a random send to increase contract nonce")
        txtransfer = erc20.functions.transferFrom(
            ADDRS["validator"], ADDRS["validator2"], 1
        ).build_transaction({"from": ADDRS["validator"]})
        txreceipt = send_transaction(geth, txtransfer, KEYS["validator"])
        assert txreceipt.status == 1, "should success"

        print("send again max_int to community address using gravity bridge")
        txreceipt = send_to_cosmos(
            gravity.contract, erc20, geth, recipient, max_int, KEYS["validator"]
        )
        assert txreceipt.status == 1, "should success"

        # Wait enough for orchestrator to relay the event
        wait_for_new_blocks(cli, 30)

        # check that the bridge has not been desactivated
        activate = cli.query_gravity_params()["params"]["bridge_active"]
        assert activate is True

        # check that balance is still same
        assert cli.balance(eth_to_bech32(recipient), denom=denom) == max_int


def test_gravity_non_cosmos_denom(gravity):
    if gravity.elysium.enable_auto_deployment:
        return

    elysium_cli = gravity.elysium.cosmos_cli()
    # deploy test erc20 contract
    erc20 = deploy_contract(
        gravity.geth,
        CONTRACTS["TestERC20A"],
    )
    print("send to elysium frc20")
    recipient = HexBytes(ADDRS["community"])
    balance = erc20.caller.balanceOf(ADDRS["validator"])
    assert balance == 100000000000000000000000000
    amount = 100
    txreceipt = send_to_cosmos(
        gravity.contract, erc20, gravity.geth, recipient, amount, KEYS["validator"]
    )
    assert txreceipt.status == 1, "should success"
    assert erc20.caller.balanceOf(ADDRS["validator"]) == balance - amount

    denom = f"gravity{erc20.address}"

    def check_gravity_native_tokens():
        "check the balance of gravity native token"
        return elysium_cli.balance(eth_to_bech32(recipient), denom=denom) == amount

    wait_for_fn("send-to-gravity-native", check_gravity_native_tokens)

    # Deploy a bad cosmos erc20 token with single character
    # Cosmos denom must be 3 ~ 128 characters long and support letters,
    # followed by either a letter, a number or a separator
    # ('/', ':', '.', '_' or '-').
    print("Deploy cosmos erc20 contract on ethereum")
    tx_receipt = deploy_erc20(
        gravity.contract, gravity.geth, "A", "A", "DOG", 6, KEYS["validator"]
    )
    assert tx_receipt.status == 1, "should success"

    # Wait enough for orchestrator to relay the event
    wait_for_new_blocks(elysium_cli, 30)

    # Send again token to elysium and verify that the network is not stopped
    print("send to elysium frc20")
    recipient = HexBytes(ADDRS["community"])
    balance = erc20.caller.balanceOf(ADDRS["validator"])
    txreceipt = send_to_cosmos(
        gravity.contract, erc20, gravity.geth, recipient, amount, KEYS["validator"]
    )
    assert txreceipt.status == 1, "should success"
    assert erc20.caller.balanceOf(ADDRS["validator"]) == balance - amount

    denom = f"gravity{erc20.address}"

    def check_gravity_native_tokens():
        "check the balance of gravity native token"
        return elysium_cli.balance(eth_to_bech32(recipient), denom=denom) == 2 * amount

    wait_for_fn("send-to-gravity-native", check_gravity_native_tokens)
