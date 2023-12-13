const hre = require('hardhat');

async function main() {
  const store = await hre.ethers.deployContract('Store');
  await store.waitForDeployment();
  console.log(`Store deployed to ${store.target}`);
}

main().catch((error) => {
  console.error(error);
  process.exit(1);
});