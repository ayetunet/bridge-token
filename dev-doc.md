# Ayetu Token Bridge Smart Contract Developer Documentation

## Overview

The `AyetuToken` contract is an ERC-20 compliant token on the Ethereum blockchain, featuring enhanced functionalities like minting, burning, and a specialized `bridgeBurn` mechanism for cross-chain interactions. It's designed with upgradeability in mind using the UUPS pattern. The contract is based on the reference OpenZeppelin ERC-20 token contracts. It is designed to act as the wrapped token component of a single-token bridge system for the Ayetu blockchain.

## Contract Features

- **ERC-20 Compliance**: Implements all standard ERC-20 token functionalities.
- **Mintable**: Allows designated addresses (owner or minter) to mint new tokens.
- **Burnable**: Permits token holders to burn their tokens, reducing the total supply.
- **BridgeBurn Functionality**: Specialized function for burning tokens and facilitating cross-chain operations.
- **Upgradeable**: Implemented using the UUPS upgradeable pattern for future enhancements.

## Key Functions

### Standard ERC-20 Functions

Implemented as per the OpenZeppelin `ERC20` standard, providing basic token functionalities:

- `transfer(address recipient, uint256 amount)`: Moves `amount` tokens from the caller's account to `recipient`.
- `approve(address spender, uint256 amount)`: Sets `amount` as the allowance of `spender` over the caller's tokens.
- `transferFrom(address sender, address recipient, uint256 amount)`: Moves `amount` tokens from `sender` to `recipient` using the allowance mechanism.
- `balanceOf(address account)`: Returns the amount of tokens owned by `account`.
- `allowance(address owner, address spender)`: Returns the remaining number of tokens that `spender` is allowed to spend on behalf of `owner`.

### Minting

- `mint(address to, uint256 amount)`: Mints `amount` tokens and assigns them to `to`, increasing the total supply. Can only be called by the owner or a designated minter.

### Burning

- `burn(uint256 amount)`: Destroys `amount` tokens from the caller's account, reducing the total supply.
- `burnFrom(address account, uint256 amount)`: Allows a spender to burn `amount` tokens from `account`, based on the allowance set by `account`.

### BridgeBurn

- `bridgeBurn(string memory recipient, uint256 amount)`: Burns `amount` tokens from the owner's account for cross-chain operations. Emits a `BridgeBurn` event with the specified `recipient` and `amount`. Callable only by the contract owner.

## Events

- `Transfer(address indexed from, address indexed to, uint256 value)`: Emitted when `value` tokens are moved from one account to another.
- `Approval(address indexed owner, address indexed spender, uint256 value)`: Emitted when a new allowance is set by a call to `approve`.
- `MinterChanged(address indexed newMinter)`: Emitted when a new minter is designated by the owner.
- `BridgeBurn(address indexed owner, string recipient, uint256 amount)`: Custom event for the `bridgeBurn` function, indicating cross-chain token burns.

## Roles and Access Control

- **Owner**: Initially set to the contract deployer. The owner has the authority to designate minters, call `bridgeBurn`, and manage upgrade operations.
- **Minter**: An address designated by the owner with the authority to mint new tokens.

## Upgradeability

This contract is upgradeable using the UUPS pattern, allowing future improvements and bug fixes without losing the current state or deployed address.

## Modifications from Previous Version

- **Direct Initialization**: Token name and symbol are now directly set in the `initialize` function.
- **NatSpec Comments**: Added comprehensive documentation comments for better readability and maintainability.
- **Access Control Modifier**: Introduced `onlyMinterOrOwner` modifier for the `mint` function to allow both the owner and the designated minter to mint new tokens.

## Security Considerations

- **Testing and Audits**: Extensive testing and potential audits are recommended to ensure the contract's security, especially for custom functionalities like `mint` and `bridgeBurn`.
- **Solidity Version**: The contract should be kept up-to-date with the latest Solidity version compatible with OpenZeppelin contracts for security and efficiency.

## Future Enhancements

- **Security Features**: Depending on use cases, consider adding features like pausing functionality, time-locked upgrades, or reentrancy guards for enhanced security.

This documentation should be updated alongside any future changes to the smart contract to ensure accuracy and clarity for developers and auditors.

## Copyright and License

This smart contract is based on reference contracts by OpenZeppelin. Modifications are copyright 2024 Douglas Horn. Available for use under MIT license.
