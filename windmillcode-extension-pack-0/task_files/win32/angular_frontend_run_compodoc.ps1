Param (
    [Parameter(Mandatory=$true)] [string] $workspaceLocation=""

)


cd $workspaceLocation
cd .\\apps\\zero\\frontend\\AngularApp\\;
yarn compodoc:build-and-serve
