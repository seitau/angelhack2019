pragma solidity ^0.5.0;

import "../installed_contracts/openzeppelin-solidity/contracts/token/ERC721/ERC721Full.sol";
import "../installed_contracts/openzeppelin-solidity/contracts/token/ERC721/ERC721Mintable.sol";

contract DisasterToken is ERC721Full, ERC721Mintable {

    struct Certificate {
        string title;
        string seriousness;
        string imageRef;
        string category;
        string date;
        string estimatedMoney;
    }
    mapping(uint256 => Certificate) private _certificates;

    constructor (string memory name, string memory symbol) public ERC721Full(name, symbol) {
        // solhint-disable-previous-line no-empty-blocks
    }

    function getCertificate(uint256 tokenId) public view returns (
        string memory,
        string memory,
        string memory,
        string memory,
        string memory,
        string memory
    ) {
        return (_certificates[tokenId].title,
            _certificates[tokenId].seriousness,
            _certificates[tokenId].imageRef,
            _certificates[tokenId].category,
            _certificates[tokenId].date,
            _certificates[tokenId].estimatedMoney);
    }

    function mintWithCertificate(
        address to,
        uint256 tokenId,
        string memory title,
        string memory seriousness,
        string memory category,
        string memory date
    ) public onlyMinter returns (bool) {
        _certificates[tokenId].title = title;
        _certificates[tokenId].seriousness = seriousness;
        _certificates[tokenId].category = category;
        _certificates[tokenId].date = date;
        return mint(to, tokenId);
    }

}
