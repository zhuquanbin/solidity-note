const ABCToken = artifacts.require("ABCToken");

contract('ABCToken', function(accounts) {
    let owner   = accounts[0];
    let wallet1 = accounts[1];
    let wallet2 = accounts[2];
    let TestInstance = null;
    
    beforeEach('setup contract for each test', async() => {
        TestInstance = await ABCToken.new("AbcCoin","ABC", 10000, 8);
    })

    it('1) 检测初始化参数', async() => {
        assert.equal(await TestInstance.name(),"AbcCoin");
        assert.equal(await TestInstance.symbol(), "ABC");
        assert.equal(await TestInstance.totalSupply(), 10 ** 12);
        assert.equal(await TestInstance.owner(), owner);
    })

    it('2) 钱包转账测试', async() => {
        let balance1 = await TestInstance.balanceOf.call(owner)
        assert.equal(balance1.toNumber(), 10 ** 12);

        await TestInstance.transfer(wallet1, 1.2 * 10**8);
        
        let balance2 = await TestInstance.balanceOf(owner);
        assert.equal(balance2.toNumber(), 10 ** 12 - 1.2 * 10**8);

        let balance3 = await TestInstance.balanceOf(wallet1);
        assert.equal(balance3.toNumber(), 1.2 * 10**8);
    })

    it('3) 代币铸造权限检测', async() => {

        try{
            await TestInstance.mint(wallet1, 0.3 * 10**8, {from: wallet2});
        }catch(error){
            const revertFound = error.message.search('revert') >= 0;
            assert(revertFound, `Expected "revert", got ${error} instead`);
        }
        let balance1 = await TestInstance.balanceOf(wallet1);
        assert.equal(balance1.toNumber(), 0);
    })

    it('4) 代币铸造测试', async() => {

        try{
            await TestInstance.mint(wallet1, 0.3 * 10**8);
        }catch(error){
            const revertFound = error.message.search('revert') >= 0;
            assert(revertFound, `Expected "revert", got ${error} instead`);
        }
        let balance1 = await TestInstance.balanceOf(wallet1);
        assert.equal(balance1.toNumber(), 3 * 10**7);

        assert.equal(await TestInstance.totalSupply(), 10 ** 12 + 0.3 * 10 ** 8);
    })

    it('5) 代币销毁权限检测', async() => {
        await TestInstance.transfer(wallet1, 5 * 10**8);
        try{
            await TestInstance.burn(wallet1, 2 * 10**8, {from: wallet2});
        }catch(error){
            const revertFound = error.message.search('revert') >= 0;
            assert(revertFound, `Expected "revert", got ${error} instead`);
        }
        let balance1 = await TestInstance.balanceOf(wallet1);
        assert.equal(balance1.toNumber(), 5 * 10 ** 8);
    })

    it('6) 代币销毁测试', async() => {
        await TestInstance.transfer(wallet1, 5 * 10**8);
        try{
            await TestInstance.burn(wallet1, 2 * 10**8);
        }catch(error){
            const revertFound = error.message.search('revert') >= 0;
            assert(revertFound, `Expected "revert", got ${error} instead`);
        }
        let balance1 = await TestInstance.balanceOf(wallet1);
        assert.equal(balance1.toNumber(), 3 * 10**8);

        assert.equal(await TestInstance.totalSupply(), 10 ** 12 - 2 * 10 ** 8);
    })
});
