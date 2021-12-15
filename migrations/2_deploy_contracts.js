// Deploy Identity
var Identity = artifacts.require("./Identity.sol");

module.exports = function(deployer) {
        deployer.deploy(Identity); //"参数在第二个变量携带"
};
