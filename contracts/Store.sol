// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

contract Store {
    mapping(address => uint) public count;

    function add(uint num) public {
        count[msg.sender] += num;
    }

    // Getter function for the cupcake balance of a user
    function getCount(address userAddress) public view returns (uint) {
        return count[userAddress];
    }
}