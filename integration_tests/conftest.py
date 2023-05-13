import os
import sys
from pathlib import Path

import pytest

from .network import setup_elysium, setup_custom_elysium, setup_geth

dir = os.path.dirname(os.path.realpath(__file__))
sys.path.append(dir + "/protobuf")


def pytest_configure(config):
    config.addinivalue_line("markers", "slow: marks tests as slow")
    config.addinivalue_line("markers", "gravity: gravity bridge test cases")


@pytest.fixture(scope="session")
def suspend_capture(pytestconfig):
    """
    used to pause in testing

    Example:
    ```
    def test_simple(suspend_capture):
        with suspend_capture:
            # read user input
            print(input())
    ```
    """

    class SuspendGuard:
        def __init__(self):
            self.capmanager = pytestconfig.pluginmanager.getplugin("capturemanager")

        def __enter__(self):
            self.capmanager.suspend_global_capture(in_=True)

        def __exit__(self, _1, _2, _3):
            self.capmanager.resume_global_capture()

    yield SuspendGuard()


@pytest.fixture(scope="session", params=[True])
def elysium(request, tmp_path_factory):
    enable_indexer = request.param
    if enable_indexer:
        path = tmp_path_factory.mktemp("indexer")
        yield from setup_custom_elysium(
            path, 27000, Path(__file__).parent / "configs/enable-indexer.jsonnet"
        )
    else:
        path = tmp_path_factory.mktemp("elysium")
        yield from setup_elysium(path, 26650)


@pytest.fixture(scope="session")
def geth(tmp_path_factory):
    path = tmp_path_factory.mktemp("geth")
    yield from setup_geth(path, 8545)


@pytest.fixture(scope="session", params=["elysium", "geth", "elysium-ws"])
def cluster(request, elysium, geth):
    """
    run on both elysium and geth
    """
    provider = request.param
    if provider == "elysium":
        yield elysium
    elif provider == "geth":
        yield geth
    elif provider == "elysium-ws":
        elysium_ws = elysium.copy()
        elysium_ws.use_websocket()
        yield elysium_ws
    else:
        raise NotImplementedError
