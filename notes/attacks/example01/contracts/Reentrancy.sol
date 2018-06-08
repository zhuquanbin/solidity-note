pragma solidity ^0.4.23;


contract Reentrancy {
    address public owner;
    mapping (address => uint256) public balances;

    uint256 amountValue = 0;

    constructor() public {
        owner = msg.sender;
    }

    function balanceOf(address addr) public view returns(uint256) {
        return balances[addr];
    }

    function withdraw() public {
        // At this point, the caller's code is executed, and can call withdraw again
        require(msg.sender.call.value(balances[msg.sender])()); 
        balances[msg.sender] = 0;
    }

    function put() public payable {
        balances[msg.sender] += msg.value;
    }

    function () external payable {
        balances[msg.sender] += msg.value;
    }
}
