const Identity = artifacts.require("Identity");

contract("Identity", (accounts) => {
    it("Identity", async () => {
    const instance = await Identity.deployed()
    await instance.register("Hello world")
    })
})
