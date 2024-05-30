const {expect} = require("chai");
const {ethers} = require("hardhat");
describe("Fundraising contract", function () {
	let contract;
	this.beforeEach(async function () {
		const Payment = await ethers.getContractFactory("Payment");
		const payment = await Payment.deploy();

		await payment.deployed();
		contract = payment;
	});
	it("Should accept etherium", async function () {
		const [owner] = await ethers.getSigners();
		// Send 50 eth to contract
		await owner.sendTransaction({
			to: contract.address,
			value: ethers.utils.parseEther("50"),
		});

		// Contract balance must be 50 now
		expect(await ethers.provider.getBalance(contract.address)).equal(
			ethers.utils.parseEther("50")
		);
	});
	it("Should create fundraisings", async function () {
		const [owner, addr1, addr2] = await ethers.getSigners();
		const timeStamp = (await ethers.provider.getBlock("latest")).timestamp;
		expect(
			await contract.openNewFundraising(
				"Some description",
				ethers.utils.parseEther("4"),
				timeStamp + 5000
			)
		).ok;
		expect(await contract.connect(addr1).fundraisings(0)).to.be.ok;
		expect(
			await contract.openNewFundraising(
				"Some another",
				ethers.utils.parseEther("3"),
				timeStamp + 5000
			)
		).ok;
		expect(await contract.connect(addr2).fundraisings(1)).to.be.ok;
		expect(
			await contract.openNewFundraising(
				"Some description",
				ethers.utils.parseEther("2"),
				timeStamp + 5000
			)
		).ok;
		expect(contract.fundraisings(2)).to.be.ok;
	});
	it("Fundraisings should accept donations", async function () {
		const [owner, addr1, addr2] = await ethers.getSigners();
		const timeStamp = (await ethers.provider.getBlock("latest")).timestamp;
		await contract.openNewFundraising(
			"Some description",
			ethers.utils.parseEther("4"),
			timeStamp + 5000
		);
		await contract.connect(addr1).donate(0, {
			value: ethers.utils.parseEther("1"),
		});
		let fund = await contract.fundraisings(0);
		let current = Object.values(fund)[12];

		expect(current).equal(ethers.utils.parseEther("1"));

		await contract.connect(addr2).donate(0, {
			value: ethers.utils.parseEther("2"),
		});
		fund = await contract.fundraisings(0);
		current = Object.values(fund)[12];

		expect(current).equal(ethers.utils.parseEther("3"));
	});
	it("Should close fundraising(must fail if already closed)", async function () {
		const owner = ethers.getSigner();
		const timeStamp = (await ethers.provider.getBlock("latest")).timestamp;
		await contract.openNewFundraising(
			"Some description",
			ethers.utils.parseEther("4"),
			timeStamp + 99999999
		);
		expect(contract.closeFundraising(0)).ok;
		await expect(contract.closeFundraising(0)).reverted;
		const fund = await contract.fundraisings(0);
		expect(Object.values(fund)[15]).false;
	});
	it("Shouldn't close fundraising if address is not an owner and closing time not reached", async function () {
		const [owner, addr1] = await ethers.getSigners();
		const timeStamp = (await ethers.provider.getBlock("latest")).timestamp;
		await contract.openNewFundraising(
			"Some description",
			ethers.utils.parseEther("4"),
			timeStamp + 99999999
		);
		await expect(contract.connect(addr1).closeFundraising(0)).revertedWith(
			"You are not an owner or the time is not expired!"
		);
	});
	it("Donation must be reverted if fundraising is closed", async function () {
		const owner = ethers.getSigner();
		const timeStamp = (await ethers.provider.getBlock("latest")).timestamp;
		await contract.openNewFundraising(
			"Some description",
			ethers.utils.parseEther("4"),
			timeStamp + 9999999
		);
		await contract.closeFundraising(0);
		await expect(contract.donate(0, {value: ethers.utils.parseEther("1")}))
			.reverted;
	});
	it("Donation must be reverted if fundraising time expired", async function () {
		const owner = ethers.getSigner();
		const timeStamp = (await ethers.provider.getBlock("latest")).timestamp;
		await contract.openNewFundraising(
			"Some description",
			ethers.utils.parseEther("4"),
			timeStamp + 10000
		);
		ethers.provider.send("evm_increaseTime", [1000]);
		ethers.provider.send("evm_mine");
		expect(contract.donate(0, {value: ethers.utils.parseEther("1")})).ok;
		ethers.provider.send("evm_increaseTime", [10000]);
		ethers.provider.send("evm_mine");
		await expect(contract.donate(0, {value: ethers.utils.parseEther("1")}))
			.reverted;
	});
	it("Fundraising must make a refund after being closed and if goal is not reached", async function () {
		const [owner, addr1, addr2] = await ethers.getSigners();
		const timeStamp = (await ethers.provider.getBlock("latest")).timestamp;
		await contract.openNewFundraising(
			"Some description",
			ethers.utils.parseEther("6"),
			timeStamp + 10000
		);
		await contract.connect(addr1).donate(0, {
			value: ethers.utils.parseEther("2"),
		});
		await contract.connect(addr2).donate(0, {
			value: ethers.utils.parseEther("3"),
		});
		const payBalance1 = await ethers.provider.getBalance(addr1.address);
		const payBalance2 = await ethers.provider.getBalance(addr2.address);
		await contract.closeFundraising(0);
		expect(await ethers.provider.getBalance(addr1.address)).greaterThan(
			payBalance1
		);
		expect(await ethers.provider.getBalance(addr2.address)).greaterThan(
			payBalance2
		);
	});
	it("Fundraising must transfer etherium to owner if it is closed and goal reached", async function () {
		const [owner, addr1, addr2] = await ethers.getSigners();
		const timeStamp = (await ethers.provider.getBlock("latest")).timestamp;
		const initialValue = await ethers.provider.getBalance(owner.address);
		await contract.openNewFundraising(
			"Some description",
			ethers.utils.parseEther("5"),
			timeStamp + 10000
		);
		await contract.connect(addr1).donate(0, {
			value: ethers.utils.parseEther("2"),
		});
		await contract.connect(addr2).donate(0, {
			value: ethers.utils.parseEther("3"),
		});
		let fund = await contract.fundraisings(0);
		await contract.closeFundraising(0);
		fund = await contract.fundraisings(0);
		expect(await ethers.provider.getBalance(owner.address)).greaterThan(
			initialValue
		);
	});
});
