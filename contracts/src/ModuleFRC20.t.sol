pragma solidity ^0.6.8;

import "ds-test/test.sol";

import "./ModuleFRC20.sol";

contract ModuleFRC20Test is DSTest {
    ModuleFRC20 token;

    function setUp() public {
        token = new ModuleFRC20("gravity0x0", 0);
    }

    function test_basic_sanity() public {
        assertEq(uint(token.decimals()), uint(0));
        assertEq(token.native_denom(), "gravity0x0");
    }

    function testFail_mint_by_elysium_module() public {
        token.mint_by_elysium_module(0x208AE63c976d145AB328afdcE251c7051D8E452D, 100);
    }

    function testFail_burn_by_elysium_module() public {
        token.burn_by_elysium_module(0x208AE63c976d145AB328afdcE251c7051D8E452D, 100);
    }
}
