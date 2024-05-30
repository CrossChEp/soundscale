// SPDX-License-Identifier: GPL-3.0
pragma solidity 0.8.15;

struct Fundraising {
    uint256 id;
    address payable owner;
    string description;
    uint256 goal; // in wei
    uint256 current; // in wei
    address payable[] payers;
    uint256 closingTime;
    bool goalReached;
    bool isOpened;
}

contract Payment {
    address immutable contractOwner;

    Fundraising[] public fundraisings;
    // fund_id => mapping(address => contribution in wei)
    mapping(uint256 => mapping(address => uint256)) contributions;
    event FundraisingOpened(
        address owner,
        uint256 goal,
        string description,
        uint256 closingTime
    );
    event DonationAccepted(
        address sender,
        uint256 amount,
        uint256 fundraisingId
    );
    event Received(address from, uint256 value);
    event Refunded(address to, uint256 value, uint256 fundraisingId);
    modifier checkIfOpen(uint256 fund_id) {
        if (
            fundraisings[fund_id].closingTime <= block.timestamp ||
            !fundraisings[fund_id].isOpened
        ) {
            revert("Fundraising is closed!");
        }

        _;
    }

    constructor() payable {
        contractOwner = msg.sender;
    }

    receive() external payable {
        if (msg.sender != contractOwner) {
            revert(
                'You are not allowed to send eth directly into contract. If you want to make a donation use method "donation"'
            );
        }
        emit Received(msg.sender, msg.value);
    }

    function openNewFundraising(
        string calldata _description,
        uint256 _goal,
        uint256 _closingTime
    ) public {
        require(
            bytes(_description).length != 0,
            "Description can not be null!"
        );
        require(_goal != 0, "Goal can not be null!");
        require(
            _closingTime > block.timestamp,
            "Closing time must be greater then current!"
        );
        require(_goal <= 6 * 10 ** 18, "Maximal goal must be lower then 6 ETH");
        require(
            _goal >= 588000000000000,
            "Minimal goal must be higher then 0,000588 ETH"
        );

        address payable[] memory _payers;

        Fundraising memory fund = Fundraising({
            id: fundraisings.length,
            owner: payable(msg.sender),
            description: _description,
            goal: _goal,
            current: 0,
            payers: _payers,
            closingTime: _closingTime,
            isOpened: true,
            goalReached: false
        });
        fundraisings.push(fund);
        emit FundraisingOpened(msg.sender, _goal, _description, _closingTime);
    }

    function donate(
        uint256 fundraisingId
    ) public payable checkIfOpen(fundraisingId) {
        require(
            fundraisingId < fundraisings.length,
            "Given fundraising id is not exists!"
        );
        require(
            fundraisings[fundraisingId].isOpened == true,
            "Fundraising is closed!"
        );
        require(
            msg.value > 6000000000000,
            "Sending value must be greater then 0,000006 ETH!"
        );

        fundraisings[fundraisingId].current += msg.value;
        if (contributions[fundraisingId][msg.sender] == 0) {
            fundraisings[fundraisingId].payers.push(payable(msg.sender));
        }

        contributions[fundraisingId][msg.sender] += msg.value;
        emit DonationAccepted(msg.sender, msg.value, fundraisingId);
    }

    function makeRefund(uint256 fundraisingId) internal {
        Fundraising storage fund = fundraisings[fundraisingId];

        for (
            uint256 payer_index = 0;
            payer_index < fund.payers.length;
            ++payer_index
        ) {
            address payable payer = fund.payers[payer_index];
            payer.transfer(contributions[fund.id][payer]);
            emit Refunded(payer, contributions[fund.id][payer], fundraisingId);
        }
    }

    function closeFundraising(uint256 _id) public {
        require(fundraisings[_id].isOpened, "Fundraising is already closed!");
        Fundraising storage fund = fundraisings[_id];

        if (fund.closingTime <= block.timestamp) {
            fund.isOpened = false;
            if (fund.goal <= fund.current) {
                fund.goalReached = true;
                fund.owner.transfer(fund.current);
            } else {
                makeRefund(_id);
            }
        } else if (fund.owner == msg.sender) {
            fund.isOpened = false;
            if (fund.goal <= fund.current) {
                fund.goalReached = true;
                fund.owner.transfer(fund.current);
            } else {
                makeRefund(_id);
            }
        } else {
            revert("You are not an owner or the time is not expired!");
        }
    }
}
