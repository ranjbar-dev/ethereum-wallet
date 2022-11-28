pragma solidity ^0.8.0;

contract ERC20 {
    string public constant name = "";
    string public constant symbol = "";
    uint8 public constant decimals = 0;

    function totalSupply() virtual public view returns (uint256) {}
    function balanceOf(address tokenOwner) virtual public view returns (uint256 balance) {}
    function allowance(address tokenOwner, address spender) virtual public view returns (uint256 remaining) {}
    function transfer(address to, uint256 tokens) virtual public returns (bool success) {}
    function approve(address spender, uint256 tokens) virtual public returns (bool success) {}
    function transferFrom(address from, address to, uint256 tokens) virtual public returns (bool success) {}

    event Transfer(address indexed from, address indexed to, uint256 tokens);
    event Approval(address indexed tokenOwner, address indexed spender, uint256 tokens);
}
