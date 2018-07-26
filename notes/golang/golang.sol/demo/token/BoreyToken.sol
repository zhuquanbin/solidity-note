pragma solidity ^0.4.23;

contract BoreyToken {
    event Transfer(address indexed _from, address indexed _to, uint256 _value);

    address  public owner;
    mapping (address => uint)  public balanceOf;

    constructor(uint256 supply) public {
        owner = msg.sender;
        balanceOf[msg.sender] = supply;
    }

    function transfer(address _to, uint256 _value) public returns (bool) {
        require(_to != address(0));
        require(_value <= balanceOf[msg.sender]);

        balanceOf[msg.sender] = balanceOf[msg.sender] - _value;
        balanceOf[_to] = balanceOf[_to] + _value;
        emit Transfer(msg.sender, _to, _value);
        return true;
    }   
}
