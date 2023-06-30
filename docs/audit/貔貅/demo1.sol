// 极简貔貅ERC20代币，只能买，不能卖
contract HoneyPot is ERC20, Ownable {
    address public pair;

    // 构造函数：初始化代币名称和代号
    constructor() ERC20("HoneyPot", "Pi Xiu") {
        address factory = 0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f; // goerli uniswap v2 factory
        address tokenA = address(this); // 貔貅代币地址
        address tokenB = 0xB4FBF271143F4FBf7B91A5ded31805e42b2208d6; //  goerli WETH
        (address token0, address token1) = tokenA < tokenB ? (tokenA, tokenB) : (tokenB, tokenA); //将tokenA和tokenB按大小排序
        bytes32 salt = keccak256(abi.encodePacked(token0, token1));
        // calculate pair address
        pair = address(uint160(uint(keccak256(abi.encodePacked(
        hex'ff',
        factory,
        salt,
        hex'96e8ac4277198ff8b6f785478aa9a39f403cb768dd02cbee326c3e7da348845f'
        )))));
    }
    
    /**
     * 铸造函数，只有合约所有者可以调用
     */
    function mint(address to, uint amount) public onlyOwner {
        _mint(to, amount);
    }

    /**
     * @dev See {ERC20-_beforeTokenTransfer}.
     * 貔貅函数：只有合约拥有者可以卖出
     */
    function _beforeTokenTransfer(
        address from,
        address to,
        uint256 amount
    ) internal virtual override {
        super._beforeTokenTransfer(from, to, amount);
        // 当转账的目标地址为 LP 时，会revert
        if(to == pair){
            require(from == owner(), "Can not Transfer");
        }
    }
}

/*
https://www.playbtc.cn/story/89942.html

‘貔貅币合约最大的特点就是：只有合约拥有者能够卖出代币，其他人无法卖出。对此，需要对应使用到三种函数：构造函数、铸造函数以及转账函数。

    构造函数：初始化代币的名称和代号，并根据去中心化交易所的的原理计算LP合约地址，这个地址会在 _beforeTokenTransfer() 函数中用到。

    铸造函数mint()：仅 owner 地址（合约拥有者）可以调用，用于铸造貔貅代币。

    _beforeTokenTransfer()：ERC20代币在被转账前会调用的函数。在其中，我们限制了当转账的目标地址 to 为 LP 的时候，也就是其他持有者卖出的时候，交易会 revert；只有调用者为owner的时候能够成功。这也是貔貅合约的核心。

*/