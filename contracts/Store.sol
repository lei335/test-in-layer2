// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

contract Store {
    mapping(address => uint) public count;

    uint256 time = 1703062800;

    function add(uint num) public {
        count[msg.sender] += num;
    }

    // Getter function for the cupcake balance of a user
    function getCount(address userAddress) public view returns (uint) {
        return count[userAddress];
    }

    function testTime() public {
        uint256 nowtime = block.timestamp;
        if(nowtime<time+600){
            count[msg.sender] += 12;
        }
    }
}