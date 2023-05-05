// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.4;

contract Coin {
    // 关键字 "public" 使变量可以从其他合约中访问。
    address public minter;
    mapping (address => uint) public balances;

    // 事件允许客户端对您声明的特定合约变化做出反应
    event Sent(address from, address to, uint amount);

    // 构造函数代码只有在合约创建时运行
    constructor() {
        minter = msg.sender; // msg.sender 当前合约调用者
    }

    // 向一个地址发送一定数量的新创建的代币
    // 但只能由合约创建者调用
    function mint(address receiver, uint amount) public {
        require(msg.sender == minter);
        balances[receiver] += amount;
    }

    // 错误类型变量允许您提供关于操作失败原因的信息。
    // 它们会返回给函数的调用者。
    error InsufficientBalance(uint requested, uint available);

    // 从任何调用者那里发送一定数量的代币到一个地址
    function send(address receiver, uint amount) public {
        if (amount > balances[msg.sender])
            revert InsufficientBalance({
                requested: amount,
                available: balances[msg.sender]
            });

        balances[msg.sender] -= amount;
        balances[receiver] += amount;
        emit Sent(msg.sender, receiver, amount);
    }
}

