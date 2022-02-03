mkdir -p temp
cd temp
svn checkout https://github.com/NetBlockGit/smart-contracts/trunk/contracts \
                    --username=$GITPOD_GIT_USER_NAME --password=$GITHUB_TOKEN
solc --abi contracts/Blocklist.sol -o out
abigen --pkg blocklist --out=../generated/smartcontract/blocklist/blocklist.go --abi out/Blocklist.abi 
cd -
rm -rf temp