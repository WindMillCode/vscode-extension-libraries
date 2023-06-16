Param (
    [Parameter(Mandatory=$true)] [string] $workspaceLocation=""

)

cd $workspaceLocation
git add .;git commit -m'[CHECKPOINT] before upgrading to next angular version';

cd apps\frontend\AngularApp
$inputText = npx ng update
# Split the input into individual lines
$inputLines = $inputText.Split([Environment]::NewLine)

# Extract the package update information
$packagesToUpdate = $inputLines | Select-String -Pattern "ng update @"

# Build the ng update command
$updateCommand = "npx ng update"
foreach ($package in $packagesToUpdate) {
    $packageGroup = ($package -split "->")[0].Trim()
    $packageName = ($packageGroup -split " ")[0].Trim()
    $updateCommand += " $packageName  "
}

# Remove the trailing "&& " from the command
# $updateCommand = $updateCommand.Substring(0, $updateCommand.Length - 3)
# $updateCommand += " --allow-dirty"
Invoke-Expression $updateCommand


Invoke-Expression "yarn upgrade --dev @faker-js/faker @windmillcode/angular-templates  webpack-bundle-analyzer browserify"
Invoke-Expression "yarn upgrade @windmillcode/wml-components-base  @rxweb/reactive-form-validators @fortawesome/fontawesome-free @compodoc/compodoc  @sentry/angular-ivy @sentry/tracing"
