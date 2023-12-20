const hre = require('hardhat');

async function main() {
  const store = await hre.ethers.deployContract('Store');
  await store.waitForDeployment();
  console.log(`Store deployed to ${store.target}`);

  // send transaction
  const addr = '0x9c570A8e0783Adf20b982172999Eccad59a17518';
  const tx = await store.add(2);
  console.log('tx:', tx.hash);

  await sleep(20000);
  const num = await store.getCount(addr);
  console.log('count:', num);

  // send transaction
  const tx2 = await store.testTime();
  console.log('tx:', tx2.hash);

  await sleep(20000);
  const num2 = await store.getCount(addr);
  console.log('count:', num2);
}

function sleep(ms) {
  return new Promise(resolveFunc => setTimeout(resolveFunc, ms));
}

main().catch((error) => {
  console.error(error);
  process.exit(1);
});