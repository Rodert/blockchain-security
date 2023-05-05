// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.4.16 <0.9.0;

contract SimpleStorage {
    uint storedData;

    function set(uint x) public {
        storedData = x;
    }

    function get() public view returns (uint) {
        return storedData;
    }
}


// --- ---

// 入门

// Solidity意义上的合约是代码（其 函数）和数据（其 状态）的集合， 驻留在以太坊区块链的一个特定地址。

// SPDX-License-Identifier: GPL-3.0 // 开源软件协议
pragma solidity >=0.4.16 <0.9.0; // 编码版本

contract SimpleStorage {
    uint storedData; // 定义一个状态（因为是写在连上，所以不仅仅是一个变量）

    function set(uint x) public {
        storedData = x;
    }

    function get() public view returns (uint) {
        return storedData;
    }
}




