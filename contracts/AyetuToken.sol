// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0; // Ensure compatibility with the latest Solidity version recommended by OpenZeppelin.

import "@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol";
import "@openzeppelin/contracts-upgradeable/token/ERC20/extensions/ERC20BurnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";

/**
 * @title AyetuToken
 * @dev ERC-20 token with minting, burning, and bridge burn functionalities for cross-chain interactions.
 * Implements an upgradeable contract using the UUPS pattern.
 */
contract AyetuToken is Initializable, ERC20Upgradeable, ERC20BurnableUpgradeable, OwnableUpgradeable, UUPSUpgradeable {
    address private _minter; // Minter address.

    event MinterChanged(address indexed newMinter);
    event BridgeBurn(address indexed owner, string recipient, uint256 amount);

    function initialize() public initializer {
        __ERC20_init("Ayetu", "ATU"); 
        __Ownable_init();
    }

    /**
     * @dev Restricts function access to either the owner or the designated minter.
     */
    modifier onlyMinterOrOwner() {
        require(_minter == _msgSender() || owner() == _msgSender(), "Caller is not the minter or the owner");
        _;
    }

    /**
     * @dev Sets a new minter address. Can only be called by the contract owner.
     * @param newMinter The address to be assigned as the new minter.
     */
    function setMinter(address newMinter) public onlyOwner {
        _minter = newMinter;
        emit MinterChanged(newMinter);
    }

    /**
     * @dev Mints new tokens to the specified address. Can only be called by the owner or the designated minter.
     * @param to The address to mint tokens to.
     * @param amount The amount of tokens to mint.
     */
    function mint(address to, uint256 amount) public onlyMinterOrOwner {
        _mint(to, amount);
    }

    /**
     * @dev Burns tokens and emits a custom event for cross-chain interactions. Can only be called by the contract owner.
     * @param recipient The identifier for the cross-chain recipient.
     * @param amount The amount of tokens to burn.
     */
    function bridgeBurn(string memory recipient, uint256 amount) public onlyOwner {
        require(bytes(recipient).length > 0, "Recipient cannot be empty");
        emit BridgeBurn(_msgSender(), recipient, amount);
        _burn(_msgSender(), amount);
    }

    /**
     * @dev Authorizes an upgrade to a new implementation. Restricted to the contract owner.
     * @param newImplementation The address of the new contract implementation.
     */
    function _authorizeUpgrade(address newImplementation) internal override onlyOwner {}

    // Additional security practices like reentrancy guards, pausable functionality,
    // or time-locks for critical operations can be considered based on contract usage.
}
