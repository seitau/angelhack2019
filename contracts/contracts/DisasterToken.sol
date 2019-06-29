pragma solidity ^0.5.0;

import "../installed_contracts/openzeppelin-solidity/contracts/token/ERC721/ERC721Full.sol";

contract DisasterToken is ERC721Full {
    constructor (string memory name, string memory symbol) public ERC721Full(name, symbol) {
        // solhint-disable-previous-line no-empty-blocks
    }
}
