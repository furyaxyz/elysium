import json

from .utils import approve_proposal


def test_gov_update_params(elysium, tmp_path):
    cli = elysium.cosmos_cli()

    proposal = tmp_path / "proposal.json"
    # governance module account as signer
    signer = "did:fury:e10d07y265gmmuvt4z0w9aw880jnsr700jdufnyd"
    proposal_src = {
        "messages": [
            {
                "@type": "/elysium.MsgUpdateParams",
                "authority": signer,
                "params": {
                    "elysium_admin": "did:fury:e12luku6uxehhak02py4rcz65zu0swh7wjsrw0pp",
                    "enable_auto_deployment": False,
                    "ibc_ely_denom": "ibc/6411AE2ADA1E73DB59DB151"
                    "A8988F9B7D5E7E233D8414DB6817F8F1A01600000",
                    "ibc_timeout": "96400000000000",
                },
            }
        ],
        "deposit": "1basetely",
    }
    proposal.write_text(json.dumps(proposal_src))
    rsp = cli.submit_gov_proposal(proposal, from_="community")

    assert rsp["code"] == 0, rsp["raw_log"]
    approve_proposal(elysium, rsp)
    print("check params have been updated now")
    rsp = cli.query_params()
    print("params", rsp)
    assert rsp == {
        "elysium_admin": "did:fury:e12luku6uxehhak02py4rcz65zu0swh7wjsrw0pp",
        "enable_auto_deployment": False,
        "ibc_ely_denom": "ibc/6411AE2ADA1E73DB59DB151"
        "A8988F9B7D5E7E233D8414DB6817F8F1A01600000",
        "ibc_timeout": "96400000000000",
    }
