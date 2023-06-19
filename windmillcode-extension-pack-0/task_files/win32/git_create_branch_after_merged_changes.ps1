Param (
     [string] $workspaceLocation=$PSScriptRoot + '\..\..\..\',
    [string] $sourceBranch ="dev",
    [string] $deleteBranch ="" ,
    [string] $createBranch =""
)




$utilsFile = $PSScriptRoot + '\utils.ps1'
. $utilsFile;

cd $workspaceLocation

$currentBranch = git rev-parse --abbrev-ref HEAD


if ( $deleteBranch -eq "") {
    $deleteBranch = Read-Host -Prompt "the local branch to delete: "
    if ( $deleteBranch -eq "") {
        $deleteBranch =  $currentBranch
    }
}

if ( $createBranch -eq "") {
    $createBranch = Read-Host -Prompt "the local branch to create: "
    if ( $createBranch -eq "") {
        $createBranch =  $currentBranch
    }
}






$command0 = "git checkout $sourceBranch; git pull origin $sourceBranch; git branch -D $deleteBranch; git checkout -b $createBranch"

Invoke-Expression -Command $command0

