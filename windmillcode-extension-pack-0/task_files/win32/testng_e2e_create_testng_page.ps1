Param (
     [string] $workspaceLocation=$PSScriptRoot + '\..\..\..\',
    [string] $pageName=""
)

try {
    $utilsFile = $PSScriptRoot + '\utils.ps1'
    . $utilsFile;

    cd $workspaceLocation
    $pageFolder  = "apps\testing\testng\src\main\java\pages"
    $testFolder =  "apps\testing\testng\src\test\java\e2e"

    cd $pageFolder
    if ( $pageName -eq "") {
        $pageName = Read-Host -Prompt "the name of the page on the website: "
        if ( $pageName -eq "") {
            throw "Your must provide a test name"
        }
    }

    $MyDir = $pageName.ToLower()
    mkdir $pageName.ToLower()
    $myPrefix = (Get-Culture).TextInfo.ToTitleCase($MyDir)

    $myAct = ".\"+$MyDir+"\"+$myPrefix+"ActController.java"
    $myPage = ".\"+$MyDir+"\"+$myPrefix+"Page.java"
    $myVerify = ".\"+$MyDir+"\"+$myPrefix+"VerifyController.java"

    cp   ".\template\TemplateActController.java"    $myAct;
    cp   ".\template\TemplatePage.java"             $myPage;
    cp   ".\template\TemplateVerifyController.java" $myVerify;

    cd $workspaceLocation

    cd $testFolder
    $myTest = $myPrefix + "Test.java"
    cp "TemplateTest.java" $myTest
}
catch {
    Write-Host "An error occurred: $($_.Exception.Message)"
    exit 1
}
