
pragma solidity ^0.4.24;

import "./lib/safeMath.sol";
import "./lib/set.sol";

contract Identity {
    using Set for Set.Data;
    using SafeMath for uint256;

    address owner; // owner has permissions to modify parameters
    bool public enabled = true; // if upgrade contract, then the old contract should be disabled

    Set.Data private users; // users
    mapping (address => string) private identities; // identities
    uint256 count = 0; // users count

    event NewIdentity(address who, string identity);
    event UpdateIdentity(address who, string identity);
    event RemoveIdentity(address who);

    modifier onlyOwner() {require(msg.sender == owner);_;}
    modifier onlyEnabled() {require(enabled);_;}

    constructor () public {
        owner = msg.sender;
    }

    function register(string content) public payable onlyEnabled {
        identities[msg.sender] = content;
        if (!users.contains(msg.sender)) {
            users.insert(msg.sender);
            count += 1;
            emit NewIdentity(msg.sender, content);
        } else {
            emit UpdateIdentity(msg.sender, content);
        }
    }

    function remove() public payable onlyEnabled {
        require(users.contains(msg.sender));
        users.remove(msg.sender);
        count -= 1;
        emit RemoveIdentity(msg.sender);
    }

    function get(address addr) public view returns (string) {
        require(users.contains(addr));
        return identities[addr];
    }

    function size() public view returns (uint256) {
        return count;
    }

    // owner can enable and disable rnode contract
    function enableContract() public onlyOwner {
        enabled = true;
    }

    function disableContract() public onlyOwner {
        enabled = false;
    }
}
