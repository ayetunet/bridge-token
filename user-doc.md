# Ayetu Token Bridge User Guide

## Introduction

Welcome to the Ayetu Token Bridge user guide. AyetuToken (ATU) is an ERC-20 compliant token on the Ethereum blockchain, designed to facilitate cross-chain operations between Ethereum and the Ayetu Mainnet Blockchain. This document provides essential information for interacting with the AyetuToken, focusing on its unique `bridgeBurn` functionality.

## Key Features

- **ERC-20 Compliance**: Ayetu Token Bridge adheres to the standard ERC-20 token functionalities, allowing for basic operations like transfer, approve, and balance checks.
- **BridgeBurn**: A specialized function for burning ATU tokens on Ethereum and facilitating their transfer to the Ayetu Mainnet Blockchain.

## Using AyetuToken

### Standard Token Operations

As an ERC-20 token, AyetuToken supports standard operations which can be performed using any standard Ethereum wallet or interface that supports ERC-20 tokens:

- **Transfer**: Send ATU tokens to another Ethereum address.
- **Approve**: Allow a third-party contract or address to spend a specific amount of your ATU tokens.
- **Check Balance**: View your ATU token balance.

### BridgeBurn Functionality

The `bridgeBurn` function is a unique feature of AyetuToken designed for users who wish to transfer their ATU tokens from the Ethereum blockchain to the Ayetu Mainnet Blockchain. This function burns ATU tokens on Ethereum and signals the Ayetu Mainnet to issue the equivalent amount of native tokens to a specified Ayetu account.

#### Important Notes for Users:

- **Ayetu API**: It is strongly recommended to use the Ayetu API for interacting with the `bridgeBurn` functionality. The API ensures that the `recipient` data corresponds to a valid Ayetu account, minimizing the risk of token loss.
- **Direct Interaction**: Advanced users may interact directly with the `bridgeBurn` action through Ethereum interfaces (such as wallets or Web3 tools), but this is advised against due to the risk of tokens being lost if sent to an invalid Ayetu account.
- **Recipient Data**: When using `bridgeBurn`, ensure that the `recipient` data is accurate. The smart contract only validates the presence of data in the `recipient` field, not its validity as an Ayetu account.

### Performing a BridgeBurn

To use the `bridgeBurn` functionality through the Ayetu API:

1. **Access the Ayetu API**: Use the interface provided by Ayetu to initiate a bridge burn operation. The API will guide you through the necessary steps.
2. **Enter Recipient Details**: Input the Ayetu Mainnet account details to which you want to transfer your tokens.
3. **Confirm and Execute**: Follow the API's instructions to confirm and execute the bridge burn. The API handles the validation of the Ayetu account and the interaction with the smart contract.

## Risks and Considerations

- **Token Loss**: Directly interacting with the `bridgeBurn` function without validating the Ayetu account can result in irreversible token loss. Always prefer using the Ayetu API for such operations.
- **Smart Contract Security**: While the AyetuToken contract is designed with security in mind, users should remain aware of the general risks associated with smart contracts and token transactions on the Ethereum blockchain.

## Support

For assistance or more information about using AyetuToken and the Ayetu API, please visit [Ayetu Support](#) or contact our support team through the official channels.

This guide aims to provide intermediate computer users with the information needed to safely and effectively interact with AyetuToken. For any advanced operations or inquiries, professional advice or direct support from the Ayetu team is recommended.
