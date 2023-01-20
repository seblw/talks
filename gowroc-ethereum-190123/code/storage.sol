// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

contract Storage {
    uint public value;

    event ValueChanged(address indexed sender, uint indexed value);

    function setValue(uint _value) public {
        value = _value;
        emit ValueChanged(msg.sender, _value);
    }
}
