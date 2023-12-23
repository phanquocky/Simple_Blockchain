// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

// ERC721 interface
interface IERC721 {
    event Transfer(
        address indexed _from,
        address indexed _to,
        uint256 indexed _tokenId
    );
    event Approval(
        address indexed _owner,
        address indexed _approved,
        uint256 indexed _tokenId
    );
    event ApprovalForAll(
        address indexed _owner,
        address indexed _operator,
        bool _approved
    );

    function balanceOf(address _owner) external view returns (uint256);

    function ownerOf(uint256 _tokenId) external view returns (address);

    function safeTransferFrom(
        address _from,
        address _to,
        uint256 _tokenId,
        bytes memory data
    ) external payable;

    function safeTransferFrom(
        address _from,
        address _to,
        uint256 _tokenId
    ) external payable;

    function transferFrom(
        address _from,
        address _to,
        uint256 _tokenId
    ) external payable;

    function approve(address _approved, uint256 _tokenId) external payable;

    function setApprovalForAll(address _operator, bool _approved) external;

    function getApproved(uint256 _tokenId) external view returns (address);

    function isApprovedForAll(
        address _owner,
        address _operator
    ) external view returns (bool);
}

// IERC721Receiver is interface for smartcontract want to receive ERC721 token
interface IERC721Receiver {
    function onERC721Received(
        address _operator,
        address _from,
        uint _tokenId,
        bytes calldata _data
    ) external returns (bytes4);
}

// ErrorZeroAddress error is error when address is zero
error ErrorZeroAddress();

// Error when don't have permission except owner
error ErrorUnauthorize();

contract ERC721 is IERC721 {
    mapping(uint => address) internal _ownerOf;
    mapping(address => uint) internal _balanceOf;
    mapping(uint => address) internal _approvals;
    mapping(address => mapping(address => bool)) public isApprovedForAll;

    modifier notAddressZero(address _addr) {
        if (_addr == address(0)) {
            revert ErrorZeroAddress();
        }
        _;
    }

    modifier isOwner(address _owner, address _addr) {
        if (_addr != _owner) {
            revert ErrorUnauthorize();
        }
        _;
    }

    function balanceOf(
        address _owner
    ) external view notAddressZero(_owner) returns (uint256) {
        return _balanceOf[_owner];
    }

    function ownerOf(uint256 _tokenId) external view returns (address) {
        address owner = _ownerOf[_tokenId];
        if (owner == address(0)) {
            revert ErrorZeroAddress();
        }

        return owner;
    }

    function setApprovalForAll(address _operator, bool _approved) external {
        isApprovedForAll[msg.sender][_operator] = _approved;

        emit ApprovalForAll(msg.sender, _operator, _approved);
    }

    function approve(address _approved, uint256 _tokenId) external payable {
        address owner = _ownerOf[_tokenId];
        if (owner == address(0)) {
            revert ErrorZeroAddress();
        }

        if (owner != msg.sender) {
            revert ErrorUnauthorize();
        }

        _approvals[_tokenId] = _approved;

        emit Approval(msg.sender, _approved, _tokenId);
    }

    function getApproved(
        uint256 _tokenId
    ) external view notAddressZero(_ownerOf[_tokenId]) returns (address) {
        return _approvals[_tokenId];
    }

    function _isApprovedOrOwner(
        address _owner,
        address _spender,
        uint _tokenId
    ) internal view returns (bool) {
        return
            _spender == _owner ||
            isApprovedForAll[_owner][_spender] ||
            _spender == _approvals[_tokenId];
    }

    function transferFrom(
        address _from,
        address _to,
        uint256 _tokenId
    ) public payable notAddressZero(_to) isOwner(_ownerOf[_tokenId], _from) {
        require(
            _isApprovedOrOwner(_from, msg.sender, _tokenId),
            "not authorize!"
        );

        _balanceOf[_from]--;
        _balanceOf[_to]++;
        _ownerOf[_tokenId] = _to;

        delete _approvals[_tokenId];

        emit Transfer(_from, _to, _tokenId);
    }

    function safeTransferFrom(
        address _from,
        address _to,
        uint256 _tokenId
    ) external payable {
        transferFrom(_from, _to, _tokenId);

        require(
            _to.code.length == 0 ||
                IERC721Receiver(_to).onERC721Received(
                    msg.sender,
                    _from,
                    _tokenId,
                    ""
                ) ==
                IERC721Receiver.onERC721Received.selector,
            "unsafe recipient"
        );
    }

    // @dev safeTransferFrom check if the receiver is a contract it will
    // execute onERC721Received function of it and then check it with
    // selector of onERC721Received function
    function safeTransferFrom(
        address _from,
        address _to,
        uint256 _tokenId,
        bytes memory data
    ) external payable {
        transferFrom(_from, _to, _tokenId);

        require(
            _to.code.length == 0 ||
                IERC721Receiver(_to).onERC721Received(
                    msg.sender,
                    _from,
                    _tokenId,
                    data
                ) ==
                IERC721Receiver.onERC721Received.selector,
            "unsafe recipient"
        );
    }

    function _mint(address _to, uint _tokenId) internal notAddressZero(_to) {
        require(_ownerOf[_tokenId] == address(0), "token exists");

        _balanceOf[_to]++;
        _ownerOf[_tokenId] = _to;

        emit Transfer(address(0), _to, _tokenId);
    }

    function _burn(uint _tokenId) internal {
        address owner = _ownerOf[_tokenId];
        if (owner == address(0)) {
            revert ErrorZeroAddress();
        }

        if (owner != msg.sender) {
            revert ErrorUnauthorize();
        }

        _balanceOf[owner]--;
        delete _ownerOf[_tokenId];
        delete _approvals[_tokenId];

        emit Transfer(owner, address(0), _tokenId);
    }
}

contract NFT is ERC721 {
    // MetadataUpdate event when tokenId change URI
    event MetadataUpdate(uint _tokenId);

    uint private _tokenCount;
    // Optional mapping for token URIs
    mapping(uint256 tokenId => string) private _tokenURIs;

    function _setTokenURI(
        uint256 tokenId,
        string memory _tokenURI
    ) internal virtual {
        _tokenURIs[tokenId] = _tokenURI;
        emit MetadataUpdate(tokenId);
    }

    // mint function mint tokenURI to address and return tokenID
    function mint(
        address _to,
        string memory _tokenURI
    ) external returns (uint) {
        _tokenCount++;
        _mint(_to, _tokenCount);
        _tokenURIs[_tokenCount] = _tokenURI;

        return _tokenCount;
    }

    function burn(uint _tokenId) external {
        require(msg.sender == _ownerOf[_tokenId], "not wner");
        _burn(_tokenId);
    }

    function tokenURI(uint256 _tokenId) public view returns (string memory) {
        return _tokenURIs[_tokenId];
    }
}
