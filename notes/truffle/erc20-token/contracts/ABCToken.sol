pragma solidity ^0.4.23;

import "./utils/SafeMath.sol";
import "./utils/ERC20Token.sol";

// 代币支付的以太坊智能服务 https://ethfans.org/posts/ethereum-smart-service-payment-with-tokens

contract ABCToken is ERC20Token{
    using SafeMath for uint256;
    
    event Mint(address indexed to, uint256 amount);
    event MintFinished();
    event Burn(address indexed burner, uint256 value);

    bool public mintingFinished = false;

    modifier onlyOwner() {
        require(msg.sender == owner);
        _;
    }

    modifier canMint() {
        require(!mintingFinished);
        _;
    }

    constructor(string _name, string _symbol, uint256 _supply, uint8 _decimails) 
        ERC20Token(_name, _symbol, _supply, _decimails) public {
    }

    /**
    * @dev Function to mint tokens
    * @param _to The address that will receive the minted tokens.
    * @param _amount The amount of tokens to mint.
    * @return A boolean that indicates if the operation was successful.
    */
    function mint(address _to, uint256 _amount) onlyOwner public returns (bool) {
        // uint256 amount_wei = _amount.mul(uint256(10) ** decimals_);
        totalSupply_  = totalSupply_.add(_amount);
        balances[_to] = balances[_to].add(_amount);
        emit Mint(_to, _amount);
        emit Transfer(address(0), _to, _amount);
        return true;
    }

    /**
    * @dev Function to stop minting new tokens.
    * @return True if the operation was successful.
    */
    function finishMinting() onlyOwner canMint public returns (bool) {
        mintingFinished = true;
        emit MintFinished();
        return true;
    }
    
    /**
    * @dev Burns a specific amount of tokens.
    * @param _value The amount of token to be burned.
    */
    function burn(address _who, uint256 _value) onlyOwner public returns (bool) {
        require(_value <= balances[_who]);
        // no need to require value <= totalSupply, since that would imply the
        // sender's balance is greater than the totalSupply, which *should* be an assertion failure

        balances[_who] = balances[_who].sub(_value);
        totalSupply_   = totalSupply_.sub(_value);

        emit Burn(_who, _value);
        emit Transfer(_who, address(0), _value);
        return true;
    }

    /**
    * @dev Don't accept ETH
    */
    function () public payable {
        revert();
    }

}
