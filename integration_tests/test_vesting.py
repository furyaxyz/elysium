import time


def test_create_account(elysium):
    """
    test create vesting account tx works:
    """
    cli = elysium.cosmos_cli()
    src = "vesting"
    addr = cli.create_account(src)["address"]
    denom = "basetely"
    balance = cli.balance(addr, denom)
    assert balance == 0
    amount = 10000
    fee = 4000000000000000
    amt = f"{amount}{denom}"
    end_time = int(time.time()) + 3000
    fees = f"{fee}{denom}"
    res = cli.create_vesting_account(addr, amt, end_time, from_="validator", fees=fees)
    assert res["code"] == 0, res["raw_log"]
    balance = cli.balance(addr, denom)
    assert balance == amount
