const DisasterToken = artifacts.require("./DisasterToken.sol");

module.exports = async (deployer, network, accounts) => {
    const name = "DisasterToken";
    const symbol = "DIT";
    deployer.deploy(DisasterToken, name, symbol);
};
