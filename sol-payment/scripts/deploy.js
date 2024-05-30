const {ethers} = require("hardhat");

async function main() {
	const [deployer] = await ethers.getSigners();

	console.log("Deploying contracts with the account:", deployer.address);

	console.log("Account balance:", (await deployer.getBalance()).toString());

	const Payment = await ethers.getContractFactory("Payment");
	const payment = await Payment.deploy();

	console.log("Contract address:", payment.address);
}

main()
	.then(() => process.exit(0))
	.catch((error) => {
		console.error(error);
		process.exit(1);
	});
