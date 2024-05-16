// SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;

contract StringStorage {
    address private owner;
    mapping(string => uint256) private stringMap;

    event StringAdded(string, uint);

    modifier onlyOwner() {
        require(msg.sender == owner, "Only owner can call this function");
        _;
    }

    constructor() {
        owner = msg.sender;
    }

    function addString(string memory _str, uint _createdTime ) external onlyOwner {
        stringMap[_str]= _createdTime;
        emit StringAdded(_str, _createdTime);
    }

    function getString(string memory _str) external view returns (uint) {
        return stringMap[_str];
    }

}